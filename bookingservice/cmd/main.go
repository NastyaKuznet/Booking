package main

import (
	api "bookingservice/order"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
)

const connection = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"

func main() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	conString := fmt.Sprintf(connection, "postgres", "5432", "postgres", "postgres", "postgres", "disable")
	conn, err := pgxpool.New(context.Background(), conString)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterBookingServiceServer(s, &Server{db: conn})
	log.Println("Starting bookingservice server...")
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type Server struct {
	api.UnimplementedBookingServiceServer
	db *pgxpool.Pool
}

// Реализация методов gRPC
func (s *Server) GetAllRooms(ctx context.Context, req *api.GetAllRoomsRequest) (*api.GetAllRoomsResponse, error) {
	rows, err := s.db.Query(ctx, `SELECT id, room_number, description, available, price FROM rooms`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []*api.Room
	for rows.Next() {
		var room api.Room
		if err := rows.Scan(&room.Id, &room.RoomNumber, &room.Description, &room.Available, &room.Price); err != nil {
			return nil, err
		}
		rooms = append(rooms, &room)
	}

	return &api.GetAllRoomsResponse{Rooms: rooms}, nil
}

func (s *Server) GetAvailableRooms(ctx context.Context, req *api.GetAvailableRoomsRequest) (*api.GetAvailableRoomsResponse, error) {
	rows, err := s.db.Query(ctx, `
		SELECT id, room_number, description, available, price 
		FROM rooms 
		WHERE available = true 
		AND id NOT IN (
			SELECT room_id 
			FROM bookings 
			WHERE (start_date <= $1 AND end_date >= $2)
		)`, req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []*api.Room
	for rows.Next() {
		var room api.Room
		if err := rows.Scan(&room.Id, &room.RoomNumber, &room.Description, &room.Available, &room.Price); err != nil {
			return nil, err
		}
		rooms = append(rooms, &room)
	}

	return &api.GetAvailableRoomsResponse{Rooms: rooms}, nil
}

func (s *Server) BookRoom(ctx context.Context, req *api.BookRoomRequest) (*api.BookRoomResponse, error) {
	var bookingID int64
	err := s.db.QueryRow(ctx, `
		INSERT INTO bookings (room_id, user_id, start_date, end_date) 
		VALUES ($1, $2, $3, $4)
		RETURNING id`, req.RoomId, req.UserId, req.StartDate, req.EndDate).Scan(&bookingID)
	if err != nil {
		return &api.BookRoomResponse{Success: false, Message: "Failed to book room"}, err
	}

	return &api.BookRoomResponse{Success: true, Message: "Room booked successfully", BookingId: bookingID}, nil
}

func (s *Server) CancelBooking(ctx context.Context, req *api.CancelBookingRequest) (*api.CancelBookingResponse, error) {
	_, err := s.db.Exec(ctx, `DELETE FROM bookings WHERE id = $1`, req.BookingId)
	if err != nil {
		return &api.CancelBookingResponse{Success: false, Message: "Failed to cancel booking"}, err
	}

	return &api.CancelBookingResponse{Success: true, Message: "Booking canceled successfully"}, nil
}
