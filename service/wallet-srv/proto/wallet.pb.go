// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallet.proto

package fotune_srv_wallet

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Timestamp from public import google/protobuf/timestamp.proto
type Timestamp = timestamp.Timestamp

type StrategyRunNotifyReq struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StrategyRunNotifyReq) Reset()         { *m = StrategyRunNotifyReq{} }
func (m *StrategyRunNotifyReq) String() string { return proto.CompactTextString(m) }
func (*StrategyRunNotifyReq) ProtoMessage()    {}
func (*StrategyRunNotifyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{0}
}

func (m *StrategyRunNotifyReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StrategyRunNotifyReq.Unmarshal(m, b)
}
func (m *StrategyRunNotifyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StrategyRunNotifyReq.Marshal(b, m, deterministic)
}
func (m *StrategyRunNotifyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StrategyRunNotifyReq.Merge(m, src)
}
func (m *StrategyRunNotifyReq) XXX_Size() int {
	return xxx_messageInfo_StrategyRunNotifyReq.Size(m)
}
func (m *StrategyRunNotifyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StrategyRunNotifyReq.DiscardUnknown(m)
}

var xxx_messageInfo_StrategyRunNotifyReq proto.InternalMessageInfo

func (m *StrategyRunNotifyReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetTotalRebateReq struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTotalRebateReq) Reset()         { *m = GetTotalRebateReq{} }
func (m *GetTotalRebateReq) String() string { return proto.CompactTextString(m) }
func (*GetTotalRebateReq) ProtoMessage()    {}
func (*GetTotalRebateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{1}
}

func (m *GetTotalRebateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTotalRebateReq.Unmarshal(m, b)
}
func (m *GetTotalRebateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTotalRebateReq.Marshal(b, m, deterministic)
}
func (m *GetTotalRebateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTotalRebateReq.Merge(m, src)
}
func (m *GetTotalRebateReq) XXX_Size() int {
	return xxx_messageInfo_GetTotalRebateReq.Size(m)
}
func (m *GetTotalRebateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTotalRebateReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTotalRebateReq proto.InternalMessageInfo

func (m *GetTotalRebateReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type IfcRecord struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Volume               string   `protobuf:"bytes,2,opt,name=volume,proto3" json:"volume,omitempty"`
	TypeMsg              string   `protobuf:"bytes,3,opt,name=type_msg,json=typeMsg,proto3" json:"type_msg,omitempty"`
	Date                 string   `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IfcRecord) Reset()         { *m = IfcRecord{} }
func (m *IfcRecord) String() string { return proto.CompactTextString(m) }
func (*IfcRecord) ProtoMessage()    {}
func (*IfcRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{2}
}

func (m *IfcRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IfcRecord.Unmarshal(m, b)
}
func (m *IfcRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IfcRecord.Marshal(b, m, deterministic)
}
func (m *IfcRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IfcRecord.Merge(m, src)
}
func (m *IfcRecord) XXX_Size() int {
	return xxx_messageInfo_IfcRecord.Size(m)
}
func (m *IfcRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_IfcRecord.DiscardUnknown(m)
}

var xxx_messageInfo_IfcRecord proto.InternalMessageInfo

func (m *IfcRecord) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *IfcRecord) GetVolume() string {
	if m != nil {
		return m.Volume
	}
	return ""
}

func (m *IfcRecord) GetTypeMsg() string {
	if m != nil {
		return m.TypeMsg
	}
	return ""
}

func (m *IfcRecord) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

type GetTotalRebateResp struct {
	Total                string       `protobuf:"bytes,1,opt,name=total,proto3" json:"total,omitempty"`
	Record               []*IfcRecord `protobuf:"bytes,2,rep,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetTotalRebateResp) Reset()         { *m = GetTotalRebateResp{} }
func (m *GetTotalRebateResp) String() string { return proto.CompactTextString(m) }
func (*GetTotalRebateResp) ProtoMessage()    {}
func (*GetTotalRebateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{3}
}

func (m *GetTotalRebateResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTotalRebateResp.Unmarshal(m, b)
}
func (m *GetTotalRebateResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTotalRebateResp.Marshal(b, m, deterministic)
}
func (m *GetTotalRebateResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTotalRebateResp.Merge(m, src)
}
func (m *GetTotalRebateResp) XXX_Size() int {
	return xxx_messageInfo_GetTotalRebateResp.Size(m)
}
func (m *GetTotalRebateResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTotalRebateResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetTotalRebateResp proto.InternalMessageInfo

func (m *GetTotalRebateResp) GetTotal() string {
	if m != nil {
		return m.Total
	}
	return ""
}

func (m *GetTotalRebateResp) GetRecord() []*IfcRecord {
	if m != nil {
		return m.Record
	}
	return nil
}

type AddIfcBalanceReq struct {
	UserMasterId         string   `protobuf:"bytes,1,opt,name=user_master_id,json=userMasterId,proto3" json:"user_master_id,omitempty"`
	InUserId             string   `protobuf:"bytes,2,opt,name=in_user_id,json=inUserId,proto3" json:"in_user_id,omitempty"`
	Volume               float64  `protobuf:"fixed64,3,opt,name=volume,proto3" json:"volume,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Exchange             string   `protobuf:"bytes,5,opt,name=exchange,proto3" json:"exchange,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddIfcBalanceReq) Reset()         { *m = AddIfcBalanceReq{} }
func (m *AddIfcBalanceReq) String() string { return proto.CompactTextString(m) }
func (*AddIfcBalanceReq) ProtoMessage()    {}
func (*AddIfcBalanceReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{4}
}

func (m *AddIfcBalanceReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddIfcBalanceReq.Unmarshal(m, b)
}
func (m *AddIfcBalanceReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddIfcBalanceReq.Marshal(b, m, deterministic)
}
func (m *AddIfcBalanceReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddIfcBalanceReq.Merge(m, src)
}
func (m *AddIfcBalanceReq) XXX_Size() int {
	return xxx_messageInfo_AddIfcBalanceReq.Size(m)
}
func (m *AddIfcBalanceReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddIfcBalanceReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddIfcBalanceReq proto.InternalMessageInfo

func (m *AddIfcBalanceReq) GetUserMasterId() string {
	if m != nil {
		return m.UserMasterId
	}
	return ""
}

func (m *AddIfcBalanceReq) GetInUserId() string {
	if m != nil {
		return m.InUserId
	}
	return ""
}

func (m *AddIfcBalanceReq) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *AddIfcBalanceReq) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AddIfcBalanceReq) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

type WithdrawalReq struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Coin                 string   `protobuf:"bytes,2,opt,name=coin,proto3" json:"coin,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Volume               float64  `protobuf:"fixed64,4,opt,name=volume,proto3" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WithdrawalReq) Reset()         { *m = WithdrawalReq{} }
func (m *WithdrawalReq) String() string { return proto.CompactTextString(m) }
func (*WithdrawalReq) ProtoMessage()    {}
func (*WithdrawalReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{5}
}

func (m *WithdrawalReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WithdrawalReq.Unmarshal(m, b)
}
func (m *WithdrawalReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WithdrawalReq.Marshal(b, m, deterministic)
}
func (m *WithdrawalReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WithdrawalReq.Merge(m, src)
}
func (m *WithdrawalReq) XXX_Size() int {
	return xxx_messageInfo_WithdrawalReq.Size(m)
}
func (m *WithdrawalReq) XXX_DiscardUnknown() {
	xxx_messageInfo_WithdrawalReq.DiscardUnknown(m)
}

var xxx_messageInfo_WithdrawalReq proto.InternalMessageInfo

func (m *WithdrawalReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *WithdrawalReq) GetCoin() string {
	if m != nil {
		return m.Coin
	}
	return ""
}

func (m *WithdrawalReq) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *WithdrawalReq) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

type ConvertCoinTipsReq struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	UserId               string   `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConvertCoinTipsReq) Reset()         { *m = ConvertCoinTipsReq{} }
func (m *ConvertCoinTipsReq) String() string { return proto.CompactTextString(m) }
func (*ConvertCoinTipsReq) ProtoMessage()    {}
func (*ConvertCoinTipsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{6}
}

func (m *ConvertCoinTipsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConvertCoinTipsReq.Unmarshal(m, b)
}
func (m *ConvertCoinTipsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConvertCoinTipsReq.Marshal(b, m, deterministic)
}
func (m *ConvertCoinTipsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConvertCoinTipsReq.Merge(m, src)
}
func (m *ConvertCoinTipsReq) XXX_Size() int {
	return xxx_messageInfo_ConvertCoinTipsReq.Size(m)
}
func (m *ConvertCoinTipsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ConvertCoinTipsReq.DiscardUnknown(m)
}

var xxx_messageInfo_ConvertCoinTipsReq proto.InternalMessageInfo

func (m *ConvertCoinTipsReq) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *ConvertCoinTipsReq) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *ConvertCoinTipsReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ConvertCoinTipsResp struct {
	Describe             string   `protobuf:"bytes,1,opt,name=describe,proto3" json:"describe,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConvertCoinTipsResp) Reset()         { *m = ConvertCoinTipsResp{} }
func (m *ConvertCoinTipsResp) String() string { return proto.CompactTextString(m) }
func (*ConvertCoinTipsResp) ProtoMessage()    {}
func (*ConvertCoinTipsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{7}
}

func (m *ConvertCoinTipsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConvertCoinTipsResp.Unmarshal(m, b)
}
func (m *ConvertCoinTipsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConvertCoinTipsResp.Marshal(b, m, deterministic)
}
func (m *ConvertCoinTipsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConvertCoinTipsResp.Merge(m, src)
}
func (m *ConvertCoinTipsResp) XXX_Size() int {
	return xxx_messageInfo_ConvertCoinTipsResp.Size(m)
}
func (m *ConvertCoinTipsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ConvertCoinTipsResp.DiscardUnknown(m)
}

var xxx_messageInfo_ConvertCoinTipsResp proto.InternalMessageInfo

func (m *ConvertCoinTipsResp) GetDescribe() string {
	if m != nil {
		return m.Describe
	}
	return ""
}

type ConvertCoinReq struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Volume               float64  `protobuf:"fixed64,3,opt,name=volume,proto3" json:"volume,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConvertCoinReq) Reset()         { *m = ConvertCoinReq{} }
func (m *ConvertCoinReq) String() string { return proto.CompactTextString(m) }
func (*ConvertCoinReq) ProtoMessage()    {}
func (*ConvertCoinReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{8}
}

func (m *ConvertCoinReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConvertCoinReq.Unmarshal(m, b)
}
func (m *ConvertCoinReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConvertCoinReq.Marshal(b, m, deterministic)
}
func (m *ConvertCoinReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConvertCoinReq.Merge(m, src)
}
func (m *ConvertCoinReq) XXX_Size() int {
	return xxx_messageInfo_ConvertCoinReq.Size(m)
}
func (m *ConvertCoinReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ConvertCoinReq.DiscardUnknown(m)
}

var xxx_messageInfo_ConvertCoinReq proto.InternalMessageInfo

func (m *ConvertCoinReq) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *ConvertCoinReq) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *ConvertCoinReq) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *ConvertCoinReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ConvertCoinResp struct {
	Describe             string   `protobuf:"bytes,1,opt,name=describe,proto3" json:"describe,omitempty"`
	Volume               float64  `protobuf:"fixed64,2,opt,name=volume,proto3" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConvertCoinResp) Reset()         { *m = ConvertCoinResp{} }
func (m *ConvertCoinResp) String() string { return proto.CompactTextString(m) }
func (*ConvertCoinResp) ProtoMessage()    {}
func (*ConvertCoinResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{9}
}

func (m *ConvertCoinResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConvertCoinResp.Unmarshal(m, b)
}
func (m *ConvertCoinResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConvertCoinResp.Marshal(b, m, deterministic)
}
func (m *ConvertCoinResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConvertCoinResp.Merge(m, src)
}
func (m *ConvertCoinResp) XXX_Size() int {
	return xxx_messageInfo_ConvertCoinResp.Size(m)
}
func (m *ConvertCoinResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ConvertCoinResp.DiscardUnknown(m)
}

var xxx_messageInfo_ConvertCoinResp proto.InternalMessageInfo

func (m *ConvertCoinResp) GetDescribe() string {
	if m != nil {
		return m.Describe
	}
	return ""
}

func (m *ConvertCoinResp) GetVolume() float64 {
	if m != nil {
		return m.Volume
	}
	return 0
}

type UsdtDepositAddrResp struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UsdtDepositAddrResp) Reset()         { *m = UsdtDepositAddrResp{} }
func (m *UsdtDepositAddrResp) String() string { return proto.CompactTextString(m) }
func (*UsdtDepositAddrResp) ProtoMessage()    {}
func (*UsdtDepositAddrResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{10}
}

func (m *UsdtDepositAddrResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UsdtDepositAddrResp.Unmarshal(m, b)
}
func (m *UsdtDepositAddrResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UsdtDepositAddrResp.Marshal(b, m, deterministic)
}
func (m *UsdtDepositAddrResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UsdtDepositAddrResp.Merge(m, src)
}
func (m *UsdtDepositAddrResp) XXX_Size() int {
	return xxx_messageInfo_UsdtDepositAddrResp.Size(m)
}
func (m *UsdtDepositAddrResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UsdtDepositAddrResp.DiscardUnknown(m)
}

var xxx_messageInfo_UsdtDepositAddrResp proto.InternalMessageInfo

func (m *UsdtDepositAddrResp) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type UidReq struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UidReq) Reset()         { *m = UidReq{} }
func (m *UidReq) String() string { return proto.CompactTextString(m) }
func (*UidReq) ProtoMessage()    {}
func (*UidReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{11}
}

func (m *UidReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UidReq.Unmarshal(m, b)
}
func (m *UidReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UidReq.Marshal(b, m, deterministic)
}
func (m *UidReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UidReq.Merge(m, src)
}
func (m *UidReq) XXX_Size() int {
	return xxx_messageInfo_UidReq.Size(m)
}
func (m *UidReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UidReq.DiscardUnknown(m)
}

var xxx_messageInfo_UidReq proto.InternalMessageInfo

func (m *UidReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type TransferReq struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	FromCoin             string   `protobuf:"bytes,2,opt,name=fromCoin,proto3" json:"fromCoin,omitempty"`
	ToCoin               string   `protobuf:"bytes,3,opt,name=toCoin,proto3" json:"toCoin,omitempty"`
	FromCoinAmount       float64  `protobuf:"fixed64,4,opt,name=fromCoinAmount,proto3" json:"fromCoinAmount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferReq) Reset()         { *m = TransferReq{} }
func (m *TransferReq) String() string { return proto.CompactTextString(m) }
func (*TransferReq) ProtoMessage()    {}
func (*TransferReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{12}
}

func (m *TransferReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferReq.Unmarshal(m, b)
}
func (m *TransferReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferReq.Marshal(b, m, deterministic)
}
func (m *TransferReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferReq.Merge(m, src)
}
func (m *TransferReq) XXX_Size() int {
	return xxx_messageInfo_TransferReq.Size(m)
}
func (m *TransferReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferReq.DiscardUnknown(m)
}

var xxx_messageInfo_TransferReq proto.InternalMessageInfo

func (m *TransferReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *TransferReq) GetFromCoin() string {
	if m != nil {
		return m.FromCoin
	}
	return ""
}

func (m *TransferReq) GetToCoin() string {
	if m != nil {
		return m.ToCoin
	}
	return ""
}

func (m *TransferReq) GetFromCoinAmount() float64 {
	if m != nil {
		return m.FromCoinAmount
	}
	return 0
}

type WalletBalanceResp struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Symbol               string   `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Total                string   `protobuf:"bytes,3,opt,name=total,proto3" json:"total,omitempty"`
	Available            string   `protobuf:"bytes,4,opt,name=available,proto3" json:"available,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WalletBalanceResp) Reset()         { *m = WalletBalanceResp{} }
func (m *WalletBalanceResp) String() string { return proto.CompactTextString(m) }
func (*WalletBalanceResp) ProtoMessage()    {}
func (*WalletBalanceResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{13}
}

func (m *WalletBalanceResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletBalanceResp.Unmarshal(m, b)
}
func (m *WalletBalanceResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletBalanceResp.Marshal(b, m, deterministic)
}
func (m *WalletBalanceResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletBalanceResp.Merge(m, src)
}
func (m *WalletBalanceResp) XXX_Size() int {
	return xxx_messageInfo_WalletBalanceResp.Size(m)
}
func (m *WalletBalanceResp) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletBalanceResp.DiscardUnknown(m)
}

var xxx_messageInfo_WalletBalanceResp proto.InternalMessageInfo

func (m *WalletBalanceResp) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *WalletBalanceResp) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *WalletBalanceResp) GetTotal() string {
	if m != nil {
		return m.Total
	}
	return ""
}

func (m *WalletBalanceResp) GetAvailable() string {
	if m != nil {
		return m.Available
	}
	return ""
}

func init() {
	proto.RegisterType((*StrategyRunNotifyReq)(nil), "StrategyRunNotifyReq")
	proto.RegisterType((*GetTotalRebateReq)(nil), "GetTotalRebateReq")
	proto.RegisterType((*IfcRecord)(nil), "ifcRecord")
	proto.RegisterType((*GetTotalRebateResp)(nil), "GetTotalRebateResp")
	proto.RegisterType((*AddIfcBalanceReq)(nil), "AddIfcBalanceReq")
	proto.RegisterType((*WithdrawalReq)(nil), "withdrawalReq")
	proto.RegisterType((*ConvertCoinTipsReq)(nil), "convertCoinTipsReq")
	proto.RegisterType((*ConvertCoinTipsResp)(nil), "convertCoinTipsResp")
	proto.RegisterType((*ConvertCoinReq)(nil), "convertCoinReq")
	proto.RegisterType((*ConvertCoinResp)(nil), "convertCoinResp")
	proto.RegisterType((*UsdtDepositAddrResp)(nil), "usdtDepositAddrResp")
	proto.RegisterType((*UidReq)(nil), "uidReq")
	proto.RegisterType((*TransferReq)(nil), "transferReq")
	proto.RegisterType((*WalletBalanceResp)(nil), "walletBalanceResp")
}

func init() { proto.RegisterFile("wallet.proto", fileDescriptor_b88fd140af4deb6f) }

var fileDescriptor_b88fd140af4deb6f = []byte{
	// 790 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdf, 0x6f, 0x22, 0x37,
	0x10, 0x86, 0x84, 0x03, 0x32, 0x01, 0x72, 0x0c, 0x5c, 0x4a, 0xe9, 0x49, 0xbd, 0xae, 0xaa, 0xea,
	0x1e, 0x5a, 0xa3, 0x4b, 0xd5, 0xbe, 0x5c, 0x55, 0x29, 0x47, 0xae, 0x27, 0x1e, 0xee, 0xd4, 0x12,
	0xa2, 0x48, 0x7d, 0x41, 0x66, 0xd7, 0x80, 0xa5, 0xdd, 0xf5, 0xd6, 0xf6, 0x92, 0xf2, 0xd8, 0x7f,
	0xa4, 0xff, 0x67, 0xdf, 0x2a, 0x7b, 0x7f, 0xb0, 0x0b, 0x09, 0xed, 0x9b, 0xbf, 0xf1, 0x8c, 0x67,
	0xe6, 0xb3, 0xbf, 0x31, 0xb4, 0x1e, 0xa8, 0xef, 0x33, 0x4d, 0x22, 0x29, 0xb4, 0x18, 0x7e, 0xb1,
	0x12, 0x62, 0xe5, 0xb3, 0x91, 0x45, 0x8b, 0x78, 0x39, 0x62, 0x41, 0xa4, 0xb7, 0xe9, 0xe6, 0x97,
	0xfb, 0x9b, 0x9a, 0x07, 0x4c, 0x69, 0x1a, 0x44, 0x89, 0x83, 0x33, 0x82, 0xfe, 0xad, 0x96, 0x54,
	0xb3, 0xd5, 0x76, 0x1a, 0x87, 0x9f, 0x84, 0xe6, 0xcb, 0xed, 0x94, 0xfd, 0x81, 0x9f, 0x41, 0x23,
	0x56, 0x4c, 0xce, 0xb9, 0x37, 0xa8, 0xbe, 0xaa, 0xbe, 0x3e, 0x9b, 0xd6, 0x0d, 0x9c, 0x78, 0xce,
	0xb7, 0xd0, 0xfd, 0xc0, 0xf4, 0x4c, 0x68, 0xea, 0x4f, 0xd9, 0x82, 0x6a, 0x76, 0xd4, 0x7b, 0x0d,
	0x67, 0x7c, 0xe9, 0x4e, 0x99, 0x2b, 0xa4, 0x87, 0x7d, 0x78, 0x16, 0xad, 0x45, 0xc8, 0x52, 0x9f,
	0x04, 0xe0, 0x25, 0xd4, 0x37, 0xc2, 0x8f, 0x03, 0x36, 0x38, 0x49, 0x42, 0x13, 0x84, 0x9f, 0x43,
	0x53, 0x6f, 0x23, 0x36, 0x0f, 0xd4, 0x6a, 0x70, 0x6a, 0x77, 0x1a, 0x06, 0x7f, 0x54, 0x2b, 0x44,
	0xa8, 0x79, 0x54, 0xb3, 0x41, 0xcd, 0x9a, 0xed, 0xda, 0xf9, 0x04, 0xb8, 0x5f, 0x97, 0x8a, 0x4c,
	0x4a, 0x6d, 0x4c, 0x59, 0x4a, 0x0b, 0xd0, 0x81, 0xba, 0xb4, 0x25, 0x0d, 0x4e, 0x5e, 0x9d, 0xbe,
	0x3e, 0xbf, 0x02, 0x92, 0x17, 0x39, 0x4d, 0x77, 0x9c, 0xbf, 0xab, 0xf0, 0xfc, 0xda, 0xf3, 0x26,
	0x4b, 0xf7, 0x1d, 0xf5, 0x69, 0xe8, 0xda, 0x3e, 0xbf, 0x86, 0x8e, 0xed, 0x33, 0xa0, 0x4a, 0x17,
	0xdb, 0x6d, 0x19, 0xeb, 0x47, 0x6b, 0x9c, 0x78, 0xf8, 0x12, 0x80, 0x87, 0xf3, 0x8c, 0x90, 0xa4,
	0xab, 0x26, 0x0f, 0xef, 0x2c, 0x25, 0x85, 0x7e, 0x4d, 0x57, 0xd5, 0xbc, 0x5f, 0x84, 0x9a, 0xe9,
	0x2f, 0x6b, 0xca, 0xac, 0x71, 0x08, 0x4d, 0xf6, 0xa7, 0xbb, 0xa6, 0xe1, 0x8a, 0x0d, 0x9e, 0x25,
	0xe7, 0x64, 0xd8, 0x09, 0xa1, 0xfd, 0xc0, 0xf5, 0xda, 0x93, 0xf4, 0xc1, 0xb4, 0xfc, 0xf4, 0x25,
	0x98, 0x93, 0x5d, 0xc1, 0xc3, 0xb4, 0x12, 0xbb, 0xc6, 0x01, 0x34, 0xa8, 0xe7, 0x49, 0xa6, 0x54,
	0x46, 0x6e, 0x0a, 0x0b, 0xf5, 0xd5, 0x8a, 0xf5, 0x39, 0xbf, 0x01, 0xba, 0x22, 0xdc, 0x30, 0xa9,
	0xc7, 0x82, 0x87, 0x33, 0x1e, 0x29, 0x93, 0x14, 0xa1, 0xb6, 0x94, 0x22, 0x48, 0x33, 0xda, 0x35,
	0x76, 0xe0, 0x44, 0x8b, 0x34, 0xdb, 0x89, 0x16, 0xc5, 0xc2, 0x4e, 0x4b, 0xaf, 0xe3, 0x0d, 0xf4,
	0x0e, 0x8e, 0x54, 0x91, 0xe9, 0xda, 0x63, 0xca, 0x95, 0x7c, 0x91, 0x3d, 0x95, 0x1c, 0x3b, 0x0c,
	0x3a, 0x85, 0x90, 0xff, 0x5b, 0xc1, 0x53, 0x9c, 0x17, 0x2a, 0xab, 0x95, 0x2a, 0x7b, 0x0f, 0x17,
	0xa5, 0x34, 0xc7, 0xab, 0xda, 0x7b, 0xc3, 0x3b, 0xce, 0x46, 0xd0, 0x8b, 0x95, 0xa7, 0x6f, 0x58,
	0x24, 0x14, 0xd7, 0xd7, 0x9e, 0x27, 0xed, 0x51, 0x05, 0xf2, 0xab, 0x25, 0xf2, 0x9d, 0xaf, 0xa0,
	0x1e, 0x73, 0xef, 0xa8, 0xa4, 0xfe, 0xaa, 0xc2, 0xb9, 0x96, 0x34, 0x54, 0x4b, 0x26, 0x8d, 0xe3,
	0x25, 0xa4, 0x3b, 0x7b, 0xb7, 0x3e, 0x84, 0xa6, 0xe1, 0x62, 0xbc, 0xbb, 0xf9, 0x1c, 0x9b, 0x18,
	0x2d, 0xec, 0x4e, 0x7a, 0x21, 0x09, 0xc2, 0x6f, 0xa0, 0x93, 0xf9, 0x5c, 0x07, 0x22, 0x0e, 0x75,
	0xfa, 0x06, 0xf6, 0xac, 0x4e, 0x0c, 0xdd, 0x64, 0x06, 0xe5, 0xda, 0x48, 0xb5, 0xc6, 0xb5, 0x9f,
	0xcb, 0xdb, 0x02, 0x93, 0x4a, 0x6d, 0x83, 0x85, 0xf0, 0x33, 0x79, 0x27, 0x68, 0xa7, 0xcc, 0xd3,
	0xa2, 0x32, 0x5f, 0xc2, 0x19, 0xdd, 0x50, 0xee, 0xd3, 0x85, 0x9f, 0x29, 0x61, 0x67, 0xb8, 0xfa,
	0xa7, 0x06, 0xed, 0x7b, 0x9b, 0xf7, 0x96, 0xc9, 0x0d, 0x77, 0x19, 0xbe, 0x81, 0xd6, 0x58, 0x32,
	0xaa, 0x59, 0x62, 0xc6, 0x06, 0x49, 0xe8, 0x1b, 0x5e, 0x92, 0x64, 0xf2, 0x91, 0x6c, 0xf2, 0x91,
	0xf7, 0x66, 0x2c, 0x3a, 0x15, 0xbc, 0x82, 0xe6, 0x2c, 0xa5, 0x0f, 0x5b, 0xa4, 0xc0, 0xe4, 0x91,
	0x98, 0xef, 0xa0, 0xf5, 0x81, 0xe9, 0x24, 0xc7, 0xe4, 0x97, 0xf1, 0x2e, 0x0d, 0x92, 0x03, 0x1e,
	0x9c, 0x0a, 0x12, 0x68, 0xe7, 0xee, 0x77, 0xb7, 0x37, 0xb3, 0xff, 0xf2, 0xff, 0xc1, 0xce, 0xae,
	0xbb, 0xf2, 0x4b, 0xd9, 0x05, 0xf5, 0xc9, 0x23, 0x8f, 0xc8, 0xa9, 0xe0, 0xcf, 0x70, 0x31, 0x2e,
	0xcb, 0x07, 0x7b, 0xe4, 0x50, 0xa3, 0xc3, 0x3e, 0x79, 0x44, 0x65, 0x96, 0x89, 0xf3, 0x42, 0x3c,
	0x5e, 0x90, 0xb2, 0xb2, 0x86, 0xcf, 0xc9, 0x9e, 0x06, 0x9c, 0x0a, 0xfe, 0x08, 0x70, 0x9f, 0x4f,
	0x1d, 0xec, 0x90, 0xd2, 0x08, 0x3a, 0xc2, 0xe0, 0x4f, 0xd0, 0x2e, 0x4d, 0x53, 0xec, 0x92, 0xfd,
	0xe9, 0x7a, 0x24, 0xfa, 0x2d, 0x74, 0xca, 0xc3, 0x1d, 0x91, 0x1c, 0xfc, 0x42, 0xc3, 0x1e, 0x39,
	0xfc, 0x01, 0x9c, 0x0a, 0xde, 0x40, 0xf7, 0xe0, 0x8b, 0xc3, 0x17, 0xe4, 0xb1, 0x6f, 0xef, 0xe9,
	0x12, 0xde, 0xbd, 0xf8, 0xbd, 0x47, 0xde, 0x2e, 0x85, 0x8e, 0x43, 0x36, 0x57, 0x72, 0x33, 0x4f,
	0xee, 0xf1, 0xd7, 0xea, 0xa2, 0x6e, 0x5d, 0xbf, 0xff, 0x37, 0x00, 0x00, 0xff, 0xff, 0x17, 0x33,
	0xf4, 0xa7, 0x96, 0x07, 0x00, 0x00,
}