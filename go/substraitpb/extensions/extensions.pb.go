// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: substrait/extensions/extensions.proto

package extensions

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SimpleExtensionURI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A surrogate key used in the context of a single plan used to reference the
	// URI associated with an extension.
	ExtensionUriAnchor uint32 `protobuf:"varint,1,opt,name=extension_uri_anchor,json=extensionUriAnchor,proto3" json:"extension_uri_anchor,omitempty"`
	// The URI where this extension YAML can be retrieved. This is the "namespace"
	// of this extension.
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *SimpleExtensionURI) Reset() {
	*x = SimpleExtensionURI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substrait_extensions_extensions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleExtensionURI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleExtensionURI) ProtoMessage() {}

func (x *SimpleExtensionURI) ProtoReflect() protoreflect.Message {
	mi := &file_substrait_extensions_extensions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleExtensionURI.ProtoReflect.Descriptor instead.
func (*SimpleExtensionURI) Descriptor() ([]byte, []int) {
	return file_substrait_extensions_extensions_proto_rawDescGZIP(), []int{0}
}

func (x *SimpleExtensionURI) GetExtensionUriAnchor() uint32 {
	if x != nil {
		return x.ExtensionUriAnchor
	}
	return 0
}

func (x *SimpleExtensionURI) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

// Describes a mapping between a specific extension entity and the uri where
// that extension can be found.
type SimpleExtensionDeclaration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MappingType:
	//
	//	*SimpleExtensionDeclaration_ExtensionType_
	//	*SimpleExtensionDeclaration_ExtensionTypeVariation_
	//	*SimpleExtensionDeclaration_ExtensionFunction_
	MappingType isSimpleExtensionDeclaration_MappingType `protobuf_oneof:"mapping_type"`
}

func (x *SimpleExtensionDeclaration) Reset() {
	*x = SimpleExtensionDeclaration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substrait_extensions_extensions_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleExtensionDeclaration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleExtensionDeclaration) ProtoMessage() {}

func (x *SimpleExtensionDeclaration) ProtoReflect() protoreflect.Message {
	mi := &file_substrait_extensions_extensions_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleExtensionDeclaration.ProtoReflect.Descriptor instead.
func (*SimpleExtensionDeclaration) Descriptor() ([]byte, []int) {
	return file_substrait_extensions_extensions_proto_rawDescGZIP(), []int{1}
}

func (m *SimpleExtensionDeclaration) GetMappingType() isSimpleExtensionDeclaration_MappingType {
	if m != nil {
		return m.MappingType
	}
	return nil
}

func (x *SimpleExtensionDeclaration) GetExtensionType() *SimpleExtensionDeclaration_ExtensionType {
	if x, ok := x.GetMappingType().(*SimpleExtensionDeclaration_ExtensionType_); ok {
		return x.ExtensionType
	}
	return nil
}

func (x *SimpleExtensionDeclaration) GetExtensionTypeVariation() *SimpleExtensionDeclaration_ExtensionTypeVariation {
	if x, ok := x.GetMappingType().(*SimpleExtensionDeclaration_ExtensionTypeVariation_); ok {
		return x.ExtensionTypeVariation
	}
	return nil
}

func (x *SimpleExtensionDeclaration) GetExtensionFunction() *SimpleExtensionDeclaration_ExtensionFunction {
	if x, ok := x.GetMappingType().(*SimpleExtensionDeclaration_ExtensionFunction_); ok {
		return x.ExtensionFunction
	}
	return nil
}

type isSimpleExtensionDeclaration_MappingType interface {
	isSimpleExtensionDeclaration_MappingType()
}

type SimpleExtensionDeclaration_ExtensionType_ struct {
	ExtensionType *SimpleExtensionDeclaration_ExtensionType `protobuf:"bytes,1,opt,name=extension_type,json=extensionType,proto3,oneof"`
}

type SimpleExtensionDeclaration_ExtensionTypeVariation_ struct {
	ExtensionTypeVariation *SimpleExtensionDeclaration_ExtensionTypeVariation `protobuf:"bytes,2,opt,name=extension_type_variation,json=extensionTypeVariation,proto3,oneof"`
}

type SimpleExtensionDeclaration_ExtensionFunction_ struct {
	ExtensionFunction *SimpleExtensionDeclaration_ExtensionFunction `protobuf:"bytes,3,opt,name=extension_function,json=extensionFunction,proto3,oneof"`
}

func (*SimpleExtensionDeclaration_ExtensionType_) isSimpleExtensionDeclaration_MappingType() {}

func (*SimpleExtensionDeclaration_ExtensionTypeVariation_) isSimpleExtensionDeclaration_MappingType() {
}

func (*SimpleExtensionDeclaration_ExtensionFunction_) isSimpleExtensionDeclaration_MappingType() {}

// A generic object that can be used to embed additional extension information
// into the serialized substrait plan.
type AdvancedExtension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An optimization is helpful information that don't influence semantics. May
	// be ignored by a consumer.
	Optimization []*anypb.Any `protobuf:"bytes,1,rep,name=optimization,proto3" json:"optimization,omitempty"`
	// An enhancement alter semantics. Cannot be ignored by a consumer.
	Enhancement *anypb.Any `protobuf:"bytes,2,opt,name=enhancement,proto3" json:"enhancement,omitempty"`
}

func (x *AdvancedExtension) Reset() {
	*x = AdvancedExtension{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substrait_extensions_extensions_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvancedExtension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvancedExtension) ProtoMessage() {}

func (x *AdvancedExtension) ProtoReflect() protoreflect.Message {
	mi := &file_substrait_extensions_extensions_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvancedExtension.ProtoReflect.Descriptor instead.
func (*AdvancedExtension) Descriptor() ([]byte, []int) {
	return file_substrait_extensions_extensions_proto_rawDescGZIP(), []int{2}
}

func (x *AdvancedExtension) GetOptimization() []*anypb.Any {
	if x != nil {
		return x.Optimization
	}
	return nil
}

func (x *AdvancedExtension) GetEnhancement() *anypb.Any {
	if x != nil {
		return x.Enhancement
	}
	return nil
}

// Describes a Type
type SimpleExtensionDeclaration_ExtensionType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// references the extension_uri_anchor defined for a specific extension URI.
	ExtensionUriReference uint32 `protobuf:"varint,1,opt,name=extension_uri_reference,json=extensionUriReference,proto3" json:"extension_uri_reference,omitempty"`
	// A surrogate key used in the context of a single plan to reference a
	// specific extension type
	TypeAnchor uint32 `protobuf:"varint,2,opt,name=type_anchor,json=typeAnchor,proto3" json:"type_anchor,omitempty"`
	// the name of the type in the defined extension YAML.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SimpleExtensionDeclaration_ExtensionType) Reset() {
	*x = SimpleExtensionDeclaration_ExtensionType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substrait_extensions_extensions_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleExtensionDeclaration_ExtensionType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleExtensionDeclaration_ExtensionType) ProtoMessage() {}

func (x *SimpleExtensionDeclaration_ExtensionType) ProtoReflect() protoreflect.Message {
	mi := &file_substrait_extensions_extensions_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleExtensionDeclaration_ExtensionType.ProtoReflect.Descriptor instead.
func (*SimpleExtensionDeclaration_ExtensionType) Descriptor() ([]byte, []int) {
	return file_substrait_extensions_extensions_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SimpleExtensionDeclaration_ExtensionType) GetExtensionUriReference() uint32 {
	if x != nil {
		return x.ExtensionUriReference
	}
	return 0
}

func (x *SimpleExtensionDeclaration_ExtensionType) GetTypeAnchor() uint32 {
	if x != nil {
		return x.TypeAnchor
	}
	return 0
}

func (x *SimpleExtensionDeclaration_ExtensionType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SimpleExtensionDeclaration_ExtensionTypeVariation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// references the extension_uri_anchor defined for a specific extension URI.
	ExtensionUriReference uint32 `protobuf:"varint,1,opt,name=extension_uri_reference,json=extensionUriReference,proto3" json:"extension_uri_reference,omitempty"`
	// A surrogate key used in the context of a single plan to reference a
	// specific type variation
	TypeVariationAnchor uint32 `protobuf:"varint,2,opt,name=type_variation_anchor,json=typeVariationAnchor,proto3" json:"type_variation_anchor,omitempty"`
	// the name of the type in the defined extension YAML.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SimpleExtensionDeclaration_ExtensionTypeVariation) Reset() {
	*x = SimpleExtensionDeclaration_ExtensionTypeVariation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substrait_extensions_extensions_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleExtensionDeclaration_ExtensionTypeVariation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleExtensionDeclaration_ExtensionTypeVariation) ProtoMessage() {}

func (x *SimpleExtensionDeclaration_ExtensionTypeVariation) ProtoReflect() protoreflect.Message {
	mi := &file_substrait_extensions_extensions_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleExtensionDeclaration_ExtensionTypeVariation.ProtoReflect.Descriptor instead.
func (*SimpleExtensionDeclaration_ExtensionTypeVariation) Descriptor() ([]byte, []int) {
	return file_substrait_extensions_extensions_proto_rawDescGZIP(), []int{1, 1}
}

func (x *SimpleExtensionDeclaration_ExtensionTypeVariation) GetExtensionUriReference() uint32 {
	if x != nil {
		return x.ExtensionUriReference
	}
	return 0
}

func (x *SimpleExtensionDeclaration_ExtensionTypeVariation) GetTypeVariationAnchor() uint32 {
	if x != nil {
		return x.TypeVariationAnchor
	}
	return 0
}

func (x *SimpleExtensionDeclaration_ExtensionTypeVariation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SimpleExtensionDeclaration_ExtensionFunction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// references the extension_uri_anchor defined for a specific extension URI.
	ExtensionUriReference uint32 `protobuf:"varint,1,opt,name=extension_uri_reference,json=extensionUriReference,proto3" json:"extension_uri_reference,omitempty"`
	// A surrogate key used in the context of a single plan to reference a
	// specific function
	FunctionAnchor uint32 `protobuf:"varint,2,opt,name=function_anchor,json=functionAnchor,proto3" json:"function_anchor,omitempty"`
	// A function signature compound name
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SimpleExtensionDeclaration_ExtensionFunction) Reset() {
	*x = SimpleExtensionDeclaration_ExtensionFunction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_substrait_extensions_extensions_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimpleExtensionDeclaration_ExtensionFunction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleExtensionDeclaration_ExtensionFunction) ProtoMessage() {}

func (x *SimpleExtensionDeclaration_ExtensionFunction) ProtoReflect() protoreflect.Message {
	mi := &file_substrait_extensions_extensions_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleExtensionDeclaration_ExtensionFunction.ProtoReflect.Descriptor instead.
func (*SimpleExtensionDeclaration_ExtensionFunction) Descriptor() ([]byte, []int) {
	return file_substrait_extensions_extensions_proto_rawDescGZIP(), []int{1, 2}
}

func (x *SimpleExtensionDeclaration_ExtensionFunction) GetExtensionUriReference() uint32 {
	if x != nil {
		return x.ExtensionUriReference
	}
	return 0
}

func (x *SimpleExtensionDeclaration_ExtensionFunction) GetFunctionAnchor() uint32 {
	if x != nil {
		return x.FunctionAnchor
	}
	return 0
}

func (x *SimpleExtensionDeclaration_ExtensionFunction) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_substrait_extensions_extensions_proto protoreflect.FileDescriptor

var file_substrait_extensions_extensions_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x12, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x52, 0x49, 0x12, 0x30,
	0x0a, 0x14, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x69, 0x5f,
	0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x69, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x69, 0x22, 0xb4, 0x06, 0x0a, 0x1a, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x67, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x73, 0x75, 0x62, 0x73,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x44, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x18, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x16, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x73, 0x0a, 0x12, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x73,
	0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x44, 0x65, 0x63, 0x6c, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x11, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x7c, 0x0a, 0x0d, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x69, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x55, 0x72, 0x69, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x79, 0x70, 0x65, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x1a, 0x98, 0x01, 0x0a, 0x16, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36,
	0x0a, 0x17, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x69, 0x5f,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x15, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x69, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x74, 0x79, 0x70, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x88,
	0x01, 0x0a, 0x11, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x17, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x75, 0x72, 0x69, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x55, 0x72, 0x69, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x6e, 0x63, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x6d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x11, 0x41, 0x64,
	0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x38, 0x0a, 0x0c, 0x6f, 0x70, 0x74, 0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0c, 0x6f, 0x70, 0x74,
	0x69, 0x6d, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x0b, 0x65, 0x6e, 0x68,
	0x61, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0b, 0x65, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x42, 0xe2, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72,
	0x61, 0x69, 0x74, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x0f,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75,
	0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74,
	0x72, 0x61, 0x69, 0x74, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x70, 0x62, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0xa2, 0x02, 0x03, 0x53, 0x45, 0x58, 0xaa, 0x02, 0x14,
	0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0xca, 0x02, 0x14, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74,
	0x5c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0xe2, 0x02, 0x20, 0x53, 0x75,
	0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x5c, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x15, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x61, 0x69, 0x74, 0x3a, 0x3a, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_substrait_extensions_extensions_proto_rawDescOnce sync.Once
	file_substrait_extensions_extensions_proto_rawDescData = file_substrait_extensions_extensions_proto_rawDesc
)

func file_substrait_extensions_extensions_proto_rawDescGZIP() []byte {
	file_substrait_extensions_extensions_proto_rawDescOnce.Do(func() {
		file_substrait_extensions_extensions_proto_rawDescData = protoimpl.X.CompressGZIP(file_substrait_extensions_extensions_proto_rawDescData)
	})
	return file_substrait_extensions_extensions_proto_rawDescData
}

var file_substrait_extensions_extensions_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_substrait_extensions_extensions_proto_goTypes = []interface{}{
	(*SimpleExtensionURI)(nil),                                // 0: substrait.extensions.SimpleExtensionURI
	(*SimpleExtensionDeclaration)(nil),                        // 1: substrait.extensions.SimpleExtensionDeclaration
	(*AdvancedExtension)(nil),                                 // 2: substrait.extensions.AdvancedExtension
	(*SimpleExtensionDeclaration_ExtensionType)(nil),          // 3: substrait.extensions.SimpleExtensionDeclaration.ExtensionType
	(*SimpleExtensionDeclaration_ExtensionTypeVariation)(nil), // 4: substrait.extensions.SimpleExtensionDeclaration.ExtensionTypeVariation
	(*SimpleExtensionDeclaration_ExtensionFunction)(nil),      // 5: substrait.extensions.SimpleExtensionDeclaration.ExtensionFunction
	(*anypb.Any)(nil),                                         // 6: google.protobuf.Any
}
var file_substrait_extensions_extensions_proto_depIdxs = []int32{
	3, // 0: substrait.extensions.SimpleExtensionDeclaration.extension_type:type_name -> substrait.extensions.SimpleExtensionDeclaration.ExtensionType
	4, // 1: substrait.extensions.SimpleExtensionDeclaration.extension_type_variation:type_name -> substrait.extensions.SimpleExtensionDeclaration.ExtensionTypeVariation
	5, // 2: substrait.extensions.SimpleExtensionDeclaration.extension_function:type_name -> substrait.extensions.SimpleExtensionDeclaration.ExtensionFunction
	6, // 3: substrait.extensions.AdvancedExtension.optimization:type_name -> google.protobuf.Any
	6, // 4: substrait.extensions.AdvancedExtension.enhancement:type_name -> google.protobuf.Any
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_substrait_extensions_extensions_proto_init() }
func file_substrait_extensions_extensions_proto_init() {
	if File_substrait_extensions_extensions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_substrait_extensions_extensions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleExtensionURI); i {
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
		file_substrait_extensions_extensions_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleExtensionDeclaration); i {
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
		file_substrait_extensions_extensions_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvancedExtension); i {
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
		file_substrait_extensions_extensions_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleExtensionDeclaration_ExtensionType); i {
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
		file_substrait_extensions_extensions_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleExtensionDeclaration_ExtensionTypeVariation); i {
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
		file_substrait_extensions_extensions_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimpleExtensionDeclaration_ExtensionFunction); i {
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
	file_substrait_extensions_extensions_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SimpleExtensionDeclaration_ExtensionType_)(nil),
		(*SimpleExtensionDeclaration_ExtensionTypeVariation_)(nil),
		(*SimpleExtensionDeclaration_ExtensionFunction_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_substrait_extensions_extensions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_substrait_extensions_extensions_proto_goTypes,
		DependencyIndexes: file_substrait_extensions_extensions_proto_depIdxs,
		MessageInfos:      file_substrait_extensions_extensions_proto_msgTypes,
	}.Build()
	File_substrait_extensions_extensions_proto = out.File
	file_substrait_extensions_extensions_proto_rawDesc = nil
	file_substrait_extensions_extensions_proto_goTypes = nil
	file_substrait_extensions_extensions_proto_depIdxs = nil
}
