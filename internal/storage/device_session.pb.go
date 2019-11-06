// Code generated by protoc-gen-go. DO NOT EDIT.
// source: device_session.proto

package storage

import (
	fmt "fmt"
	common "github.com/brocaar/chirpstack-api/go/common"
	proto "github.com/golang/protobuf/proto"
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

type DeviceSessionPBChannel struct {
	// Frequency (Hz).
	Frequency uint32 `protobuf:"varint,1,opt,name=frequency,proto3" json:"frequency,omitempty"`
	// Min. data-rate.
	MinDr uint32 `protobuf:"varint,2,opt,name=min_dr,json=minDr,proto3" json:"min_dr,omitempty"`
	// Max. data-rate.
	MaxDr                uint32   `protobuf:"varint,3,opt,name=max_dr,json=maxDr,proto3" json:"max_dr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceSessionPBChannel) Reset()         { *m = DeviceSessionPBChannel{} }
func (m *DeviceSessionPBChannel) String() string { return proto.CompactTextString(m) }
func (*DeviceSessionPBChannel) ProtoMessage()    {}
func (*DeviceSessionPBChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_958563bbc6ebadf7, []int{0}
}

func (m *DeviceSessionPBChannel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceSessionPBChannel.Unmarshal(m, b)
}
func (m *DeviceSessionPBChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceSessionPBChannel.Marshal(b, m, deterministic)
}
func (m *DeviceSessionPBChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceSessionPBChannel.Merge(m, src)
}
func (m *DeviceSessionPBChannel) XXX_Size() int {
	return xxx_messageInfo_DeviceSessionPBChannel.Size(m)
}
func (m *DeviceSessionPBChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceSessionPBChannel.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceSessionPBChannel proto.InternalMessageInfo

func (m *DeviceSessionPBChannel) GetFrequency() uint32 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *DeviceSessionPBChannel) GetMinDr() uint32 {
	if m != nil {
		return m.MinDr
	}
	return 0
}

func (m *DeviceSessionPBChannel) GetMaxDr() uint32 {
	if m != nil {
		return m.MaxDr
	}
	return 0
}

type DeviceSessionPBUplinkADRHistory struct {
	// Uplink frame-counter.
	FCnt uint32 `protobuf:"varint,1,opt,name=f_cnt,json=fCnt,proto3" json:"f_cnt,omitempty"`
	// Max SNR (of deduplicated frames received by one or multiple gateways).
	MaxSnr float32 `protobuf:"fixed32,2,opt,name=max_snr,json=maxSnr,proto3" json:"max_snr,omitempty"`
	// TX Power (as known by the network-server).
	TxPowerIndex uint32 `protobuf:"varint,3,opt,name=tx_power_index,json=txPowerIndex,proto3" json:"tx_power_index,omitempty"`
	// Number of receiving gateways.
	GatewayCount         uint32   `protobuf:"varint,4,opt,name=gateway_count,json=gatewayCount,proto3" json:"gateway_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceSessionPBUplinkADRHistory) Reset()         { *m = DeviceSessionPBUplinkADRHistory{} }
func (m *DeviceSessionPBUplinkADRHistory) String() string { return proto.CompactTextString(m) }
func (*DeviceSessionPBUplinkADRHistory) ProtoMessage()    {}
func (*DeviceSessionPBUplinkADRHistory) Descriptor() ([]byte, []int) {
	return fileDescriptor_958563bbc6ebadf7, []int{1}
}

func (m *DeviceSessionPBUplinkADRHistory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceSessionPBUplinkADRHistory.Unmarshal(m, b)
}
func (m *DeviceSessionPBUplinkADRHistory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceSessionPBUplinkADRHistory.Marshal(b, m, deterministic)
}
func (m *DeviceSessionPBUplinkADRHistory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceSessionPBUplinkADRHistory.Merge(m, src)
}
func (m *DeviceSessionPBUplinkADRHistory) XXX_Size() int {
	return xxx_messageInfo_DeviceSessionPBUplinkADRHistory.Size(m)
}
func (m *DeviceSessionPBUplinkADRHistory) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceSessionPBUplinkADRHistory.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceSessionPBUplinkADRHistory proto.InternalMessageInfo

func (m *DeviceSessionPBUplinkADRHistory) GetFCnt() uint32 {
	if m != nil {
		return m.FCnt
	}
	return 0
}

func (m *DeviceSessionPBUplinkADRHistory) GetMaxSnr() float32 {
	if m != nil {
		return m.MaxSnr
	}
	return 0
}

func (m *DeviceSessionPBUplinkADRHistory) GetTxPowerIndex() uint32 {
	if m != nil {
		return m.TxPowerIndex
	}
	return 0
}

func (m *DeviceSessionPBUplinkADRHistory) GetGatewayCount() uint32 {
	if m != nil {
		return m.GatewayCount
	}
	return 0
}

type DeviceSessionPB struct {
	// ID of the device-profile.
	DeviceProfileId string `protobuf:"bytes,1,opt,name=device_profile_id,json=deviceProfileId,proto3" json:"device_profile_id,omitempty"`
	// ID of the service-profile.
	ServiceProfileId string `protobuf:"bytes,2,opt,name=service_profile_id,json=serviceProfileId,proto3" json:"service_profile_id,omitempty"`
	// ID of the routing-profile.
	RoutingProfileId string `protobuf:"bytes,3,opt,name=routing_profile_id,json=routingProfileId,proto3" json:"routing_profile_id,omitempty"`
	// Device address.
	DevAddr []byte `protobuf:"bytes,4,opt,name=dev_addr,json=devAddr,proto3" json:"dev_addr,omitempty"`
	// Device EUI.
	DevEui []byte `protobuf:"bytes,5,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	// Join EUI.
	JoinEui []byte `protobuf:"bytes,6,opt,name=join_eui,json=joinEui,proto3" json:"join_eui,omitempty"`
	// FNwkSIntKey.
	FNwkSIntKey []byte `protobuf:"bytes,7,opt,name=f_nwk_s_int_key,json=fNwkSIntKey,proto3" json:"f_nwk_s_int_key,omitempty"`
	// SNwkSIntKey.
	SNwkSIntKey []byte `protobuf:"bytes,8,opt,name=s_nwk_s_int_key,json=sNwkSIntKey,proto3" json:"s_nwk_s_int_key,omitempty"`
	// NwkSEncKey.
	NwkSEncKey []byte `protobuf:"bytes,9,opt,name=nwk_s_enc_key,json=nwkSEncKey,proto3" json:"nwk_s_enc_key,omitempty"`
	// AppSKeyEnvelope contains the (encrypted) AppSKey key-envelope.
	AppSKeyEnvelope *common.KeyEnvelope `protobuf:"bytes,45,opt,name=app_s_key_envelope,json=appSKeyEnvelope,proto3" json:"app_s_key_envelope,omitempty"`
	// Uplink frame-counter.
	FCntUp uint32 `protobuf:"varint,10,opt,name=f_cnt_up,json=fCntUp,proto3" json:"f_cnt_up,omitempty"`
	// Downlink frame-counter (network-server).
	NFCntDown uint32 `protobuf:"varint,11,opt,name=n_f_cnt_down,json=nFCntDown,proto3" json:"n_f_cnt_down,omitempty"`
	// Uplink frame-counter (application-server).
	// Note: this frame-counter is managed by the application-server.
	AFCntDown uint32 `protobuf:"varint,12,opt,name=a_f_cnt_down,json=aFCntDown,proto3" json:"a_f_cnt_down,omitempty"`
	// Frame-counter holding the last confirmed downlink frame (n_f_cnt_down or a_f_cnt_down).
	ConfFCnt uint32 `protobuf:"varint,39,opt,name=conf_f_cnt,json=confFCnt,proto3" json:"conf_f_cnt,omitempty"`
	// Skip uplink frame-counter validation.
	SkipFCntCheck bool `protobuf:"varint,13,opt,name=skip_f_cnt_check,json=skipFCntCheck,proto3" json:"skip_f_cnt_check,omitempty"`
	// RX Delay.
	RxDelay uint32 `protobuf:"varint,14,opt,name=rx_delay,json=rxDelay,proto3" json:"rx_delay,omitempty"`
	// RX1 data-rate offset.
	Rx1DrOffset uint32 `protobuf:"varint,15,opt,name=rx1_dr_offset,json=rx1DrOffset,proto3" json:"rx1_dr_offset,omitempty"`
	// RX2 data-rate.
	Rx2Dr uint32 `protobuf:"varint,16,opt,name=rx2_dr,json=rx2Dr,proto3" json:"rx2_dr,omitempty"`
	// RX2 frequency.
	Rx2Frequency uint32 `protobuf:"varint,17,opt,name=rx2_frequency,json=rx2Frequency,proto3" json:"rx2_frequency,omitempty"`
	// TXPowerIndex which the node is using. The possible values are defined
	// by the lorawan/band package and are region specific. By default it is
	// assumed that the node is using TXPower 0. This value is controlled by
	// the ADR engine.
	TxPowerIndex uint32 `protobuf:"varint,18,opt,name=tx_power_index,json=txPowerIndex,proto3" json:"tx_power_index,omitempty"`
	// DR defines the (last known) data-rate at which the node is operating.
	// This value is controlled by the ADR engine.
	Dr uint32 `protobuf:"varint,19,opt,name=dr,proto3" json:"dr,omitempty"`
	// ADR defines if the device has ADR enabled.
	Adr bool `protobuf:"varint,20,opt,name=adr,proto3" json:"adr,omitempty"`
	// MaxSupportedTXPowerIndex defines the maximum supported tx-power index
	// by the node, or 0 when not set.
	MaxSupportedTxPowerIndex uint32 `protobuf:"varint,21,opt,name=max_supported_tx_power_index,json=maxSupportedTxPowerIndex,proto3" json:"max_supported_tx_power_index,omitempty"`
	// NbTrans defines the number of transmissions for each unconfirmed uplink
	// frame. In case of 0, the default value is used.
	// This value is controlled by the ADR engine.
	NbTrans uint32 `protobuf:"varint,23,opt,name=nb_trans,json=nbTrans,proto3" json:"nb_trans,omitempty"`
	// Channels that are activated on the device.
	EnabledUplinkChannels []uint32 `protobuf:"varint,24,rep,packed,name=enabled_uplink_channels,json=enabledUplinkChannels,proto3" json:"enabled_uplink_channels,omitempty"`
	// Extra uplink channels, configured by the user.
	ExtraUplinkChannels map[uint32]*DeviceSessionPBChannel `protobuf:"bytes,25,rep,name=extra_uplink_channels,json=extraUplinkChannels,proto3" json:"extra_uplink_channels,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Frequency of each channel.
	ChannelFrequencies []uint32 `protobuf:"varint,26,rep,packed,name=channel_frequencies,json=channelFrequencies,proto3" json:"channel_frequencies,omitempty"`
	// Uplink history for ADR (last 20 uplink transmissions).
	UplinkAdrHistory []*DeviceSessionPBUplinkADRHistory `protobuf:"bytes,27,rep,name=uplink_adr_history,json=uplinkAdrHistory,proto3" json:"uplink_adr_history,omitempty"`
	// Last device-status requested timestamp (Unix ns)
	LastDeviceStatusRequestTimeUnixNs int64 `protobuf:"varint,29,opt,name=last_device_status_request_time_unix_ns,json=lastDeviceStatusRequestTimeUnixNs,proto3" json:"last_device_status_request_time_unix_ns,omitempty"`
	// Last downlink timestamp (Unix ns).
	LastDownlinkTxTimestampUnixNs int64 `protobuf:"varint,32,opt,name=last_downlink_tx_timestamp_unix_ns,json=lastDownlinkTxTimestampUnixNs,proto3" json:"last_downlink_tx_timestamp_unix_ns,omitempty"`
	// Class-B beacon is locked.
	BeaconLocked bool `protobuf:"varint,33,opt,name=beacon_locked,json=beaconLocked,proto3" json:"beacon_locked,omitempty"`
	// Class-B ping-slot nb.
	PingSlotNb uint32 `protobuf:"varint,34,opt,name=ping_slot_nb,json=pingSlotNb,proto3" json:"ping_slot_nb,omitempty"`
	// Class-B ping-slot data-rate.
	PingSlotDr uint32 `protobuf:"varint,35,opt,name=ping_slot_dr,json=pingSlotDr,proto3" json:"ping_slot_dr,omitempty"`
	// Class-B ping-slot tx frequency.
	PingSlotFrequency uint32 `protobuf:"varint,36,opt,name=ping_slot_frequency,json=pingSlotFrequency,proto3" json:"ping_slot_frequency,omitempty"`
	// LoRaWAN mac-version.
	MacVersion string `protobuf:"bytes,37,opt,name=mac_version,json=macVersion,proto3" json:"mac_version,omitempty"`
	// MinSupportedTXPowerIndex defines the minimum supported tx-power index
	// by the node (default 0).
	MinSupportedTxPowerIndex uint32 `protobuf:"varint,38,opt,name=min_supported_tx_power_index,json=minSupportedTxPowerIndex,proto3" json:"min_supported_tx_power_index,omitempty"`
	// RejoinRequestEnabled defines if the rejoin-request is enabled on the
	// device.
	RejoinRequestEnabled bool `protobuf:"varint,44,opt,name=rejoin_request_enabled,json=rejoinRequestEnabled,proto3" json:"rejoin_request_enabled,omitempty"`
	// RejoinRequestMaxCountN defines the 2^(C+4) uplink message interval for
	// the rejoin-request.
	RejoinRequestMaxCountN uint32 `protobuf:"varint,40,opt,name=rejoin_request_max_count_n,json=rejoinRequestMaxCountN,proto3" json:"rejoin_request_max_count_n,omitempty"`
	// RejoinRequestMaxTimeN defines the 2^(T+10) time interval (seconds)
	// for the rejoin-request.
	RejoinRequestMaxTimeN uint32 `protobuf:"varint,41,opt,name=rejoin_request_max_time_n,json=rejoinRequestMaxTimeN,proto3" json:"rejoin_request_max_time_n,omitempty"`
	// Rejoin counter (RJCount0).
	// This counter is reset to 0 after each successful join-accept.
	RejoinCount_0 uint32 `protobuf:"varint,42,opt,name=rejoin_count_0,json=rejoinCount0,proto3" json:"rejoin_count_0,omitempty"`
	// Pending rejoin device-session contains a device-session which has not
	// yet been activated by the device (by sending a first uplink).
	PendingRejoinDeviceSession []byte `protobuf:"bytes,43,opt,name=pending_rejoin_device_session,json=pendingRejoinDeviceSession,proto3" json:"pending_rejoin_device_session,omitempty"`
	// Device reference altitude for geolocation.
	ReferenceAltitude float64 `protobuf:"fixed64,46,opt,name=reference_altitude,json=referenceAltitude,proto3" json:"reference_altitude,omitempty"`
	// UplinkDwellTime.
	UplinkDwellTime_400Ms bool `protobuf:"varint,47,opt,name=uplink_dwell_time_400ms,json=uplinkDwellTime400ms,proto3" json:"uplink_dwell_time_400ms,omitempty"`
	// DownlinkDwellTime.
	DownlinkDwellTime_400Ms bool `protobuf:"varint,48,opt,name=downlink_dwell_time_400ms,json=downlinkDwellTime400ms,proto3" json:"downlink_dwell_time_400ms,omitempty"`
	// Uplink max. EIRP index.
	UplinkMaxEirpIndex   uint32   `protobuf:"varint,49,opt,name=uplink_max_eirp_index,json=uplinkMaxEirpIndex,proto3" json:"uplink_max_eirp_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceSessionPB) Reset()         { *m = DeviceSessionPB{} }
func (m *DeviceSessionPB) String() string { return proto.CompactTextString(m) }
func (*DeviceSessionPB) ProtoMessage()    {}
func (*DeviceSessionPB) Descriptor() ([]byte, []int) {
	return fileDescriptor_958563bbc6ebadf7, []int{2}
}

func (m *DeviceSessionPB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceSessionPB.Unmarshal(m, b)
}
func (m *DeviceSessionPB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceSessionPB.Marshal(b, m, deterministic)
}
func (m *DeviceSessionPB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceSessionPB.Merge(m, src)
}
func (m *DeviceSessionPB) XXX_Size() int {
	return xxx_messageInfo_DeviceSessionPB.Size(m)
}
func (m *DeviceSessionPB) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceSessionPB.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceSessionPB proto.InternalMessageInfo

func (m *DeviceSessionPB) GetDeviceProfileId() string {
	if m != nil {
		return m.DeviceProfileId
	}
	return ""
}

func (m *DeviceSessionPB) GetServiceProfileId() string {
	if m != nil {
		return m.ServiceProfileId
	}
	return ""
}

func (m *DeviceSessionPB) GetRoutingProfileId() string {
	if m != nil {
		return m.RoutingProfileId
	}
	return ""
}

func (m *DeviceSessionPB) GetDevAddr() []byte {
	if m != nil {
		return m.DevAddr
	}
	return nil
}

func (m *DeviceSessionPB) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *DeviceSessionPB) GetJoinEui() []byte {
	if m != nil {
		return m.JoinEui
	}
	return nil
}

func (m *DeviceSessionPB) GetFNwkSIntKey() []byte {
	if m != nil {
		return m.FNwkSIntKey
	}
	return nil
}

func (m *DeviceSessionPB) GetSNwkSIntKey() []byte {
	if m != nil {
		return m.SNwkSIntKey
	}
	return nil
}

func (m *DeviceSessionPB) GetNwkSEncKey() []byte {
	if m != nil {
		return m.NwkSEncKey
	}
	return nil
}

func (m *DeviceSessionPB) GetAppSKeyEnvelope() *common.KeyEnvelope {
	if m != nil {
		return m.AppSKeyEnvelope
	}
	return nil
}

func (m *DeviceSessionPB) GetFCntUp() uint32 {
	if m != nil {
		return m.FCntUp
	}
	return 0
}

func (m *DeviceSessionPB) GetNFCntDown() uint32 {
	if m != nil {
		return m.NFCntDown
	}
	return 0
}

func (m *DeviceSessionPB) GetAFCntDown() uint32 {
	if m != nil {
		return m.AFCntDown
	}
	return 0
}

func (m *DeviceSessionPB) GetConfFCnt() uint32 {
	if m != nil {
		return m.ConfFCnt
	}
	return 0
}

func (m *DeviceSessionPB) GetSkipFCntCheck() bool {
	if m != nil {
		return m.SkipFCntCheck
	}
	return false
}

func (m *DeviceSessionPB) GetRxDelay() uint32 {
	if m != nil {
		return m.RxDelay
	}
	return 0
}

func (m *DeviceSessionPB) GetRx1DrOffset() uint32 {
	if m != nil {
		return m.Rx1DrOffset
	}
	return 0
}

func (m *DeviceSessionPB) GetRx2Dr() uint32 {
	if m != nil {
		return m.Rx2Dr
	}
	return 0
}

func (m *DeviceSessionPB) GetRx2Frequency() uint32 {
	if m != nil {
		return m.Rx2Frequency
	}
	return 0
}

func (m *DeviceSessionPB) GetTxPowerIndex() uint32 {
	if m != nil {
		return m.TxPowerIndex
	}
	return 0
}

func (m *DeviceSessionPB) GetDr() uint32 {
	if m != nil {
		return m.Dr
	}
	return 0
}

func (m *DeviceSessionPB) GetAdr() bool {
	if m != nil {
		return m.Adr
	}
	return false
}

func (m *DeviceSessionPB) GetMaxSupportedTxPowerIndex() uint32 {
	if m != nil {
		return m.MaxSupportedTxPowerIndex
	}
	return 0
}

func (m *DeviceSessionPB) GetNbTrans() uint32 {
	if m != nil {
		return m.NbTrans
	}
	return 0
}

func (m *DeviceSessionPB) GetEnabledUplinkChannels() []uint32 {
	if m != nil {
		return m.EnabledUplinkChannels
	}
	return nil
}

func (m *DeviceSessionPB) GetExtraUplinkChannels() map[uint32]*DeviceSessionPBChannel {
	if m != nil {
		return m.ExtraUplinkChannels
	}
	return nil
}

func (m *DeviceSessionPB) GetChannelFrequencies() []uint32 {
	if m != nil {
		return m.ChannelFrequencies
	}
	return nil
}

func (m *DeviceSessionPB) GetUplinkAdrHistory() []*DeviceSessionPBUplinkADRHistory {
	if m != nil {
		return m.UplinkAdrHistory
	}
	return nil
}

func (m *DeviceSessionPB) GetLastDeviceStatusRequestTimeUnixNs() int64 {
	if m != nil {
		return m.LastDeviceStatusRequestTimeUnixNs
	}
	return 0
}

func (m *DeviceSessionPB) GetLastDownlinkTxTimestampUnixNs() int64 {
	if m != nil {
		return m.LastDownlinkTxTimestampUnixNs
	}
	return 0
}

func (m *DeviceSessionPB) GetBeaconLocked() bool {
	if m != nil {
		return m.BeaconLocked
	}
	return false
}

func (m *DeviceSessionPB) GetPingSlotNb() uint32 {
	if m != nil {
		return m.PingSlotNb
	}
	return 0
}

func (m *DeviceSessionPB) GetPingSlotDr() uint32 {
	if m != nil {
		return m.PingSlotDr
	}
	return 0
}

func (m *DeviceSessionPB) GetPingSlotFrequency() uint32 {
	if m != nil {
		return m.PingSlotFrequency
	}
	return 0
}

func (m *DeviceSessionPB) GetMacVersion() string {
	if m != nil {
		return m.MacVersion
	}
	return ""
}

func (m *DeviceSessionPB) GetMinSupportedTxPowerIndex() uint32 {
	if m != nil {
		return m.MinSupportedTxPowerIndex
	}
	return 0
}

func (m *DeviceSessionPB) GetRejoinRequestEnabled() bool {
	if m != nil {
		return m.RejoinRequestEnabled
	}
	return false
}

func (m *DeviceSessionPB) GetRejoinRequestMaxCountN() uint32 {
	if m != nil {
		return m.RejoinRequestMaxCountN
	}
	return 0
}

func (m *DeviceSessionPB) GetRejoinRequestMaxTimeN() uint32 {
	if m != nil {
		return m.RejoinRequestMaxTimeN
	}
	return 0
}

func (m *DeviceSessionPB) GetRejoinCount_0() uint32 {
	if m != nil {
		return m.RejoinCount_0
	}
	return 0
}

func (m *DeviceSessionPB) GetPendingRejoinDeviceSession() []byte {
	if m != nil {
		return m.PendingRejoinDeviceSession
	}
	return nil
}

func (m *DeviceSessionPB) GetReferenceAltitude() float64 {
	if m != nil {
		return m.ReferenceAltitude
	}
	return 0
}

func (m *DeviceSessionPB) GetUplinkDwellTime_400Ms() bool {
	if m != nil {
		return m.UplinkDwellTime_400Ms
	}
	return false
}

func (m *DeviceSessionPB) GetDownlinkDwellTime_400Ms() bool {
	if m != nil {
		return m.DownlinkDwellTime_400Ms
	}
	return false
}

func (m *DeviceSessionPB) GetUplinkMaxEirpIndex() uint32 {
	if m != nil {
		return m.UplinkMaxEirpIndex
	}
	return 0
}

type DeviceGatewayRXInfoSetPB struct {
	// Device EUI.
	DevEui []byte `protobuf:"bytes,1,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	// Data-rate.
	Dr uint32 `protobuf:"varint,2,opt,name=dr,proto3" json:"dr,omitempty"`
	// Items contains set items.
	Items                []*DeviceGatewayRXInfoPB `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *DeviceGatewayRXInfoSetPB) Reset()         { *m = DeviceGatewayRXInfoSetPB{} }
func (m *DeviceGatewayRXInfoSetPB) String() string { return proto.CompactTextString(m) }
func (*DeviceGatewayRXInfoSetPB) ProtoMessage()    {}
func (*DeviceGatewayRXInfoSetPB) Descriptor() ([]byte, []int) {
	return fileDescriptor_958563bbc6ebadf7, []int{3}
}

func (m *DeviceGatewayRXInfoSetPB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceGatewayRXInfoSetPB.Unmarshal(m, b)
}
func (m *DeviceGatewayRXInfoSetPB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceGatewayRXInfoSetPB.Marshal(b, m, deterministic)
}
func (m *DeviceGatewayRXInfoSetPB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceGatewayRXInfoSetPB.Merge(m, src)
}
func (m *DeviceGatewayRXInfoSetPB) XXX_Size() int {
	return xxx_messageInfo_DeviceGatewayRXInfoSetPB.Size(m)
}
func (m *DeviceGatewayRXInfoSetPB) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceGatewayRXInfoSetPB.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceGatewayRXInfoSetPB proto.InternalMessageInfo

func (m *DeviceGatewayRXInfoSetPB) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *DeviceGatewayRXInfoSetPB) GetDr() uint32 {
	if m != nil {
		return m.Dr
	}
	return 0
}

func (m *DeviceGatewayRXInfoSetPB) GetItems() []*DeviceGatewayRXInfoPB {
	if m != nil {
		return m.Items
	}
	return nil
}

type DeviceGatewayRXInfoPB struct {
	// Gateway ID.
	GatewayId []byte `protobuf:"bytes,1,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
	// RSSI.
	Rssi int32 `protobuf:"varint,2,opt,name=rssi,proto3" json:"rssi,omitempty"`
	// LoRa SNR.
	LoraSnr float64 `protobuf:"fixed64,3,opt,name=lora_snr,json=loraSnr,proto3" json:"lora_snr,omitempty"`
	// Board.
	Board uint32 `protobuf:"varint,4,opt,name=board,proto3" json:"board,omitempty"`
	// Antenna.
	Antenna uint32 `protobuf:"varint,5,opt,name=antenna,proto3" json:"antenna,omitempty"`
	// Gateway specific context.
	Context              []byte   `protobuf:"bytes,6,opt,name=context,proto3" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceGatewayRXInfoPB) Reset()         { *m = DeviceGatewayRXInfoPB{} }
func (m *DeviceGatewayRXInfoPB) String() string { return proto.CompactTextString(m) }
func (*DeviceGatewayRXInfoPB) ProtoMessage()    {}
func (*DeviceGatewayRXInfoPB) Descriptor() ([]byte, []int) {
	return fileDescriptor_958563bbc6ebadf7, []int{4}
}

func (m *DeviceGatewayRXInfoPB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceGatewayRXInfoPB.Unmarshal(m, b)
}
func (m *DeviceGatewayRXInfoPB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceGatewayRXInfoPB.Marshal(b, m, deterministic)
}
func (m *DeviceGatewayRXInfoPB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceGatewayRXInfoPB.Merge(m, src)
}
func (m *DeviceGatewayRXInfoPB) XXX_Size() int {
	return xxx_messageInfo_DeviceGatewayRXInfoPB.Size(m)
}
func (m *DeviceGatewayRXInfoPB) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceGatewayRXInfoPB.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceGatewayRXInfoPB proto.InternalMessageInfo

func (m *DeviceGatewayRXInfoPB) GetGatewayId() []byte {
	if m != nil {
		return m.GatewayId
	}
	return nil
}

func (m *DeviceGatewayRXInfoPB) GetRssi() int32 {
	if m != nil {
		return m.Rssi
	}
	return 0
}

func (m *DeviceGatewayRXInfoPB) GetLoraSnr() float64 {
	if m != nil {
		return m.LoraSnr
	}
	return 0
}

func (m *DeviceGatewayRXInfoPB) GetBoard() uint32 {
	if m != nil {
		return m.Board
	}
	return 0
}

func (m *DeviceGatewayRXInfoPB) GetAntenna() uint32 {
	if m != nil {
		return m.Antenna
	}
	return 0
}

func (m *DeviceGatewayRXInfoPB) GetContext() []byte {
	if m != nil {
		return m.Context
	}
	return nil
}

func init() {
	proto.RegisterType((*DeviceSessionPBChannel)(nil), "storage.DeviceSessionPBChannel")
	proto.RegisterType((*DeviceSessionPBUplinkADRHistory)(nil), "storage.DeviceSessionPBUplinkADRHistory")
	proto.RegisterType((*DeviceSessionPB)(nil), "storage.DeviceSessionPB")
	proto.RegisterMapType((map[uint32]*DeviceSessionPBChannel)(nil), "storage.DeviceSessionPB.ExtraUplinkChannelsEntry")
	proto.RegisterType((*DeviceGatewayRXInfoSetPB)(nil), "storage.DeviceGatewayRXInfoSetPB")
	proto.RegisterType((*DeviceGatewayRXInfoPB)(nil), "storage.DeviceGatewayRXInfoPB")
}

func init() { proto.RegisterFile("device_session.proto", fileDescriptor_958563bbc6ebadf7) }

var fileDescriptor_958563bbc6ebadf7 = []byte{
	// 1326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x56, 0xed, 0x56, 0x1b, 0x37,
	0x13, 0x3e, 0xc6, 0xe1, 0x6b, 0xc0, 0x01, 0xc4, 0x97, 0xe0, 0x0d, 0x2f, 0x8e, 0x93, 0xf7, 0x8d,
	0x9b, 0x26, 0x04, 0x68, 0xd2, 0x93, 0xe6, 0x47, 0x4f, 0x09, 0x26, 0x2d, 0x27, 0x0d, 0xe5, 0x2c,
	0x24, 0xa7, 0xff, 0x74, 0xe4, 0x95, 0x4c, 0x54, 0xaf, 0xb5, 0x5b, 0xad, 0x8c, 0xd7, 0xb7, 0xd2,
	0x9b, 0xe8, 0x95, 0xf5, 0x1e, 0x7a, 0x34, 0x92, 0x6d, 0x6c, 0xe0, 0x97, 0xad, 0xe7, 0x79, 0x66,
	0xa4, 0x1d, 0xcd, 0x8c, 0x06, 0xd6, 0x84, 0xbc, 0x56, 0xb1, 0x64, 0xb9, 0xcc, 0x73, 0x95, 0xea,
	0xbd, 0xcc, 0xa4, 0x36, 0x25, 0xb3, 0xb9, 0x4d, 0x0d, 0xbf, 0x92, 0xdb, 0x9b, 0x3c, 0x53, 0xaf,
	0xe2, 0xb4, 0xd3, 0x49, 0x75, 0xf8, 0xf1, 0x8a, 0x9a, 0x80, 0x8d, 0x06, 0x5a, 0x5e, 0x78, 0xc3,
	0xf3, 0xf7, 0xc7, 0x5f, 0xb9, 0xd6, 0x32, 0x21, 0x8f, 0x60, 0xbe, 0x65, 0xe4, 0x9f, 0x5d, 0xa9,
	0xe3, 0x3e, 0x2d, 0x55, 0x4b, 0xf5, 0x4a, 0x34, 0x02, 0xc8, 0x3a, 0xcc, 0x74, 0x94, 0x66, 0xc2,
	0xd0, 0x29, 0xa4, 0xa6, 0x3b, 0x4a, 0x37, 0x0c, 0xc2, 0xbc, 0x70, 0x70, 0x39, 0xc0, 0xbc, 0x68,
	0x98, 0xda, 0x5f, 0x25, 0xd8, 0x9d, 0xd8, 0xe6, 0x73, 0x96, 0x28, 0xdd, 0x3e, 0x6a, 0x44, 0xbf,
	0x28, 0x77, 0xc8, 0x3e, 0x59, 0x85, 0xe9, 0x16, 0x8b, 0xb5, 0x0d, 0x7b, 0x3d, 0x68, 0x1d, 0x6b,
	0x4b, 0x36, 0x61, 0xd6, 0xf9, 0xcb, 0xb5, 0xdf, 0x67, 0x2a, 0x72, 0xee, 0x2f, 0xb4, 0x21, 0x4f,
	0xe1, 0xa1, 0x2d, 0x58, 0x96, 0xf6, 0xa4, 0x61, 0x4a, 0x0b, 0x59, 0x84, 0x0d, 0x17, 0x6d, 0x71,
	0xee, 0xc0, 0x53, 0x87, 0x91, 0x27, 0x50, 0xb9, 0xe2, 0x56, 0xf6, 0x78, 0x9f, 0xc5, 0x69, 0x57,
	0x5b, 0xfa, 0xc0, 0x8b, 0x02, 0x78, 0xec, 0xb0, 0xda, 0x3f, 0xcb, 0xb0, 0x34, 0x71, 0x38, 0xf2,
	0x1c, 0x56, 0x42, 0x40, 0x33, 0x93, 0xb6, 0x54, 0x22, 0x99, 0x12, 0x78, 0xb0, 0xf9, 0x68, 0xc9,
	0x13, 0xe7, 0x1e, 0x3f, 0x15, 0xe4, 0x05, 0x90, 0x5c, 0x9a, 0x49, 0xf1, 0x14, 0x8a, 0x97, 0x03,
	0x33, 0xa6, 0x36, 0x69, 0xd7, 0x2a, 0x7d, 0x75, 0x53, 0x5d, 0xf6, 0xea, 0xc0, 0x8c, 0xd4, 0x5b,
	0x30, 0x27, 0xe4, 0x35, 0xe3, 0x42, 0x18, 0x3c, 0xfb, 0x62, 0x34, 0x2b, 0xe4, 0xf5, 0x91, 0x10,
	0xc6, 0x85, 0xc6, 0x51, 0xb2, 0xab, 0xe8, 0x34, 0x32, 0x33, 0x42, 0x5e, 0x9f, 0x74, 0x95, 0xb3,
	0xf9, 0x23, 0x55, 0x1a, 0x99, 0x19, 0x6f, 0xe3, 0xd6, 0x8e, 0x7a, 0x0a, 0x4b, 0x2d, 0xa6, 0x7b,
	0x6d, 0x96, 0x33, 0xa5, 0x2d, 0x6b, 0xcb, 0x3e, 0x9d, 0x45, 0xc5, 0x42, 0xeb, 0xac, 0xd7, 0xbe,
	0x38, 0xd5, 0xf6, 0xa3, 0xec, 0x3b, 0x55, 0x3e, 0xa1, 0x9a, 0xf3, 0xaa, 0xfc, 0x86, 0xea, 0x31,
	0x54, 0xbc, 0x46, 0xea, 0x18, 0x35, 0xf3, 0xa8, 0x01, 0xdd, 0x6b, 0x5f, 0x9c, 0xe8, 0xd8, 0x49,
	0x7e, 0x02, 0xc2, 0xb3, 0x8c, 0xe5, 0x8e, 0x66, 0x52, 0x5f, 0xcb, 0x24, 0xcd, 0x24, 0x7d, 0x59,
	0x2d, 0xd5, 0x17, 0x0e, 0x57, 0xf7, 0x42, 0x1e, 0x7e, 0x94, 0xfd, 0x93, 0x40, 0x45, 0x4b, 0x3c,
	0xcb, 0x2e, 0x6e, 0x00, 0x84, 0xc2, 0x1c, 0x26, 0x05, 0xeb, 0x66, 0x14, 0xf0, 0xee, 0x66, 0x5c,
	0x5e, 0x7c, 0xce, 0xc8, 0x2e, 0x2c, 0x6a, 0xe6, 0x39, 0x91, 0xf6, 0x34, 0x5d, 0xf0, 0x19, 0xaa,
	0x3f, 0x1c, 0x6b, 0xdb, 0x48, 0x7b, 0xda, 0x09, 0xf8, 0x4d, 0xc1, 0xa2, 0x17, 0xf0, 0xa1, 0xe0,
	0x11, 0x40, 0x9c, 0xea, 0x96, 0xd7, 0xd0, 0x67, 0x48, 0xcf, 0x39, 0xc4, 0x29, 0xc8, 0x33, 0x58,
	0xce, 0xdb, 0x2a, 0x0b, 0x1e, 0xe2, 0xaf, 0x32, 0x6e, 0xd3, 0x4a, 0xb5, 0x54, 0x9f, 0x8b, 0x2a,
	0x0e, 0x77, 0x9a, 0x63, 0x07, 0xba, 0x70, 0x9b, 0x82, 0x09, 0x99, 0xf0, 0x3e, 0x7d, 0x88, 0x4e,
	0x66, 0x4d, 0xd1, 0x70, 0x4b, 0x52, 0x83, 0x8a, 0x29, 0x0e, 0x98, 0x30, 0x2c, 0x6d, 0xb5, 0x72,
	0x69, 0xe9, 0x12, 0xf2, 0x0b, 0xa6, 0x38, 0x68, 0x98, 0xdf, 0x10, 0x72, 0x15, 0x63, 0x8a, 0x43,
	0x57, 0x31, 0xcb, 0xbe, 0x62, 0x4c, 0x71, 0xd8, 0x30, 0x2e, 0x73, 0x1d, 0x3c, 0xaa, 0xc0, 0x15,
	0x9f, 0xb9, 0xa6, 0x38, 0xfc, 0x30, 0x2c, 0xc2, 0xdb, 0x45, 0x40, 0xee, 0x28, 0x82, 0x87, 0x30,
	0x25, 0x0c, 0x5d, 0x45, 0x66, 0x4a, 0x18, 0xb2, 0x0c, 0x65, 0x2e, 0x0c, 0x5d, 0xc3, 0x8f, 0x71,
	0x7f, 0xc9, 0x8f, 0xf0, 0x08, 0xab, 0xac, 0x9b, 0x65, 0xa9, 0xb1, 0x52, 0xb0, 0x09, 0xaf, 0xeb,
	0x68, 0x4b, 0x5d, 0xe9, 0x0d, 0x24, 0x97, 0x37, 0x77, 0xd8, 0x82, 0x39, 0xdd, 0x64, 0xd6, 0x70,
	0x9d, 0xd3, 0x4d, 0x1f, 0x02, 0xdd, 0xbc, 0x74, 0x4b, 0xf2, 0x3d, 0x6c, 0x4a, 0xcd, 0x9b, 0x89,
	0x14, 0xac, 0x8b, 0x15, 0xcf, 0x62, 0xdf, 0x5f, 0x72, 0x4a, 0xab, 0xe5, 0x7a, 0x25, 0x5a, 0x0f,
	0xb4, 0xef, 0x07, 0xa1, 0xf9, 0xe4, 0x44, 0xc2, 0xba, 0x2c, 0xac, 0xe1, 0xb7, 0xac, 0xb6, 0xaa,
	0xe5, 0xfa, 0xc2, 0xe1, 0xc1, 0x5e, 0xe8, 0x6c, 0x7b, 0x13, 0x95, 0xbb, 0x77, 0xe2, 0xac, 0xc6,
	0x9d, 0x9d, 0x68, 0x6b, 0xfa, 0xd1, 0xaa, 0xbc, 0xcd, 0x90, 0x57, 0xb0, 0x1a, 0x3c, 0x0f, 0x43,
	0xad, 0x64, 0x4e, 0xb7, 0xf1, 0x68, 0x24, 0x50, 0x1f, 0x46, 0x0c, 0xf9, 0x02, 0x24, 0x9c, 0x88,
	0x0b, 0xc3, 0xbe, 0xfa, 0xde, 0x45, 0xff, 0x83, 0x87, 0xaa, 0xdf, 0x77, 0xa8, 0xc9, 0x5e, 0x17,
	0x2d, 0x7b, 0x1f, 0x47, 0xc2, 0x0c, 0xba, 0x5f, 0x04, 0xcf, 0x12, 0x9e, 0x5b, 0x36, 0x68, 0xe3,
	0x96, 0xdb, 0x6e, 0xce, 0x70, 0xe3, 0xdc, 0x32, 0xab, 0x3a, 0x92, 0x75, 0xb5, 0x2a, 0x98, 0xce,
	0xe9, 0x4e, 0xb5, 0x54, 0x2f, 0x47, 0x8f, 0x9d, 0x3c, 0xec, 0x83, 0xe2, 0xc8, 0x6b, 0x2f, 0x55,
	0x47, 0x7e, 0xd6, 0xaa, 0x38, 0xcb, 0xc9, 0x29, 0xd4, 0xbc, 0xcf, 0xb4, 0xa7, 0xf1, 0xc8, 0xb6,
	0x40, 0x4f, 0xb9, 0xe5, 0x9d, 0x6c, 0xe8, 0xae, 0x8a, 0xee, 0x76, 0xd0, 0x5d, 0x10, 0x5e, 0x16,
	0x97, 0x03, 0x59, 0x70, 0xf5, 0x04, 0x2a, 0x4d, 0xc9, 0xe3, 0x54, 0xb3, 0x24, 0x8d, 0xdb, 0x52,
	0xd0, 0xc7, 0x98, 0x3d, 0x8b, 0x1e, 0xfc, 0x15, 0x31, 0x52, 0x85, 0xc5, 0xcc, 0xf5, 0xb5, 0x3c,
	0x49, 0x2d, 0xd3, 0x4d, 0x5a, 0xc3, 0x54, 0x00, 0x87, 0x5d, 0x24, 0xa9, 0x3d, 0x6b, 0x8e, 0x2b,
	0x84, 0xa1, 0x4f, 0xc6, 0x15, 0x0d, 0x43, 0xf6, 0x60, 0x75, 0xa4, 0x18, 0x65, 0xff, 0x53, 0x14,
	0xae, 0x0c, 0x84, 0xa3, 0x12, 0xd8, 0x85, 0x85, 0x0e, 0x8f, 0xd9, 0xb5, 0x34, 0x2e, 0xd4, 0xf4,
	0x7f, 0xd8, 0x47, 0xa1, 0xc3, 0xe3, 0x2f, 0x1e, 0xc1, 0xdc, 0x56, 0xfa, 0xfe, 0xdc, 0xfe, 0x7f,
	0xc8, 0x6d, 0xa5, 0xef, 0xce, 0xed, 0xd7, 0xb0, 0x61, 0x24, 0xf6, 0xd3, 0xc1, 0x65, 0x84, 0x84,
	0xa5, 0x2f, 0x30, 0x04, 0x6b, 0x9e, 0x0d, 0xd1, 0x3f, 0xf1, 0x1c, 0x79, 0x07, 0xdb, 0x13, 0x56,
	0xae, 0xc0, 0xf0, 0x0d, 0x62, 0x9a, 0xd6, 0x71, 0xcf, 0x8d, 0x31, 0xcb, 0x4f, 0xbc, 0xc0, 0xe7,
	0xe8, 0x8c, 0xbc, 0x85, 0xad, 0x3b, 0x6c, 0x31, 0x05, 0x34, 0xfd, 0x06, 0x4d, 0xd7, 0x27, 0x4d,
	0xdd, 0x7d, 0x9d, 0xb9, 0x7e, 0x10, 0x2c, 0xfd, 0x4e, 0xfb, 0xf4, 0x79, 0xe8, 0x1a, 0x88, 0xa2,
	0xff, 0x7d, 0x72, 0x04, 0x3b, 0x99, 0xd4, 0xc2, 0x45, 0x39, 0xa8, 0xc7, 0x67, 0x07, 0xfa, 0x2d,
	0x36, 0xf2, 0xed, 0x20, 0x8a, 0x50, 0x33, 0x96, 0xd1, 0xe4, 0x25, 0x10, 0x23, 0x5b, 0xd2, 0x48,
	0x1d, 0x4b, 0xc6, 0x13, 0xab, 0x6c, 0x57, 0x48, 0xba, 0x57, 0x2d, 0xd5, 0x4b, 0xd1, 0xca, 0x90,
	0x39, 0x0a, 0x04, 0x79, 0x03, 0x9b, 0xa1, 0x68, 0x44, 0x4f, 0x26, 0x89, 0xff, 0x96, 0xd7, 0xfb,
	0xfb, 0x9d, 0x9c, 0xbe, 0xf2, 0x41, 0xf4, 0x74, 0xc3, 0xb1, 0xee, 0x53, 0x90, 0x23, 0x3f, 0xc0,
	0xd6, 0x30, 0x75, 0x6f, 0x19, 0xee, 0xa3, 0xe1, 0xc6, 0x40, 0x30, 0x61, 0x7a, 0x00, 0xeb, 0x61,
	0x47, 0x17, 0x3b, 0xa9, 0x4c, 0x16, 0xae, 0xfb, 0x00, 0x03, 0x12, 0x6a, 0xf8, 0x13, 0x2f, 0x4e,
	0x94, 0xc9, 0xf0, 0xa2, 0xb7, 0xaf, 0x80, 0xde, 0xd7, 0x3b, 0x5c, 0xcb, 0x74, 0x2f, 0x9c, 0x9f,
	0x4c, 0xdc, 0x5f, 0xf2, 0x06, 0xa6, 0xaf, 0x79, 0xd2, 0x95, 0xf8, 0xce, 0x2f, 0x1c, 0xee, 0xde,
	0x57, 0xfa, 0xc1, 0x4f, 0xe4, 0xd5, 0xef, 0xa6, 0xde, 0x96, 0x6a, 0x7d, 0xa0, 0x5e, 0xf4, 0xb3,
	0x9f, 0x42, 0xa2, 0xdf, 0x4f, 0x75, 0x2b, 0xbd, 0x90, 0xf6, 0xfc, 0xfd, 0xcd, 0x47, 0xbd, 0x34,
	0xf6, 0xa8, 0xfb, 0x26, 0x3e, 0x35, 0x6c, 0xe2, 0xaf, 0x61, 0x5a, 0x59, 0xd9, 0xc9, 0x69, 0x19,
	0x5b, 0xcf, 0x7f, 0x27, 0xf6, 0x1f, 0x73, 0x7d, 0xfe, 0x3e, 0xf2, 0xe2, 0xda, 0xdf, 0x25, 0x58,
	0xbf, 0x53, 0x40, 0x76, 0x00, 0x06, 0x93, 0x52, 0x98, 0x74, 0x16, 0xa3, 0xf9, 0x80, 0x9c, 0x0a,
	0x42, 0xe0, 0x81, 0xc9, 0x73, 0x85, 0x07, 0x98, 0x8e, 0xf0, 0xbf, 0xeb, 0xfa, 0x49, 0x6a, 0x38,
	0x0e, 0x67, 0x65, 0xbc, 0xfa, 0x59, 0xb7, 0x76, 0xd3, 0xd9, 0x1a, 0x4c, 0x37, 0x53, 0x6e, 0x44,
	0x98, 0xb7, 0xfc, 0x82, 0x50, 0x98, 0xe5, 0xda, 0x4a, 0xad, 0x39, 0x4e, 0x2c, 0x95, 0x68, 0xb0,
	0x74, 0x4c, 0x9c, 0x6a, 0x2b, 0x0b, 0x3b, 0x98, 0x58, 0xc2, 0xb2, 0x39, 0x83, 0x63, 0xea, 0x77,
	0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x3f, 0x46, 0xb2, 0xe0, 0x0a, 0x00, 0x00,
}
