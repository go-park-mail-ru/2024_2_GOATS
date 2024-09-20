// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: validation.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ValidateRegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email           string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password        string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	PasswordConfirm string                 `protobuf:"bytes,3,opt,name=password_confirm,json=passwordConfirm,proto3" json:"password_confirm,omitempty"`
	Sex             int32                  `protobuf:"varint,4,opt,name=sex,proto3" json:"sex,omitempty"`
	Bitrhdate       *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=bitrhdate,proto3" json:"bitrhdate,omitempty"`
}

func (x *ValidateRegistrationRequest) Reset() {
	*x = ValidateRegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateRegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRegistrationRequest) ProtoMessage() {}

func (x *ValidateRegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_validation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateRegistrationRequest.ProtoReflect.Descriptor instead.
func (*ValidateRegistrationRequest) Descriptor() ([]byte, []int) {
	return file_validation_proto_rawDescGZIP(), []int{0}
}

func (x *ValidateRegistrationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ValidateRegistrationRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ValidateRegistrationRequest) GetPasswordConfirm() string {
	if x != nil {
		return x.PasswordConfirm
	}
	return ""
}

func (x *ValidateRegistrationRequest) GetSex() int32 {
	if x != nil {
		return x.Sex
	}
	return 0
}

func (x *ValidateRegistrationRequest) GetBitrhdate() *timestamppb.Timestamp {
	if x != nil {
		return x.Bitrhdate
	}
	return nil
}

type ErrorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ErrorMessage) Reset() {
	*x = ErrorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorMessage) ProtoMessage() {}

func (x *ErrorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_validation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorMessage.ProtoReflect.Descriptor instead.
func (*ErrorMessage) Descriptor() ([]byte, []int) {
	return file_validation_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorMessage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ErrorMessage) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ValidationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool            `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Errors  []*ErrorMessage `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *ValidationResponse) Reset() {
	*x = ValidationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_validation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationResponse) ProtoMessage() {}

func (x *ValidationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_validation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationResponse.ProtoReflect.Descriptor instead.
func (*ValidationResponse) Descriptor() ([]byte, []int) {
	return file_validation_proto_rawDescGZIP(), []int{2}
}

func (x *ValidationResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ValidationResponse) GetErrors() []*ErrorMessage {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_validation_proto protoreflect.FileDescriptor

var file_validation_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc6, 0x01, 0x0a, 0x1b, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x38,
	0x0a, 0x09, 0x62, 0x69, 0x74, 0x72, 0x68, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x62,
	0x69, 0x74, 0x72, 0x68, 0x64, 0x61, 0x74, 0x65, 0x22, 0x38, 0x0a, 0x0c, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x60, 0x0a, 0x12, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x30, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x06, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x32, 0x6d, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x5f, 0x0a, 0x14, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_validation_proto_rawDescOnce sync.Once
	file_validation_proto_rawDescData = file_validation_proto_rawDesc
)

func file_validation_proto_rawDescGZIP() []byte {
	file_validation_proto_rawDescOnce.Do(func() {
		file_validation_proto_rawDescData = protoimpl.X.CompressGZIP(file_validation_proto_rawDescData)
	})
	return file_validation_proto_rawDescData
}

var file_validation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_validation_proto_goTypes = []any{
	(*ValidateRegistrationRequest)(nil), // 0: validation.ValidateRegistrationRequest
	(*ErrorMessage)(nil),                // 1: validation.ErrorMessage
	(*ValidationResponse)(nil),          // 2: validation.ValidationResponse
	(*timestamppb.Timestamp)(nil),       // 3: google.protobuf.Timestamp
}
var file_validation_proto_depIdxs = []int32{
	3, // 0: validation.ValidateRegistrationRequest.bitrhdate:type_name -> google.protobuf.Timestamp
	1, // 1: validation.ValidationResponse.errors:type_name -> validation.ErrorMessage
	0, // 2: validation.Validation.ValidateRegistration:input_type -> validation.ValidateRegistrationRequest
	2, // 3: validation.Validation.ValidateRegistration:output_type -> validation.ValidationResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_validation_proto_init() }
func file_validation_proto_init() {
	if File_validation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_validation_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ValidateRegistrationRequest); i {
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
		file_validation_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ErrorMessage); i {
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
		file_validation_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ValidationResponse); i {
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
			RawDescriptor: file_validation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_validation_proto_goTypes,
		DependencyIndexes: file_validation_proto_depIdxs,
		MessageInfos:      file_validation_proto_msgTypes,
	}.Build()
	File_validation_proto = out.File
	file_validation_proto_rawDesc = nil
	file_validation_proto_goTypes = nil
	file_validation_proto_depIdxs = nil
}
