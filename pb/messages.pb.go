// Code generated by protoc-gen-go.
// source: messages.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	Employee
	Vacation
	GetAllRequest
	GetByBadgeNumberRequest
	EmployeeRequest
	EmployeeResponse
	AddPhotoRequest
	AddPhotoResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Employee struct {
	Id                  int32       `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	BadgeNumber         int32       `protobuf:"varint,2,opt,name=badgeNumber" json:"badgeNumber,omitempty"`
	FirstName           string      `protobuf:"bytes,3,opt,name=firstName" json:"firstName,omitempty"`
	LastName            string      `protobuf:"bytes,4,opt,name=lastName" json:"lastName,omitempty"`
	VacationAccrualRate float32     `protobuf:"fixed32,5,opt,name=vacationAccrualRate" json:"vacationAccrualRate,omitempty"`
	VacationAccrued     float32     `protobuf:"fixed32,6,opt,name=vacationAccrued" json:"vacationAccrued,omitempty"`
	Vacations           []*Vacation `protobuf:"bytes,7,rep,name=vacations" json:"vacations,omitempty"`
}

func (m *Employee) Reset()                    { *m = Employee{} }
func (m *Employee) String() string            { return proto.CompactTextString(m) }
func (*Employee) ProtoMessage()               {}
func (*Employee) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Employee) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Employee) GetBadgeNumber() int32 {
	if m != nil {
		return m.BadgeNumber
	}
	return 0
}

func (m *Employee) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Employee) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Employee) GetVacationAccrualRate() float32 {
	if m != nil {
		return m.VacationAccrualRate
	}
	return 0
}

func (m *Employee) GetVacationAccrued() float32 {
	if m != nil {
		return m.VacationAccrued
	}
	return 0
}

func (m *Employee) GetVacations() []*Vacation {
	if m != nil {
		return m.Vacations
	}
	return nil
}

type Vacation struct {
	Id          int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	StartDate   int64   `protobuf:"varint,2,opt,name=startDate" json:"startDate,omitempty"`
	Duration    float32 `protobuf:"fixed32,3,opt,name=duration" json:"duration,omitempty"`
	IsCancelled bool    `protobuf:"varint,4,opt,name=isCancelled" json:"isCancelled,omitempty"`
}

func (m *Vacation) Reset()                    { *m = Vacation{} }
func (m *Vacation) String() string            { return proto.CompactTextString(m) }
func (*Vacation) ProtoMessage()               {}
func (*Vacation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Vacation) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Vacation) GetStartDate() int64 {
	if m != nil {
		return m.StartDate
	}
	return 0
}

func (m *Vacation) GetDuration() float32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Vacation) GetIsCancelled() bool {
	if m != nil {
		return m.IsCancelled
	}
	return false
}

type GetAllRequest struct {
}

func (m *GetAllRequest) Reset()                    { *m = GetAllRequest{} }
func (m *GetAllRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAllRequest) ProtoMessage()               {}
func (*GetAllRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type GetByBadgeNumberRequest struct {
	BadgeNumber int32 `protobuf:"varint,1,opt,name=badgeNumber" json:"badgeNumber,omitempty"`
}

func (m *GetByBadgeNumberRequest) Reset()                    { *m = GetByBadgeNumberRequest{} }
func (m *GetByBadgeNumberRequest) String() string            { return proto.CompactTextString(m) }
func (*GetByBadgeNumberRequest) ProtoMessage()               {}
func (*GetByBadgeNumberRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetByBadgeNumberRequest) GetBadgeNumber() int32 {
	if m != nil {
		return m.BadgeNumber
	}
	return 0
}

type EmployeeRequest struct {
	Employee *Employee `protobuf:"bytes,1,opt,name=employee" json:"employee,omitempty"`
}

func (m *EmployeeRequest) Reset()                    { *m = EmployeeRequest{} }
func (m *EmployeeRequest) String() string            { return proto.CompactTextString(m) }
func (*EmployeeRequest) ProtoMessage()               {}
func (*EmployeeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *EmployeeRequest) GetEmployee() *Employee {
	if m != nil {
		return m.Employee
	}
	return nil
}

type EmployeeResponse struct {
	Employee *Employee `protobuf:"bytes,1,opt,name=employee" json:"employee,omitempty"`
}

func (m *EmployeeResponse) Reset()                    { *m = EmployeeResponse{} }
func (m *EmployeeResponse) String() string            { return proto.CompactTextString(m) }
func (*EmployeeResponse) ProtoMessage()               {}
func (*EmployeeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *EmployeeResponse) GetEmployee() *Employee {
	if m != nil {
		return m.Employee
	}
	return nil
}

type AddPhotoRequest struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *AddPhotoRequest) Reset()                    { *m = AddPhotoRequest{} }
func (m *AddPhotoRequest) String() string            { return proto.CompactTextString(m) }
func (*AddPhotoRequest) ProtoMessage()               {}
func (*AddPhotoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AddPhotoRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type AddPhotoResponse struct {
	IsOk bool `protobuf:"varint,1,opt,name=isOk" json:"isOk,omitempty"`
}

func (m *AddPhotoResponse) Reset()                    { *m = AddPhotoResponse{} }
func (m *AddPhotoResponse) String() string            { return proto.CompactTextString(m) }
func (*AddPhotoResponse) ProtoMessage()               {}
func (*AddPhotoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AddPhotoResponse) GetIsOk() bool {
	if m != nil {
		return m.IsOk
	}
	return false
}

func init() {
	proto.RegisterType((*Employee)(nil), "Employee")
	proto.RegisterType((*Vacation)(nil), "Vacation")
	proto.RegisterType((*GetAllRequest)(nil), "GetAllRequest")
	proto.RegisterType((*GetByBadgeNumberRequest)(nil), "GetByBadgeNumberRequest")
	proto.RegisterType((*EmployeeRequest)(nil), "EmployeeRequest")
	proto.RegisterType((*EmployeeResponse)(nil), "EmployeeResponse")
	proto.RegisterType((*AddPhotoRequest)(nil), "AddPhotoRequest")
	proto.RegisterType((*AddPhotoResponse)(nil), "AddPhotoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EmployeeService service

type EmployeeServiceClient interface {
	GetByBadgeNumber(ctx context.Context, in *GetByBadgeNumberRequest, opts ...grpc.CallOption) (*EmployeeResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (EmployeeService_GetAllClient, error)
	Save(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*EmployeeResponse, error)
	SaveAll(ctx context.Context, opts ...grpc.CallOption) (EmployeeService_SaveAllClient, error)
	AddPhoto(ctx context.Context, opts ...grpc.CallOption) (EmployeeService_AddPhotoClient, error)
}

type employeeServiceClient struct {
	cc *grpc.ClientConn
}

func NewEmployeeServiceClient(cc *grpc.ClientConn) EmployeeServiceClient {
	return &employeeServiceClient{cc}
}

func (c *employeeServiceClient) GetByBadgeNumber(ctx context.Context, in *GetByBadgeNumberRequest, opts ...grpc.CallOption) (*EmployeeResponse, error) {
	out := new(EmployeeResponse)
	err := grpc.Invoke(ctx, "/EmployeeService/GetByBadgeNumber", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (EmployeeService_GetAllClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EmployeeService_serviceDesc.Streams[0], c.cc, "/EmployeeService/GetAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &employeeServiceGetAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EmployeeService_GetAllClient interface {
	Recv() (*EmployeeResponse, error)
	grpc.ClientStream
}

type employeeServiceGetAllClient struct {
	grpc.ClientStream
}

func (x *employeeServiceGetAllClient) Recv() (*EmployeeResponse, error) {
	m := new(EmployeeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *employeeServiceClient) Save(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*EmployeeResponse, error) {
	out := new(EmployeeResponse)
	err := grpc.Invoke(ctx, "/EmployeeService/Save", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *employeeServiceClient) SaveAll(ctx context.Context, opts ...grpc.CallOption) (EmployeeService_SaveAllClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EmployeeService_serviceDesc.Streams[1], c.cc, "/EmployeeService/SaveAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &employeeServiceSaveAllClient{stream}
	return x, nil
}

type EmployeeService_SaveAllClient interface {
	Send(*EmployeeRequest) error
	Recv() (*EmployeeResponse, error)
	grpc.ClientStream
}

type employeeServiceSaveAllClient struct {
	grpc.ClientStream
}

func (x *employeeServiceSaveAllClient) Send(m *EmployeeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *employeeServiceSaveAllClient) Recv() (*EmployeeResponse, error) {
	m := new(EmployeeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *employeeServiceClient) AddPhoto(ctx context.Context, opts ...grpc.CallOption) (EmployeeService_AddPhotoClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EmployeeService_serviceDesc.Streams[2], c.cc, "/EmployeeService/AddPhoto", opts...)
	if err != nil {
		return nil, err
	}
	x := &employeeServiceAddPhotoClient{stream}
	return x, nil
}

type EmployeeService_AddPhotoClient interface {
	Send(*AddPhotoRequest) error
	CloseAndRecv() (*AddPhotoResponse, error)
	grpc.ClientStream
}

type employeeServiceAddPhotoClient struct {
	grpc.ClientStream
}

func (x *employeeServiceAddPhotoClient) Send(m *AddPhotoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *employeeServiceAddPhotoClient) CloseAndRecv() (*AddPhotoResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AddPhotoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for EmployeeService service

type EmployeeServiceServer interface {
	GetByBadgeNumber(context.Context, *GetByBadgeNumberRequest) (*EmployeeResponse, error)
	GetAll(*GetAllRequest, EmployeeService_GetAllServer) error
	Save(context.Context, *EmployeeRequest) (*EmployeeResponse, error)
	SaveAll(EmployeeService_SaveAllServer) error
	AddPhoto(EmployeeService_AddPhotoServer) error
}

func RegisterEmployeeServiceServer(s *grpc.Server, srv EmployeeServiceServer) {
	s.RegisterService(&_EmployeeService_serviceDesc, srv)
}

func _EmployeeService_GetByBadgeNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByBadgeNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).GetByBadgeNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EmployeeService/GetByBadgeNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).GetByBadgeNumber(ctx, req.(*GetByBadgeNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_GetAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAllRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EmployeeServiceServer).GetAll(m, &employeeServiceGetAllServer{stream})
}

type EmployeeService_GetAllServer interface {
	Send(*EmployeeResponse) error
	grpc.ServerStream
}

type employeeServiceGetAllServer struct {
	grpc.ServerStream
}

func (x *employeeServiceGetAllServer) Send(m *EmployeeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EmployeeService_Save_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).Save(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EmployeeService/Save",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).Save(ctx, req.(*EmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmployeeService_SaveAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EmployeeServiceServer).SaveAll(&employeeServiceSaveAllServer{stream})
}

type EmployeeService_SaveAllServer interface {
	Send(*EmployeeResponse) error
	Recv() (*EmployeeRequest, error)
	grpc.ServerStream
}

type employeeServiceSaveAllServer struct {
	grpc.ServerStream
}

func (x *employeeServiceSaveAllServer) Send(m *EmployeeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *employeeServiceSaveAllServer) Recv() (*EmployeeRequest, error) {
	m := new(EmployeeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EmployeeService_AddPhoto_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EmployeeServiceServer).AddPhoto(&employeeServiceAddPhotoServer{stream})
}

type EmployeeService_AddPhotoServer interface {
	SendAndClose(*AddPhotoResponse) error
	Recv() (*AddPhotoRequest, error)
	grpc.ServerStream
}

type employeeServiceAddPhotoServer struct {
	grpc.ServerStream
}

func (x *employeeServiceAddPhotoServer) SendAndClose(m *AddPhotoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *employeeServiceAddPhotoServer) Recv() (*AddPhotoRequest, error) {
	m := new(AddPhotoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _EmployeeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "EmployeeService",
	HandlerType: (*EmployeeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByBadgeNumber",
			Handler:    _EmployeeService_GetByBadgeNumber_Handler,
		},
		{
			MethodName: "Save",
			Handler:    _EmployeeService_Save_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAll",
			Handler:       _EmployeeService_GetAll_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SaveAll",
			Handler:       _EmployeeService_SaveAll_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "AddPhoto",
			Handler:       _EmployeeService_AddPhoto_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "messages.proto",
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x8f, 0x94, 0x30,
	0x18, 0x0e, 0x0c, 0x3b, 0x0b, 0xef, 0xea, 0x80, 0xf5, 0x20, 0xd9, 0x78, 0x98, 0x90, 0xac, 0x92,
	0x18, 0xc9, 0x3a, 0x5e, 0x34, 0x1e, 0xcc, 0x8c, 0x1a, 0x6f, 0xab, 0xe9, 0x26, 0x1e, 0xbc, 0x15,
	0x78, 0x5d, 0x89, 0xb0, 0x20, 0x2d, 0x24, 0xfb, 0x43, 0xfc, 0xb5, 0x5e, 0x6c, 0x2b, 0x1f, 0x33,
	0xcc, 0x6c, 0x32, 0xb7, 0xf6, 0xf9, 0x68, 0xfb, 0x3c, 0x2f, 0xc0, 0xa2, 0x40, 0xce, 0xd9, 0x0d,
	0xf2, 0xa8, 0xaa, 0x4b, 0x51, 0x06, 0x7f, 0x0d, 0xb0, 0x3f, 0x15, 0x55, 0x5e, 0xde, 0x21, 0x92,
	0x05, 0x98, 0x59, 0xea, 0x1b, 0x4b, 0x23, 0x3c, 0xa1, 0x72, 0x45, 0x96, 0x70, 0x16, 0xb3, 0xf4,
	0x06, 0xaf, 0x9a, 0x22, 0xc6, 0xda, 0x37, 0x35, 0xb1, 0x0d, 0x91, 0xa7, 0xe0, 0xfc, 0xc8, 0x6a,
	0x2e, 0xae, 0x58, 0x81, 0xfe, 0x4c, 0xf2, 0x0e, 0x1d, 0x01, 0x72, 0x0e, 0x76, 0xce, 0x3a, 0xd2,
	0xd2, 0xe4, 0xb0, 0x27, 0x97, 0xf0, 0xb8, 0x65, 0x09, 0x13, 0x59, 0x79, 0xbb, 0x4e, 0x92, 0xba,
	0x61, 0x39, 0x65, 0x02, 0xfd, 0x13, 0x29, 0x33, 0xe9, 0x21, 0x8a, 0x84, 0xe0, 0xee, 0xc0, 0x98,
	0xfa, 0x73, 0xad, 0x9e, 0xc2, 0xe4, 0x39, 0x38, 0x3d, 0xc4, 0xfd, 0xd3, 0xe5, 0x2c, 0x3c, 0x5b,
	0x39, 0xd1, 0xb7, 0x0e, 0xa1, 0x23, 0x17, 0xb4, 0x60, 0xf7, 0xf0, 0x5e, 0x78, 0x19, 0x8d, 0x0b,
	0x56, 0x8b, 0x8f, 0xea, 0x59, 0x2a, 0xfa, 0x8c, 0x8e, 0x80, 0x8a, 0x96, 0x36, 0xb5, 0x76, 0xea,
	0xdc, 0x26, 0x1d, 0xf6, 0xaa, 0xb6, 0x8c, 0x7f, 0x60, 0xb7, 0x09, 0xe6, 0xb9, 0x7c, 0xa4, 0x4a,
	0x6e, 0xd3, 0x6d, 0x28, 0x70, 0xe1, 0xe1, 0x67, 0x14, 0xeb, 0x3c, 0xa7, 0xf8, 0xbb, 0x41, 0x2e,
	0x82, 0x77, 0xf0, 0x44, 0x02, 0x9b, 0xbb, 0xcd, 0xd8, 0x6d, 0x47, 0x4d, 0x87, 0x60, 0xec, 0x0d,
	0x21, 0x78, 0x03, 0x6e, 0x3f, 0xc2, 0xde, 0x74, 0x01, 0x36, 0x76, 0x90, 0x76, 0xa8, 0x02, 0x06,
	0xcd, 0x40, 0x05, 0x6f, 0xc1, 0x1b, 0x9d, 0xbc, 0x92, 0x95, 0xe0, 0xb1, 0xd6, 0x0b, 0x70, 0xd7,
	0x69, 0xfa, 0xf5, 0xa7, 0xfc, 0x88, 0xfa, 0x4b, 0x09, 0x58, 0x29, 0x13, 0x4c, 0xbb, 0x1e, 0x50,
	0xbd, 0x0e, 0x9e, 0x81, 0x37, 0xca, 0xba, 0x1b, 0xa4, 0x2e, 0xe3, 0x5f, 0x7e, 0x69, 0x9d, 0x4d,
	0xf5, 0x7a, 0xf5, 0xc7, 0x1c, 0x43, 0x5c, 0x63, 0xdd, 0x66, 0x09, 0x92, 0xf7, 0xe0, 0x4d, 0x4b,
	0x21, 0x7e, 0x74, 0x4f, 0x4f, 0xe7, 0x8f, 0xa2, 0xbd, 0x28, 0x2f, 0x61, 0xfe, 0xbf, 0x66, 0xb2,
	0x88, 0x76, 0xfa, 0x3e, 0x20, 0xbe, 0x34, 0xc8, 0x0b, 0xb0, 0xae, 0x59, 0x8b, 0xc4, 0x8b, 0x26,
	0x75, 0x1e, 0x3a, 0x7b, 0x05, 0xa7, 0x4a, 0xac, 0x0e, 0x3f, 0x46, 0x1f, 0x1a, 0xf2, 0x82, 0x57,
	0x60, 0xf7, 0x65, 0x48, 0xd3, 0xa4, 0x3e, 0x69, 0x9a, 0x36, 0x15, 0x1a, 0x1b, 0xeb, 0xbb, 0x59,
	0xc5, 0xf1, 0x5c, 0xff, 0xac, 0xaf, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x93, 0xce, 0xba,
	0xbe, 0x03, 0x00, 0x00,
}
