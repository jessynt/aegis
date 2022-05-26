// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: internal/proto/model.proto

package proto

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PropertyType int32

const (
	PropertyTypeUnknown  PropertyType = 0
	PropertyTypeInteger  PropertyType = 1
	PropertyTypeDouble   PropertyType = 2
	PropertyTypeString   PropertyType = 3
	PropertyTypeBool     PropertyType = 4
	PropertyTypeDate     PropertyType = 5
	PropertyTypeDateTime PropertyType = 6
)

var PropertyType_name = map[int32]string{
	0: "PropertyTypeUnknown",
	1: "PropertyTypeInteger",
	2: "PropertyTypeDouble",
	3: "PropertyTypeString",
	4: "PropertyTypeBool",
	5: "PropertyTypeDate",
	6: "PropertyTypeDateTime",
}

var PropertyType_value = map[string]int32{
	"PropertyTypeUnknown":  0,
	"PropertyTypeInteger":  1,
	"PropertyTypeDouble":   2,
	"PropertyTypeString":   3,
	"PropertyTypeBool":     4,
	"PropertyTypeDate":     5,
	"PropertyTypeDateTime": 6,
}

func (PropertyType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f221f5c250c1a5e4, []int{0}
}

type AggregateType int32

const (
	AggregateTypeUnknown       AggregateType = 0
	AggregateTypeCount         AggregateType = 1
	AggregateTypeDistinctCount AggregateType = 2
	AggregateTypeAverage       AggregateType = 3
	AggregateTypeSum           AggregateType = 4
	AggregateTypeMin           AggregateType = 5
	AggregateTypeMax           AggregateType = 6
	AggregateTypeSD            AggregateType = 7
	AggregateTypeVariance      AggregateType = 8
	AggregateTypeDeviation     AggregateType = 9
	AggregateTypeMedian        AggregateType = 10
)

var AggregateType_name = map[int32]string{
	0:  "AggregateTypeUnknown",
	1:  "AggregateTypeCount",
	2:  "AggregateTypeDistinctCount",
	3:  "AggregateTypeAverage",
	4:  "AggregateTypeSum",
	5:  "AggregateTypeMin",
	6:  "AggregateTypeMax",
	7:  "AggregateTypeSD",
	8:  "AggregateTypeVariance",
	9:  "AggregateTypeDeviation",
	10: "AggregateTypeMedian",
}

var AggregateType_value = map[string]int32{
	"AggregateTypeUnknown":       0,
	"AggregateTypeCount":         1,
	"AggregateTypeDistinctCount": 2,
	"AggregateTypeAverage":       3,
	"AggregateTypeSum":           4,
	"AggregateTypeMin":           5,
	"AggregateTypeMax":           6,
	"AggregateTypeSD":            7,
	"AggregateTypeVariance":      8,
	"AggregateTypeDeviation":     9,
	"AggregateTypeMedian":        10,
}

func (AggregateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f221f5c250c1a5e4, []int{1}
}

type AggregateIntervalType int32

const (
	AggregateIntervalTypeUnknown AggregateIntervalType = 0
	AggregateIntervalTypeYear    AggregateIntervalType = 1
	AggregateIntervalTypeMonth   AggregateIntervalType = 2
	AggregateIntervalTypeDay     AggregateIntervalType = 3
	AggregateIntervalTypeHour    AggregateIntervalType = 4
	AggregateIntervalTypeMinute  AggregateIntervalType = 5
	AggregateIntervalTypeSecond  AggregateIntervalType = 6
)

var AggregateIntervalType_name = map[int32]string{
	0: "AggregateIntervalTypeUnknown",
	1: "AggregateIntervalTypeYear",
	2: "AggregateIntervalTypeMonth",
	3: "AggregateIntervalTypeDay",
	4: "AggregateIntervalTypeHour",
	5: "AggregateIntervalTypeMinute",
	6: "AggregateIntervalTypeSecond",
}

var AggregateIntervalType_value = map[string]int32{
	"AggregateIntervalTypeUnknown": 0,
	"AggregateIntervalTypeYear":    1,
	"AggregateIntervalTypeMonth":   2,
	"AggregateIntervalTypeDay":     3,
	"AggregateIntervalTypeHour":    4,
	"AggregateIntervalTypeMinute":  5,
	"AggregateIntervalTypeSecond":  6,
}

func (AggregateIntervalType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f221f5c250c1a5e4, []int{2}
}

type OperatorType int32

const (
	OperatorTypeUnknown OperatorType = 0
	OperatorTypeAdd     OperatorType = 1
	OperatorTypeSub     OperatorType = 2
	OperatorTypeMul     OperatorType = 3
	OperatorTypeDiv     OperatorType = 4
)

var OperatorType_name = map[int32]string{
	0: "OperatorTypeUnknown",
	1: "OperatorTypeAdd",
	2: "OperatorTypeSub",
	3: "OperatorTypeMul",
	4: "OperatorTypeDiv",
}

var OperatorType_value = map[string]int32{
	"OperatorTypeUnknown": 0,
	"OperatorTypeAdd":     1,
	"OperatorTypeSub":     2,
	"OperatorTypeMul":     3,
	"OperatorTypeDiv":     4,
}

func (OperatorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f221f5c250c1a5e4, []int{3}
}

type ValidateType int32

const (
	ValidateTypeUnknown  ValidateType = 0
	ValidateTypeNumber   ValidateType = 1
	ValidateTypeString   ValidateType = 2
	ValidateTypeBool     ValidateType = 3
	ValidateTypeRegex    ValidateType = 4
	ValidateTypeDatetime ValidateType = 5
)

var ValidateType_name = map[int32]string{
	0: "ValidateTypeUnknown",
	1: "ValidateTypeNumber",
	2: "ValidateTypeString",
	3: "ValidateTypeBool",
	4: "ValidateTypeRegex",
	5: "ValidateTypeDatetime",
}

var ValidateType_value = map[string]int32{
	"ValidateTypeUnknown":  0,
	"ValidateTypeNumber":   1,
	"ValidateTypeString":   2,
	"ValidateTypeBool":     3,
	"ValidateTypeRegex":    4,
	"ValidateTypeDatetime": 5,
}

func (ValidateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f221f5c250c1a5e4, []int{4}
}

type RiskType int32

const (
	RiskUnknown RiskType = 0
	RiskPass    RiskType = 1
	RiskWarning RiskType = 2
	RiskDeny    RiskType = 3
)

var RiskType_name = map[int32]string{
	0: "RiskUnknown",
	1: "RiskPass",
	2: "RiskWarning",
	3: "RiskDeny",
}

var RiskType_value = map[string]int32{
	"RiskUnknown": 0,
	"RiskPass":    1,
	"RiskWarning": 2,
	"RiskDeny":    3,
}

func (RiskType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f221f5c250c1a5e4, []int{5}
}

type CollectionType int32

const (
	CollectionTypeUnknown   CollectionType = 0
	CollectionTypeBlackList CollectionType = 1
	CollectionTypeWhiteList CollectionType = 2
)

var CollectionType_name = map[int32]string{
	0: "CollectionTypeUnknown",
	1: "CollectionTypeBlackList",
	2: "CollectionTypeWhiteList",
}

var CollectionType_value = map[string]int32{
	"CollectionTypeUnknown":   0,
	"CollectionTypeBlackList": 1,
	"CollectionTypeWhiteList": 2,
}

func (CollectionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f221f5c250c1a5e4, []int{6}
}

type RiskObject struct {
	RickType RiskType `protobuf:"varint,1,opt,name=rick_type,json=rickType,proto3,enum=aegis.proto.RiskType" json:"rick_type,omitempty"`
	Score    float64  `protobuf:"fixed64,2,opt,name=score,proto3" json:"score,omitempty"`
}

func (m *RiskObject) Reset()      { *m = RiskObject{} }
func (*RiskObject) ProtoMessage() {}
func (*RiskObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_f221f5c250c1a5e4, []int{0}
}
func (m *RiskObject) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RiskObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RiskObject.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RiskObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RiskObject.Merge(m, src)
}
func (m *RiskObject) XXX_Size() int {
	return m.Size()
}
func (m *RiskObject) XXX_DiscardUnknown() {
	xxx_messageInfo_RiskObject.DiscardUnknown(m)
}

var xxx_messageInfo_RiskObject proto.InternalMessageInfo

func (m *RiskObject) GetRickType() RiskType {
	if m != nil {
		return m.RickType
	}
	return RiskUnknown
}

func (m *RiskObject) GetScore() float64 {
	if m != nil {
		return m.Score
	}
	return 0
}

func init() {
	proto.RegisterEnum("aegis.proto.PropertyType", PropertyType_name, PropertyType_value)
	proto.RegisterEnum("aegis.proto.AggregateType", AggregateType_name, AggregateType_value)
	proto.RegisterEnum("aegis.proto.AggregateIntervalType", AggregateIntervalType_name, AggregateIntervalType_value)
	proto.RegisterEnum("aegis.proto.OperatorType", OperatorType_name, OperatorType_value)
	proto.RegisterEnum("aegis.proto.ValidateType", ValidateType_name, ValidateType_value)
	proto.RegisterEnum("aegis.proto.RiskType", RiskType_name, RiskType_value)
	proto.RegisterEnum("aegis.proto.CollectionType", CollectionType_name, CollectionType_value)
	proto.RegisterType((*RiskObject)(nil), "aegis.proto.RiskObject")
}

func init() { proto.RegisterFile("internal/proto/model.proto", fileDescriptor_f221f5c250c1a5e4) }

var fileDescriptor_f221f5c250c1a5e4 = []byte{
	// 650 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xbd, 0x52, 0xdb, 0x40,
	0x10, 0x80, 0x7d, 0xf2, 0x4f, 0x60, 0x71, 0xe0, 0x72, 0xb6, 0xc1, 0xfc, 0x44, 0x61, 0x52, 0x31,
	0x2e, 0x60, 0x86, 0x3c, 0x01, 0xa0, 0x82, 0xcc, 0x84, 0xc0, 0xd8, 0x04, 0x26, 0x69, 0x32, 0x67,
	0x79, 0x47, 0x5c, 0x2c, 0xdf, 0x79, 0xce, 0x27, 0x07, 0x4f, 0x9a, 0x3c, 0x42, 0x9e, 0x20, 0x45,
	0xaa, 0x3c, 0x41, 0x9e, 0x21, 0x25, 0x25, 0x65, 0x30, 0x4d, 0x4a, 0xba, 0xb4, 0x19, 0x49, 0x78,
	0x2c, 0x09, 0x53, 0x59, 0xfb, 0x7d, 0x3b, 0xeb, 0xbd, 0xd3, 0xae, 0x60, 0x4d, 0x48, 0x83, 0x5a,
	0x72, 0x7f, 0xa7, 0xaf, 0x95, 0x51, 0x3b, 0x3d, 0xd5, 0x41, 0x7f, 0x3b, 0x7a, 0x66, 0x0b, 0x1c,
	0x3d, 0x31, 0x88, 0x83, 0x97, 0x67, 0x00, 0x4d, 0x31, 0xe8, 0x1e, 0xb7, 0x3f, 0xa1, 0x6b, 0xd8,
	0x2e, 0xcc, 0x6b, 0xe1, 0x76, 0x3f, 0x9a, 0x51, 0x1f, 0xeb, 0x64, 0x93, 0x6c, 0x2d, 0xee, 0xd6,
	0xb6, 0x13, 0xe9, 0xdb, 0x61, 0xee, 0xe9, 0xa8, 0x8f, 0xcd, 0xb9, 0x30, 0x2f, 0x7c, 0x62, 0x55,
	0x28, 0x0e, 0x5c, 0xa5, 0xb1, 0x6e, 0x6d, 0x92, 0x2d, 0xd2, 0x8c, 0x83, 0xc6, 0x2f, 0x02, 0xe5,
	0x13, 0xad, 0xfa, 0xa8, 0xcd, 0x28, 0x4a, 0x5b, 0x81, 0x4a, 0x32, 0x7e, 0x27, 0xbb, 0x52, 0x7d,
	0x96, 0x34, 0x97, 0x15, 0xaf, 0xa5, 0x41, 0x0f, 0x35, 0x25, 0x6c, 0x19, 0x58, 0x52, 0x38, 0x2a,
	0x68, 0xfb, 0x48, 0xad, 0x2c, 0x6f, 0x19, 0x2d, 0xa4, 0x47, 0xf3, 0xac, 0x0a, 0x34, 0xc9, 0xf7,
	0x95, 0xf2, 0x69, 0x21, 0x4b, 0x1d, 0x6e, 0x90, 0x16, 0x59, 0x1d, 0xaa, 0x59, 0x7a, 0x2a, 0x7a,
	0x48, 0x4b, 0x8d, 0x1f, 0x16, 0x3c, 0xdd, 0xf3, 0x3c, 0x8d, 0x5e, 0xc8, 0xc2, 0xce, 0xeb, 0x50,
	0x4d, 0x81, 0x69, 0xeb, 0xcb, 0xc0, 0x52, 0xe6, 0x40, 0x05, 0xd2, 0x50, 0xc2, 0x6c, 0x58, 0x4b,
	0x71, 0x47, 0x0c, 0x8c, 0x90, 0xae, 0x89, 0xbd, 0xf5, 0xa0, 0xe2, 0xde, 0x10, 0x35, 0xf7, 0x30,
	0x3e, 0x43, 0xca, 0xb4, 0x82, 0x5e, 0x7c, 0x86, 0x14, 0x3d, 0x12, 0x92, 0x16, 0x1f, 0x52, 0x7e,
	0x49, 0x4b, 0xac, 0x02, 0x4b, 0xe9, 0x0a, 0x0e, 0x7d, 0xc2, 0x56, 0xa1, 0x96, 0x82, 0x67, 0x5c,
	0x0b, 0x2e, 0x5d, 0xa4, 0x73, 0x6c, 0x0d, 0x96, 0xd3, 0xbd, 0xe2, 0x50, 0x70, 0x23, 0x94, 0xa4,
	0xf3, 0xe1, 0xab, 0x49, 0xff, 0x03, 0x76, 0x04, 0x97, 0x14, 0x1a, 0xff, 0x48, 0xa2, 0x60, 0xf8,
	0xc6, 0xf4, 0x90, 0xfb, 0xd1, 0x65, 0x6d, 0xc2, 0xc6, 0x4c, 0x31, 0xbd, 0xb4, 0xe7, 0xb0, 0x3a,
	0x33, 0xe3, 0x3d, 0x72, 0x9d, 0xb9, 0xbb, 0xa4, 0x3e, 0x52, 0xd2, 0x5c, 0x50, 0x8b, 0x6d, 0x40,
	0x7d, 0xa6, 0x77, 0xf8, 0x88, 0xe6, 0x1f, 0x2d, 0x7e, 0xa8, 0x02, 0x4d, 0x0b, 0xec, 0x05, 0xac,
	0xcf, 0x2e, 0x2e, 0x64, 0x10, 0xcd, 0xc5, 0x63, 0x09, 0x2d, 0x74, 0x95, 0xec, 0xd0, 0x52, 0xe3,
	0x0b, 0x94, 0x8f, 0xfb, 0xa8, 0xb9, 0x51, 0x7a, 0x32, 0xd6, 0xc9, 0x78, 0x7a, 0xcc, 0x0a, 0x2c,
	0x25, 0xc5, 0x5e, 0xa7, 0x43, 0x49, 0x16, 0xb6, 0x82, 0x36, 0xb5, 0xb2, 0xf0, 0x28, 0xf0, 0x69,
	0x3e, 0x0b, 0x1d, 0x31, 0xa4, 0x85, 0xc6, 0x77, 0x02, 0xe5, 0x33, 0xee, 0x8b, 0xce, 0x64, 0x34,
	0x57, 0xa0, 0x92, 0x8c, 0x53, 0x93, 0x99, 0x14, 0x6f, 0x83, 0x5e, 0x7b, 0xb2, 0x53, 0x49, 0x7e,
	0xbf, 0x3b, 0x56, 0x38, 0x4b, 0x49, 0x1e, 0xed, 0x4e, 0x9e, 0xd5, 0xe0, 0x59, 0x92, 0x36, 0xd1,
	0xc3, 0x4b, 0x5a, 0x08, 0xc7, 0x37, 0x89, 0xc3, 0xe5, 0x31, 0xe1, 0xf2, 0x14, 0x1b, 0x87, 0x30,
	0x37, 0xf9, 0x42, 0xb0, 0x25, 0x58, 0x08, 0x9f, 0xa7, 0x3d, 0x95, 0x63, 0x79, 0xc2, 0x07, 0x03,
	0x4a, 0x26, 0xfa, 0x9c, 0x6b, 0x19, 0xb7, 0x70, 0xaf, 0x1d, 0x94, 0x23, 0x9a, 0x6f, 0x20, 0x2c,
	0x1e, 0x28, 0xdf, 0x47, 0x37, 0x1c, 0xc5, 0xa8, 0xde, 0x2a, 0xd4, 0xd2, 0x64, 0x5a, 0x79, 0x1d,
	0x56, 0xd2, 0x6a, 0xdf, 0xe7, 0x6e, 0xf7, 0x8d, 0x18, 0x84, 0xcb, 0xf8, 0x40, 0x9e, 0x5f, 0x08,
	0x83, 0x91, 0xb4, 0xf6, 0x9d, 0xab, 0x1b, 0x3b, 0x77, 0x7d, 0x63, 0xe7, 0xee, 0x6e, 0x6c, 0xf2,
	0x75, 0x6c, 0x93, 0x9f, 0x63, 0x9b, 0xfc, 0x1e, 0xdb, 0xe4, 0x6a, 0x6c, 0x93, 0x3f, 0x63, 0x9b,
	0xfc, 0x1d, 0xdb, 0xb9, 0xbb, 0xb1, 0x4d, 0xbe, 0xdd, 0xda, 0xb9, 0xab, 0x5b, 0x3b, 0x77, 0x7d,
	0x6b, 0xe7, 0x3e, 0x2c, 0xa6, 0xbf, 0xae, 0xed, 0x52, 0xf4, 0xf3, 0xea, 0x7f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x64, 0xe4, 0x7b, 0x11, 0x76, 0x05, 0x00, 0x00,
}

func (x PropertyType) String() string {
	s, ok := PropertyType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x AggregateType) String() string {
	s, ok := AggregateType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x AggregateIntervalType) String() string {
	s, ok := AggregateIntervalType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x OperatorType) String() string {
	s, ok := OperatorType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ValidateType) String() string {
	s, ok := ValidateType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x RiskType) String() string {
	s, ok := RiskType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x CollectionType) String() string {
	s, ok := CollectionType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *RiskObject) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RiskObject)
	if !ok {
		that2, ok := that.(RiskObject)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.RickType != that1.RickType {
		return false
	}
	if this.Score != that1.Score {
		return false
	}
	return true
}
func (this *RiskObject) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&proto.RiskObject{")
	s = append(s, "RickType: "+fmt.Sprintf("%#v", this.RickType)+",\n")
	s = append(s, "Score: "+fmt.Sprintf("%#v", this.Score)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringModel(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *RiskObject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RiskObject) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RiskObject) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Score != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Score))))
		i--
		dAtA[i] = 0x11
	}
	if m.RickType != 0 {
		i = encodeVarintModel(dAtA, i, uint64(m.RickType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintModel(dAtA []byte, offset int, v uint64) int {
	offset -= sovModel(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RiskObject) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RickType != 0 {
		n += 1 + sovModel(uint64(m.RickType))
	}
	if m.Score != 0 {
		n += 9
	}
	return n
}

func sovModel(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModel(x uint64) (n int) {
	return sovModel(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *RiskObject) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RiskObject{`,
		`RickType:` + fmt.Sprintf("%v", this.RickType) + `,`,
		`Score:` + fmt.Sprintf("%v", this.Score) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringModel(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *RiskObject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModel
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RiskObject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RiskObject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RickType", wireType)
			}
			m.RickType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RickType |= RiskType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Score = float64(math.Float64frombits(v))
		default:
			iNdEx = preIndex
			skippy, err := skipModel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthModel
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthModel
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipModel(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModel
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowModel
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowModel
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthModel
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModel
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModel
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModel        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModel          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModel = fmt.Errorf("proto: unexpected end of group")
)