// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: commands.proto

package rpc

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

// StripeCLIClient is the client API for StripeCLI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StripeCLIClient interface {
	// Resend an event given an event ID. Like `stripe events resend`.
	EventsResend(ctx context.Context, in *EventsResendRequest, opts ...grpc.CallOption) (*EventsResendResponse, error)
	// Retrieve the default fixture of given triggering event.
	Fixture(ctx context.Context, in *FixtureRequest, opts ...grpc.CallOption) (*FixtureResponse, error)
	// Receive webhook events from the Stripe API to your local machine. Like `stripe listen`.
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (StripeCLI_ListenClient, error)
	// Get a link to log in to the Stripe CLI. The client will have to open the browser to complete
	// the login. Use `LoginStatus` after this method to wait for success. Like `stripe login`.
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// Successfully returns when login has succeeded, or returns an error if login has failed or timed
	// out. Use this method after `Login` to check for success.
	LoginStatus(ctx context.Context, in *LoginStatusRequest, opts ...grpc.CallOption) (*LoginStatusResponse, error)
	// Get a realtime stream of API logs. Like `stripe logs tail`.
	LogsTail(ctx context.Context, in *LogsTailRequest, opts ...grpc.CallOption) (StripeCLI_LogsTailClient, error)
	// Get a list of available configs for a given Stripe sample.
	SampleConfigs(ctx context.Context, in *SampleConfigsRequest, opts ...grpc.CallOption) (*SampleConfigsResponse, error)
	// Clone a Stripe sample. Like `stripe samples create`.
	SampleCreate(ctx context.Context, in *SampleCreateRequest, opts ...grpc.CallOption) (*SampleCreateResponse, error)
	// Get a list of available Stripe samples. Like `stripe samples list`.
	SamplesList(ctx context.Context, in *SamplesListRequest, opts ...grpc.CallOption) (*SamplesListResponse, error)
	// Trigger a webhook event. Like `stripe trigger`.
	Trigger(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*TriggerResponse, error)
	// Get a list of supported events for `Trigger`.
	TriggersList(ctx context.Context, in *TriggersListRequest, opts ...grpc.CallOption) (*TriggersListResponse, error)
	// Get the version of the Stripe CLI. Like `stripe version`.
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
	// Create a new webhook endpoint
	WebhookEndpointCreate(ctx context.Context, in *WebhookEndpointCreateRequest, opts ...grpc.CallOption) (*WebhookEndpointCreateResponse, error)
	// Get the list of webhook endpoints.
	WebhookEndpointsList(ctx context.Context, in *WebhookEndpointsListRequest, opts ...grpc.CallOption) (*WebhookEndpointsListResponse, error)
}

type stripeCLIClient struct {
	cc grpc.ClientConnInterface
}

func NewStripeCLIClient(cc grpc.ClientConnInterface) StripeCLIClient {
	return &stripeCLIClient{cc}
}

func (c *stripeCLIClient) EventsResend(ctx context.Context, in *EventsResendRequest, opts ...grpc.CallOption) (*EventsResendResponse, error) {
	out := new(EventsResendResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/EventsResend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeCLIClient) Fixture(ctx context.Context, in *FixtureRequest, opts ...grpc.CallOption) (*FixtureResponse, error) {
	out := new(FixtureResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/Fixture", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeCLIClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (StripeCLI_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &StripeCLI_ServiceDesc.Streams[0], "/rpc.StripeCLI/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &stripeCLIListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StripeCLI_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type stripeCLIListenClient struct {
	grpc.ClientStream
}

func (x *stripeCLIListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *stripeCLIClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeCLIClient) LoginStatus(ctx context.Context, in *LoginStatusRequest, opts ...grpc.CallOption) (*LoginStatusResponse, error) {
	out := new(LoginStatusResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/LoginStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeCLIClient) LogsTail(ctx context.Context, in *LogsTailRequest, opts ...grpc.CallOption) (StripeCLI_LogsTailClient, error) {
	stream, err := c.cc.NewStream(ctx, &StripeCLI_ServiceDesc.Streams[1], "/rpc.StripeCLI/LogsTail", opts...)
	if err != nil {
		return nil, err
	}
	x := &stripeCLILogsTailClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StripeCLI_LogsTailClient interface {
	Recv() (*LogsTailResponse, error)
	grpc.ClientStream
}

type stripeCLILogsTailClient struct {
	grpc.ClientStream
}

func (x *stripeCLILogsTailClient) Recv() (*LogsTailResponse, error) {
	m := new(LogsTailResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *stripeCLIClient) SampleConfigs(ctx context.Context, in *SampleConfigsRequest, opts ...grpc.CallOption) (*SampleConfigsResponse, error) {
	out := new(SampleConfigsResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/SampleConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeCLIClient) SampleCreate(ctx context.Context, in *SampleCreateRequest, opts ...grpc.CallOption) (*SampleCreateResponse, error) {
	out := new(SampleCreateResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/SampleCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeCLIClient) SamplesList(ctx context.Context, in *SamplesListRequest, opts ...grpc.CallOption) (*SamplesListResponse, error) {
	out := new(SamplesListResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/SamplesList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeCLIClient) Trigger(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*TriggerResponse, error) {
	out := new(TriggerResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/Trigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeCLIClient) TriggersList(ctx context.Context, in *TriggersListRequest, opts ...grpc.CallOption) (*TriggersListResponse, error) {
	out := new(TriggersListResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/TriggersList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeCLIClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeCLIClient) WebhookEndpointCreate(ctx context.Context, in *WebhookEndpointCreateRequest, opts ...grpc.CallOption) (*WebhookEndpointCreateResponse, error) {
	out := new(WebhookEndpointCreateResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/WebhookEndpointCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stripeCLIClient) WebhookEndpointsList(ctx context.Context, in *WebhookEndpointsListRequest, opts ...grpc.CallOption) (*WebhookEndpointsListResponse, error) {
	out := new(WebhookEndpointsListResponse)
	err := c.cc.Invoke(ctx, "/rpc.StripeCLI/WebhookEndpointsList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StripeCLIServer is the server API for StripeCLI service.
// All implementations should embed UnimplementedStripeCLIServer
// for forward compatibility
type StripeCLIServer interface {
	// Resend an event given an event ID. Like `stripe events resend`.
	EventsResend(context.Context, *EventsResendRequest) (*EventsResendResponse, error)
	// Retrieve the default fixture of given triggering event.
	Fixture(context.Context, *FixtureRequest) (*FixtureResponse, error)
	// Receive webhook events from the Stripe API to your local machine. Like `stripe listen`.
	Listen(*ListenRequest, StripeCLI_ListenServer) error
	// Get a link to log in to the Stripe CLI. The client will have to open the browser to complete
	// the login. Use `LoginStatus` after this method to wait for success. Like `stripe login`.
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// Successfully returns when login has succeeded, or returns an error if login has failed or timed
	// out. Use this method after `Login` to check for success.
	LoginStatus(context.Context, *LoginStatusRequest) (*LoginStatusResponse, error)
	// Get a realtime stream of API logs. Like `stripe logs tail`.
	LogsTail(*LogsTailRequest, StripeCLI_LogsTailServer) error
	// Get a list of available configs for a given Stripe sample.
	SampleConfigs(context.Context, *SampleConfigsRequest) (*SampleConfigsResponse, error)
	// Clone a Stripe sample. Like `stripe samples create`.
	SampleCreate(context.Context, *SampleCreateRequest) (*SampleCreateResponse, error)
	// Get a list of available Stripe samples. Like `stripe samples list`.
	SamplesList(context.Context, *SamplesListRequest) (*SamplesListResponse, error)
	// Trigger a webhook event. Like `stripe trigger`.
	Trigger(context.Context, *TriggerRequest) (*TriggerResponse, error)
	// Get a list of supported events for `Trigger`.
	TriggersList(context.Context, *TriggersListRequest) (*TriggersListResponse, error)
	// Get the version of the Stripe CLI. Like `stripe version`.
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
	// Create a new webhook endpoint
	WebhookEndpointCreate(context.Context, *WebhookEndpointCreateRequest) (*WebhookEndpointCreateResponse, error)
	// Get the list of webhook endpoints.
	WebhookEndpointsList(context.Context, *WebhookEndpointsListRequest) (*WebhookEndpointsListResponse, error)
}

// UnimplementedStripeCLIServer should be embedded to have forward compatible implementations.
type UnimplementedStripeCLIServer struct {
}

func (UnimplementedStripeCLIServer) EventsResend(context.Context, *EventsResendRequest) (*EventsResendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EventsResend not implemented")
}
func (UnimplementedStripeCLIServer) Fixture(context.Context, *FixtureRequest) (*FixtureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fixture not implemented")
}
func (UnimplementedStripeCLIServer) Listen(*ListenRequest, StripeCLI_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedStripeCLIServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedStripeCLIServer) LoginStatus(context.Context, *LoginStatusRequest) (*LoginStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginStatus not implemented")
}
func (UnimplementedStripeCLIServer) LogsTail(*LogsTailRequest, StripeCLI_LogsTailServer) error {
	return status.Errorf(codes.Unimplemented, "method LogsTail not implemented")
}
func (UnimplementedStripeCLIServer) SampleConfigs(context.Context, *SampleConfigsRequest) (*SampleConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SampleConfigs not implemented")
}
func (UnimplementedStripeCLIServer) SampleCreate(context.Context, *SampleCreateRequest) (*SampleCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SampleCreate not implemented")
}
func (UnimplementedStripeCLIServer) SamplesList(context.Context, *SamplesListRequest) (*SamplesListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SamplesList not implemented")
}
func (UnimplementedStripeCLIServer) Trigger(context.Context, *TriggerRequest) (*TriggerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Trigger not implemented")
}
func (UnimplementedStripeCLIServer) TriggersList(context.Context, *TriggersListRequest) (*TriggersListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggersList not implemented")
}
func (UnimplementedStripeCLIServer) Version(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedStripeCLIServer) WebhookEndpointCreate(context.Context, *WebhookEndpointCreateRequest) (*WebhookEndpointCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WebhookEndpointCreate not implemented")
}
func (UnimplementedStripeCLIServer) WebhookEndpointsList(context.Context, *WebhookEndpointsListRequest) (*WebhookEndpointsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WebhookEndpointsList not implemented")
}

// UnsafeStripeCLIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StripeCLIServer will
// result in compilation errors.
type UnsafeStripeCLIServer interface {
	mustEmbedUnimplementedStripeCLIServer()
}

func RegisterStripeCLIServer(s grpc.ServiceRegistrar, srv StripeCLIServer) {
	s.RegisterService(&StripeCLI_ServiceDesc, srv)
}

func _StripeCLI_EventsResend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventsResendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).EventsResend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/EventsResend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).EventsResend(ctx, req.(*EventsResendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeCLI_Fixture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FixtureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).Fixture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/Fixture",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).Fixture(ctx, req.(*FixtureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeCLI_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StripeCLIServer).Listen(m, &stripeCLIListenServer{stream})
}

type StripeCLI_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type stripeCLIListenServer struct {
	grpc.ServerStream
}

func (x *stripeCLIListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StripeCLI_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeCLI_LoginStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).LoginStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/LoginStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).LoginStatus(ctx, req.(*LoginStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeCLI_LogsTail_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogsTailRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StripeCLIServer).LogsTail(m, &stripeCLILogsTailServer{stream})
}

type StripeCLI_LogsTailServer interface {
	Send(*LogsTailResponse) error
	grpc.ServerStream
}

type stripeCLILogsTailServer struct {
	grpc.ServerStream
}

func (x *stripeCLILogsTailServer) Send(m *LogsTailResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StripeCLI_SampleConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SampleConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).SampleConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/SampleConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).SampleConfigs(ctx, req.(*SampleConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeCLI_SampleCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SampleCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).SampleCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/SampleCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).SampleCreate(ctx, req.(*SampleCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeCLI_SamplesList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SamplesListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).SamplesList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/SamplesList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).SamplesList(ctx, req.(*SamplesListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeCLI_Trigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).Trigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/Trigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).Trigger(ctx, req.(*TriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeCLI_TriggersList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggersListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).TriggersList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/TriggersList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).TriggersList(ctx, req.(*TriggersListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeCLI_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeCLI_WebhookEndpointCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookEndpointCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).WebhookEndpointCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/WebhookEndpointCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).WebhookEndpointCreate(ctx, req.(*WebhookEndpointCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StripeCLI_WebhookEndpointsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookEndpointsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StripeCLIServer).WebhookEndpointsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.StripeCLI/WebhookEndpointsList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StripeCLIServer).WebhookEndpointsList(ctx, req.(*WebhookEndpointsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StripeCLI_ServiceDesc is the grpc.ServiceDesc for StripeCLI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StripeCLI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.StripeCLI",
	HandlerType: (*StripeCLIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EventsResend",
			Handler:    _StripeCLI_EventsResend_Handler,
		},
		{
			MethodName: "Fixture",
			Handler:    _StripeCLI_Fixture_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _StripeCLI_Login_Handler,
		},
		{
			MethodName: "LoginStatus",
			Handler:    _StripeCLI_LoginStatus_Handler,
		},
		{
			MethodName: "SampleConfigs",
			Handler:    _StripeCLI_SampleConfigs_Handler,
		},
		{
			MethodName: "SampleCreate",
			Handler:    _StripeCLI_SampleCreate_Handler,
		},
		{
			MethodName: "SamplesList",
			Handler:    _StripeCLI_SamplesList_Handler,
		},
		{
			MethodName: "Trigger",
			Handler:    _StripeCLI_Trigger_Handler,
		},
		{
			MethodName: "TriggersList",
			Handler:    _StripeCLI_TriggersList_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _StripeCLI_Version_Handler,
		},
		{
			MethodName: "WebhookEndpointCreate",
			Handler:    _StripeCLI_WebhookEndpointCreate_Handler,
		},
		{
			MethodName: "WebhookEndpointsList",
			Handler:    _StripeCLI_WebhookEndpointsList_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _StripeCLI_Listen_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LogsTail",
			Handler:       _StripeCLI_LogsTail_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "commands.proto",
}
