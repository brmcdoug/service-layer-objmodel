// Code generated by protoc-gen-go.
// source: sl_global.proto
// DO NOT EDIT!

package service_layer

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

// Global Event Notification types.
type SLGlobalNotifType int32

const (
	// Reserved. 0x0
	SLGlobalNotifType_SL_GLOBAL_EVENT_TYPE_RESERVED SLGlobalNotifType = 0
	// Error. ErrStatus field elaborates on the message. 0x1
	SLGlobalNotifType_SL_GLOBAL_EVENT_TYPE_ERROR SLGlobalNotifType = 1
	// HeartBeat. 0x2
	SLGlobalNotifType_SL_GLOBAL_EVENT_TYPE_HEARTBEAT SLGlobalNotifType = 2
	// Version. SLInitMsgRsp field elaborates on the server version. 0x3
	SLGlobalNotifType_SL_GLOBAL_EVENT_TYPE_VERSION SLGlobalNotifType = 3
)

var SLGlobalNotifType_name = map[int32]string{
	0: "SL_GLOBAL_EVENT_TYPE_RESERVED",
	1: "SL_GLOBAL_EVENT_TYPE_ERROR",
	2: "SL_GLOBAL_EVENT_TYPE_HEARTBEAT",
	3: "SL_GLOBAL_EVENT_TYPE_VERSION",
}
var SLGlobalNotifType_value = map[string]int32{
	"SL_GLOBAL_EVENT_TYPE_RESERVED":  0,
	"SL_GLOBAL_EVENT_TYPE_ERROR":     1,
	"SL_GLOBAL_EVENT_TYPE_HEARTBEAT": 2,
	"SL_GLOBAL_EVENT_TYPE_VERSION":   3,
}

func (x SLGlobalNotifType) String() string {
	return proto.EnumName(SLGlobalNotifType_name, int32(x))
}
func (SLGlobalNotifType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

// Initialization message sent to the server.
// If the client and server are running compatible version numbers, a
// connection will be made and the server response will be received
// with a successful status code.
type SLInitMsg struct {
	// Client's Major version of service-layer API (refer to sl_version.proto)
	MajorVer uint32 `protobuf:"varint,1,opt,name=MajorVer,json=majorVer" json:"MajorVer,omitempty"`
	// Minor Version
	MinorVer uint32 `protobuf:"varint,2,opt,name=MinorVer,json=minorVer" json:"MinorVer,omitempty"`
	// Sub-Version
	SubVer uint32 `protobuf:"varint,3,opt,name=SubVer,json=subVer" json:"SubVer,omitempty"`
}

func (m *SLInitMsg) Reset()                    { *m = SLInitMsg{} }
func (m *SLInitMsg) String() string            { return proto.CompactTextString(m) }
func (*SLInitMsg) ProtoMessage()               {}
func (*SLInitMsg) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

// Server's response to the SLInitMsg.
// On Success (ErrStatus), the session with the server is established
// and the client is allowed to proceed.
type SLInitMsgRsp struct {
	// Server's version of service-layer API (refer to sl_version.proto)
	// Major-number revisions are NOT backwards compatible,
	// unless otherwise specified. The Server may reject a session if there
	// is a version number mismatch or non-backwards compatibility.
	MajorVer uint32 `protobuf:"varint,1,opt,name=MajorVer,json=majorVer" json:"MajorVer,omitempty"`
	// Minor Version
	MinorVer uint32 `protobuf:"varint,2,opt,name=MinorVer,json=minorVer" json:"MinorVer,omitempty"`
	// Sub-Version
	SubVer uint32 `protobuf:"varint,3,opt,name=SubVer,json=subVer" json:"SubVer,omitempty"`
}

func (m *SLInitMsgRsp) Reset()                    { *m = SLInitMsgRsp{} }
func (m *SLInitMsgRsp) String() string            { return proto.CompactTextString(m) }
func (*SLInitMsgRsp) ProtoMessage()               {}
func (*SLInitMsgRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

// Globals query message.
type SLGlobalNotif struct {
	// Event Type.
	EventType SLGlobalNotifType `protobuf:"varint,1,opt,name=EventType,json=eventType,enum=service_layer.SLGlobalNotifType" json:"EventType,omitempty"`
	// Status code, interpreted based on the Event Type.
	//
	//   case EventType == SL_GLOBAL_EVENT_TYPE_ERROR:
	//       case ErrStatus == SL_NOTIF_TERM:
	//          => Another client is attempting to take over the session.
	//             This session will be closed.
	//       case ErrStatus == (some error from SLErrorStatus)
	//          => Client must look into the specific error message returned.
	//
	//   case EventType == SL_GLOBAL_EVENT_TYPE_HEARTBEAT:
	//       case ErrStatus == SL_SUCCESS
	//          => Client can safely ignore this heartbeat message.
	//
	//   case EventType == SL_GLOBAL_EVENT_TYPE_VERSION:
	//       case ErrStatus == SL_SUCCESS
	//          => Client version accepted.
	//       case ErrStatus == SL_INIT_STATE_READY
	//          => Client version accepted.
	//             Any previous state was sucessfully recovered.
	//       case ErrStatus == SL_INIT_STATE_CLEAR
	//          => Client version accepted. Any previous state was lost.
	//             Client must replay all previous objects to server.
	//       case ErrStatus == SL_UNSUPPORTED_VER
	//          => Client and Server version mismatch. The client is not
	//             allowed to proceed, and the channel will be closed.
	//       case ErrStatus == (some error from SLErrorStatus)
	//          => Client must either try again, or look into the specific
	//             error message returned.
	ErrStatus *SLErrorStatus `protobuf:"bytes,2,opt,name=ErrStatus,json=errStatus" json:"ErrStatus,omitempty"`
	// Further info based on EventType.
	//
	// Types that are valid to be assigned to Event:
	//	*SLGlobalNotif_InitRspMsg
	Event isSLGlobalNotif_Event `protobuf_oneof:"Event"`
}

func (m *SLGlobalNotif) Reset()                    { *m = SLGlobalNotif{} }
func (m *SLGlobalNotif) String() string            { return proto.CompactTextString(m) }
func (*SLGlobalNotif) ProtoMessage()               {}
func (*SLGlobalNotif) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

type isSLGlobalNotif_Event interface {
	isSLGlobalNotif_Event()
}

type SLGlobalNotif_InitRspMsg struct {
	InitRspMsg *SLInitMsgRsp `protobuf:"bytes,3,opt,name=InitRspMsg,json=initRspMsg,oneof"`
}

func (*SLGlobalNotif_InitRspMsg) isSLGlobalNotif_Event() {}

func (m *SLGlobalNotif) GetEvent() isSLGlobalNotif_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *SLGlobalNotif) GetErrStatus() *SLErrorStatus {
	if m != nil {
		return m.ErrStatus
	}
	return nil
}

func (m *SLGlobalNotif) GetInitRspMsg() *SLInitMsgRsp {
	if x, ok := m.GetEvent().(*SLGlobalNotif_InitRspMsg); ok {
		return x.InitRspMsg
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SLGlobalNotif) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SLGlobalNotif_OneofMarshaler, _SLGlobalNotif_OneofUnmarshaler, _SLGlobalNotif_OneofSizer, []interface{}{
		(*SLGlobalNotif_InitRspMsg)(nil),
	}
}

func _SLGlobalNotif_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SLGlobalNotif)
	// Event
	switch x := m.Event.(type) {
	case *SLGlobalNotif_InitRspMsg:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.InitRspMsg); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SLGlobalNotif.Event has unexpected type %T", x)
	}
	return nil
}

func _SLGlobalNotif_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SLGlobalNotif)
	switch tag {
	case 3: // Event.InitRspMsg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SLInitMsgRsp)
		err := b.DecodeMessage(msg)
		m.Event = &SLGlobalNotif_InitRspMsg{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SLGlobalNotif_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SLGlobalNotif)
	// Event
	switch x := m.Event.(type) {
	case *SLGlobalNotif_InitRspMsg:
		s := proto.Size(x.InitRspMsg)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Globals query message.
type SLGlobalsGetMsg struct {
}

func (m *SLGlobalsGetMsg) Reset()                    { *m = SLGlobalsGetMsg{} }
func (m *SLGlobalsGetMsg) String() string            { return proto.CompactTextString(m) }
func (*SLGlobalsGetMsg) ProtoMessage()               {}
func (*SLGlobalsGetMsg) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

// Platform specific globals Response.
type SLGlobalsGetMsgRsp struct {
	// Corresponding error code
	ErrStatus *SLErrorStatus `protobuf:"bytes,1,opt,name=ErrStatus,json=errStatus" json:"ErrStatus,omitempty"`
	// Maximum vrf name length.
	MaxVrfNameLength uint32 `protobuf:"varint,2,opt,name=MaxVrfNameLength,json=maxVrfNameLength" json:"MaxVrfNameLength,omitempty"`
	// Maximum interface name length.
	MaxInterfaceNameLength uint32 `protobuf:"varint,3,opt,name=MaxInterfaceNameLength,json=maxInterfaceNameLength" json:"MaxInterfaceNameLength,omitempty"`
	// Maximum paths per Route/ILM Entry.
	MaxPathsPerEntry uint32 `protobuf:"varint,4,opt,name=MaxPathsPerEntry,json=maxPathsPerEntry" json:"MaxPathsPerEntry,omitempty"`
	// Maximum primary paths per Route/ILM Entry.
	MaxPrimaryPathPerEntry uint32 `protobuf:"varint,5,opt,name=MaxPrimaryPathPerEntry,json=maxPrimaryPathPerEntry" json:"MaxPrimaryPathPerEntry,omitempty"`
	// Maximum backup paths per Route/ILM Entry.
	MaxBackupPathPerEntry uint32 `protobuf:"varint,6,opt,name=MaxBackupPathPerEntry,json=maxBackupPathPerEntry" json:"MaxBackupPathPerEntry,omitempty"`
	// Maximum MPLS labels per Route/ILM Entry.
	MaxMplsLabelsPerPath uint32 `protobuf:"varint,7,opt,name=MaxMplsLabelsPerPath,json=maxMplsLabelsPerPath" json:"MaxMplsLabelsPerPath,omitempty"`
	// Minimum Primary path id number.
	MinPrimaryPathIdNum uint32 `protobuf:"varint,8,opt,name=MinPrimaryPathIdNum,json=minPrimaryPathIdNum" json:"MinPrimaryPathIdNum,omitempty"`
	// Maximum Primary path id number.
	MaxPrimaryPathIdNum uint32 `protobuf:"varint,9,opt,name=MaxPrimaryPathIdNum,json=maxPrimaryPathIdNum" json:"MaxPrimaryPathIdNum,omitempty"`
	// Minimum Pure Backup path id number.
	MinBackupPathIdNum uint32 `protobuf:"varint,10,opt,name=MinBackupPathIdNum,json=minBackupPathIdNum" json:"MinBackupPathIdNum,omitempty"`
	// Maximum Pure Backup path id number.
	MaxBackupPathIdNum uint32 `protobuf:"varint,11,opt,name=MaxBackupPathIdNum,json=maxBackupPathIdNum" json:"MaxBackupPathIdNum,omitempty"`
	// Maximum number of remote addresses
	MaxRemoteAddressNum uint32 `protobuf:"varint,12,opt,name=MaxRemoteAddressNum,json=maxRemoteAddressNum" json:"MaxRemoteAddressNum,omitempty"`
}

func (m *SLGlobalsGetMsgRsp) Reset()                    { *m = SLGlobalsGetMsgRsp{} }
func (m *SLGlobalsGetMsgRsp) String() string            { return proto.CompactTextString(m) }
func (*SLGlobalsGetMsgRsp) ProtoMessage()               {}
func (*SLGlobalsGetMsgRsp) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *SLGlobalsGetMsgRsp) GetErrStatus() *SLErrorStatus {
	if m != nil {
		return m.ErrStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*SLInitMsg)(nil), "service_layer.SLInitMsg")
	proto.RegisterType((*SLInitMsgRsp)(nil), "service_layer.SLInitMsgRsp")
	proto.RegisterType((*SLGlobalNotif)(nil), "service_layer.SLGlobalNotif")
	proto.RegisterType((*SLGlobalsGetMsg)(nil), "service_layer.SLGlobalsGetMsg")
	proto.RegisterType((*SLGlobalsGetMsgRsp)(nil), "service_layer.SLGlobalsGetMsgRsp")
	proto.RegisterEnum("service_layer.SLGlobalNotifType", SLGlobalNotifType_name, SLGlobalNotifType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for SLGlobal service

type SLGlobalClient interface {
	// Initialize the connection, and setup a notification channel.
	// This MUST be the first call to setup the Service Layer connection.
	//
	// The caller MUST maintain the notification channel to be able to
	// communicate with the server.
	// If this channel is not properly established and maintained, all other
	// RPC requests are rejected.
	//
	// The caller must send its version information as part of the SLInitMsg
	// message. The server will reply with SL_GLOBAL_EVENT_TYPE_VERSION
	// that tells the caller whether he can proceed or not.
	// Refer to message SLGlobalNotif below for further details.
	//
	// After the version handshake, the notification channel is used for
	// "push" event notifications, such as:
	//    - SLGlobalNotif.EventType = SL_GLOBAL_EVENT_TYPE_HEARTBEAT
	//        heartbeat notification messages are sent to the client on
	//        a periodic basis.
	//    Refer to SLGlobalNotif definition for further info.
	SLGlobalInitNotif(ctx context.Context, in *SLInitMsg, opts ...grpc.CallOption) (SLGlobal_SLGlobalInitNotifClient, error)
	// Get platform specific globals
	SLGlobalsGet(ctx context.Context, in *SLGlobalsGetMsg, opts ...grpc.CallOption) (*SLGlobalsGetMsgRsp, error)
}

type sLGlobalClient struct {
	cc *grpc.ClientConn
}

func NewSLGlobalClient(cc *grpc.ClientConn) SLGlobalClient {
	return &sLGlobalClient{cc}
}

func (c *sLGlobalClient) SLGlobalInitNotif(ctx context.Context, in *SLInitMsg, opts ...grpc.CallOption) (SLGlobal_SLGlobalInitNotifClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SLGlobal_serviceDesc.Streams[0], c.cc, "/service_layer.SLGlobal/SLGlobalInitNotif", opts...)
	if err != nil {
		return nil, err
	}
	x := &sLGlobalSLGlobalInitNotifClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SLGlobal_SLGlobalInitNotifClient interface {
	Recv() (*SLGlobalNotif, error)
	grpc.ClientStream
}

type sLGlobalSLGlobalInitNotifClient struct {
	grpc.ClientStream
}

func (x *sLGlobalSLGlobalInitNotifClient) Recv() (*SLGlobalNotif, error) {
	m := new(SLGlobalNotif)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sLGlobalClient) SLGlobalsGet(ctx context.Context, in *SLGlobalsGetMsg, opts ...grpc.CallOption) (*SLGlobalsGetMsgRsp, error) {
	out := new(SLGlobalsGetMsgRsp)
	err := grpc.Invoke(ctx, "/service_layer.SLGlobal/SLGlobalsGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SLGlobal service

type SLGlobalServer interface {
	// Initialize the connection, and setup a notification channel.
	// This MUST be the first call to setup the Service Layer connection.
	//
	// The caller MUST maintain the notification channel to be able to
	// communicate with the server.
	// If this channel is not properly established and maintained, all other
	// RPC requests are rejected.
	//
	// The caller must send its version information as part of the SLInitMsg
	// message. The server will reply with SL_GLOBAL_EVENT_TYPE_VERSION
	// that tells the caller whether he can proceed or not.
	// Refer to message SLGlobalNotif below for further details.
	//
	// After the version handshake, the notification channel is used for
	// "push" event notifications, such as:
	//    - SLGlobalNotif.EventType = SL_GLOBAL_EVENT_TYPE_HEARTBEAT
	//        heartbeat notification messages are sent to the client on
	//        a periodic basis.
	//    Refer to SLGlobalNotif definition for further info.
	SLGlobalInitNotif(*SLInitMsg, SLGlobal_SLGlobalInitNotifServer) error
	// Get platform specific globals
	SLGlobalsGet(context.Context, *SLGlobalsGetMsg) (*SLGlobalsGetMsgRsp, error)
}

func RegisterSLGlobalServer(s *grpc.Server, srv SLGlobalServer) {
	s.RegisterService(&_SLGlobal_serviceDesc, srv)
}

func _SLGlobal_SLGlobalInitNotif_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SLInitMsg)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SLGlobalServer).SLGlobalInitNotif(m, &sLGlobalSLGlobalInitNotifServer{stream})
}

type SLGlobal_SLGlobalInitNotifServer interface {
	Send(*SLGlobalNotif) error
	grpc.ServerStream
}

type sLGlobalSLGlobalInitNotifServer struct {
	grpc.ServerStream
}

func (x *sLGlobalSLGlobalInitNotifServer) Send(m *SLGlobalNotif) error {
	return x.ServerStream.SendMsg(m)
}

func _SLGlobal_SLGlobalsGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SLGlobalsGetMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SLGlobalServer).SLGlobalsGet(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _SLGlobal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service_layer.SLGlobal",
	HandlerType: (*SLGlobalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SLGlobalsGet",
			Handler:    _SLGlobal_SLGlobalsGet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SLGlobalInitNotif",
			Handler:       _SLGlobal_SLGlobalInitNotif_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor4 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x94, 0x5d, 0x6b, 0xdb, 0x3e,
	0x14, 0xc6, 0xeb, 0xbe, 0xa4, 0xc9, 0x69, 0xfb, 0xaf, 0xab, 0xbe, 0x10, 0xf2, 0xef, 0x4a, 0x9b,
	0xab, 0xd1, 0x8b, 0x50, 0xb2, 0xb1, 0x8b, 0xc1, 0x06, 0x09, 0x13, 0x6d, 0xc0, 0x49, 0x33, 0x39,
	0x04, 0xc6, 0x60, 0xc6, 0x49, 0xd5, 0xd6, 0x9b, 0x65, 0x07, 0x49, 0x29, 0xcb, 0x77, 0xd9, 0x47,
	0xd8, 0xb7, 0xd9, 0x97, 0xd9, 0xe5, 0x64, 0xd9, 0x71, 0xea, 0xd8, 0x2d, 0x0c, 0x76, 0x29, 0x3d,
	0xcf, 0xcf, 0xe7, 0x39, 0x4a, 0xce, 0x81, 0x5d, 0xe1, 0x3b, 0x77, 0x7e, 0x38, 0x72, 0xfd, 0xc6,
	0x84, 0x87, 0x32, 0x44, 0x3b, 0x82, 0xf2, 0x07, 0x6f, 0x4c, 0x1d, 0xdf, 0x9d, 0x51, 0x5e, 0x3b,
	0x54, 0xfa, 0x38, 0x64, 0x2c, 0x0c, 0x1c, 0x39, 0x9b, 0x50, 0x11, 0xbb, 0xea, 0x9f, 0xa1, 0x62,
	0x5b, 0x9d, 0xc0, 0x93, 0x5d, 0x71, 0x87, 0x6a, 0x50, 0xee, 0xba, 0x5f, 0x43, 0x3e, 0xa4, 0xbc,
	0x6a, 0x9c, 0x1a, 0x2f, 0x77, 0x48, 0x99, 0x25, 0x67, 0xad, 0x79, 0x41, 0xac, 0xad, 0x26, 0x5a,
	0x72, 0x46, 0x47, 0x50, 0xb2, 0xa7, 0xa3, 0x48, 0x59, 0xd3, 0x4a, 0x49, 0xe8, 0x53, 0xfd, 0x0b,
	0x6c, 0xa7, 0x1f, 0x27, 0x62, 0xf2, 0xcf, 0xbf, 0xff, 0xcb, 0x80, 0x1d, 0xdb, 0xba, 0xd4, 0x5d,
	0xf7, 0x42, 0xe9, 0xdd, 0xa2, 0xf7, 0x50, 0xc1, 0x0f, 0x34, 0x90, 0x03, 0xd5, 0xa2, 0x2e, 0xf1,
	0x5f, 0xf3, 0xb4, 0x91, 0x79, 0x88, 0x46, 0x06, 0x88, 0x7c, 0xa4, 0x42, 0xe7, 0x08, 0x7a, 0xab,
	0x78, 0xce, 0x6d, 0xe9, 0xca, 0xa9, 0xd0, 0x31, 0xb6, 0x9a, 0xc7, 0x39, 0x5e, 0x39, 0xc2, 0xc4,
	0xa3, 0xd8, 0xb9, 0x1d, 0xbd, 0x03, 0x88, 0x7a, 0x55, 0x8d, 0xaa, 0x76, 0x75, 0xd2, 0xad, 0xe6,
	0xff, 0x39, 0x78, 0xf1, 0x1c, 0x57, 0x2b, 0x04, 0xbc, 0x14, 0x68, 0x6f, 0xc2, 0x86, 0x8e, 0x5e,
	0xdf, 0x83, 0xdd, 0x79, 0x46, 0x71, 0x49, 0x23, 0x73, 0xfd, 0xf7, 0x3a, 0xa0, 0xa5, 0xbb, 0xe8,
	0x3d, 0x33, 0x69, 0x8d, 0xbf, 0x4b, 0x7b, 0x0e, 0x66, 0xd7, 0xfd, 0x3e, 0xe4, 0xb7, 0x3d, 0x97,
	0x51, 0x8b, 0x06, 0x77, 0xf2, 0x3e, 0x79, 0x77, 0x93, 0x2d, 0xdd, 0xa3, 0x37, 0x70, 0xa4, 0xbc,
	0x9d, 0x40, 0x52, 0x7e, 0xeb, 0x8e, 0xe9, 0x23, 0x22, 0xfe, 0x3d, 0x8e, 0x58, 0xa1, 0x9a, 0xd4,
	0xe8, 0xbb, 0xf2, 0x5e, 0xf4, 0x29, 0xc7, 0x81, 0xe4, 0xb3, 0xea, 0x7a, 0x5a, 0x23, 0x73, 0x9f,
	0xd4, 0xe8, 0x73, 0x8f, 0xb9, 0x7c, 0x16, 0x49, 0x29, 0xb1, 0x91, 0xd6, 0x28, 0x50, 0xd1, 0x6b,
	0x38, 0x54, 0x5c, 0xdb, 0x1d, 0x7f, 0x9b, 0x4e, 0x32, 0x58, 0x49, 0x63, 0x87, 0xac, 0x48, 0x44,
	0x4d, 0x38, 0x50, 0x54, 0x77, 0xe2, 0x0b, 0xcb, 0x1d, 0x51, 0x3f, 0x8a, 0x11, 0xe9, 0xd5, 0x4d,
	0x0d, 0x1d, 0xb0, 0x02, 0x0d, 0x5d, 0xc0, 0xbe, 0xfa, 0x87, 0x3e, 0xca, 0xd0, 0xb9, 0xe9, 0x4d,
	0x59, 0xb5, 0xac, 0x91, 0x7d, 0x96, 0x97, 0x34, 0x91, 0x49, 0x1d, 0x13, 0x95, 0x84, 0xc8, 0x4b,
	0xa8, 0x01, 0x48, 0xd5, 0x58, 0x04, 0x8e, 0x01, 0xd0, 0x00, 0x62, 0x39, 0x45, 0xfb, 0x1f, 0x37,
	0x18, 0xfb, 0xb7, 0x12, 0x7f, 0x4e, 0x49, 0x12, 0x11, 0xca, 0x42, 0x49, 0x5b, 0x37, 0x37, 0x9c,
	0x0a, 0x11, 0x01, 0xdb, 0x69, 0xa2, 0x65, 0xe9, 0xfc, 0x87, 0x01, 0x7b, 0xb9, 0x91, 0x41, 0x67,
	0xf0, 0xc2, 0xb6, 0x9c, 0x4b, 0xeb, 0xba, 0xdd, 0xb2, 0x1c, 0x3c, 0xc4, 0xbd, 0x81, 0x33, 0xf8,
	0xd4, 0xc7, 0x0e, 0xc1, 0x36, 0x26, 0x43, 0xfc, 0xc1, 0x5c, 0x41, 0x27, 0x50, 0x2b, 0xb4, 0x60,
	0x42, 0xae, 0x89, 0x69, 0xa0, 0x3a, 0x9c, 0x14, 0xea, 0x57, 0xb8, 0x45, 0x06, 0x6d, 0xdc, 0x1a,
	0x98, 0xab, 0xe8, 0x14, 0x8e, 0x0b, 0x3d, 0x43, 0x4c, 0xec, 0xce, 0x75, 0xcf, 0x5c, 0x6b, 0xfe,
	0x34, 0xa0, 0x3c, 0x8f, 0x87, 0xba, 0x8b, 0xa8, 0xd1, 0x98, 0xc5, 0x2b, 0xa1, 0xfa, 0xd4, 0x08,
	0xd6, 0x8e, 0x9f, 0xdb, 0x0c, 0x17, 0x06, 0xfa, 0x18, 0xad, 0xaf, 0xc5, 0xd0, 0xa1, 0x93, 0x27,
	0xfc, 0xc9, 0x44, 0xd6, 0xce, 0x9e, 0xd7, 0xd5, 0xc4, 0x8e, 0x4a, 0x7a, 0xeb, 0xbe, 0xfa, 0x13,
	0x00, 0x00, 0xff, 0xff, 0xfe, 0xdf, 0x07, 0xca, 0xae, 0x05, 0x00, 0x00,
}
