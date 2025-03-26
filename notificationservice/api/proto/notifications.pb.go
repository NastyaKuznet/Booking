// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.29.4
// source: api/proto/notifications.proto

package notifications

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Запрос на отправку уведомления
type BookingNotificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId       string `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`                     // ID номера
	GuestName    string `protobuf:"bytes,2,opt,name=guest_name,json=guestName,proto3" json:"guest_name,omitempty"`            // Имя гостя
	CheckInDate  string `protobuf:"bytes,3,opt,name=check_in_date,json=checkInDate,proto3" json:"check_in_date,omitempty"`    // Дата заезда
	CheckOutDate string `protobuf:"bytes,4,opt,name=check_out_date,json=checkOutDate,proto3" json:"check_out_date,omitempty"` // Дата выезда
}

func (x *BookingNotificationRequest) Reset() {
	*x = BookingNotificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_notifications_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingNotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingNotificationRequest) ProtoMessage() {}

func (x *BookingNotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_notifications_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingNotificationRequest.ProtoReflect.Descriptor instead.
func (*BookingNotificationRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_notifications_proto_rawDescGZIP(), []int{0}
}

func (x *BookingNotificationRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *BookingNotificationRequest) GetGuestName() string {
	if x != nil {
		return x.GuestName
	}
	return ""
}

func (x *BookingNotificationRequest) GetCheckInDate() string {
	if x != nil {
		return x.CheckInDate
	}
	return ""
}

func (x *BookingNotificationRequest) GetCheckOutDate() string {
	if x != nil {
		return x.CheckOutDate
	}
	return ""
}

// Ответ после отправки уведомления
type BookingNotificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // Успешность отправки
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`  // Сообщение
}

func (x *BookingNotificationResponse) Reset() {
	*x = BookingNotificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_notifications_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingNotificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingNotificationResponse) ProtoMessage() {}

func (x *BookingNotificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_notifications_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingNotificationResponse.ProtoReflect.Descriptor instead.
func (*BookingNotificationResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_notifications_proto_rawDescGZIP(), []int{1}
}

func (x *BookingNotificationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *BookingNotificationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_proto_notifications_proto protoreflect.FileDescriptor

var file_api_proto_notifications_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x9e,
	0x01, 0x0a, 0x1a, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x75, 0x65, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x49, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x65, 0x22,
	0x51, 0x0a, 0x1b, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x87, 0x01, 0x0a, 0x13, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x70, 0x0a, 0x17, 0x53, 0x65,
	0x6e, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2a, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10,
	0x2e, 0x2f, 0x3b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_notifications_proto_rawDescOnce sync.Once
	file_api_proto_notifications_proto_rawDescData = file_api_proto_notifications_proto_rawDesc
)

func file_api_proto_notifications_proto_rawDescGZIP() []byte {
	file_api_proto_notifications_proto_rawDescOnce.Do(func() {
		file_api_proto_notifications_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_notifications_proto_rawDescData)
	})
	return file_api_proto_notifications_proto_rawDescData
}

var file_api_proto_notifications_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_proto_notifications_proto_goTypes = []interface{}{
	(*BookingNotificationRequest)(nil),  // 0: notifications.BookingNotificationRequest
	(*BookingNotificationResponse)(nil), // 1: notifications.BookingNotificationResponse
}
var file_api_proto_notifications_proto_depIdxs = []int32{
	0, // 0: notifications.NotificationService.SendBookingNotification:input_type -> notifications.BookingNotificationRequest
	1, // 1: notifications.NotificationService.SendBookingNotification:output_type -> notifications.BookingNotificationResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_notifications_proto_init() }
func file_api_proto_notifications_proto_init() {
	if File_api_proto_notifications_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_notifications_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingNotificationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_notifications_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingNotificationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_notifications_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_notifications_proto_goTypes,
		DependencyIndexes: file_api_proto_notifications_proto_depIdxs,
		MessageInfos:      file_api_proto_notifications_proto_msgTypes,
	}.Build()
	File_api_proto_notifications_proto = out.File
	file_api_proto_notifications_proto_rawDesc = nil
	file_api_proto_notifications_proto_goTypes = nil
	file_api_proto_notifications_proto_depIdxs = nil
}
