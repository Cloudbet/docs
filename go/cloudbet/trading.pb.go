// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.5
// source: cloudbet/trading.proto

package cloudbet

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

// Side of a selection signals whether a selection
// is available for back or lay side betting
type Side int32

const (
	// selection offers back side betting
	Side_BACK Side = 0
	// selection offers lay side betting
	Side_LAY Side = 1
)

// Enum value maps for Side.
var (
	Side_name = map[int32]string{
		0: "BACK",
		1: "LAY",
	}
	Side_value = map[string]int32{
		"BACK": 0,
		"LAY":  1,
	}
)

func (x Side) Enum() *Side {
	p := new(Side)
	*p = x
	return p
}

func (x Side) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Side) Descriptor() protoreflect.EnumDescriptor {
	return file_cloudbet_trading_proto_enumTypes[0].Descriptor()
}

func (Side) Type() protoreflect.EnumType {
	return &file_cloudbet_trading_proto_enumTypes[0]
}

func (x Side) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Side.Descriptor instead.
func (Side) EnumDescriptor() ([]byte, []int) {
	return file_cloudbet_trading_proto_rawDescGZIP(), []int{0}
}

type AcceptPriceChange int32

const (
	// not accepting any price change
	AcceptPriceChange_NONE AcceptPriceChange = 0
	// price accepted could be lower than input
	AcceptPriceChange_ALL AcceptPriceChange = 1
	// accept price change go higher or no change
	AcceptPriceChange_BETTER AcceptPriceChange = 2
)

// Enum value maps for AcceptPriceChange.
var (
	AcceptPriceChange_name = map[int32]string{
		0: "NONE",
		1: "ALL",
		2: "BETTER",
	}
	AcceptPriceChange_value = map[string]int32{
		"NONE":   0,
		"ALL":    1,
		"BETTER": 2,
	}
)

func (x AcceptPriceChange) Enum() *AcceptPriceChange {
	p := new(AcceptPriceChange)
	*p = x
	return p
}

func (x AcceptPriceChange) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AcceptPriceChange) Descriptor() protoreflect.EnumDescriptor {
	return file_cloudbet_trading_proto_enumTypes[1].Descriptor()
}

func (AcceptPriceChange) Type() protoreflect.EnumType {
	return &file_cloudbet_trading_proto_enumTypes[1]
}

func (x AcceptPriceChange) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AcceptPriceChange.Descriptor instead.
func (AcceptPriceChange) EnumDescriptor() ([]byte, []int) {
	return file_cloudbet_trading_proto_rawDescGZIP(), []int{1}
}

// Public Bet Request
//
// swagger:model PublicBetRequest
type PublicBetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reference ID, randomly generated request id to allow idempotent calls. Required to be in the UUID format.
	//
	// required: true
	// example: 0e81877b-4b10-4587-a4cf-89f4336ac843
	ReferenceId string `protobuf:"bytes,1,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
	// Stake in number
	//
	// required: true
	// example: 1.1
	Stake string `protobuf:"bytes,2,opt,name=stake,proto3" json:"stake,omitempty"`
	// Price as number
	//
	// required: true
	// example: 1.86
	Price string `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	// Event ID
	//
	// required: true
	// example: 5945044
	EventId string `protobuf:"bytes,4,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	// Market URL, is compiled from the feed data as `marketKey/outcome?params`, if no params (empty string) are present omit the query string and format as `marketKey/outcome`
	// Special Note about handicap markets:
	// For handicap markets a line is identified by the same params. Selections can be grouped by market key and params alone. Home and away outcomes then have the same market URL for the same handicap lines. The handicap value is dictated by the home team value and inverted on the away side.
	//
	// required: true
	// example: soccer.total_goals/over?total=3
	MarketUrl string `protobuf:"bytes,5,opt,name=market_url,json=marketUrl,proto3" json:"market_url,omitempty"`
	// Side, default to be BACK
	//
	// required: false
	Side Side `protobuf:"varint,6,opt,name=side,proto3,enum=cloudbet.Side" json:"side,omitempty"`
	// Currency for given stake
	//
	// required: true
	// example: PLAY_EUR
	Currency string `protobuf:"bytes,7,opt,name=currency,proto3" json:"currency,omitempty"`
	// Accept better price, default to be no price change
	//
	// required: false
	// example: BETTER
	AcceptPriceChange AcceptPriceChange `protobuf:"varint,12,opt,name=accept_price_change,json=acceptPriceChange,proto3,enum=cloudbet.AcceptPriceChange" json:"accept_price_change,omitempty"`
}

func (x *PublicBetRequest) Reset() {
	*x = PublicBetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudbet_trading_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicBetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicBetRequest) ProtoMessage() {}

func (x *PublicBetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudbet_trading_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicBetRequest.ProtoReflect.Descriptor instead.
func (*PublicBetRequest) Descriptor() ([]byte, []int) {
	return file_cloudbet_trading_proto_rawDescGZIP(), []int{0}
}

func (x *PublicBetRequest) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *PublicBetRequest) GetStake() string {
	if x != nil {
		return x.Stake
	}
	return ""
}

func (x *PublicBetRequest) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *PublicBetRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *PublicBetRequest) GetMarketUrl() string {
	if x != nil {
		return x.MarketUrl
	}
	return ""
}

func (x *PublicBetRequest) GetSide() Side {
	if x != nil {
		return x.Side
	}
	return Side_BACK
}

func (x *PublicBetRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *PublicBetRequest) GetAcceptPriceChange() AcceptPriceChange {
	if x != nil {
		return x.AcceptPriceChange
	}
	return AcceptPriceChange_NONE
}

// Public Bet Response
//
// swagger:model PublicBetResponse
type PublicBetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Reference ID
	//
	// required: true
	// example: 0e81877b-4b10-4587-a4cf-89f4336ac843
	ReferenceId string `protobuf:"bytes,1,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
	// Price as number
	//
	// required: true
	// example: 1.86
	Price string `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	// Event ID
	//
	// required: true
	// example: 5945044
	EventId string `protobuf:"bytes,4,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	// Market URL
	//
	// required: true
	// example: soccer.total_goals/over?total=3
	MarketUrl string `protobuf:"bytes,5,opt,name=market_url,json=marketUrl,proto3" json:"market_url,omitempty"`
	// Side, default to be BACK
	//
	// required: false
	Side Side `protobuf:"varint,6,opt,name=side,proto3,enum=cloudbet.Side" json:"side,omitempty"`
	// Currency for given stake
	//
	// required: true
	// example: PLAY_EUR
	Currency string `protobuf:"bytes,7,opt,name=currency,proto3" json:"currency,omitempty"`
	// Stake in number
	//
	// required: true
	// example: 1.1
	Stake string `protobuf:"bytes,8,opt,name=stake,proto3" json:"stake,omitempty"`
	// Create time, time in string format "2006-01-02T15:04:05Z07:00" (RFC3339)
	//
	// required: true
	// example: "2006-01-02T15:04:05Z07:00"
	CreateTime string `protobuf:"bytes,18,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Bet status
	//
	// required: true
	Status Status `protobuf:"varint,9,opt,name=status,proto3,enum=cloudbet.Status" json:"status,omitempty"`
	// Return Amount as number, pending bets will have 0
	//
	// required: true
	// example: 0.0
	ReturnAmount string `protobuf:"bytes,12,opt,name=return_amount,json=returnAmount,proto3" json:"return_amount,omitempty"`
	// Event name in english
	//
	// required: true
	// example: TSG Hoffenheim V Borussia Monchengladbach
	EventName string `protobuf:"bytes,13,opt,name=event_name,json=eventName,proto3" json:"event_name,omitempty"`
	// Sports Key
	//
	// required: true
	// example: soccer
	SportsKey string `protobuf:"bytes,14,opt,name=sports_key,json=sportsKey,proto3" json:"sports_key,omitempty"`
	// Competition ID
	//
	// required: true
	// example: 32
	CompetitionId int64 `protobuf:"varint,15,opt,name=competition_id,json=competitionId,proto3" json:"competition_id,omitempty"`
	// Category Key
	//
	// required: true
	// example: germany
	CategoryKey string `protobuf:"bytes,16,opt,name=category_key,json=categoryKey,proto3" json:"category_key,omitempty"`
	// Customer Reference
	//
	// required: true
	CustomerReference string `protobuf:"bytes,11,opt,name=customer_reference,json=customerReference,proto3" json:"customer_reference,omitempty"`
	// Error Details, indicates some descriptive reasoning on bet rejection, for `MALFORMED_REQUEST`
	//
	// required: false
	Error string `protobuf:"bytes,17,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *PublicBetResponse) Reset() {
	*x = PublicBetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudbet_trading_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicBetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicBetResponse) ProtoMessage() {}

func (x *PublicBetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloudbet_trading_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicBetResponse.ProtoReflect.Descriptor instead.
func (*PublicBetResponse) Descriptor() ([]byte, []int) {
	return file_cloudbet_trading_proto_rawDescGZIP(), []int{1}
}

func (x *PublicBetResponse) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *PublicBetResponse) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *PublicBetResponse) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *PublicBetResponse) GetMarketUrl() string {
	if x != nil {
		return x.MarketUrl
	}
	return ""
}

func (x *PublicBetResponse) GetSide() Side {
	if x != nil {
		return x.Side
	}
	return Side_BACK
}

func (x *PublicBetResponse) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *PublicBetResponse) GetStake() string {
	if x != nil {
		return x.Stake
	}
	return ""
}

func (x *PublicBetResponse) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *PublicBetResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_INTERNAL_SERVER_ERROR
}

func (x *PublicBetResponse) GetReturnAmount() string {
	if x != nil {
		return x.ReturnAmount
	}
	return ""
}

func (x *PublicBetResponse) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

func (x *PublicBetResponse) GetSportsKey() string {
	if x != nil {
		return x.SportsKey
	}
	return ""
}

func (x *PublicBetResponse) GetCompetitionId() int64 {
	if x != nil {
		return x.CompetitionId
	}
	return 0
}

func (x *PublicBetResponse) GetCategoryKey() string {
	if x != nil {
		return x.CategoryKey
	}
	return ""
}

func (x *PublicBetResponse) GetCustomerReference() string {
	if x != nil {
		return x.CustomerReference
	}
	return ""
}

func (x *PublicBetResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

// Public Bet History Response
// swagger:model PublicBetHistoryResponse
type PublicBetHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Bets
	//
	// required: true
	Bets []*PublicBetResponse `protobuf:"bytes,1,rep,name=bets,proto3" json:"bets,omitempty"`
	// Total number of bets used for pagination
	//
	// required: true
	// example: 654
	TotalBets int64 `protobuf:"varint,2,opt,name=total_bets,json=totalBets,proto3" json:"total_bets,omitempty"`
}

func (x *PublicBetHistoryResponse) Reset() {
	*x = PublicBetHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudbet_trading_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicBetHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicBetHistoryResponse) ProtoMessage() {}

func (x *PublicBetHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloudbet_trading_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicBetHistoryResponse.ProtoReflect.Descriptor instead.
func (*PublicBetHistoryResponse) Descriptor() ([]byte, []int) {
	return file_cloudbet_trading_proto_rawDescGZIP(), []int{2}
}

func (x *PublicBetHistoryResponse) GetBets() []*PublicBetResponse {
	if x != nil {
		return x.Bets
	}
	return nil
}

func (x *PublicBetHistoryResponse) GetTotalBets() int64 {
	if x != nil {
		return x.TotalBets
	}
	return 0
}

var File_cloudbet_trading_proto protoreflect.FileDescriptor

var file_cloudbet_trading_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62, 0x65, 0x74, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62,
	0x65, 0x74, 0x1a, 0x17, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62, 0x65, 0x74, 0x2f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x02, 0x0a, 0x10,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x42, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x04, 0x73, 0x69, 0x64,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62,
	0x65, 0x74, 0x2e, 0x53, 0x69, 0x64, 0x65, 0x52, 0x04, 0x73, 0x69, 0x64, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x4b, 0x0a, 0x13, 0x61, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62, 0x65,
	0x74, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x11, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x99, 0x04, 0x0a, 0x11, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x42, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x12,
	0x22, 0x0a, 0x04, 0x73, 0x69, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62, 0x65, 0x74, 0x2e, 0x53, 0x69, 0x64, 0x65, 0x52, 0x04, 0x73,
	0x69, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x73, 0x74, 0x61, 0x6b, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62, 0x65,
	0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73,
	0x4b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x63, 0x6f, 0x6d,
	0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x2d, 0x0a,
	0x12, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x6a, 0x0a, 0x18, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x42, 0x65, 0x74, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f,
	0x0a, 0x04, 0x62, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x62, 0x65, 0x74, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x42, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x62, 0x65, 0x74, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x65, 0x74, 0x73, 0x2a, 0x19,
	0x0a, 0x04, 0x53, 0x69, 0x64, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x41, 0x43, 0x4b, 0x10, 0x00,
	0x12, 0x07, 0x0a, 0x03, 0x4c, 0x41, 0x59, 0x10, 0x01, 0x2a, 0x32, 0x0a, 0x11, 0x41, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x08,
	0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4c, 0x4c, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x45, 0x54, 0x54, 0x45, 0x52, 0x10, 0x02, 0x42, 0x11, 0x5a,
	0x0f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62, 0x65, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloudbet_trading_proto_rawDescOnce sync.Once
	file_cloudbet_trading_proto_rawDescData = file_cloudbet_trading_proto_rawDesc
)

func file_cloudbet_trading_proto_rawDescGZIP() []byte {
	file_cloudbet_trading_proto_rawDescOnce.Do(func() {
		file_cloudbet_trading_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudbet_trading_proto_rawDescData)
	})
	return file_cloudbet_trading_proto_rawDescData
}

var file_cloudbet_trading_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cloudbet_trading_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cloudbet_trading_proto_goTypes = []interface{}{
	(Side)(0),                        // 0: cloudbet.Side
	(AcceptPriceChange)(0),           // 1: cloudbet.AcceptPriceChange
	(*PublicBetRequest)(nil),         // 2: cloudbet.PublicBetRequest
	(*PublicBetResponse)(nil),        // 3: cloudbet.PublicBetResponse
	(*PublicBetHistoryResponse)(nil), // 4: cloudbet.PublicBetHistoryResponse
	(Status)(0),                      // 5: cloudbet.Status
}
var file_cloudbet_trading_proto_depIdxs = []int32{
	0, // 0: cloudbet.PublicBetRequest.side:type_name -> cloudbet.Side
	1, // 1: cloudbet.PublicBetRequest.accept_price_change:type_name -> cloudbet.AcceptPriceChange
	0, // 2: cloudbet.PublicBetResponse.side:type_name -> cloudbet.Side
	5, // 3: cloudbet.PublicBetResponse.status:type_name -> cloudbet.Status
	3, // 4: cloudbet.PublicBetHistoryResponse.bets:type_name -> cloudbet.PublicBetResponse
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_cloudbet_trading_proto_init() }
func file_cloudbet_trading_proto_init() {
	if File_cloudbet_trading_proto != nil {
		return
	}
	file_cloudbet_response_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cloudbet_trading_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicBetRequest); i {
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
		file_cloudbet_trading_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicBetResponse); i {
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
		file_cloudbet_trading_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicBetHistoryResponse); i {
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
			RawDescriptor: file_cloudbet_trading_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloudbet_trading_proto_goTypes,
		DependencyIndexes: file_cloudbet_trading_proto_depIdxs,
		EnumInfos:         file_cloudbet_trading_proto_enumTypes,
		MessageInfos:      file_cloudbet_trading_proto_msgTypes,
	}.Build()
	File_cloudbet_trading_proto = out.File
	file_cloudbet_trading_proto_rawDesc = nil
	file_cloudbet_trading_proto_goTypes = nil
	file_cloudbet_trading_proto_depIdxs = nil
}