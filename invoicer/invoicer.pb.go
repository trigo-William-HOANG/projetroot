// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.20.0
// source: invoicer.proto

package invoicer

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

type InvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *InvoiceRequest) Reset() {
	*x = InvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoicer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceRequest) ProtoMessage() {}

func (x *InvoiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_invoicer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvoiceRequest.ProtoReflect.Descriptor instead.
func (*InvoiceRequest) Descriptor() ([]byte, []int) {
	return file_invoicer_proto_rawDescGZIP(), []int{0}
}

func (x *InvoiceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type App struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdApp          string `protobuf:"bytes,1,opt,name=idApp,proto3" json:"idApp,omitempty"`
	NameApp        string `protobuf:"bytes,2,opt,name=nameApp,proto3" json:"nameApp,omitempty"`
	DescriptionApp string `protobuf:"bytes,3,opt,name=descriptionApp,proto3" json:"descriptionApp,omitempty"`
	AppUrl         string `protobuf:"bytes,4,opt,name=appUrl,proto3" json:"appUrl,omitempty"`
	ImageAppUrl    string `protobuf:"bytes,5,opt,name=imageAppUrl,proto3" json:"imageAppUrl,omitempty"`
}

func (x *App) Reset() {
	*x = App{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoicer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *App) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*App) ProtoMessage() {}

func (x *App) ProtoReflect() protoreflect.Message {
	mi := &file_invoicer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use App.ProtoReflect.Descriptor instead.
func (*App) Descriptor() ([]byte, []int) {
	return file_invoicer_proto_rawDescGZIP(), []int{1}
}

func (x *App) GetIdApp() string {
	if x != nil {
		return x.IdApp
	}
	return ""
}

func (x *App) GetNameApp() string {
	if x != nil {
		return x.NameApp
	}
	return ""
}

func (x *App) GetDescriptionApp() string {
	if x != nil {
		return x.DescriptionApp
	}
	return ""
}

func (x *App) GetAppUrl() string {
	if x != nil {
		return x.AppUrl
	}
	return ""
}

func (x *App) GetImageAppUrl() string {
	if x != nil {
		return x.ImageAppUrl
	}
	return ""
}

type InvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Apps []*App `protobuf:"bytes,1,rep,name=apps,proto3" json:"apps,omitempty"` // Return multiple apps
}

func (x *InvoiceResponse) Reset() {
	*x = InvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoicer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceResponse) ProtoMessage() {}

func (x *InvoiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_invoicer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvoiceResponse.ProtoReflect.Descriptor instead.
func (*InvoiceResponse) Descriptor() ([]byte, []int) {
	return file_invoicer_proto_rawDescGZIP(), []int{2}
}

func (x *InvoiceResponse) GetApps() []*App {
	if x != nil {
		return x.Apps
	}
	return nil
}

var File_invoicer_proto protoreflect.FileDescriptor

var file_invoicer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x72, 0x22, 0x20, 0x0a, 0x0e, 0x49, 0x6e,
	0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x97, 0x01, 0x0a,
	0x03, 0x41, 0x70, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x64, 0x41, 0x70, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x64, 0x41, 0x70, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x61,
	0x6d, 0x65, 0x41, 0x70, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x61, 0x6d,
	0x65, 0x41, 0x70, 0x70, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x70, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x70, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x70, 0x70, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70,
	0x70, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x41, 0x70, 0x70,
	0x55, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x41, 0x70, 0x70, 0x55, 0x72, 0x6c, 0x22, 0x34, 0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x61, 0x70, 0x70,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x52, 0x04, 0x61, 0x70, 0x70, 0x73, 0x32, 0x4d, 0x0a, 0x08,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x18, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65,
	0x72, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x69, 0x67, 0x6f, 0x2d,
	0x57, 0x69, 0x6c, 0x6c, 0x69, 0x61, 0x6d, 0x2d, 0x48, 0x4f, 0x41, 0x4e, 0x47, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x72, 0x6f, 0x6f, 0x74, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_invoicer_proto_rawDescOnce sync.Once
	file_invoicer_proto_rawDescData = file_invoicer_proto_rawDesc
)

func file_invoicer_proto_rawDescGZIP() []byte {
	file_invoicer_proto_rawDescOnce.Do(func() {
		file_invoicer_proto_rawDescData = protoimpl.X.CompressGZIP(file_invoicer_proto_rawDescData)
	})
	return file_invoicer_proto_rawDescData
}

var file_invoicer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_invoicer_proto_goTypes = []any{
	(*InvoiceRequest)(nil),  // 0: invoicer.InvoiceRequest
	(*App)(nil),             // 1: invoicer.App
	(*InvoiceResponse)(nil), // 2: invoicer.InvoiceResponse
}
var file_invoicer_proto_depIdxs = []int32{
	1, // 0: invoicer.InvoiceResponse.apps:type_name -> invoicer.App
	0, // 1: invoicer.Invoicer.GetInvoice:input_type -> invoicer.InvoiceRequest
	2, // 2: invoicer.Invoicer.GetInvoice:output_type -> invoicer.InvoiceResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_invoicer_proto_init() }
func file_invoicer_proto_init() {
	if File_invoicer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_invoicer_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*InvoiceRequest); i {
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
		file_invoicer_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*App); i {
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
		file_invoicer_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*InvoiceResponse); i {
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
			RawDescriptor: file_invoicer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_invoicer_proto_goTypes,
		DependencyIndexes: file_invoicer_proto_depIdxs,
		MessageInfos:      file_invoicer_proto_msgTypes,
	}.Build()
	File_invoicer_proto = out.File
	file_invoicer_proto_rawDesc = nil
	file_invoicer_proto_goTypes = nil
	file_invoicer_proto_depIdxs = nil
}