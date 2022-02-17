// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: tf_predict.proto

package tf_predict_protos

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

type ArrayDataType int32

const (
	// Not a legal value for DataType. Used to indicate a DataType field
	// has not been set.
	ArrayDataType_DT_INVALID ArrayDataType = 0
	// Data types that all computation devices are expected to be
	// capable to support.
	ArrayDataType_DT_FLOAT      ArrayDataType = 1
	ArrayDataType_DT_DOUBLE     ArrayDataType = 2
	ArrayDataType_DT_INT32      ArrayDataType = 3
	ArrayDataType_DT_UINT8      ArrayDataType = 4
	ArrayDataType_DT_INT16      ArrayDataType = 5
	ArrayDataType_DT_INT8       ArrayDataType = 6
	ArrayDataType_DT_STRING     ArrayDataType = 7
	ArrayDataType_DT_COMPLEX64  ArrayDataType = 8 // Single-precision complex
	ArrayDataType_DT_INT64      ArrayDataType = 9
	ArrayDataType_DT_BOOL       ArrayDataType = 10
	ArrayDataType_DT_QINT8      ArrayDataType = 11 // Quantized int8
	ArrayDataType_DT_QUINT8     ArrayDataType = 12 // Quantized uint8
	ArrayDataType_DT_QINT32     ArrayDataType = 13 // Quantized int32
	ArrayDataType_DT_BFLOAT16   ArrayDataType = 14 // Float32 truncated to 16 bits.  Only for cast ops.
	ArrayDataType_DT_QINT16     ArrayDataType = 15 // Quantized int16
	ArrayDataType_DT_QUINT16    ArrayDataType = 16 // Quantized uint16
	ArrayDataType_DT_UINT16     ArrayDataType = 17
	ArrayDataType_DT_COMPLEX128 ArrayDataType = 18 // Double-precision complex
	ArrayDataType_DT_HALF       ArrayDataType = 19
	ArrayDataType_DT_RESOURCE   ArrayDataType = 20
	ArrayDataType_DT_VARIANT    ArrayDataType = 21 // Arbitrary C++ data types
)

// Enum value maps for ArrayDataType.
var (
	ArrayDataType_name = map[int32]string{
		0:  "DT_INVALID",
		1:  "DT_FLOAT",
		2:  "DT_DOUBLE",
		3:  "DT_INT32",
		4:  "DT_UINT8",
		5:  "DT_INT16",
		6:  "DT_INT8",
		7:  "DT_STRING",
		8:  "DT_COMPLEX64",
		9:  "DT_INT64",
		10: "DT_BOOL",
		11: "DT_QINT8",
		12: "DT_QUINT8",
		13: "DT_QINT32",
		14: "DT_BFLOAT16",
		15: "DT_QINT16",
		16: "DT_QUINT16",
		17: "DT_UINT16",
		18: "DT_COMPLEX128",
		19: "DT_HALF",
		20: "DT_RESOURCE",
		21: "DT_VARIANT",
	}
	ArrayDataType_value = map[string]int32{
		"DT_INVALID":    0,
		"DT_FLOAT":      1,
		"DT_DOUBLE":     2,
		"DT_INT32":      3,
		"DT_UINT8":      4,
		"DT_INT16":      5,
		"DT_INT8":       6,
		"DT_STRING":     7,
		"DT_COMPLEX64":  8,
		"DT_INT64":      9,
		"DT_BOOL":       10,
		"DT_QINT8":      11,
		"DT_QUINT8":     12,
		"DT_QINT32":     13,
		"DT_BFLOAT16":   14,
		"DT_QINT16":     15,
		"DT_QUINT16":    16,
		"DT_UINT16":     17,
		"DT_COMPLEX128": 18,
		"DT_HALF":       19,
		"DT_RESOURCE":   20,
		"DT_VARIANT":    21,
	}
)

func (x ArrayDataType) Enum() *ArrayDataType {
	p := new(ArrayDataType)
	*p = x
	return p
}

func (x ArrayDataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ArrayDataType) Descriptor() protoreflect.EnumDescriptor {
	return file_tf_predict_proto_enumTypes[0].Descriptor()
}

func (ArrayDataType) Type() protoreflect.EnumType {
	return &file_tf_predict_proto_enumTypes[0]
}

func (x ArrayDataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ArrayDataType.Descriptor instead.
func (ArrayDataType) EnumDescriptor() ([]byte, []int) {
	return file_tf_predict_proto_rawDescGZIP(), []int{0}
}

// Dimensions of an array
type ArrayShape struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dim []int64 `protobuf:"varint,1,rep,packed,name=dim,proto3" json:"dim,omitempty"`
}

func (x *ArrayShape) Reset() {
	*x = ArrayShape{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tf_predict_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrayShape) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrayShape) ProtoMessage() {}

func (x *ArrayShape) ProtoReflect() protoreflect.Message {
	mi := &file_tf_predict_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrayShape.ProtoReflect.Descriptor instead.
func (*ArrayShape) Descriptor() ([]byte, []int) {
	return file_tf_predict_proto_rawDescGZIP(), []int{0}
}

func (x *ArrayShape) GetDim() []int64 {
	if x != nil {
		return x.Dim
	}
	return nil
}

// Protocol buffer representing an array
type ArrayProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data Type.
	Dtype ArrayDataType `protobuf:"varint,1,opt,name=dtype,proto3,enum=tf_predict_protos.ArrayDataType" json:"dtype,omitempty"`
	// Shape of the array.
	ArrayShape *ArrayShape `protobuf:"bytes,2,opt,name=array_shape,json=arrayShape,proto3" json:"array_shape,omitempty"`
	// DT_FLOAT.
	FloatVal []float32 `protobuf:"fixed32,3,rep,packed,name=float_val,json=floatVal,proto3" json:"float_val,omitempty"`
	// DT_DOUBLE.
	DoubleVal []float64 `protobuf:"fixed64,4,rep,packed,name=double_val,json=doubleVal,proto3" json:"double_val,omitempty"`
	// DT_INT32, DT_INT16, DT_INT8, DT_UINT8.
	IntVal []int32 `protobuf:"varint,5,rep,packed,name=int_val,json=intVal,proto3" json:"int_val,omitempty"`
	// DT_STRING.
	StringVal [][]byte `protobuf:"bytes,6,rep,name=string_val,json=stringVal,proto3" json:"string_val,omitempty"`
	// DT_INT64.
	Int64Val []int64 `protobuf:"varint,7,rep,packed,name=int64_val,json=int64Val,proto3" json:"int64_val,omitempty"`
	// DT_BOOL.
	BoolVal []bool `protobuf:"varint,8,rep,packed,name=bool_val,json=boolVal,proto3" json:"bool_val,omitempty"`
}

func (x *ArrayProto) Reset() {
	*x = ArrayProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tf_predict_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrayProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrayProto) ProtoMessage() {}

func (x *ArrayProto) ProtoReflect() protoreflect.Message {
	mi := &file_tf_predict_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrayProto.ProtoReflect.Descriptor instead.
func (*ArrayProto) Descriptor() ([]byte, []int) {
	return file_tf_predict_proto_rawDescGZIP(), []int{1}
}

func (x *ArrayProto) GetDtype() ArrayDataType {
	if x != nil {
		return x.Dtype
	}
	return ArrayDataType_DT_INVALID
}

func (x *ArrayProto) GetArrayShape() *ArrayShape {
	if x != nil {
		return x.ArrayShape
	}
	return nil
}

func (x *ArrayProto) GetFloatVal() []float32 {
	if x != nil {
		return x.FloatVal
	}
	return nil
}

func (x *ArrayProto) GetDoubleVal() []float64 {
	if x != nil {
		return x.DoubleVal
	}
	return nil
}

func (x *ArrayProto) GetIntVal() []int32 {
	if x != nil {
		return x.IntVal
	}
	return nil
}

func (x *ArrayProto) GetStringVal() [][]byte {
	if x != nil {
		return x.StringVal
	}
	return nil
}

func (x *ArrayProto) GetInt64Val() []int64 {
	if x != nil {
		return x.Int64Val
	}
	return nil
}

func (x *ArrayProto) GetBoolVal() []bool {
	if x != nil {
		return x.BoolVal
	}
	return nil
}

type InputEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string      `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *ArrayProto `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *InputEntry) Reset() {
	*x = InputEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tf_predict_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputEntry) ProtoMessage() {}

func (x *InputEntry) ProtoReflect() protoreflect.Message {
	mi := &file_tf_predict_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputEntry.ProtoReflect.Descriptor instead.
func (*InputEntry) Descriptor() ([]byte, []int) {
	return file_tf_predict_proto_rawDescGZIP(), []int{2}
}

func (x *InputEntry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *InputEntry) GetValue() *ArrayProto {
	if x != nil {
		return x.Value
	}
	return nil
}

// PredictRequest specifies which TensorFlow model to run, as well as
// how inputs are mapped to tensors and how outputs are filtered before
// returning to user.
type PredictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A named signature to evaluate. If unspecified, the default signature
	// will be used
	SignatureName string `protobuf:"bytes,1,opt,name=signature_name,json=signatureName,proto3" json:"signature_name,omitempty"`
	// Input tensors.
	// Names of input tensor are alias names. The mapping from aliases to real
	// input tensor names is expected to be stored as named generic signature
	// under the key "inputs" in the model export.
	// Each alias listed in a generic signature named "inputs" should be provided
	// exactly once in order to run the prediction.
	Inputs []*InputEntry `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs,omitempty"`
	// Output filter.
	// Names specified are alias names. The mapping from aliases to real output
	// tensor names is expected to be stored as named generic signature under
	// the key "outputs" in the model export.
	// Only tensors specified here will be run/fetched and returned, with the
	// exception that when none is specified, all tensors specified in the
	// named signature will be run/fetched and returned.
	OutputFilter []string `protobuf:"bytes,3,rep,name=output_filter,json=outputFilter,proto3" json:"output_filter,omitempty"`
}

func (x *PredictRequest) Reset() {
	*x = PredictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tf_predict_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictRequest) ProtoMessage() {}

func (x *PredictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tf_predict_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictRequest.ProtoReflect.Descriptor instead.
func (*PredictRequest) Descriptor() ([]byte, []int) {
	return file_tf_predict_proto_rawDescGZIP(), []int{3}
}

func (x *PredictRequest) GetSignatureName() string {
	if x != nil {
		return x.SignatureName
	}
	return ""
}

func (x *PredictRequest) GetInputs() []*InputEntry {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *PredictRequest) GetOutputFilter() []string {
	if x != nil {
		return x.OutputFilter
	}
	return nil
}

type OutputEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string      `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *ArrayProto `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *OutputEntry) Reset() {
	*x = OutputEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tf_predict_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutputEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutputEntry) ProtoMessage() {}

func (x *OutputEntry) ProtoReflect() protoreflect.Message {
	mi := &file_tf_predict_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutputEntry.ProtoReflect.Descriptor instead.
func (*OutputEntry) Descriptor() ([]byte, []int) {
	return file_tf_predict_proto_rawDescGZIP(), []int{4}
}

func (x *OutputEntry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *OutputEntry) GetValue() *ArrayProto {
	if x != nil {
		return x.Value
	}
	return nil
}

// Response for PredictRequest on successful run.
type PredictResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output tensors.
	Outputs []*OutputEntry `protobuf:"bytes,1,rep,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *PredictResponse) Reset() {
	*x = PredictResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tf_predict_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictResponse) ProtoMessage() {}

func (x *PredictResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tf_predict_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictResponse.ProtoReflect.Descriptor instead.
func (*PredictResponse) Descriptor() ([]byte, []int) {
	return file_tf_predict_proto_rawDescGZIP(), []int{5}
}

func (x *PredictResponse) GetOutputs() []*OutputEntry {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_tf_predict_proto protoreflect.FileDescriptor

var file_tf_predict_proto_rawDesc = []byte{
	0x0a, 0x10, 0x74, 0x66, 0x5f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x74, 0x66, 0x5f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x22, 0x0a, 0x0a, 0x41, 0x72, 0x72, 0x61, 0x79, 0x53, 0x68,
	0x61, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x03, 0x64, 0x69, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03,
	0x42, 0x02, 0x10, 0x01, 0x52, 0x03, 0x64, 0x69, 0x6d, 0x22, 0xc4, 0x02, 0x0a, 0x0a, 0x41, 0x72,
	0x72, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x36, 0x0a, 0x05, 0x64, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x74, 0x66, 0x5f, 0x70, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x72, 0x72, 0x61,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x3e, 0x0a, 0x0b, 0x61, 0x72, 0x72, 0x61, 0x79, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x66, 0x5f, 0x70, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x53,
	0x68, 0x61, 0x70, 0x65, 0x52, 0x0a, 0x61, 0x72, 0x72, 0x61, 0x79, 0x53, 0x68, 0x61, 0x70, 0x65,
	0x12, 0x1f, 0x0a, 0x09, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x02, 0x42, 0x02, 0x10, 0x01, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61,
	0x6c, 0x12, 0x21, 0x0a, 0x0a, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x01, 0x42, 0x02, 0x10, 0x01, 0x52, 0x09, 0x64, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x56, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x07, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x05, 0x42, 0x02, 0x10, 0x01, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x56, 0x61,
	0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x12, 0x1f, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x03, 0x42, 0x02, 0x10, 0x01, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61,
	0x6c, 0x12, 0x1d, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x08, 0x42, 0x02, 0x10, 0x01, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c,
	0x22, 0x53, 0x0a, 0x0a, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x74, 0x66, 0x5f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x35, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x74, 0x66, 0x5f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x54, 0x0a, 0x0b, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x33, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x66,
	0x5f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x4b, 0x0a, 0x0f, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x66, 0x5f, 0x70, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x2a, 0xdb,
	0x02, 0x0a, 0x0d, 0x41, 0x72, 0x72, 0x61, 0x79, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x54, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x44, 0x54, 0x5f, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10, 0x01, 0x12, 0x0d,
	0x0a, 0x09, 0x44, 0x54, 0x5f, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a,
	0x08, 0x44, 0x54, 0x5f, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x44,
	0x54, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x38, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x54, 0x5f,
	0x49, 0x4e, 0x54, 0x31, 0x36, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x54, 0x5f, 0x49, 0x4e,
	0x54, 0x38, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x54, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e,
	0x47, 0x10, 0x07, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45,
	0x58, 0x36, 0x34, 0x10, 0x08, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x54, 0x5f, 0x49, 0x4e, 0x54, 0x36,
	0x34, 0x10, 0x09, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x54, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x0a,
	0x12, 0x0c, 0x0a, 0x08, 0x44, 0x54, 0x5f, 0x51, 0x49, 0x4e, 0x54, 0x38, 0x10, 0x0b, 0x12, 0x0d,
	0x0a, 0x09, 0x44, 0x54, 0x5f, 0x51, 0x55, 0x49, 0x4e, 0x54, 0x38, 0x10, 0x0c, 0x12, 0x0d, 0x0a,
	0x09, 0x44, 0x54, 0x5f, 0x51, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x0d, 0x12, 0x0f, 0x0a, 0x0b,
	0x44, 0x54, 0x5f, 0x42, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x31, 0x36, 0x10, 0x0e, 0x12, 0x0d, 0x0a,
	0x09, 0x44, 0x54, 0x5f, 0x51, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x10, 0x0f, 0x12, 0x0e, 0x0a, 0x0a,
	0x44, 0x54, 0x5f, 0x51, 0x55, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x10, 0x10, 0x12, 0x0d, 0x0a, 0x09,
	0x44, 0x54, 0x5f, 0x55, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x10, 0x11, 0x12, 0x11, 0x0a, 0x0d, 0x44,
	0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x58, 0x31, 0x32, 0x38, 0x10, 0x12, 0x12, 0x0b,
	0x0a, 0x07, 0x44, 0x54, 0x5f, 0x48, 0x41, 0x4c, 0x46, 0x10, 0x13, 0x12, 0x0f, 0x0a, 0x0b, 0x44,
	0x54, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x10, 0x14, 0x12, 0x0e, 0x0a, 0x0a,
	0x44, 0x54, 0x5f, 0x56, 0x41, 0x52, 0x49, 0x41, 0x4e, 0x54, 0x10, 0x15, 0x42, 0x76, 0x0a, 0x29,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6c, 0x69, 0x79, 0x75, 0x6e, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x65, 0x61, 0x73, 0x2e, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0d, 0x50, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x69, 0x2d, 0x65, 0x61, 0x73, 0x2f, 0x65, 0x61, 0x73,
	0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x61, 0x73, 0x2f,
	0x74, 0x66, 0x5f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tf_predict_proto_rawDescOnce sync.Once
	file_tf_predict_proto_rawDescData = file_tf_predict_proto_rawDesc
)

func file_tf_predict_proto_rawDescGZIP() []byte {
	file_tf_predict_proto_rawDescOnce.Do(func() {
		file_tf_predict_proto_rawDescData = protoimpl.X.CompressGZIP(file_tf_predict_proto_rawDescData)
	})
	return file_tf_predict_proto_rawDescData
}

var file_tf_predict_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tf_predict_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_tf_predict_proto_goTypes = []interface{}{
	(ArrayDataType)(0),      // 0: tf_predict_protos.ArrayDataType
	(*ArrayShape)(nil),      // 1: tf_predict_protos.ArrayShape
	(*ArrayProto)(nil),      // 2: tf_predict_protos.ArrayProto
	(*InputEntry)(nil),      // 3: tf_predict_protos.InputEntry
	(*PredictRequest)(nil),  // 4: tf_predict_protos.PredictRequest
	(*OutputEntry)(nil),     // 5: tf_predict_protos.OutputEntry
	(*PredictResponse)(nil), // 6: tf_predict_protos.PredictResponse
}
var file_tf_predict_proto_depIdxs = []int32{
	0, // 0: tf_predict_protos.ArrayProto.dtype:type_name -> tf_predict_protos.ArrayDataType
	1, // 1: tf_predict_protos.ArrayProto.array_shape:type_name -> tf_predict_protos.ArrayShape
	2, // 2: tf_predict_protos.InputEntry.value:type_name -> tf_predict_protos.ArrayProto
	3, // 3: tf_predict_protos.PredictRequest.inputs:type_name -> tf_predict_protos.InputEntry
	2, // 4: tf_predict_protos.OutputEntry.value:type_name -> tf_predict_protos.ArrayProto
	5, // 5: tf_predict_protos.PredictResponse.outputs:type_name -> tf_predict_protos.OutputEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_tf_predict_proto_init() }
func file_tf_predict_proto_init() {
	if File_tf_predict_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tf_predict_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrayShape); i {
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
		file_tf_predict_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrayProto); i {
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
		file_tf_predict_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputEntry); i {
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
		file_tf_predict_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictRequest); i {
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
		file_tf_predict_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutputEntry); i {
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
		file_tf_predict_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictResponse); i {
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
			RawDescriptor: file_tf_predict_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tf_predict_proto_goTypes,
		DependencyIndexes: file_tf_predict_proto_depIdxs,
		EnumInfos:         file_tf_predict_proto_enumTypes,
		MessageInfos:      file_tf_predict_proto_msgTypes,
	}.Build()
	File_tf_predict_proto = out.File
	file_tf_predict_proto_rawDesc = nil
	file_tf_predict_proto_goTypes = nil
	file_tf_predict_proto_depIdxs = nil
}
