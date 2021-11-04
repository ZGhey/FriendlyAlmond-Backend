// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/login/login.proto

package pbLogin

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_login_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_login_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_login_login_proto_rawDescGZIP(), []int{0}
}

type Captcha struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Image      string `protobuf:"bytes,2,opt,name=Image,proto3" json:"Image,omitempty"`
	Answer     string `protobuf:"bytes,3,opt,name=Answer,proto3" json:"Answer,omitempty"`
	StatusCode string `protobuf:"bytes,4,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
}

func (x *Captcha) Reset() {
	*x = Captcha{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_login_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Captcha) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Captcha) ProtoMessage() {}

func (x *Captcha) ProtoReflect() protoreflect.Message {
	mi := &file_proto_login_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Captcha.ProtoReflect.Descriptor instead.
func (*Captcha) Descriptor() ([]byte, []int) {
	return file_proto_login_login_proto_rawDescGZIP(), []int{1}
}

func (x *Captcha) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Captcha) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Captcha) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

func (x *Captcha) GetStatusCode() string {
	if x != nil {
		return x.StatusCode
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName       string           `protobuf:"bytes,1,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	Password        string           `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	Email           string           `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Uid             string           `protobuf:"bytes,4,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Phone           string           `protobuf:"bytes,5,opt,name=Phone,proto3" json:"Phone,omitempty"`
	Address         string           `protobuf:"bytes,6,opt,name=Address,proto3" json:"Address,omitempty"`
	Captcha         *Captcha         `protobuf:"bytes,7,opt,name=Captcha,proto3" json:"Captcha,omitempty"`
	OperationResult *OperationResult `protobuf:"bytes,8,opt,name=OperationResult,proto3" json:"OperationResult,omitempty"`
	MiddleName      string           `protobuf:"bytes,9,opt,name=MiddleName,proto3" json:"MiddleName,omitempty"`
	LastName        string           `protobuf:"bytes,10,opt,name=LastName,proto3" json:"LastName,omitempty"`
	AreaCode        string           `protobuf:"bytes,11,opt,name=AreaCode,proto3" json:"AreaCode,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_login_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_login_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_proto_login_login_proto_rawDescGZIP(), []int{2}
}

func (x *UserInfo) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *UserInfo) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UserInfo) GetCaptcha() *Captcha {
	if x != nil {
		return x.Captcha
	}
	return nil
}

func (x *UserInfo) GetOperationResult() *OperationResult {
	if x != nil {
		return x.OperationResult
	}
	return nil
}

func (x *UserInfo) GetMiddleName() string {
	if x != nil {
		return x.MiddleName
	}
	return ""
}

func (x *UserInfo) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserInfo) GetAreaCode() string {
	if x != nil {
		return x.AreaCode
	}
	return ""
}

type OperationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode string `protobuf:"bytes,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
}

func (x *OperationResult) Reset() {
	*x = OperationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_login_login_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationResult) ProtoMessage() {}

func (x *OperationResult) ProtoReflect() protoreflect.Message {
	mi := &file_proto_login_login_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationResult.ProtoReflect.Descriptor instead.
func (*OperationResult) Descriptor() ([]byte, []int) {
	return file_proto_login_login_proto_rawDescGZIP(), []int{3}
}

func (x *OperationResult) GetStatusCode() string {
	if x != nil {
		return x.StatusCode
	}
	return ""
}

var File_proto_login_login_proto protoreflect.FileDescriptor

var file_proto_login_login_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x67, 0x0a, 0x07, 0x43, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x6e, 0x73, 0x77,
	0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x22, 0xe0, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x10, 0x0a, 0x03, 0x55, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x28, 0x0a, 0x07, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x63,
	0x68, 0x61, 0x52, 0x07, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x40, 0x0a, 0x0f, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0f, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x72, 0x65,
	0x61, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x41, 0x72, 0x65,
	0x61, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x31, 0x0a, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x32, 0xba, 0x02, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x31, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x0c, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x70, 0x74,
	0x63, 0x68, 0x61, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x2b, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x0f, 0x2e, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0f, 0x2e, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x2b,
	0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0f, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0b, 0x49,
	0x73, 0x53, 0x61, 0x6d, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0f, 0x2e, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3b, 0x70, 0x62, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_login_login_proto_rawDescOnce sync.Once
	file_proto_login_login_proto_rawDescData = file_proto_login_login_proto_rawDesc
)

func file_proto_login_login_proto_rawDescGZIP() []byte {
	file_proto_login_login_proto_rawDescOnce.Do(func() {
		file_proto_login_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_login_login_proto_rawDescData)
	})
	return file_proto_login_login_proto_rawDescData
}

var file_proto_login_login_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_login_login_proto_goTypes = []interface{}{
	(*Empty)(nil),           // 0: login.Empty
	(*Captcha)(nil),         // 1: login.Captcha
	(*UserInfo)(nil),        // 2: login.UserInfo
	(*OperationResult)(nil), // 3: login.OperationResult
}
var file_proto_login_login_proto_depIdxs = []int32{
	1, // 0: login.UserInfo.Captcha:type_name -> login.Captcha
	3, // 1: login.UserInfo.OperationResult:type_name -> login.OperationResult
	0, // 2: login.Login.GenerateCaptcha:input_type -> login.Empty
	2, // 3: login.Login.Register:input_type -> login.UserInfo
	2, // 4: login.Login.Update:input_type -> login.UserInfo
	2, // 5: login.Login.Query:input_type -> login.UserInfo
	2, // 6: login.Login.Login:input_type -> login.UserInfo
	2, // 7: login.Login.IsSameEmail:input_type -> login.UserInfo
	1, // 8: login.Login.GenerateCaptcha:output_type -> login.Captcha
	3, // 9: login.Login.Register:output_type -> login.OperationResult
	3, // 10: login.Login.Update:output_type -> login.OperationResult
	2, // 11: login.Login.Query:output_type -> login.UserInfo
	2, // 12: login.Login.Login:output_type -> login.UserInfo
	3, // 13: login.Login.IsSameEmail:output_type -> login.OperationResult
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_login_login_proto_init() }
func file_proto_login_login_proto_init() {
	if File_proto_login_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_login_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_login_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Captcha); i {
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
		file_proto_login_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_proto_login_login_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationResult); i {
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
			RawDescriptor: file_proto_login_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_login_login_proto_goTypes,
		DependencyIndexes: file_proto_login_login_proto_depIdxs,
		MessageInfos:      file_proto_login_login_proto_msgTypes,
	}.Build()
	File_proto_login_login_proto = out.File
	file_proto_login_login_proto_rawDesc = nil
	file_proto_login_login_proto_goTypes = nil
	file_proto_login_login_proto_depIdxs = nil
}
