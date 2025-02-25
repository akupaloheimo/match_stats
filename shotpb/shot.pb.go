// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v3.12.4
// source: shot.proto

package shotpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Outcome int32

const (
	Outcome_UNDEFINED Outcome = 0
	Outcome_IN        Outcome = 1
	Outcome_OUT       Outcome = 2
	Outcome_NET       Outcome = 3
	Outcome_LET       Outcome = 4
)

// Enum value maps for Outcome.
var (
	Outcome_name = map[int32]string{
		0: "UNDEFINED",
		1: "IN",
		2: "OUT",
		3: "NET",
		4: "LET",
	}
	Outcome_value = map[string]int32{
		"UNDEFINED": 0,
		"IN":        1,
		"OUT":       2,
		"NET":       3,
		"LET":       4,
	}
)

func (x Outcome) Enum() *Outcome {
	p := new(Outcome)
	*p = x
	return p
}

func (x Outcome) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Outcome) Descriptor() protoreflect.EnumDescriptor {
	return file_shot_proto_enumTypes[0].Descriptor()
}

func (Outcome) Type() protoreflect.EnumType {
	return &file_shot_proto_enumTypes[0]
}

func (x Outcome) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Outcome.Descriptor instead.
func (Outcome) EnumDescriptor() ([]byte, []int) {
	return file_shot_proto_rawDescGZIP(), []int{0}
}

type Shot struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The player who made the shot
	PlayerId string `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	// The outcome of the shot
	ShotOutcome Outcome `protobuf:"varint,2,opt,name=shot_outcome,json=shotOutcome,proto3,enum=Outcome" json:"shot_outcome,omitempty"`
	// Order in rally.
	OrderInRally int64 `protobuf:"varint,3,opt,name=order_in_rally,json=orderInRally,proto3" json:"order_in_rally,omitempty"`
	// Running count of rallies. Failed serve counts as a rally.
	RallyIndex    int64 `protobuf:"varint,4,opt,name=rally_index,json=rallyIndex,proto3" json:"rally_index,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Shot) Reset() {
	*x = Shot{}
	mi := &file_shot_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Shot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shot) ProtoMessage() {}

func (x *Shot) ProtoReflect() protoreflect.Message {
	mi := &file_shot_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shot.ProtoReflect.Descriptor instead.
func (*Shot) Descriptor() ([]byte, []int) {
	return file_shot_proto_rawDescGZIP(), []int{0}
}

func (x *Shot) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *Shot) GetShotOutcome() Outcome {
	if x != nil {
		return x.ShotOutcome
	}
	return Outcome_UNDEFINED
}

func (x *Shot) GetOrderInRally() int64 {
	if x != nil {
		return x.OrderInRally
	}
	return 0
}

func (x *Shot) GetRallyIndex() int64 {
	if x != nil {
		return x.RallyIndex
	}
	return 0
}

type ShotList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Shots         []*Shot                `protobuf:"bytes,1,rep,name=shots,proto3" json:"shots,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShotList) Reset() {
	*x = ShotList{}
	mi := &file_shot_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShotList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShotList) ProtoMessage() {}

func (x *ShotList) ProtoReflect() protoreflect.Message {
	mi := &file_shot_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShotList.ProtoReflect.Descriptor instead.
func (*ShotList) Descriptor() ([]byte, []int) {
	return file_shot_proto_rawDescGZIP(), []int{1}
}

func (x *ShotList) GetShots() []*Shot {
	if x != nil {
		return x.Shots
	}
	return nil
}

var File_shot_proto protoreflect.FileDescriptor

var file_shot_proto_rawDesc = string([]byte{
	0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a,
	0x04, 0x53, 0x68, 0x6f, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x2b, 0x0a, 0x0c, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x63, 0x6f,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f,
	0x6d, 0x65, 0x52, 0x0b, 0x73, 0x68, 0x6f, 0x74, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12,
	0x24, 0x0a, 0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x5f, 0x72, 0x61, 0x6c, 0x6c,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x6e,
	0x52, 0x61, 0x6c, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x61, 0x6c, 0x6c, 0x79, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x61, 0x6c, 0x6c,
	0x79, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x27, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x05, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x05, 0x2e, 0x53, 0x68, 0x6f, 0x74, 0x52, 0x05, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x2a,
	0x3b, 0x0a, 0x07, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e,
	0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x49, 0x4e, 0x10,
	0x01, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x55, 0x54, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x45,
	0x54, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x4c, 0x45, 0x54, 0x10, 0x04, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x73, 0x68, 0x6f, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_shot_proto_rawDescOnce sync.Once
	file_shot_proto_rawDescData []byte
)

func file_shot_proto_rawDescGZIP() []byte {
	file_shot_proto_rawDescOnce.Do(func() {
		file_shot_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_shot_proto_rawDesc), len(file_shot_proto_rawDesc)))
	})
	return file_shot_proto_rawDescData
}

var file_shot_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_shot_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shot_proto_goTypes = []any{
	(Outcome)(0),     // 0: Outcome
	(*Shot)(nil),     // 1: Shot
	(*ShotList)(nil), // 2: ShotList
}
var file_shot_proto_depIdxs = []int32{
	0, // 0: Shot.shot_outcome:type_name -> Outcome
	1, // 1: ShotList.shots:type_name -> Shot
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_shot_proto_init() }
func file_shot_proto_init() {
	if File_shot_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_shot_proto_rawDesc), len(file_shot_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shot_proto_goTypes,
		DependencyIndexes: file_shot_proto_depIdxs,
		EnumInfos:         file_shot_proto_enumTypes,
		MessageInfos:      file_shot_proto_msgTypes,
	}.Build()
	File_shot_proto = out.File
	file_shot_proto_goTypes = nil
	file_shot_proto_depIdxs = nil
}
