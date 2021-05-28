// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: settings.proto

package model

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type RegistrationRequest_ContactMethod int32

const (
	RegistrationRequest_unknown RegistrationRequest_ContactMethod = 0
	RegistrationRequest_sms     RegistrationRequest_ContactMethod = 1
	RegistrationRequest_voice   RegistrationRequest_ContactMethod = 2
)

// Enum value maps for RegistrationRequest_ContactMethod.
var (
	RegistrationRequest_ContactMethod_name = map[int32]string{
		0: "unknown",
		1: "sms",
		2: "voice",
	}
	RegistrationRequest_ContactMethod_value = map[string]int32{
		"unknown": 0,
		"sms":     1,
		"voice":   2,
	}
)

func (x RegistrationRequest_ContactMethod) Enum() *RegistrationRequest_ContactMethod {
	p := new(RegistrationRequest_ContactMethod)
	*p = x
	return p
}

func (x RegistrationRequest_ContactMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegistrationRequest_ContactMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_settings_proto_enumTypes[0].Descriptor()
}

func (RegistrationRequest_ContactMethod) Type() protoreflect.EnumType {
	return &file_settings_proto_enumTypes[0]
}

func (x RegistrationRequest_ContactMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegistrationRequest_ContactMethod.Descriptor instead.
func (RegistrationRequest_ContactMethod) EnumDescriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{0, 0}
}

type RegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cc          string                            `protobuf:"bytes,1,opt,name=cc,proto3" json:"cc,omitempty"`
	PhoneNumber string                            `protobuf:"bytes,2,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Method      RegistrationRequest_ContactMethod `protobuf:"varint,3,opt,name=method,proto3,enum=whatsapp.RegistrationRequest_ContactMethod" json:"method,omitempty"`
	Cert        string                            `protobuf:"bytes,4,opt,name=cert,proto3" json:"cert,omitempty"`
	Pin         string                            `protobuf:"bytes,5,opt,name=pin,proto3" json:"pin,omitempty"`
}

func (x *RegistrationRequest) Reset() {
	*x = RegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationRequest) ProtoMessage() {}

func (x *RegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationRequest.ProtoReflect.Descriptor instead.
func (*RegistrationRequest) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{0}
}

func (x *RegistrationRequest) GetCc() string {
	if x != nil {
		return x.Cc
	}
	return ""
}

func (x *RegistrationRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *RegistrationRequest) GetMethod() RegistrationRequest_ContactMethod {
	if x != nil {
		return x.Method
	}
	return RegistrationRequest_unknown
}

func (x *RegistrationRequest) GetCert() string {
	if x != nil {
		return x.Cert
	}
	return ""
}

func (x *RegistrationRequest) GetPin() string {
	if x != nil {
		return x.Pin
	}
	return ""
}

type VerifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *VerifyRequest) Reset() {
	*x = VerifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyRequest) ProtoMessage() {}

func (x *VerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyRequest.ProtoReflect.Descriptor instead.
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type ApplicationSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallbackBackoffDelayMs    int32                         `protobuf:"varint,1,opt,name=callback_backoff_delay_ms,json=callbackBackoffDelayMs,proto3" json:"callback_backoff_delay_ms,omitempty"`
	MaxCallbackBackoffDelayMs int32                         `protobuf:"varint,2,opt,name=max_callback_backoff_delay_ms,json=maxCallbackBackoffDelayMs,proto3" json:"max_callback_backoff_delay_ms,omitempty"`
	CallbackPersist           bool                          `protobuf:"varint,3,opt,name=callback_persist,json=callbackPersist,proto3" json:"callback_persist,omitempty"`
	Media                     *ApplicationSettings_Media    `protobuf:"bytes,4,opt,name=media,proto3" json:"media,omitempty"`
	Webhooks                  *ApplicationSettings_Webhooks `protobuf:"bytes,5,opt,name=webhooks,proto3" json:"webhooks,omitempty"`
	PassThrough               bool                          `protobuf:"varint,6,opt,name=pass_through,json=passThrough,proto3" json:"pass_through,omitempty"`
	SentStatus                bool                          `protobuf:"varint,7,opt,name=sent_status,json=sentStatus,proto3" json:"sent_status,omitempty"`
	DbGarbagecollectorEnable  bool                          `protobuf:"varint,8,opt,name=db_garbagecollector_enable,json=dbGarbagecollectorEnable,proto3" json:"db_garbagecollector_enable,omitempty"`
	NotifyUserChangeNumber    bool                          `protobuf:"varint,9,opt,name=notify_user_change_number,json=notifyUserChangeNumber,proto3" json:"notify_user_change_number,omitempty"`
	ShowSecurityNotifications bool                          `protobuf:"varint,10,opt,name=show_security_notifications,json=showSecurityNotifications,proto3" json:"show_security_notifications,omitempty"`
}

func (x *ApplicationSettings) Reset() {
	*x = ApplicationSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationSettings) ProtoMessage() {}

func (x *ApplicationSettings) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationSettings.ProtoReflect.Descriptor instead.
func (*ApplicationSettings) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{2}
}

func (x *ApplicationSettings) GetCallbackBackoffDelayMs() int32 {
	if x != nil {
		return x.CallbackBackoffDelayMs
	}
	return 0
}

func (x *ApplicationSettings) GetMaxCallbackBackoffDelayMs() int32 {
	if x != nil {
		return x.MaxCallbackBackoffDelayMs
	}
	return 0
}

func (x *ApplicationSettings) GetCallbackPersist() bool {
	if x != nil {
		return x.CallbackPersist
	}
	return false
}

func (x *ApplicationSettings) GetMedia() *ApplicationSettings_Media {
	if x != nil {
		return x.Media
	}
	return nil
}

func (x *ApplicationSettings) GetWebhooks() *ApplicationSettings_Webhooks {
	if x != nil {
		return x.Webhooks
	}
	return nil
}

func (x *ApplicationSettings) GetPassThrough() bool {
	if x != nil {
		return x.PassThrough
	}
	return false
}

func (x *ApplicationSettings) GetSentStatus() bool {
	if x != nil {
		return x.SentStatus
	}
	return false
}

func (x *ApplicationSettings) GetDbGarbagecollectorEnable() bool {
	if x != nil {
		return x.DbGarbagecollectorEnable
	}
	return false
}

func (x *ApplicationSettings) GetNotifyUserChangeNumber() bool {
	if x != nil {
		return x.NotifyUserChangeNumber
	}
	return false
}

func (x *ApplicationSettings) GetShowSecurityNotifications() bool {
	if x != nil {
		return x.ShowSecurityNotifications
	}
	return false
}

type ProfileAbout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *ProfileAbout) Reset() {
	*x = ProfileAbout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileAbout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileAbout) ProtoMessage() {}

func (x *ProfileAbout) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileAbout.ProtoReflect.Descriptor instead.
func (*ProfileAbout) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{3}
}

func (x *ProfileAbout) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type BusinessProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address     string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Email       string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Vertical    string   `protobuf:"bytes,4,opt,name=vertical,proto3" json:"vertical,omitempty"`
	Websites    []string `protobuf:"bytes,5,rep,name=websites,proto3" json:"websites,omitempty"`
}

func (x *BusinessProfile) Reset() {
	*x = BusinessProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusinessProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusinessProfile) ProtoMessage() {}

func (x *BusinessProfile) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusinessProfile.ProtoReflect.Descriptor instead.
func (*BusinessProfile) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{4}
}

func (x *BusinessProfile) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *BusinessProfile) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BusinessProfile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *BusinessProfile) GetVertical() string {
	if x != nil {
		return x.Vertical
	}
	return ""
}

func (x *BusinessProfile) GetWebsites() []string {
	if x != nil {
		return x.Websites
	}
	return nil
}

type ApplicationSettings_Media struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AutoDownload []string `protobuf:"bytes,1,rep,name=auto_download,json=autoDownload,proto3" json:"auto_download,omitempty"`
}

func (x *ApplicationSettings_Media) Reset() {
	*x = ApplicationSettings_Media{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationSettings_Media) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationSettings_Media) ProtoMessage() {}

func (x *ApplicationSettings_Media) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationSettings_Media.ProtoReflect.Descriptor instead.
func (*ApplicationSettings_Media) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ApplicationSettings_Media) GetAutoDownload() []string {
	if x != nil {
		return x.AutoDownload
	}
	return nil
}

type ApplicationSettings_Webhooks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url                   string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	MaxConcurrentRequests int32  `protobuf:"varint,2,opt,name=max_concurrent_requests,json=maxConcurrentRequests,proto3" json:"max_concurrent_requests,omitempty"`
}

func (x *ApplicationSettings_Webhooks) Reset() {
	*x = ApplicationSettings_Webhooks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_settings_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplicationSettings_Webhooks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationSettings_Webhooks) ProtoMessage() {}

func (x *ApplicationSettings_Webhooks) ProtoReflect() protoreflect.Message {
	mi := &file_settings_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationSettings_Webhooks.ProtoReflect.Descriptor instead.
func (*ApplicationSettings_Webhooks) Descriptor() ([]byte, []int) {
	return file_settings_proto_rawDescGZIP(), []int{2, 1}
}

func (x *ApplicationSettings_Webhooks) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ApplicationSettings_Webhooks) GetMaxConcurrentRequests() int32 {
	if x != nil {
		return x.MaxConcurrentRequests
	}
	return 0
}

var File_settings_proto protoreflect.FileDescriptor

var file_settings_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61, 0x70, 0x70, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x01,
	0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x63, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x77, 0x68, 0x61, 0x74, 0x73,
	0x61, 0x70, 0x70, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x65, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x65, 0x72,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x70, 0x69, 0x6e, 0x22, 0x30, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x73, 0x6d, 0x73, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x10, 0x02, 0x22, 0x23, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xbd, 0x05, 0x0a, 0x13, 0x41,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x39, 0x0a, 0x19, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x62,
	0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x42,
	0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x4d, 0x73, 0x12, 0x40, 0x0a,
	0x1d, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x62, 0x61,
	0x63, 0x6b, 0x6f, 0x66, 0x66, 0x5f, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x19, 0x6d, 0x61, 0x78, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x42, 0x61, 0x63, 0x6b, 0x6f, 0x66, 0x66, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x4d, 0x73, 0x12,
	0x29, 0x0a, 0x10, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x70, 0x65, 0x72, 0x73,
	0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x63, 0x61, 0x6c, 0x6c, 0x62,
	0x61, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x05, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x77, 0x68, 0x61, 0x74,
	0x73, 0x61, 0x70, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x52, 0x05,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x42, 0x0a, 0x08, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x77, 0x68, 0x61, 0x74, 0x73, 0x61,
	0x70, 0x70, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x52,
	0x08, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x73,
	0x73, 0x5f, 0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x70, 0x61, 0x73, 0x73, 0x54, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3c, 0x0a,
	0x1a, 0x64, 0x62, 0x5f, 0x67, 0x61, 0x72, 0x62, 0x61, 0x67, 0x65, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x18, 0x64, 0x62, 0x47, 0x61, 0x72, 0x62, 0x61, 0x67, 0x65, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x39, 0x0a, 0x19, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x3e, 0x0a, 0x1b, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x19, 0x73, 0x68, 0x6f,
	0x77, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x2c, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12,
	0x23, 0x0a, 0x0d, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x54, 0x0a, 0x08, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x36, 0x0a, 0x17, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x15, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x22, 0x0a, 0x0c, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x9b,
	0x01, 0x0a, 0x0f, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x73, 0x42, 0x08, 0x5a, 0x06,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_settings_proto_rawDescOnce sync.Once
	file_settings_proto_rawDescData = file_settings_proto_rawDesc
)

func file_settings_proto_rawDescGZIP() []byte {
	file_settings_proto_rawDescOnce.Do(func() {
		file_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_settings_proto_rawDescData)
	})
	return file_settings_proto_rawDescData
}

var file_settings_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_settings_proto_goTypes = []interface{}{
	(RegistrationRequest_ContactMethod)(0), // 0: whatsapp.RegistrationRequest.ContactMethod
	(*RegistrationRequest)(nil),            // 1: whatsapp.RegistrationRequest
	(*VerifyRequest)(nil),                  // 2: whatsapp.VerifyRequest
	(*ApplicationSettings)(nil),            // 3: whatsapp.ApplicationSettings
	(*ProfileAbout)(nil),                   // 4: whatsapp.ProfileAbout
	(*BusinessProfile)(nil),                // 5: whatsapp.BusinessProfile
	(*ApplicationSettings_Media)(nil),      // 6: whatsapp.ApplicationSettings.Media
	(*ApplicationSettings_Webhooks)(nil),   // 7: whatsapp.ApplicationSettings.Webhooks
}
var file_settings_proto_depIdxs = []int32{
	0, // 0: whatsapp.RegistrationRequest.method:type_name -> whatsapp.RegistrationRequest.ContactMethod
	6, // 1: whatsapp.ApplicationSettings.media:type_name -> whatsapp.ApplicationSettings.Media
	7, // 2: whatsapp.ApplicationSettings.webhooks:type_name -> whatsapp.ApplicationSettings.Webhooks
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_settings_proto_init() }
func file_settings_proto_init() {
	if File_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationRequest); i {
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
		file_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyRequest); i {
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
		file_settings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationSettings); i {
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
		file_settings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileAbout); i {
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
		file_settings_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusinessProfile); i {
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
		file_settings_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationSettings_Media); i {
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
		file_settings_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplicationSettings_Webhooks); i {
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
			RawDescriptor: file_settings_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_settings_proto_goTypes,
		DependencyIndexes: file_settings_proto_depIdxs,
		EnumInfos:         file_settings_proto_enumTypes,
		MessageInfos:      file_settings_proto_msgTypes,
	}.Build()
	File_settings_proto = out.File
	file_settings_proto_rawDesc = nil
	file_settings_proto_goTypes = nil
	file_settings_proto_depIdxs = nil
}