// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.0
// source: orders-admin-pkg-models.proto

package orders_admin_package

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

// Shop is
type OrderAdmin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     []byte  `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name     string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MealType string  `protobuf:"bytes,3,opt,name=mealType,proto3" json:"mealType,omitempty"`
	Price    float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Day      string  `protobuf:"bytes,5,opt,name=day,proto3" json:"day,omitempty"`
}

func (x *OrderAdmin) Reset() {
	*x = OrderAdmin{}
	mi := &file_orders_admin_pkg_models_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrderAdmin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderAdmin) ProtoMessage() {}

func (x *OrderAdmin) ProtoReflect() protoreflect.Message {
	mi := &file_orders_admin_pkg_models_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderAdmin.ProtoReflect.Descriptor instead.
func (*OrderAdmin) Descriptor() ([]byte, []int) {
	return file_orders_admin_pkg_models_proto_rawDescGZIP(), []int{0}
}

func (x *OrderAdmin) GetUuid() []byte {
	if x != nil {
		return x.Uuid
	}
	return nil
}

func (x *OrderAdmin) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrderAdmin) GetMealType() string {
	if x != nil {
		return x.MealType
	}
	return ""
}

func (x *OrderAdmin) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OrderAdmin) GetDay() string {
	if x != nil {
		return x.Day
	}
	return ""
}

type OrdersAdmin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrdersAdmin []*OrderAdmin `protobuf:"bytes,1,rep,name=OrdersAdmin,proto3" json:"OrdersAdmin,omitempty"`
}

func (x *OrdersAdmin) Reset() {
	*x = OrdersAdmin{}
	mi := &file_orders_admin_pkg_models_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OrdersAdmin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdersAdmin) ProtoMessage() {}

func (x *OrdersAdmin) ProtoReflect() protoreflect.Message {
	mi := &file_orders_admin_pkg_models_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdersAdmin.ProtoReflect.Descriptor instead.
func (*OrdersAdmin) Descriptor() ([]byte, []int) {
	return file_orders_admin_pkg_models_proto_rawDescGZIP(), []int{1}
}

func (x *OrdersAdmin) GetOrdersAdmin() []*OrderAdmin {
	if x != nil {
		return x.OrdersAdmin
	}
	return nil
}

type MealDb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        []byte  `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Dummy       *Dummy  `protobuf:"bytes,5,opt,name=Dummy,proto3" json:"Dummy,omitempty"`
}

func (x *MealDb) Reset() {
	*x = MealDb{}
	mi := &file_orders_admin_pkg_models_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MealDb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MealDb) ProtoMessage() {}

func (x *MealDb) ProtoReflect() protoreflect.Message {
	mi := &file_orders_admin_pkg_models_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MealDb.ProtoReflect.Descriptor instead.
func (*MealDb) Descriptor() ([]byte, []int) {
	return file_orders_admin_pkg_models_proto_rawDescGZIP(), []int{2}
}

func (x *MealDb) GetUuid() []byte {
	if x != nil {
		return x.Uuid
	}
	return nil
}

func (x *MealDb) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MealDb) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MealDb) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *MealDb) GetDummy() *Dummy {
	if x != nil {
		return x.Dummy
	}
	return nil
}

type Dummy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dummies  []string `protobuf:"bytes,1,rep,name=dummies,proto3" json:"dummies,omitempty"`
	Trumpies []string `protobuf:"bytes,2,rep,name=trumpies,proto3" json:"trumpies,omitempty"`
	Puppies  []string `protobuf:"bytes,3,rep,name=puppies,proto3" json:"puppies,omitempty"`
}

func (x *Dummy) Reset() {
	*x = Dummy{}
	mi := &file_orders_admin_pkg_models_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Dummy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dummy) ProtoMessage() {}

func (x *Dummy) ProtoReflect() protoreflect.Message {
	mi := &file_orders_admin_pkg_models_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dummy.ProtoReflect.Descriptor instead.
func (*Dummy) Descriptor() ([]byte, []int) {
	return file_orders_admin_pkg_models_proto_rawDescGZIP(), []int{3}
}

func (x *Dummy) GetDummies() []string {
	if x != nil {
		return x.Dummies
	}
	return nil
}

func (x *Dummy) GetTrumpies() []string {
	if x != nil {
		return x.Trumpies
	}
	return nil
}

func (x *Dummy) GetPuppies() []string {
	if x != nil {
		return x.Puppies
	}
	return nil
}

type MealsDb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MealsDb []*MealDb `protobuf:"bytes,1,rep,name=MealsDb,proto3" json:"MealsDb,omitempty"`
}

func (x *MealsDb) Reset() {
	*x = MealsDb{}
	mi := &file_orders_admin_pkg_models_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MealsDb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MealsDb) ProtoMessage() {}

func (x *MealsDb) ProtoReflect() protoreflect.Message {
	mi := &file_orders_admin_pkg_models_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MealsDb.ProtoReflect.Descriptor instead.
func (*MealsDb) Descriptor() ([]byte, []int) {
	return file_orders_admin_pkg_models_proto_rawDescGZIP(), []int{4}
}

func (x *MealsDb) GetMealsDb() []*MealDb {
	if x != nil {
		return x.MealsDb
	}
	return nil
}

var File_orders_admin_pkg_models_proto protoreflect.FileDescriptor

var file_orders_admin_pkg_models_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2d, 0x70,
	0x6b, 0x67, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x78, 0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x65, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x65, 0x61, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x61,
	0x79, 0x22, 0x43, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x12, 0x34, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x8d, 0x01, 0x0a, 0x06, 0x4d, 0x65, 0x61, 0x6c, 0x44,
	0x62, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x23, 0x0a, 0x05, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x52,
	0x05, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x22, 0x57, 0x0a, 0x05, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x64, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x64, 0x75, 0x6d, 0x6d, 0x69, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x72, 0x75,
	0x6d, 0x70, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x74, 0x72, 0x75,
	0x6d, 0x70, 0x69, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x75, 0x70, 0x70, 0x69, 0x65, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x75, 0x70, 0x70, 0x69, 0x65, 0x73, 0x22,
	0x33, 0x0a, 0x07, 0x4d, 0x65, 0x61, 0x6c, 0x73, 0x44, 0x62, 0x12, 0x28, 0x0a, 0x07, 0x4d, 0x65,
	0x61, 0x6c, 0x73, 0x44, 0x62, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x4d, 0x65, 0x61, 0x6c, 0x44, 0x62, 0x52, 0x07, 0x4d, 0x65, 0x61,
	0x6c, 0x73, 0x44, 0x62, 0x42, 0x20, 0x5a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2d, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orders_admin_pkg_models_proto_rawDescOnce sync.Once
	file_orders_admin_pkg_models_proto_rawDescData = file_orders_admin_pkg_models_proto_rawDesc
)

func file_orders_admin_pkg_models_proto_rawDescGZIP() []byte {
	file_orders_admin_pkg_models_proto_rawDescOnce.Do(func() {
		file_orders_admin_pkg_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_orders_admin_pkg_models_proto_rawDescData)
	})
	return file_orders_admin_pkg_models_proto_rawDescData
}

var file_orders_admin_pkg_models_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_orders_admin_pkg_models_proto_goTypes = []any{
	(*OrderAdmin)(nil),  // 0: models.OrderAdmin
	(*OrdersAdmin)(nil), // 1: models.OrdersAdmin
	(*MealDb)(nil),      // 2: models.MealDb
	(*Dummy)(nil),       // 3: models.Dummy
	(*MealsDb)(nil),     // 4: models.MealsDb
}
var file_orders_admin_pkg_models_proto_depIdxs = []int32{
	0, // 0: models.OrdersAdmin.OrdersAdmin:type_name -> models.OrderAdmin
	3, // 1: models.MealDb.Dummy:type_name -> models.Dummy
	2, // 2: models.MealsDb.MealsDb:type_name -> models.MealDb
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_orders_admin_pkg_models_proto_init() }
func file_orders_admin_pkg_models_proto_init() {
	if File_orders_admin_pkg_models_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orders_admin_pkg_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orders_admin_pkg_models_proto_goTypes,
		DependencyIndexes: file_orders_admin_pkg_models_proto_depIdxs,
		MessageInfos:      file_orders_admin_pkg_models_proto_msgTypes,
	}.Build()
	File_orders_admin_pkg_models_proto = out.File
	file_orders_admin_pkg_models_proto_rawDesc = nil
	file_orders_admin_pkg_models_proto_goTypes = nil
	file_orders_admin_pkg_models_proto_depIdxs = nil
}
