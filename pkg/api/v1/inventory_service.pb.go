// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.3
// source: inventory_service.proto

package v1

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

type InventoryModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string                        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title  string                        `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Serial string                        `protobuf:"bytes,3,opt,name=serial,proto3" json:"serial,omitempty"`
	Color  string                        `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
	Count  []*InventoryModel_SizeCounter `protobuf:"bytes,5,rep,name=count,proto3" json:"count,omitempty"`
}

func (x *InventoryModel) Reset() {
	*x = InventoryModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryModel) ProtoMessage() {}

func (x *InventoryModel) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryModel.ProtoReflect.Descriptor instead.
func (*InventoryModel) Descriptor() ([]byte, []int) {
	return file_inventory_service_proto_rawDescGZIP(), []int{0}
}

func (x *InventoryModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InventoryModel) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *InventoryModel) GetSerial() string {
	if x != nil {
		return x.Serial
	}
	return ""
}

func (x *InventoryModel) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *InventoryModel) GetCount() []*InventoryModel_SizeCounter {
	if x != nil {
		return x.Count
	}
	return nil
}

type NewInventoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author string            `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Data   []*InventoryModel `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *NewInventoryRequest) Reset() {
	*x = NewInventoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewInventoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewInventoryRequest) ProtoMessage() {}

func (x *NewInventoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewInventoryRequest.ProtoReflect.Descriptor instead.
func (*NewInventoryRequest) Descriptor() ([]byte, []int) {
	return file_inventory_service_proto_rawDescGZIP(), []int{1}
}

func (x *NewInventoryRequest) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *NewInventoryRequest) GetData() []*InventoryModel {
	if x != nil {
		return x.Data
	}
	return nil
}

type NewInventoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NewInventoryResponse) Reset() {
	*x = NewInventoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewInventoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewInventoryResponse) ProtoMessage() {}

func (x *NewInventoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewInventoryResponse.ProtoReflect.Descriptor instead.
func (*NewInventoryResponse) Descriptor() ([]byte, []int) {
	return file_inventory_service_proto_rawDescGZIP(), []int{2}
}

func (x *NewInventoryResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type InventoryModel_SizeCounter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width float32 `protobuf:"fixed32,1,opt,name=width,proto3" json:"width,omitempty"`
	Side  string  `protobuf:"bytes,2,opt,name=side,proto3" json:"side,omitempty"`
	Count int32   `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Valid bool    `protobuf:"varint,4,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *InventoryModel_SizeCounter) Reset() {
	*x = InventoryModel_SizeCounter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_inventory_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryModel_SizeCounter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryModel_SizeCounter) ProtoMessage() {}

func (x *InventoryModel_SizeCounter) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryModel_SizeCounter.ProtoReflect.Descriptor instead.
func (*InventoryModel_SizeCounter) Descriptor() ([]byte, []int) {
	return file_inventory_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *InventoryModel_SizeCounter) GetWidth() float32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *InventoryModel_SizeCounter) GetSide() string {
	if x != nil {
		return x.Side
	}
	return ""
}

func (x *InventoryModel_SizeCounter) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *InventoryModel_SizeCounter) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

var File_inventory_service_proto protoreflect.FileDescriptor

var file_inventory_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x22, 0x86, 0x02, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x3b, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x53, 0x69, 0x7a, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x63, 0x0a, 0x0b, 0x53, 0x69, 0x7a, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x69, 0x64,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x5c, 0x0a,
	0x13, 0x4e, 0x65, 0x77, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x26, 0x0a, 0x14, 0x4e,
	0x65, 0x77, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x32, 0x65, 0x0a, 0x10, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x4e, 0x65, 0x77, 0x49, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x1e, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x4e, 0x65, 0x77, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x4e, 0x65, 0x77, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_inventory_service_proto_rawDescOnce sync.Once
	file_inventory_service_proto_rawDescData = file_inventory_service_proto_rawDesc
)

func file_inventory_service_proto_rawDescGZIP() []byte {
	file_inventory_service_proto_rawDescOnce.Do(func() {
		file_inventory_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_inventory_service_proto_rawDescData)
	})
	return file_inventory_service_proto_rawDescData
}

var file_inventory_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_inventory_service_proto_goTypes = []interface{}{
	(*InventoryModel)(nil),             // 0: inventory.InventoryModel
	(*NewInventoryRequest)(nil),        // 1: inventory.NewInventoryRequest
	(*NewInventoryResponse)(nil),       // 2: inventory.NewInventoryResponse
	(*InventoryModel_SizeCounter)(nil), // 3: inventory.InventoryModel.SizeCounter
}
var file_inventory_service_proto_depIdxs = []int32{
	3, // 0: inventory.InventoryModel.count:type_name -> inventory.InventoryModel.SizeCounter
	0, // 1: inventory.NewInventoryRequest.data:type_name -> inventory.InventoryModel
	1, // 2: inventory.InventoryService.NewInventory:input_type -> inventory.NewInventoryRequest
	2, // 3: inventory.InventoryService.NewInventory:output_type -> inventory.NewInventoryResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_inventory_service_proto_init() }
func file_inventory_service_proto_init() {
	if File_inventory_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_inventory_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryModel); i {
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
		file_inventory_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewInventoryRequest); i {
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
		file_inventory_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewInventoryResponse); i {
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
		file_inventory_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InventoryModel_SizeCounter); i {
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
			RawDescriptor: file_inventory_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inventory_service_proto_goTypes,
		DependencyIndexes: file_inventory_service_proto_depIdxs,
		MessageInfos:      file_inventory_service_proto_msgTypes,
	}.Build()
	File_inventory_service_proto = out.File
	file_inventory_service_proto_rawDesc = nil
	file_inventory_service_proto_goTypes = nil
	file_inventory_service_proto_depIdxs = nil
}
