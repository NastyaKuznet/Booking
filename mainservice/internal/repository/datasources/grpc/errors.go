package grpc

import "errors"

var (
	ErrRoomsNotFound = errors.New("rooms not found")
	ErrInternal      = errors.New("internal error")
)
