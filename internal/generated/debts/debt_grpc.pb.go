// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.25.1
// source: debts/debt.proto

package debts

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	DebtsService_AddClient_FullMethodName            = "/debts.DebtsService/AddClient"
	DebtsService_GetClient_FullMethodName            = "/debts.DebtsService/GetClient"
	DebtsService_UpdateClient_FullMethodName         = "/debts.DebtsService/UpdateClient"
	DebtsService_GetAllClients_FullMethodName        = "/debts.DebtsService/GetAllClients"
	DebtsService_CreateDebts_FullMethodName          = "/debts.DebtsService/CreateDebts"
	DebtsService_GetDebts_FullMethodName             = "/debts.DebtsService/GetDebts"
	DebtsService_PayDebts_FullMethodName             = "/debts.DebtsService/PayDebts"
	DebtsService_GetListDebts_FullMethodName         = "/debts.DebtsService/GetListDebts"
	DebtsService_GetClientDebts_FullMethodName       = "/debts.DebtsService/GetClientDebts"
	DebtsService_GetUserTotalDebtSum_FullMethodName  = "/debts.DebtsService/GetUserTotalDebtSum"
	DebtsService_GetTotalDebtSum_FullMethodName      = "/debts.DebtsService/GetTotalDebtSum"
	DebtsService_GetPayment_FullMethodName           = "/debts.DebtsService/GetPayment"
	DebtsService_GetPaymentsByDebtsId_FullMethodName = "/debts.DebtsService/GetPaymentsByDebtsId"
	DebtsService_GetPayments_FullMethodName          = "/debts.DebtsService/GetPayments"
	DebtsService_GetUserPayments_FullMethodName      = "/debts.DebtsService/GetUserPayments"
	DebtsService_GetDebtsForExel_FullMethodName      = "/debts.DebtsService/GetDebtsForExel"
)

// DebtsServiceClient is the client API for DebtsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// ----------------------- Service --------------------
type DebtsServiceClient interface {
	AddClient(ctx context.Context, in *CreateClients, opts ...grpc.CallOption) (*Client, error)
	GetClient(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*Client, error)
	UpdateClient(ctx context.Context, in *ClientUpdate, opts ...grpc.CallOption) (*Client, error)
	GetAllClients(ctx context.Context, in *FilterClient, opts ...grpc.CallOption) (*ClientList, error)
	CreateDebts(ctx context.Context, in *DebtsRequest, opts ...grpc.CallOption) (*Debts, error)
	GetDebts(ctx context.Context, in *DebtsID, opts ...grpc.CallOption) (*Debts, error)
	PayDebts(ctx context.Context, in *PayDebtsReq, opts ...grpc.CallOption) (*Debts, error)
	GetListDebts(ctx context.Context, in *FilterDebts, opts ...grpc.CallOption) (*DebtsList, error)
	GetClientDebts(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*DebtsList, error)
	GetUserTotalDebtSum(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*SumMoney, error)
	GetTotalDebtSum(ctx context.Context, in *CompanyID, opts ...grpc.CallOption) (*SumMoney, error)
	GetPayment(ctx context.Context, in *PaymentID, opts ...grpc.CallOption) (*Payment, error)
	GetPaymentsByDebtsId(ctx context.Context, in *PayDebtsID, opts ...grpc.CallOption) (*PaymentList, error)
	GetPayments(ctx context.Context, in *FilterPayment, opts ...grpc.CallOption) (*PaymentList, error)
	GetUserPayments(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*UserPaymentsRes, error)
	GetDebtsForExel(ctx context.Context, in *FilterExelDebt, opts ...grpc.CallOption) (*ExelDebtsList, error)
}

type debtsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDebtsServiceClient(cc grpc.ClientConnInterface) DebtsServiceClient {
	return &debtsServiceClient{cc}
}

func (c *debtsServiceClient) AddClient(ctx context.Context, in *CreateClients, opts ...grpc.CallOption) (*Client, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Client)
	err := c.cc.Invoke(ctx, DebtsService_AddClient_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debtsServiceClient) GetClient(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*Client, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Client)
	err := c.cc.Invoke(ctx, DebtsService_GetClient_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debtsServiceClient) UpdateClient(ctx context.Context, in *ClientUpdate, opts ...grpc.CallOption) (*Client, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Client)
	err := c.cc.Invoke(ctx, DebtsService_UpdateClient_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debtsServiceClient) GetAllClients(ctx context.Context, in *FilterClient, opts ...grpc.CallOption) (*ClientList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClientList)
	err := c.cc.Invoke(ctx, DebtsService_GetAllClients_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debtsServiceClient) CreateDebts(ctx context.Context, in *DebtsRequest, opts ...grpc.CallOption) (*Debts, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Debts)
	err := c.cc.Invoke(ctx, DebtsService_CreateDebts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debtsServiceClient) GetDebts(ctx context.Context, in *DebtsID, opts ...grpc.CallOption) (*Debts, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Debts)
	err := c.cc.Invoke(ctx, DebtsService_GetDebts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debtsServiceClient) PayDebts(ctx context.Context, in *PayDebtsReq, opts ...grpc.CallOption) (*Debts, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Debts)
	err := c.cc.Invoke(ctx, DebtsService_PayDebts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debtsServiceClient) GetListDebts(ctx context.Context, in *FilterDebts, opts ...grpc.CallOption) (*DebtsList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DebtsList)
	err := c.cc.Invoke(ctx, DebtsService_GetListDebts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debtsServiceClient) GetClientDebts(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*DebtsList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DebtsList)
	err := c.cc.Invoke(ctx, DebtsService_GetClientDebts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debtsServiceClient) GetUserTotalDebtSum(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*SumMoney, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SumMoney)
	err := c.cc.Invoke(ctx, DebtsService_GetUserTotalDebtSum_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debtsServiceClient) GetTotalDebtSum(ctx context.Context, in *CompanyID, opts ...grpc.CallOption) (*SumMoney, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SumMoney)
	err := c.cc.Invoke(ctx, DebtsService_GetTotalDebtSum_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debtsServiceClient) GetPayment(ctx context.Context, in *PaymentID, opts ...grpc.CallOption) (*Payment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Payment)
	err := c.cc.Invoke(ctx, DebtsService_GetPayment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debtsServiceClient) GetPaymentsByDebtsId(ctx context.Context, in *PayDebtsID, opts ...grpc.CallOption) (*PaymentList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PaymentList)
	err := c.cc.Invoke(ctx, DebtsService_GetPaymentsByDebtsId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debtsServiceClient) GetPayments(ctx context.Context, in *FilterPayment, opts ...grpc.CallOption) (*PaymentList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PaymentList)
	err := c.cc.Invoke(ctx, DebtsService_GetPayments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debtsServiceClient) GetUserPayments(ctx context.Context, in *ClientID, opts ...grpc.CallOption) (*UserPaymentsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserPaymentsRes)
	err := c.cc.Invoke(ctx, DebtsService_GetUserPayments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debtsServiceClient) GetDebtsForExel(ctx context.Context, in *FilterExelDebt, opts ...grpc.CallOption) (*ExelDebtsList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExelDebtsList)
	err := c.cc.Invoke(ctx, DebtsService_GetDebtsForExel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DebtsServiceServer is the server API for DebtsService service.
// All implementations must embed UnimplementedDebtsServiceServer
// for forward compatibility
//
// ----------------------- Service --------------------
type DebtsServiceServer interface {
	AddClient(context.Context, *CreateClients) (*Client, error)
	GetClient(context.Context, *ClientID) (*Client, error)
	UpdateClient(context.Context, *ClientUpdate) (*Client, error)
	GetAllClients(context.Context, *FilterClient) (*ClientList, error)
	CreateDebts(context.Context, *DebtsRequest) (*Debts, error)
	GetDebts(context.Context, *DebtsID) (*Debts, error)
	PayDebts(context.Context, *PayDebtsReq) (*Debts, error)
	GetListDebts(context.Context, *FilterDebts) (*DebtsList, error)
	GetClientDebts(context.Context, *ClientID) (*DebtsList, error)
	GetUserTotalDebtSum(context.Context, *ClientID) (*SumMoney, error)
	GetTotalDebtSum(context.Context, *CompanyID) (*SumMoney, error)
	GetPayment(context.Context, *PaymentID) (*Payment, error)
	GetPaymentsByDebtsId(context.Context, *PayDebtsID) (*PaymentList, error)
	GetPayments(context.Context, *FilterPayment) (*PaymentList, error)
	GetUserPayments(context.Context, *ClientID) (*UserPaymentsRes, error)
	GetDebtsForExel(context.Context, *FilterExelDebt) (*ExelDebtsList, error)
	mustEmbedUnimplementedDebtsServiceServer()
}

// UnimplementedDebtsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDebtsServiceServer struct {
}

func (UnimplementedDebtsServiceServer) AddClient(context.Context, *CreateClients) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddClient not implemented")
}
func (UnimplementedDebtsServiceServer) GetClient(context.Context, *ClientID) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClient not implemented")
}
func (UnimplementedDebtsServiceServer) UpdateClient(context.Context, *ClientUpdate) (*Client, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClient not implemented")
}
func (UnimplementedDebtsServiceServer) GetAllClients(context.Context, *FilterClient) (*ClientList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllClients not implemented")
}
func (UnimplementedDebtsServiceServer) CreateDebts(context.Context, *DebtsRequest) (*Debts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDebts not implemented")
}
func (UnimplementedDebtsServiceServer) GetDebts(context.Context, *DebtsID) (*Debts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDebts not implemented")
}
func (UnimplementedDebtsServiceServer) PayDebts(context.Context, *PayDebtsReq) (*Debts, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayDebts not implemented")
}
func (UnimplementedDebtsServiceServer) GetListDebts(context.Context, *FilterDebts) (*DebtsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListDebts not implemented")
}
func (UnimplementedDebtsServiceServer) GetClientDebts(context.Context, *ClientID) (*DebtsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientDebts not implemented")
}
func (UnimplementedDebtsServiceServer) GetUserTotalDebtSum(context.Context, *ClientID) (*SumMoney, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserTotalDebtSum not implemented")
}
func (UnimplementedDebtsServiceServer) GetTotalDebtSum(context.Context, *CompanyID) (*SumMoney, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTotalDebtSum not implemented")
}
func (UnimplementedDebtsServiceServer) GetPayment(context.Context, *PaymentID) (*Payment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayment not implemented")
}
func (UnimplementedDebtsServiceServer) GetPaymentsByDebtsId(context.Context, *PayDebtsID) (*PaymentList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentsByDebtsId not implemented")
}
func (UnimplementedDebtsServiceServer) GetPayments(context.Context, *FilterPayment) (*PaymentList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPayments not implemented")
}
func (UnimplementedDebtsServiceServer) GetUserPayments(context.Context, *ClientID) (*UserPaymentsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPayments not implemented")
}
func (UnimplementedDebtsServiceServer) GetDebtsForExel(context.Context, *FilterExelDebt) (*ExelDebtsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDebtsForExel not implemented")
}
func (UnimplementedDebtsServiceServer) mustEmbedUnimplementedDebtsServiceServer() {}

// UnsafeDebtsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DebtsServiceServer will
// result in compilation errors.
type UnsafeDebtsServiceServer interface {
	mustEmbedUnimplementedDebtsServiceServer()
}

func RegisterDebtsServiceServer(s grpc.ServiceRegistrar, srv DebtsServiceServer) {
	s.RegisterService(&DebtsService_ServiceDesc, srv)
}

func _DebtsService_AddClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClients)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebtsServiceServer).AddClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebtsService_AddClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebtsServiceServer).AddClient(ctx, req.(*CreateClients))
	}
	return interceptor(ctx, in, info, handler)
}

func _DebtsService_GetClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebtsServiceServer).GetClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebtsService_GetClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebtsServiceServer).GetClient(ctx, req.(*ClientID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DebtsService_UpdateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebtsServiceServer).UpdateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebtsService_UpdateClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebtsServiceServer).UpdateClient(ctx, req.(*ClientUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _DebtsService_GetAllClients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterClient)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebtsServiceServer).GetAllClients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebtsService_GetAllClients_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebtsServiceServer).GetAllClients(ctx, req.(*FilterClient))
	}
	return interceptor(ctx, in, info, handler)
}

func _DebtsService_CreateDebts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DebtsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebtsServiceServer).CreateDebts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebtsService_CreateDebts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebtsServiceServer).CreateDebts(ctx, req.(*DebtsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DebtsService_GetDebts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DebtsID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebtsServiceServer).GetDebts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebtsService_GetDebts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebtsServiceServer).GetDebts(ctx, req.(*DebtsID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DebtsService_PayDebts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayDebtsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebtsServiceServer).PayDebts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebtsService_PayDebts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebtsServiceServer).PayDebts(ctx, req.(*PayDebtsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DebtsService_GetListDebts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterDebts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebtsServiceServer).GetListDebts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebtsService_GetListDebts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebtsServiceServer).GetListDebts(ctx, req.(*FilterDebts))
	}
	return interceptor(ctx, in, info, handler)
}

func _DebtsService_GetClientDebts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebtsServiceServer).GetClientDebts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebtsService_GetClientDebts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebtsServiceServer).GetClientDebts(ctx, req.(*ClientID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DebtsService_GetUserTotalDebtSum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebtsServiceServer).GetUserTotalDebtSum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebtsService_GetUserTotalDebtSum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebtsServiceServer).GetUserTotalDebtSum(ctx, req.(*ClientID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DebtsService_GetTotalDebtSum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebtsServiceServer).GetTotalDebtSum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebtsService_GetTotalDebtSum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebtsServiceServer).GetTotalDebtSum(ctx, req.(*CompanyID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DebtsService_GetPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebtsServiceServer).GetPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebtsService_GetPayment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebtsServiceServer).GetPayment(ctx, req.(*PaymentID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DebtsService_GetPaymentsByDebtsId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayDebtsID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebtsServiceServer).GetPaymentsByDebtsId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebtsService_GetPaymentsByDebtsId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebtsServiceServer).GetPaymentsByDebtsId(ctx, req.(*PayDebtsID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DebtsService_GetPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterPayment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebtsServiceServer).GetPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebtsService_GetPayments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebtsServiceServer).GetPayments(ctx, req.(*FilterPayment))
	}
	return interceptor(ctx, in, info, handler)
}

func _DebtsService_GetUserPayments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebtsServiceServer).GetUserPayments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebtsService_GetUserPayments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebtsServiceServer).GetUserPayments(ctx, req.(*ClientID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DebtsService_GetDebtsForExel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterExelDebt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DebtsServiceServer).GetDebtsForExel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DebtsService_GetDebtsForExel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DebtsServiceServer).GetDebtsForExel(ctx, req.(*FilterExelDebt))
	}
	return interceptor(ctx, in, info, handler)
}

// DebtsService_ServiceDesc is the grpc.ServiceDesc for DebtsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DebtsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "debts.DebtsService",
	HandlerType: (*DebtsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddClient",
			Handler:    _DebtsService_AddClient_Handler,
		},
		{
			MethodName: "GetClient",
			Handler:    _DebtsService_GetClient_Handler,
		},
		{
			MethodName: "UpdateClient",
			Handler:    _DebtsService_UpdateClient_Handler,
		},
		{
			MethodName: "GetAllClients",
			Handler:    _DebtsService_GetAllClients_Handler,
		},
		{
			MethodName: "CreateDebts",
			Handler:    _DebtsService_CreateDebts_Handler,
		},
		{
			MethodName: "GetDebts",
			Handler:    _DebtsService_GetDebts_Handler,
		},
		{
			MethodName: "PayDebts",
			Handler:    _DebtsService_PayDebts_Handler,
		},
		{
			MethodName: "GetListDebts",
			Handler:    _DebtsService_GetListDebts_Handler,
		},
		{
			MethodName: "GetClientDebts",
			Handler:    _DebtsService_GetClientDebts_Handler,
		},
		{
			MethodName: "GetUserTotalDebtSum",
			Handler:    _DebtsService_GetUserTotalDebtSum_Handler,
		},
		{
			MethodName: "GetTotalDebtSum",
			Handler:    _DebtsService_GetTotalDebtSum_Handler,
		},
		{
			MethodName: "GetPayment",
			Handler:    _DebtsService_GetPayment_Handler,
		},
		{
			MethodName: "GetPaymentsByDebtsId",
			Handler:    _DebtsService_GetPaymentsByDebtsId_Handler,
		},
		{
			MethodName: "GetPayments",
			Handler:    _DebtsService_GetPayments_Handler,
		},
		{
			MethodName: "GetUserPayments",
			Handler:    _DebtsService_GetUserPayments_Handler,
		},
		{
			MethodName: "GetDebtsForExel",
			Handler:    _DebtsService_GetDebtsForExel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "debts/debt.proto",
}
