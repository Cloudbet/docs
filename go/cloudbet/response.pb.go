// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.5
// source: cloudbet/response.proto

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

type Status int32

const (
	// unexpected error at server side. Our engineering team is informed of the issue, please try again
	Status_INTERNAL_SERVER_ERROR Status = 0
	// duplicated request with same Reference ID was posted, this is due to idempotent request handling. If you want to resubmit this bet, please add a new Reference ID
	Status_DUPLICATE_REQUEST Status = 1
	// the request was not sent as per the expected request structure
	Status_MALFORMED_REQUEST Status = 2
	// bet price requested was above the current market price
	Status_PRICE_ABOVE_MARKET Status = 3
	// account doesn't have sufficient funds in the requested currency
	Status_INSUFFICIENT_FUNDS Status = 29
	// stake requested was above the current maximum stake on a selection
	Status_STAKE_ABOVE_MAX Status = 5
	// stake requested was below the current minimum stake on a selection
	Status_STAKE_BELOW_MIN Status = 6
	// your current liability limit on this event was exceeded, please try again later as limits are raised towards start time and during live betting
	Status_LIABILITY_LIMIT_EXCEEDED Status = 7
	// you attempted to bet on an inactive selection
	Status_MARKET_SUSPENDED Status = 8
	// your bet was accepted successfully
	Status_ACCEPTED Status = 9
	// your bet is being processed by the system. Please check the bet status again periodically to get bet status updates
	Status_PENDING_ACCEPTANCE Status = 10
	// your current account settings don't allow you to bet on this event. Restrictions will be lifted automatically as your account attains tenure and trust. Please contact customer support if you believe you qualify and we will review your account.
	Status_RESTRICTED Status = 12
	// your account needs to be verified using our KYC procedures. Please contact customer support for more details.
	Status_VERIFICATION_REQUIRED Status = 27
	// you won the bet
	Status_WIN Status = 20
	// you lost the bet
	Status_LOSS Status = 21
	// market not applicable to result, e.g. draw on 2way, handicap
	Status_PUSH Status = 22
	// half win, e.g. on a handicap market
	Status_HALF_WIN Status = 23
	// half loss, e.g. on a handicap market
	Status_HALF_LOSS Status = 24
	// partial win, including dead heat result
	Status_PARTIAL Status = 25
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0:  "INTERNAL_SERVER_ERROR",
		1:  "DUPLICATE_REQUEST",
		2:  "MALFORMED_REQUEST",
		3:  "PRICE_ABOVE_MARKET",
		29: "INSUFFICIENT_FUNDS",
		5:  "STAKE_ABOVE_MAX",
		6:  "STAKE_BELOW_MIN",
		7:  "LIABILITY_LIMIT_EXCEEDED",
		8:  "MARKET_SUSPENDED",
		9:  "ACCEPTED",
		10: "PENDING_ACCEPTANCE",
		12: "RESTRICTED",
		27: "VERIFICATION_REQUIRED",
		20: "WIN",
		21: "LOSS",
		22: "PUSH",
		23: "HALF_WIN",
		24: "HALF_LOSS",
		25: "PARTIAL",
	}
	Status_value = map[string]int32{
		"INTERNAL_SERVER_ERROR":    0,
		"DUPLICATE_REQUEST":        1,
		"MALFORMED_REQUEST":        2,
		"PRICE_ABOVE_MARKET":       3,
		"INSUFFICIENT_FUNDS":       29,
		"STAKE_ABOVE_MAX":          5,
		"STAKE_BELOW_MIN":          6,
		"LIABILITY_LIMIT_EXCEEDED": 7,
		"MARKET_SUSPENDED":         8,
		"ACCEPTED":                 9,
		"PENDING_ACCEPTANCE":       10,
		"RESTRICTED":               12,
		"VERIFICATION_REQUIRED":    27,
		"WIN":                      20,
		"LOSS":                     21,
		"PUSH":                     22,
		"HALF_WIN":                 23,
		"HALF_LOSS":                24,
		"PARTIAL":                  25,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_cloudbet_response_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_cloudbet_response_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_cloudbet_response_proto_rawDescGZIP(), []int{0}
}

// Error presents response errors for all error status codes.
// swagger:model Error
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Error Status
	// example: MALFORMED_REQUEST
	Status Status `protobuf:"varint,9,opt,name=status,proto3,enum=cloudbet.Status" json:"status,omitempty"`
	// Additional Error Details, if applicable for a given error
	// example: empty request body
	Error string `protobuf:"bytes,17,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudbet_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_cloudbet_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_cloudbet_response_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_INTERNAL_SERVER_ERROR
}

func (x *Error) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_cloudbet_response_proto protoreflect.FileDescriptor

var file_cloudbet_response_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62, 0x65, 0x74, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x62, 0x65, 0x74, 0x22, 0x47, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x62, 0x65, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0xf7, 0x02, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e, 0x54, 0x45, 0x52,
	0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f,
	0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x41, 0x4c,
	0x46, 0x4f, 0x52, 0x4d, 0x45, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x02,
	0x12, 0x16, 0x0a, 0x12, 0x50, 0x52, 0x49, 0x43, 0x45, 0x5f, 0x41, 0x42, 0x4f, 0x56, 0x45, 0x5f,
	0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x53, 0x55,
	0x46, 0x46, 0x49, 0x43, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x55, 0x4e, 0x44, 0x53, 0x10, 0x1d,
	0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x4b, 0x45, 0x5f, 0x41, 0x42, 0x4f, 0x56, 0x45, 0x5f,
	0x4d, 0x41, 0x58, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x4b, 0x45, 0x5f, 0x42,
	0x45, 0x4c, 0x4f, 0x57, 0x5f, 0x4d, 0x49, 0x4e, 0x10, 0x06, 0x12, 0x1c, 0x0a, 0x18, 0x4c, 0x49,
	0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x5f, 0x45, 0x58,
	0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x07, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x41, 0x52, 0x4b,
	0x45, 0x54, 0x5f, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x08, 0x12, 0x0c,
	0x0a, 0x08, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x09, 0x12, 0x16, 0x0a, 0x12,
	0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x41, 0x4e,
	0x43, 0x45, 0x10, 0x0a, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x45, 0x53, 0x54, 0x52, 0x49, 0x43, 0x54,
	0x45, 0x44, 0x10, 0x0c, 0x12, 0x19, 0x0a, 0x15, 0x56, 0x45, 0x52, 0x49, 0x46, 0x49, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x1b, 0x12,
	0x07, 0x0a, 0x03, 0x57, 0x49, 0x4e, 0x10, 0x14, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x4f, 0x53, 0x53,
	0x10, 0x15, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x55, 0x53, 0x48, 0x10, 0x16, 0x12, 0x0c, 0x0a, 0x08,
	0x48, 0x41, 0x4c, 0x46, 0x5f, 0x57, 0x49, 0x4e, 0x10, 0x17, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x41,
	0x4c, 0x46, 0x5f, 0x4c, 0x4f, 0x53, 0x53, 0x10, 0x18, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x41, 0x52,
	0x54, 0x49, 0x41, 0x4c, 0x10, 0x19, 0x42, 0x11, 0x5a, 0x0f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x62, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cloudbet_response_proto_rawDescOnce sync.Once
	file_cloudbet_response_proto_rawDescData = file_cloudbet_response_proto_rawDesc
)

func file_cloudbet_response_proto_rawDescGZIP() []byte {
	file_cloudbet_response_proto_rawDescOnce.Do(func() {
		file_cloudbet_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudbet_response_proto_rawDescData)
	})
	return file_cloudbet_response_proto_rawDescData
}

var file_cloudbet_response_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cloudbet_response_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cloudbet_response_proto_goTypes = []interface{}{
	(Status)(0),   // 0: cloudbet.Status
	(*Error)(nil), // 1: cloudbet.Error
}
var file_cloudbet_response_proto_depIdxs = []int32{
	0, // 0: cloudbet.Error.status:type_name -> cloudbet.Status
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cloudbet_response_proto_init() }
func file_cloudbet_response_proto_init() {
	if File_cloudbet_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloudbet_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_cloudbet_response_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloudbet_response_proto_goTypes,
		DependencyIndexes: file_cloudbet_response_proto_depIdxs,
		EnumInfos:         file_cloudbet_response_proto_enumTypes,
		MessageInfos:      file_cloudbet_response_proto_msgTypes,
	}.Build()
	File_cloudbet_response_proto = out.File
	file_cloudbet_response_proto_rawDesc = nil
	file_cloudbet_response_proto_goTypes = nil
	file_cloudbet_response_proto_depIdxs = nil
}