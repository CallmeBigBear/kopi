// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: kopi/mm/tx.proto

package mm

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Msg_AddDeposit_FullMethodName                   = "/kopi.mm.Msg/AddDeposit"
	Msg_CreateRedemptionRequest_FullMethodName      = "/kopi.mm.Msg/CreateRedemptionRequest"
	Msg_CancelRedemptionRequest_FullMethodName      = "/kopi.mm.Msg/CancelRedemptionRequest"
	Msg_UpdateRedemptionRequest_FullMethodName      = "/kopi.mm.Msg/UpdateRedemptionRequest"
	Msg_AddCollateral_FullMethodName                = "/kopi.mm.Msg/AddCollateral"
	Msg_RemoveCollateral_FullMethodName             = "/kopi.mm.Msg/RemoveCollateral"
	Msg_Borrow_FullMethodName                       = "/kopi.mm.Msg/Borrow"
	Msg_PartiallyRepayLoan_FullMethodName           = "/kopi.mm.Msg/PartiallyRepayLoan"
	Msg_RepayLoan_FullMethodName                    = "/kopi.mm.Msg/RepayLoan"
	Msg_UpdateCollateralDiscount_FullMethodName     = "/kopi.mm.Msg/UpdateCollateralDiscount"
	Msg_UpdateInterestRateParameters_FullMethodName = "/kopi.mm.Msg/UpdateInterestRateParameters"
	Msg_UpdateRedemptionFee_FullMethodName          = "/kopi.mm.Msg/UpdateRedemptionFee"
	Msg_UpdateProtocolShare_FullMethodName          = "/kopi.mm.Msg/UpdateProtocolShare"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	AddDeposit(ctx context.Context, in *MsgAddDeposit, opts ...grpc.CallOption) (*Void, error)
	CreateRedemptionRequest(ctx context.Context, in *MsgCreateRedemptionRequest, opts ...grpc.CallOption) (*Void, error)
	CancelRedemptionRequest(ctx context.Context, in *MsgCancelRedemptionRequest, opts ...grpc.CallOption) (*Void, error)
	UpdateRedemptionRequest(ctx context.Context, in *MsgUpdateRedemptionRequest, opts ...grpc.CallOption) (*Void, error)
	AddCollateral(ctx context.Context, in *MsgAddCollateral, opts ...grpc.CallOption) (*Void, error)
	RemoveCollateral(ctx context.Context, in *MsgRemoveCollateral, opts ...grpc.CallOption) (*Void, error)
	Borrow(ctx context.Context, in *MsgBorrow, opts ...grpc.CallOption) (*Void, error)
	PartiallyRepayLoan(ctx context.Context, in *MsgPartiallyRepayLoan, opts ...grpc.CallOption) (*Void, error)
	RepayLoan(ctx context.Context, in *MsgRepayLoan, opts ...grpc.CallOption) (*Void, error)
	UpdateCollateralDiscount(ctx context.Context, in *MsgUpdateCollateralDiscount, opts ...grpc.CallOption) (*Void, error)
	UpdateInterestRateParameters(ctx context.Context, in *MsgUpdateInterestRateParameters, opts ...grpc.CallOption) (*Void, error)
	UpdateRedemptionFee(ctx context.Context, in *MsgUpdateRedemptionFee, opts ...grpc.CallOption) (*Void, error)
	UpdateProtocolShare(ctx context.Context, in *MsgUpdateProtocolShare, opts ...grpc.CallOption) (*Void, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AddDeposit(ctx context.Context, in *MsgAddDeposit, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, Msg_AddDeposit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateRedemptionRequest(ctx context.Context, in *MsgCreateRedemptionRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, Msg_CreateRedemptionRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelRedemptionRequest(ctx context.Context, in *MsgCancelRedemptionRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, Msg_CancelRedemptionRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateRedemptionRequest(ctx context.Context, in *MsgUpdateRedemptionRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, Msg_UpdateRedemptionRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AddCollateral(ctx context.Context, in *MsgAddCollateral, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, Msg_AddCollateral_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveCollateral(ctx context.Context, in *MsgRemoveCollateral, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, Msg_RemoveCollateral_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Borrow(ctx context.Context, in *MsgBorrow, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, Msg_Borrow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) PartiallyRepayLoan(ctx context.Context, in *MsgPartiallyRepayLoan, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, Msg_PartiallyRepayLoan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RepayLoan(ctx context.Context, in *MsgRepayLoan, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, Msg_RepayLoan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateCollateralDiscount(ctx context.Context, in *MsgUpdateCollateralDiscount, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, Msg_UpdateCollateralDiscount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateInterestRateParameters(ctx context.Context, in *MsgUpdateInterestRateParameters, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, Msg_UpdateInterestRateParameters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateRedemptionFee(ctx context.Context, in *MsgUpdateRedemptionFee, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, Msg_UpdateRedemptionFee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateProtocolShare(ctx context.Context, in *MsgUpdateProtocolShare, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, Msg_UpdateProtocolShare_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	AddDeposit(context.Context, *MsgAddDeposit) (*Void, error)
	CreateRedemptionRequest(context.Context, *MsgCreateRedemptionRequest) (*Void, error)
	CancelRedemptionRequest(context.Context, *MsgCancelRedemptionRequest) (*Void, error)
	UpdateRedemptionRequest(context.Context, *MsgUpdateRedemptionRequest) (*Void, error)
	AddCollateral(context.Context, *MsgAddCollateral) (*Void, error)
	RemoveCollateral(context.Context, *MsgRemoveCollateral) (*Void, error)
	Borrow(context.Context, *MsgBorrow) (*Void, error)
	PartiallyRepayLoan(context.Context, *MsgPartiallyRepayLoan) (*Void, error)
	RepayLoan(context.Context, *MsgRepayLoan) (*Void, error)
	UpdateCollateralDiscount(context.Context, *MsgUpdateCollateralDiscount) (*Void, error)
	UpdateInterestRateParameters(context.Context, *MsgUpdateInterestRateParameters) (*Void, error)
	UpdateRedemptionFee(context.Context, *MsgUpdateRedemptionFee) (*Void, error)
	UpdateProtocolShare(context.Context, *MsgUpdateProtocolShare) (*Void, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) AddDeposit(context.Context, *MsgAddDeposit) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDeposit not implemented")
}
func (UnimplementedMsgServer) CreateRedemptionRequest(context.Context, *MsgCreateRedemptionRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRedemptionRequest not implemented")
}
func (UnimplementedMsgServer) CancelRedemptionRequest(context.Context, *MsgCancelRedemptionRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelRedemptionRequest not implemented")
}
func (UnimplementedMsgServer) UpdateRedemptionRequest(context.Context, *MsgUpdateRedemptionRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRedemptionRequest not implemented")
}
func (UnimplementedMsgServer) AddCollateral(context.Context, *MsgAddCollateral) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCollateral not implemented")
}
func (UnimplementedMsgServer) RemoveCollateral(context.Context, *MsgRemoveCollateral) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCollateral not implemented")
}
func (UnimplementedMsgServer) Borrow(context.Context, *MsgBorrow) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Borrow not implemented")
}
func (UnimplementedMsgServer) PartiallyRepayLoan(context.Context, *MsgPartiallyRepayLoan) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PartiallyRepayLoan not implemented")
}
func (UnimplementedMsgServer) RepayLoan(context.Context, *MsgRepayLoan) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepayLoan not implemented")
}
func (UnimplementedMsgServer) UpdateCollateralDiscount(context.Context, *MsgUpdateCollateralDiscount) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCollateralDiscount not implemented")
}
func (UnimplementedMsgServer) UpdateInterestRateParameters(context.Context, *MsgUpdateInterestRateParameters) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInterestRateParameters not implemented")
}
func (UnimplementedMsgServer) UpdateRedemptionFee(context.Context, *MsgUpdateRedemptionFee) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRedemptionFee not implemented")
}
func (UnimplementedMsgServer) UpdateProtocolShare(context.Context, *MsgUpdateProtocolShare) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProtocolShare not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_AddDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddDeposit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AddDeposit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddDeposit(ctx, req.(*MsgAddDeposit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateRedemptionRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateRedemptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateRedemptionRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateRedemptionRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateRedemptionRequest(ctx, req.(*MsgCreateRedemptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelRedemptionRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelRedemptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelRedemptionRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CancelRedemptionRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelRedemptionRequest(ctx, req.(*MsgCancelRedemptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateRedemptionRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateRedemptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateRedemptionRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateRedemptionRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateRedemptionRequest(ctx, req.(*MsgUpdateRedemptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AddCollateral_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddCollateral)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddCollateral(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_AddCollateral_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddCollateral(ctx, req.(*MsgAddCollateral))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveCollateral_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveCollateral)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveCollateral(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RemoveCollateral_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveCollateral(ctx, req.(*MsgRemoveCollateral))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Borrow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBorrow)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Borrow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_Borrow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Borrow(ctx, req.(*MsgBorrow))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_PartiallyRepayLoan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPartiallyRepayLoan)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PartiallyRepayLoan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_PartiallyRepayLoan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PartiallyRepayLoan(ctx, req.(*MsgPartiallyRepayLoan))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RepayLoan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRepayLoan)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RepayLoan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RepayLoan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RepayLoan(ctx, req.(*MsgRepayLoan))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateCollateralDiscount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateCollateralDiscount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateCollateralDiscount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateCollateralDiscount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateCollateralDiscount(ctx, req.(*MsgUpdateCollateralDiscount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateInterestRateParameters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateInterestRateParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateInterestRateParameters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateInterestRateParameters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateInterestRateParameters(ctx, req.(*MsgUpdateInterestRateParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateRedemptionFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateRedemptionFee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateRedemptionFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateRedemptionFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateRedemptionFee(ctx, req.(*MsgUpdateRedemptionFee))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateProtocolShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateProtocolShare)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateProtocolShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateProtocolShare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateProtocolShare(ctx, req.(*MsgUpdateProtocolShare))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kopi.mm.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDeposit",
			Handler:    _Msg_AddDeposit_Handler,
		},
		{
			MethodName: "CreateRedemptionRequest",
			Handler:    _Msg_CreateRedemptionRequest_Handler,
		},
		{
			MethodName: "CancelRedemptionRequest",
			Handler:    _Msg_CancelRedemptionRequest_Handler,
		},
		{
			MethodName: "UpdateRedemptionRequest",
			Handler:    _Msg_UpdateRedemptionRequest_Handler,
		},
		{
			MethodName: "AddCollateral",
			Handler:    _Msg_AddCollateral_Handler,
		},
		{
			MethodName: "RemoveCollateral",
			Handler:    _Msg_RemoveCollateral_Handler,
		},
		{
			MethodName: "Borrow",
			Handler:    _Msg_Borrow_Handler,
		},
		{
			MethodName: "PartiallyRepayLoan",
			Handler:    _Msg_PartiallyRepayLoan_Handler,
		},
		{
			MethodName: "RepayLoan",
			Handler:    _Msg_RepayLoan_Handler,
		},
		{
			MethodName: "UpdateCollateralDiscount",
			Handler:    _Msg_UpdateCollateralDiscount_Handler,
		},
		{
			MethodName: "UpdateInterestRateParameters",
			Handler:    _Msg_UpdateInterestRateParameters_Handler,
		},
		{
			MethodName: "UpdateRedemptionFee",
			Handler:    _Msg_UpdateRedemptionFee_Handler,
		},
		{
			MethodName: "UpdateProtocolShare",
			Handler:    _Msg_UpdateProtocolShare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kopi/mm/tx.proto",
}
