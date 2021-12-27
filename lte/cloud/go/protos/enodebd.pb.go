// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lte/protos/enodebd.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protos "magma/orc8r/lib/go/protos"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SingleEnodebStatus_StatusProperty int32

const (
	SingleEnodebStatus_OFF     SingleEnodebStatus_StatusProperty = 0
	SingleEnodebStatus_ON      SingleEnodebStatus_StatusProperty = 1
	SingleEnodebStatus_UNKNOWN SingleEnodebStatus_StatusProperty = 2
)

var SingleEnodebStatus_StatusProperty_name = map[int32]string{
	0: "OFF",
	1: "ON",
	2: "UNKNOWN",
}

var SingleEnodebStatus_StatusProperty_value = map[string]int32{
	"OFF":     0,
	"ON":      1,
	"UNKNOWN": 2,
}

func (x SingleEnodebStatus_StatusProperty) String() string {
	return proto.EnumName(SingleEnodebStatus_StatusProperty_name, int32(x))
}

func (SingleEnodebStatus_StatusProperty) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b803bda4b235859b, []int{6, 0}
}

// --------------------------------------------------------------------------
// Message Definitions for TR-069 message injection. This is used for manual
// testing of the TR-069 server.
// --------------------------------------------------------------------------
type GetParameterRequest struct {
	// Serial ID of eNodeB. Uniquely identifies the eNodeB.
	DeviceSerial string `protobuf:"bytes,1,opt,name=device_serial,json=deviceSerial,proto3" json:"device_serial,omitempty"`
	// Fully qualified parameter name, e.g:
	// InternetGatewayDevice.LANDevice.1.Hosts.
	ParameterName        string   `protobuf:"bytes,2,opt,name=parameter_name,json=parameterName,proto3" json:"parameter_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetParameterRequest) Reset()         { *m = GetParameterRequest{} }
func (m *GetParameterRequest) String() string { return proto.CompactTextString(m) }
func (*GetParameterRequest) ProtoMessage()    {}
func (*GetParameterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b803bda4b235859b, []int{0}
}

func (m *GetParameterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetParameterRequest.Unmarshal(m, b)
}
func (m *GetParameterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetParameterRequest.Marshal(b, m, deterministic)
}
func (m *GetParameterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetParameterRequest.Merge(m, src)
}
func (m *GetParameterRequest) XXX_Size() int {
	return xxx_messageInfo_GetParameterRequest.Size(m)
}
func (m *GetParameterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetParameterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetParameterRequest proto.InternalMessageInfo

func (m *GetParameterRequest) GetDeviceSerial() string {
	if m != nil {
		return m.DeviceSerial
	}
	return ""
}

func (m *GetParameterRequest) GetParameterName() string {
	if m != nil {
		return m.ParameterName
	}
	return ""
}

type NameValue struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Note: parameter value is always passed back as string. Up to calling
	// function to determine type
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NameValue) Reset()         { *m = NameValue{} }
func (m *NameValue) String() string { return proto.CompactTextString(m) }
func (*NameValue) ProtoMessage()    {}
func (*NameValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_b803bda4b235859b, []int{1}
}

func (m *NameValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameValue.Unmarshal(m, b)
}
func (m *NameValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameValue.Marshal(b, m, deterministic)
}
func (m *NameValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameValue.Merge(m, src)
}
func (m *NameValue) XXX_Size() int {
	return xxx_messageInfo_NameValue.Size(m)
}
func (m *NameValue) XXX_DiscardUnknown() {
	xxx_messageInfo_NameValue.DiscardUnknown(m)
}

var xxx_messageInfo_NameValue proto.InternalMessageInfo

func (m *NameValue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NameValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GetParameterResponse struct {
	DeviceSerial         string       `protobuf:"bytes,1,opt,name=device_serial,json=deviceSerial,proto3" json:"device_serial,omitempty"`
	Parameters           []*NameValue `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetParameterResponse) Reset()         { *m = GetParameterResponse{} }
func (m *GetParameterResponse) String() string { return proto.CompactTextString(m) }
func (*GetParameterResponse) ProtoMessage()    {}
func (*GetParameterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b803bda4b235859b, []int{2}
}

func (m *GetParameterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetParameterResponse.Unmarshal(m, b)
}
func (m *GetParameterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetParameterResponse.Marshal(b, m, deterministic)
}
func (m *GetParameterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetParameterResponse.Merge(m, src)
}
func (m *GetParameterResponse) XXX_Size() int {
	return xxx_messageInfo_GetParameterResponse.Size(m)
}
func (m *GetParameterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetParameterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetParameterResponse proto.InternalMessageInfo

func (m *GetParameterResponse) GetDeviceSerial() string {
	if m != nil {
		return m.DeviceSerial
	}
	return ""
}

func (m *GetParameterResponse) GetParameters() []*NameValue {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type SetParameterRequest struct {
	// Serial ID of eNodeB. Uniquely identifies the eNodeB.
	DeviceSerial string `protobuf:"bytes,1,opt,name=device_serial,json=deviceSerial,proto3" json:"device_serial,omitempty"`
	// Fully qualified parameter name, e.g:
	// InternetGatewayDevice.LANDevice.1.Hosts.
	ParameterName string `protobuf:"bytes,2,opt,name=parameter_name,json=parameterName,proto3" json:"parameter_name,omitempty"`
	// Data values for each data type
	//
	// Types that are valid to be assigned to Value:
	//	*SetParameterRequest_ValueInt
	//	*SetParameterRequest_ValueString
	//	*SetParameterRequest_ValueBool
	Value isSetParameterRequest_Value `protobuf_oneof:"value"`
	// Key to be used at ACS discretion to determine when parameter was last
	// updated
	ParameterKey         string   `protobuf:"bytes,6,opt,name=parameter_key,json=parameterKey,proto3" json:"parameter_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetParameterRequest) Reset()         { *m = SetParameterRequest{} }
func (m *SetParameterRequest) String() string { return proto.CompactTextString(m) }
func (*SetParameterRequest) ProtoMessage()    {}
func (*SetParameterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b803bda4b235859b, []int{3}
}

func (m *SetParameterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetParameterRequest.Unmarshal(m, b)
}
func (m *SetParameterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetParameterRequest.Marshal(b, m, deterministic)
}
func (m *SetParameterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetParameterRequest.Merge(m, src)
}
func (m *SetParameterRequest) XXX_Size() int {
	return xxx_messageInfo_SetParameterRequest.Size(m)
}
func (m *SetParameterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetParameterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetParameterRequest proto.InternalMessageInfo

func (m *SetParameterRequest) GetDeviceSerial() string {
	if m != nil {
		return m.DeviceSerial
	}
	return ""
}

func (m *SetParameterRequest) GetParameterName() string {
	if m != nil {
		return m.ParameterName
	}
	return ""
}

type isSetParameterRequest_Value interface {
	isSetParameterRequest_Value()
}

type SetParameterRequest_ValueInt struct {
	ValueInt int32 `protobuf:"varint,3,opt,name=value_int,json=valueInt,proto3,oneof"`
}

type SetParameterRequest_ValueString struct {
	ValueString string `protobuf:"bytes,4,opt,name=value_string,json=valueString,proto3,oneof"`
}

type SetParameterRequest_ValueBool struct {
	ValueBool bool `protobuf:"varint,5,opt,name=value_bool,json=valueBool,proto3,oneof"`
}

func (*SetParameterRequest_ValueInt) isSetParameterRequest_Value() {}

func (*SetParameterRequest_ValueString) isSetParameterRequest_Value() {}

func (*SetParameterRequest_ValueBool) isSetParameterRequest_Value() {}

func (m *SetParameterRequest) GetValue() isSetParameterRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *SetParameterRequest) GetValueInt() int32 {
	if x, ok := m.GetValue().(*SetParameterRequest_ValueInt); ok {
		return x.ValueInt
	}
	return 0
}

func (m *SetParameterRequest) GetValueString() string {
	if x, ok := m.GetValue().(*SetParameterRequest_ValueString); ok {
		return x.ValueString
	}
	return ""
}

func (m *SetParameterRequest) GetValueBool() bool {
	if x, ok := m.GetValue().(*SetParameterRequest_ValueBool); ok {
		return x.ValueBool
	}
	return false
}

func (m *SetParameterRequest) GetParameterKey() string {
	if m != nil {
		return m.ParameterKey
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SetParameterRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SetParameterRequest_ValueInt)(nil),
		(*SetParameterRequest_ValueString)(nil),
		(*SetParameterRequest_ValueBool)(nil),
	}
}

type EnodebIdentity struct {
	// Serial ID of eNodeB. Uniquely identifies the eNodeB.
	DeviceSerial         string   `protobuf:"bytes,1,opt,name=device_serial,json=deviceSerial,proto3" json:"device_serial,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnodebIdentity) Reset()         { *m = EnodebIdentity{} }
func (m *EnodebIdentity) String() string { return proto.CompactTextString(m) }
func (*EnodebIdentity) ProtoMessage()    {}
func (*EnodebIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_b803bda4b235859b, []int{4}
}

func (m *EnodebIdentity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnodebIdentity.Unmarshal(m, b)
}
func (m *EnodebIdentity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnodebIdentity.Marshal(b, m, deterministic)
}
func (m *EnodebIdentity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnodebIdentity.Merge(m, src)
}
func (m *EnodebIdentity) XXX_Size() int {
	return xxx_messageInfo_EnodebIdentity.Size(m)
}
func (m *EnodebIdentity) XXX_DiscardUnknown() {
	xxx_messageInfo_EnodebIdentity.DiscardUnknown(m)
}

var xxx_messageInfo_EnodebIdentity proto.InternalMessageInfo

func (m *EnodebIdentity) GetDeviceSerial() string {
	if m != nil {
		return m.DeviceSerial
	}
	return ""
}

type AllEnodebStatus struct {
	EnbStatusList        []*SingleEnodebStatus `protobuf:"bytes,1,rep,name=enb_status_list,json=enbStatusList,proto3" json:"enb_status_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AllEnodebStatus) Reset()         { *m = AllEnodebStatus{} }
func (m *AllEnodebStatus) String() string { return proto.CompactTextString(m) }
func (*AllEnodebStatus) ProtoMessage()    {}
func (*AllEnodebStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_b803bda4b235859b, []int{5}
}

func (m *AllEnodebStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllEnodebStatus.Unmarshal(m, b)
}
func (m *AllEnodebStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllEnodebStatus.Marshal(b, m, deterministic)
}
func (m *AllEnodebStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllEnodebStatus.Merge(m, src)
}
func (m *AllEnodebStatus) XXX_Size() int {
	return xxx_messageInfo_AllEnodebStatus.Size(m)
}
func (m *AllEnodebStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_AllEnodebStatus.DiscardUnknown(m)
}

var xxx_messageInfo_AllEnodebStatus proto.InternalMessageInfo

func (m *AllEnodebStatus) GetEnbStatusList() []*SingleEnodebStatus {
	if m != nil {
		return m.EnbStatusList
	}
	return nil
}

type SingleEnodebStatus struct {
	DeviceSerial         string                            `protobuf:"bytes,1,opt,name=device_serial,json=deviceSerial,proto3" json:"device_serial,omitempty"`
	IpAddress            string                            `protobuf:"bytes,2,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	Connected            SingleEnodebStatus_StatusProperty `protobuf:"varint,3,opt,name=connected,proto3,enum=magma.lte.SingleEnodebStatus_StatusProperty" json:"connected,omitempty"`
	Configured           SingleEnodebStatus_StatusProperty `protobuf:"varint,4,opt,name=configured,proto3,enum=magma.lte.SingleEnodebStatus_StatusProperty" json:"configured,omitempty"`
	OpstateEnabled       SingleEnodebStatus_StatusProperty `protobuf:"varint,5,opt,name=opstate_enabled,json=opstateEnabled,proto3,enum=magma.lte.SingleEnodebStatus_StatusProperty" json:"opstate_enabled,omitempty"`
	RfTxOn               SingleEnodebStatus_StatusProperty `protobuf:"varint,6,opt,name=rf_tx_on,json=rfTxOn,proto3,enum=magma.lte.SingleEnodebStatus_StatusProperty" json:"rf_tx_on,omitempty"`
	GpsConnected         SingleEnodebStatus_StatusProperty `protobuf:"varint,7,opt,name=gps_connected,json=gpsConnected,proto3,enum=magma.lte.SingleEnodebStatus_StatusProperty" json:"gps_connected,omitempty"`
	PtpConnected         SingleEnodebStatus_StatusProperty `protobuf:"varint,8,opt,name=ptp_connected,json=ptpConnected,proto3,enum=magma.lte.SingleEnodebStatus_StatusProperty" json:"ptp_connected,omitempty"`
	MmeConnected         SingleEnodebStatus_StatusProperty `protobuf:"varint,9,opt,name=mme_connected,json=mmeConnected,proto3,enum=magma.lte.SingleEnodebStatus_StatusProperty" json:"mme_connected,omitempty"`
	GpsLongitude         string                            `protobuf:"bytes,10,opt,name=gps_longitude,json=gpsLongitude,proto3" json:"gps_longitude,omitempty"`
	GpsLatitude          string                            `protobuf:"bytes,11,opt,name=gps_latitude,json=gpsLatitude,proto3" json:"gps_latitude,omitempty"`
	FsmState             string                            `protobuf:"bytes,12,opt,name=fsm_state,json=fsmState,proto3" json:"fsm_state,omitempty"`
	RfTxDesired          SingleEnodebStatus_StatusProperty `protobuf:"varint,13,opt,name=rf_tx_desired,json=rfTxDesired,proto3,enum=magma.lte.SingleEnodebStatus_StatusProperty" json:"rf_tx_desired,omitempty"`
	GpsAltitude          string                            `protobuf:"bytes,14,opt,name=gps_altitude,json=gpsAltitude,proto3" json:"gps_altitude,omitempty"`
	Vendor               string                            `protobuf:"bytes,15,opt,name=vendor,proto3" json:"vendor,omitempty"`
	ModelName            string                            `protobuf:"bytes,16,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	RfState              SingleEnodebStatus_StatusProperty `protobuf:"varint,17,opt,name=rf_state,json=rfState,proto3,enum=magma.lte.SingleEnodebStatus_StatusProperty" json:"rf_state,omitempty"`
	SwVersion            string                            `protobuf:"bytes,18,opt,name=sw_version,json=swVersion,proto3" json:"sw_version,omitempty"`
	Uptime               string                            `protobuf:"bytes,19,opt,name=uptime,proto3" json:"uptime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *SingleEnodebStatus) Reset()         { *m = SingleEnodebStatus{} }
func (m *SingleEnodebStatus) String() string { return proto.CompactTextString(m) }
func (*SingleEnodebStatus) ProtoMessage()    {}
func (*SingleEnodebStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_b803bda4b235859b, []int{6}
}

func (m *SingleEnodebStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SingleEnodebStatus.Unmarshal(m, b)
}
func (m *SingleEnodebStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SingleEnodebStatus.Marshal(b, m, deterministic)
}
func (m *SingleEnodebStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingleEnodebStatus.Merge(m, src)
}
func (m *SingleEnodebStatus) XXX_Size() int {
	return xxx_messageInfo_SingleEnodebStatus.Size(m)
}
func (m *SingleEnodebStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_SingleEnodebStatus.DiscardUnknown(m)
}

var xxx_messageInfo_SingleEnodebStatus proto.InternalMessageInfo

func (m *SingleEnodebStatus) GetDeviceSerial() string {
	if m != nil {
		return m.DeviceSerial
	}
	return ""
}

func (m *SingleEnodebStatus) GetIpAddress() string {
	if m != nil {
		return m.IpAddress
	}
	return ""
}

func (m *SingleEnodebStatus) GetConnected() SingleEnodebStatus_StatusProperty {
	if m != nil {
		return m.Connected
	}
	return SingleEnodebStatus_OFF
}

func (m *SingleEnodebStatus) GetConfigured() SingleEnodebStatus_StatusProperty {
	if m != nil {
		return m.Configured
	}
	return SingleEnodebStatus_OFF
}

func (m *SingleEnodebStatus) GetOpstateEnabled() SingleEnodebStatus_StatusProperty {
	if m != nil {
		return m.OpstateEnabled
	}
	return SingleEnodebStatus_OFF
}

func (m *SingleEnodebStatus) GetRfTxOn() SingleEnodebStatus_StatusProperty {
	if m != nil {
		return m.RfTxOn
	}
	return SingleEnodebStatus_OFF
}

func (m *SingleEnodebStatus) GetGpsConnected() SingleEnodebStatus_StatusProperty {
	if m != nil {
		return m.GpsConnected
	}
	return SingleEnodebStatus_OFF
}

func (m *SingleEnodebStatus) GetPtpConnected() SingleEnodebStatus_StatusProperty {
	if m != nil {
		return m.PtpConnected
	}
	return SingleEnodebStatus_OFF
}

func (m *SingleEnodebStatus) GetMmeConnected() SingleEnodebStatus_StatusProperty {
	if m != nil {
		return m.MmeConnected
	}
	return SingleEnodebStatus_OFF
}

func (m *SingleEnodebStatus) GetGpsLongitude() string {
	if m != nil {
		return m.GpsLongitude
	}
	return ""
}

func (m *SingleEnodebStatus) GetGpsLatitude() string {
	if m != nil {
		return m.GpsLatitude
	}
	return ""
}

func (m *SingleEnodebStatus) GetFsmState() string {
	if m != nil {
		return m.FsmState
	}
	return ""
}

func (m *SingleEnodebStatus) GetRfTxDesired() SingleEnodebStatus_StatusProperty {
	if m != nil {
		return m.RfTxDesired
	}
	return SingleEnodebStatus_OFF
}

func (m *SingleEnodebStatus) GetGpsAltitude() string {
	if m != nil {
		return m.GpsAltitude
	}
	return ""
}

func (m *SingleEnodebStatus) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *SingleEnodebStatus) GetModelName() string {
	if m != nil {
		return m.ModelName
	}
	return ""
}

func (m *SingleEnodebStatus) GetRfState() SingleEnodebStatus_StatusProperty {
	if m != nil {
		return m.RfState
	}
	return SingleEnodebStatus_OFF
}

func (m *SingleEnodebStatus) GetSwVersion() string {
	if m != nil {
		return m.SwVersion
	}
	return ""
}

func (m *SingleEnodebStatus) GetUptime() string {
	if m != nil {
		return m.Uptime
	}
	return ""
}

func init() {
	proto.RegisterEnum("magma.lte.SingleEnodebStatus_StatusProperty", SingleEnodebStatus_StatusProperty_name, SingleEnodebStatus_StatusProperty_value)
	proto.RegisterType((*GetParameterRequest)(nil), "magma.lte.GetParameterRequest")
	proto.RegisterType((*NameValue)(nil), "magma.lte.NameValue")
	proto.RegisterType((*GetParameterResponse)(nil), "magma.lte.GetParameterResponse")
	proto.RegisterType((*SetParameterRequest)(nil), "magma.lte.SetParameterRequest")
	proto.RegisterType((*EnodebIdentity)(nil), "magma.lte.EnodebIdentity")
	proto.RegisterType((*AllEnodebStatus)(nil), "magma.lte.AllEnodebStatus")
	proto.RegisterType((*SingleEnodebStatus)(nil), "magma.lte.SingleEnodebStatus")
}

func init() { proto.RegisterFile("lte/protos/enodebd.proto", fileDescriptor_b803bda4b235859b) }

var fileDescriptor_b803bda4b235859b = []byte{
	// 881 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x8e, 0xf3, 0x63, 0x5b, 0xc7, 0x7f, 0x29, 0x13, 0x0c, 0x8a, 0x0b, 0xaf, 0x99, 0x8b, 0x01,
	0xb9, 0x18, 0x9c, 0x2d, 0x59, 0x87, 0x6d, 0x77, 0x4e, 0x96, 0x38, 0x5d, 0x02, 0x3b, 0x95, 0xd7,
	0x6c, 0xd8, 0x8d, 0x20, 0x5b, 0xc7, 0x02, 0x31, 0x8a, 0x54, 0x49, 0x3a, 0x6d, 0x1e, 0x63, 0xcf,
	0xb5, 0xd7, 0xd9, 0x03, 0x0c, 0x22, 0x65, 0x47, 0x6e, 0xdd, 0x20, 0x31, 0xb0, 0x2b, 0x89, 0xdf,
	0xf9, 0xce, 0xa7, 0x8f, 0x87, 0x14, 0x0f, 0xc1, 0x65, 0x1a, 0x0f, 0x13, 0x29, 0xb4, 0x50, 0x87,
	0xc8, 0x45, 0x88, 0xa3, 0xb0, 0x63, 0x86, 0xc4, 0x89, 0x83, 0x28, 0x0e, 0x3a, 0x4c, 0x63, 0xb3,
	0x25, 0xe4, 0xf8, 0x47, 0x39, 0xa3, 0x29, 0x94, 0xb7, 0x74, 0x8c, 0xc7, 0xdf, 0x1e, 0x5b, 0x66,
	0x73, 0x6f, 0x21, 0x3c, 0x16, 0x71, 0x2c, 0xb8, 0x0d, 0xb5, 0x03, 0xd8, 0xe9, 0xa1, 0xbe, 0x0e,
	0x64, 0x10, 0xa3, 0x46, 0xe9, 0xe1, 0xbb, 0x29, 0x2a, 0x4d, 0x5e, 0x42, 0x2d, 0xc4, 0x54, 0xc4,
	0x57, 0x28, 0x69, 0xc0, 0xdc, 0xc2, 0x7e, 0xe1, 0xc0, 0xf1, 0xaa, 0x16, 0x1c, 0x1a, 0x8c, 0x7c,
	0x0d, 0xf5, 0x64, 0x96, 0xe8, 0xf3, 0x20, 0x46, 0x77, 0xdd, 0xb0, 0x6a, 0x73, 0xb4, 0x1f, 0xc4,
	0xd8, 0x7e, 0x05, 0x4e, 0xfa, 0xbc, 0x09, 0xd8, 0x14, 0x09, 0x81, 0x4d, 0xc3, 0xb4, 0x7a, 0xe6,
	0x9d, 0xec, 0xc2, 0xd6, 0x6d, 0x1a, 0xcc, 0xd2, 0xed, 0xa0, 0xfd, 0x0e, 0x76, 0x17, 0x9d, 0xa9,
	0x44, 0x70, 0x85, 0x8f, 0xb3, 0xf6, 0x3d, 0xc0, 0xdc, 0x84, 0x72, 0xd7, 0xf7, 0x37, 0x0e, 0x2a,
	0x47, 0xbb, 0x9d, 0x79, 0xc1, 0x3a, 0x73, 0x43, 0x5e, 0x8e, 0xd7, 0xfe, 0xb7, 0x00, 0x3b, 0xc3,
	0xff, 0xb7, 0x1a, 0xa4, 0x05, 0x8e, 0x99, 0x9f, 0x4f, 0xb9, 0x76, 0x37, 0xf6, 0x0b, 0x07, 0x5b,
	0x17, 0x6b, 0x5e, 0xd9, 0x40, 0xaf, 0x79, 0xfa, 0xa9, 0xaa, 0x0d, 0x2b, 0x2d, 0x29, 0x8f, 0xdc,
	0xcd, 0x54, 0xe3, 0x62, 0xcd, 0xab, 0x18, 0x74, 0x68, 0x40, 0xf2, 0x02, 0xc0, 0x92, 0x46, 0x42,
	0x30, 0x77, 0x6b, 0xbf, 0x70, 0x50, 0xbe, 0x58, 0xf3, 0xac, 0xee, 0x89, 0x10, 0x2c, 0x35, 0x7c,
	0xef, 0xe5, 0x2f, 0xbc, 0x73, 0x8b, 0xd6, 0xf0, 0x1c, 0xbc, 0xc4, 0xbb, 0x93, 0x52, 0x56, 0xf6,
	0xf6, 0x2b, 0xa8, 0x9f, 0x99, 0x9d, 0xf5, 0x3a, 0x44, 0xae, 0xa9, 0xbe, 0x7b, 0xd4, 0x84, 0xdb,
	0x7f, 0x40, 0xa3, 0xcb, 0x98, 0xcd, 0x1c, 0xea, 0x40, 0x4f, 0x15, 0x39, 0x83, 0x06, 0xf2, 0x91,
	0xaf, 0xcc, 0xc8, 0x67, 0x54, 0x69, 0xb7, 0x60, 0x6a, 0xdf, 0xca, 0xd5, 0x7e, 0x48, 0x79, 0xc4,
	0x30, 0x9f, 0xe7, 0xd5, 0x90, 0x67, 0xaf, 0x57, 0x54, 0xe9, 0xf6, 0x3f, 0x65, 0x20, 0x9f, 0xb2,
	0x1e, 0xb7, 0x0c, 0x2d, 0x00, 0x9a, 0xf8, 0x41, 0x18, 0x4a, 0x54, 0x2a, 0x5b, 0x02, 0x87, 0x26,
	0x5d, 0x0b, 0x90, 0x5f, 0xc1, 0x19, 0x0b, 0xce, 0x71, 0xac, 0x31, 0x34, 0xe5, 0xaf, 0x1f, 0x7d,
	0xf3, 0xa0, 0xb7, 0x8e, 0x7d, 0x5c, 0x4b, 0x91, 0xa0, 0xd4, 0x77, 0xde, 0x7d, 0x3a, 0xb9, 0x02,
	0x18, 0x0b, 0x3e, 0xa1, 0xd1, 0x54, 0x62, 0x68, 0x56, 0xea, 0xa9, 0x62, 0xb9, 0x7c, 0xf2, 0x16,
	0x1a, 0x22, 0x49, 0x2b, 0x87, 0x3e, 0xf2, 0x60, 0xc4, 0x30, 0x34, 0x2b, 0xfb, 0x54, 0xc9, 0x7a,
	0x26, 0x72, 0x66, 0x35, 0xc8, 0x39, 0x94, 0xe5, 0xc4, 0xd7, 0x1f, 0x7c, 0xc1, 0xcd, 0x2e, 0x78,
	0xaa, 0x5e, 0x51, 0x4e, 0x7e, 0xfb, 0x30, 0xe0, 0xe4, 0x0d, 0xd4, 0xa2, 0x44, 0xf9, 0xf7, 0xc5,
	0x2b, 0xad, 0x20, 0x56, 0x8d, 0x12, 0x75, 0x3a, 0xaf, 0xdf, 0x1b, 0xa8, 0x25, 0x3a, 0xc9, 0x49,
	0x96, 0x57, 0x91, 0x4c, 0x74, 0xb2, 0x20, 0x19, 0xc7, 0x98, 0x93, 0x74, 0x56, 0x91, 0x8c, 0x63,
	0xbc, 0x97, 0x7c, 0x69, 0x27, 0xce, 0x04, 0x8f, 0xa8, 0x9e, 0x86, 0xe8, 0x82, 0xdd, 0x75, 0x51,
	0xa2, 0xae, 0x66, 0x18, 0xf9, 0x0a, 0xaa, 0x86, 0x14, 0x68, 0xcb, 0xa9, 0x18, 0x4e, 0x25, 0xe5,
	0x64, 0x10, 0x79, 0x0e, 0xce, 0x44, 0xc5, 0xe6, 0xdf, 0x40, 0xb7, 0x6a, 0xe2, 0xe5, 0x89, 0x8a,
	0xd3, 0x6f, 0x23, 0xb9, 0x86, 0x9a, 0x5d, 0xa5, 0x10, 0x15, 0x4d, 0x77, 0x53, 0x6d, 0x05, 0xdf,
	0x95, 0x74, 0xa9, 0x7e, 0xb1, 0x02, 0x33, 0x47, 0x01, 0xcb, 0x1c, 0xd5, 0xe7, 0x8e, 0xba, 0x19,
	0x44, 0xbe, 0x80, 0xe2, 0x2d, 0xf2, 0x50, 0x48, 0xb7, 0x61, 0x82, 0xd9, 0x28, 0xfd, 0x85, 0x62,
	0x11, 0x22, 0xb3, 0xa7, 0xd8, 0xb6, 0xfd, 0x85, 0x0c, 0x62, 0x4e, 0xb0, 0x9e, 0xd9, 0x51, 0x76,
	0x1e, 0xcf, 0x56, 0xb0, 0x59, 0x92, 0x13, 0x3b, 0xe9, 0x16, 0x80, 0x7a, 0xef, 0xdf, 0xa2, 0x54,
	0x54, 0x70, 0x97, 0xd8, 0xef, 0xa8, 0xf7, 0x37, 0x16, 0x48, 0xed, 0x4d, 0x13, 0x4d, 0x63, 0x74,
	0x77, 0xac, 0x3d, 0x3b, 0x6a, 0x77, 0xa0, 0xbe, 0xa8, 0x48, 0x4a, 0xb0, 0x31, 0x38, 0x3f, 0xdf,
	0x5e, 0x23, 0x45, 0x58, 0x1f, 0xf4, 0xb7, 0x0b, 0xa4, 0x02, 0xa5, 0xb7, 0xfd, 0xcb, 0xfe, 0xe0,
	0xf7, 0xfe, 0xf6, 0xfa, 0xd1, 0xdf, 0x9b, 0x50, 0xb2, 0x7e, 0x42, 0x32, 0x80, 0x6a, 0xbe, 0xa9,
	0x90, 0x2f, 0x73, 0xce, 0x97, 0xf4, 0xc1, 0xe6, 0x8b, 0xcf, 0xc6, 0xb3, 0x6e, 0xd4, 0x85, 0xea,
	0xf0, 0x73, 0x82, 0x4b, 0x5a, 0x49, 0xf3, 0x59, 0x16, 0x37, 0x1d, 0xb9, 0x73, 0x23, 0x68, 0x48,
	0x7e, 0x02, 0xe7, 0x74, 0x76, 0x0c, 0x90, 0xbd, 0x5c, 0xfe, 0xe2, 0xa1, 0xbc, 0x2c, 0xf5, 0x07,
	0x28, 0x7a, 0x38, 0x12, 0x42, 0x3f, 0x31, 0xef, 0x3b, 0x70, 0x6c, 0x5e, 0x97, 0x31, 0xf2, 0x69,
	0x7c, 0x59, 0xca, 0xcf, 0xe0, 0xf4, 0x50, 0x67, 0x27, 0xf1, 0x92, 0x94, 0xe6, 0x02, 0x34, 0xb4,
	0x57, 0x90, 0x8c, 0x7e, 0x0a, 0xa4, 0x87, 0xfa, 0xe3, 0x66, 0xf1, 0x80, 0x48, 0x3a, 0x8b, 0x8f,
	0xe9, 0x97, 0xd0, 0xe8, 0xa1, 0x5e, 0x80, 0x1e, 0x98, 0xf4, 0xc3, 0x0d, 0xe7, 0xe4, 0xf9, 0x9f,
	0x7b, 0x26, 0x7e, 0x98, 0xde, 0xae, 0xc6, 0x4c, 0x4c, 0xc3, 0xc3, 0x48, 0x64, 0x17, 0xa4, 0x51,
	0xd1, 0x3c, 0x8f, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x48, 0xef, 0xc0, 0x7b, 0x09, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EnodebdClient is the client API for Enodebd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EnodebdClient interface {
	// Sends GetParameterValues message to ENodeB. TR-069 supports multiple
	// parameter names per message, but only one is supported here.
	GetParameter(ctx context.Context, in *GetParameterRequest, opts ...grpc.CallOption) (*GetParameterResponse, error)
	// Sends SetParameterValues message to ENodeB. TR-069 supports multiple
	// parameter names per message, but only one is supported here.
	SetParameter(ctx context.Context, in *SetParameterRequest, opts ...grpc.CallOption) (*protos.Void, error)
	// Configure eNodeB based on enodebd config file
	Configure(ctx context.Context, in *EnodebIdentity, opts ...grpc.CallOption) (*protos.Void, error)
	// Reboot eNodeB
	Reboot(ctx context.Context, in *EnodebIdentity, opts ...grpc.CallOption) (*protos.Void, error)
	// Reboot every connected eNodeB
	RebootAll(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*protos.Void, error)
	// Get current status
	GetStatus(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*protos.ServiceStatus, error)
	// Get status info for all connected eNodeB devices
	GetAllEnodebStatus(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*AllEnodebStatus, error)
	// Get status info of a single connected eNodeB device
	GetEnodebStatus(ctx context.Context, in *EnodebIdentity, opts ...grpc.CallOption) (*SingleEnodebStatus, error)
}

type enodebdClient struct {
	cc grpc.ClientConnInterface
}

func NewEnodebdClient(cc grpc.ClientConnInterface) EnodebdClient {
	return &enodebdClient{cc}
}

func (c *enodebdClient) GetParameter(ctx context.Context, in *GetParameterRequest, opts ...grpc.CallOption) (*GetParameterResponse, error) {
	out := new(GetParameterResponse)
	err := c.cc.Invoke(ctx, "/magma.lte.Enodebd/GetParameter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enodebdClient) SetParameter(ctx context.Context, in *SetParameterRequest, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.lte.Enodebd/SetParameter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enodebdClient) Configure(ctx context.Context, in *EnodebIdentity, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.lte.Enodebd/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enodebdClient) Reboot(ctx context.Context, in *EnodebIdentity, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.lte.Enodebd/Reboot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enodebdClient) RebootAll(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*protos.Void, error) {
	out := new(protos.Void)
	err := c.cc.Invoke(ctx, "/magma.lte.Enodebd/RebootAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enodebdClient) GetStatus(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*protos.ServiceStatus, error) {
	out := new(protos.ServiceStatus)
	err := c.cc.Invoke(ctx, "/magma.lte.Enodebd/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enodebdClient) GetAllEnodebStatus(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*AllEnodebStatus, error) {
	out := new(AllEnodebStatus)
	err := c.cc.Invoke(ctx, "/magma.lte.Enodebd/GetAllEnodebStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *enodebdClient) GetEnodebStatus(ctx context.Context, in *EnodebIdentity, opts ...grpc.CallOption) (*SingleEnodebStatus, error) {
	out := new(SingleEnodebStatus)
	err := c.cc.Invoke(ctx, "/magma.lte.Enodebd/GetEnodebStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnodebdServer is the server API for Enodebd service.
type EnodebdServer interface {
	// Sends GetParameterValues message to ENodeB. TR-069 supports multiple
	// parameter names per message, but only one is supported here.
	GetParameter(context.Context, *GetParameterRequest) (*GetParameterResponse, error)
	// Sends SetParameterValues message to ENodeB. TR-069 supports multiple
	// parameter names per message, but only one is supported here.
	SetParameter(context.Context, *SetParameterRequest) (*protos.Void, error)
	// Configure eNodeB based on enodebd config file
	Configure(context.Context, *EnodebIdentity) (*protos.Void, error)
	// Reboot eNodeB
	Reboot(context.Context, *EnodebIdentity) (*protos.Void, error)
	// Reboot every connected eNodeB
	RebootAll(context.Context, *protos.Void) (*protos.Void, error)
	// Get current status
	GetStatus(context.Context, *protos.Void) (*protos.ServiceStatus, error)
	// Get status info for all connected eNodeB devices
	GetAllEnodebStatus(context.Context, *protos.Void) (*AllEnodebStatus, error)
	// Get status info of a single connected eNodeB device
	GetEnodebStatus(context.Context, *EnodebIdentity) (*SingleEnodebStatus, error)
}

// UnimplementedEnodebdServer can be embedded to have forward compatible implementations.
type UnimplementedEnodebdServer struct {
}

func (*UnimplementedEnodebdServer) GetParameter(ctx context.Context, req *GetParameterRequest) (*GetParameterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParameter not implemented")
}
func (*UnimplementedEnodebdServer) SetParameter(ctx context.Context, req *SetParameterRequest) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetParameter not implemented")
}
func (*UnimplementedEnodebdServer) Configure(ctx context.Context, req *EnodebIdentity) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (*UnimplementedEnodebdServer) Reboot(ctx context.Context, req *EnodebIdentity) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reboot not implemented")
}
func (*UnimplementedEnodebdServer) RebootAll(ctx context.Context, req *protos.Void) (*protos.Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RebootAll not implemented")
}
func (*UnimplementedEnodebdServer) GetStatus(ctx context.Context, req *protos.Void) (*protos.ServiceStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (*UnimplementedEnodebdServer) GetAllEnodebStatus(ctx context.Context, req *protos.Void) (*AllEnodebStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllEnodebStatus not implemented")
}
func (*UnimplementedEnodebdServer) GetEnodebStatus(ctx context.Context, req *EnodebIdentity) (*SingleEnodebStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEnodebStatus not implemented")
}

func RegisterEnodebdServer(s *grpc.Server, srv EnodebdServer) {
	s.RegisterService(&_Enodebd_serviceDesc, srv)
}

func _Enodebd_GetParameter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetParameterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnodebdServer).GetParameter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Enodebd/GetParameter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnodebdServer).GetParameter(ctx, req.(*GetParameterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Enodebd_SetParameter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetParameterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnodebdServer).SetParameter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Enodebd/SetParameter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnodebdServer).SetParameter(ctx, req.(*SetParameterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Enodebd_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnodebIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnodebdServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Enodebd/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnodebdServer).Configure(ctx, req.(*EnodebIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Enodebd_Reboot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnodebIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnodebdServer).Reboot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Enodebd/Reboot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnodebdServer).Reboot(ctx, req.(*EnodebIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Enodebd_RebootAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnodebdServer).RebootAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Enodebd/RebootAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnodebdServer).RebootAll(ctx, req.(*protos.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Enodebd_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnodebdServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Enodebd/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnodebdServer).GetStatus(ctx, req.(*protos.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Enodebd_GetAllEnodebStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnodebdServer).GetAllEnodebStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Enodebd/GetAllEnodebStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnodebdServer).GetAllEnodebStatus(ctx, req.(*protos.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Enodebd_GetEnodebStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnodebIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnodebdServer).GetEnodebStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.lte.Enodebd/GetEnodebStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnodebdServer).GetEnodebStatus(ctx, req.(*EnodebIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

var _Enodebd_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.lte.Enodebd",
	HandlerType: (*EnodebdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetParameter",
			Handler:    _Enodebd_GetParameter_Handler,
		},
		{
			MethodName: "SetParameter",
			Handler:    _Enodebd_SetParameter_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _Enodebd_Configure_Handler,
		},
		{
			MethodName: "Reboot",
			Handler:    _Enodebd_Reboot_Handler,
		},
		{
			MethodName: "RebootAll",
			Handler:    _Enodebd_RebootAll_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _Enodebd_GetStatus_Handler,
		},
		{
			MethodName: "GetAllEnodebStatus",
			Handler:    _Enodebd_GetAllEnodebStatus_Handler,
		},
		{
			MethodName: "GetEnodebStatus",
			Handler:    _Enodebd_GetEnodebStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lte/protos/enodebd.proto",
}
