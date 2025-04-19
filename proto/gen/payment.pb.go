// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: payment.proto

package gen

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type KakaoReadyRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	PartnerOrderId string                 `protobuf:"bytes,1,opt,name=partner_order_id,json=partnerOrderId,proto3" json:"partner_order_id,omitempty"`
	PartnerUserId  string                 `protobuf:"bytes,2,opt,name=partner_user_id,json=partnerUserId,proto3" json:"partner_user_id,omitempty"`
	ItemName       string                 `protobuf:"bytes,3,opt,name=item_name,json=itemName,proto3" json:"item_name,omitempty"`
	Quantity       int32                  `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	TotalAmount    int64                  `protobuf:"varint,5,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	TaxFreeAmount  int64                  `protobuf:"varint,6,opt,name=tax_free_amount,json=taxFreeAmount,proto3" json:"tax_free_amount,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *KakaoReadyRequest) Reset() {
	*x = KakaoReadyRequest{}
	mi := &file_payment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KakaoReadyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KakaoReadyRequest) ProtoMessage() {}

func (x *KakaoReadyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KakaoReadyRequest.ProtoReflect.Descriptor instead.
func (*KakaoReadyRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{0}
}

func (x *KakaoReadyRequest) GetPartnerOrderId() string {
	if x != nil {
		return x.PartnerOrderId
	}
	return ""
}

func (x *KakaoReadyRequest) GetPartnerUserId() string {
	if x != nil {
		return x.PartnerUserId
	}
	return ""
}

func (x *KakaoReadyRequest) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *KakaoReadyRequest) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *KakaoReadyRequest) GetTotalAmount() int64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *KakaoReadyRequest) GetTaxFreeAmount() int64 {
	if x != nil {
		return x.TaxFreeAmount
	}
	return 0
}

type KakaoReadyResponse struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	NextRedirectAppUrl    string                 `protobuf:"bytes,1,opt,name=next_redirect_app_url,json=nextRedirectAppUrl,proto3" json:"next_redirect_app_url,omitempty"`
	NextRedirectMobileUrl string                 `protobuf:"bytes,2,opt,name=next_redirect_mobile_url,json=nextRedirectMobileUrl,proto3" json:"next_redirect_mobile_url,omitempty"`
	NextRedirectPcUrl     string                 `protobuf:"bytes,3,opt,name=next_redirect_pc_url,json=nextRedirectPcUrl,proto3" json:"next_redirect_pc_url,omitempty"`
	AndroidAppScheme      string                 `protobuf:"bytes,4,opt,name=android_app_scheme,json=androidAppScheme,proto3" json:"android_app_scheme,omitempty"`
	IosAppScheme          string                 `protobuf:"bytes,5,opt,name=ios_app_scheme,json=iosAppScheme,proto3" json:"ios_app_scheme,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *KakaoReadyResponse) Reset() {
	*x = KakaoReadyResponse{}
	mi := &file_payment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KakaoReadyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KakaoReadyResponse) ProtoMessage() {}

func (x *KakaoReadyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KakaoReadyResponse.ProtoReflect.Descriptor instead.
func (*KakaoReadyResponse) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{1}
}

func (x *KakaoReadyResponse) GetNextRedirectAppUrl() string {
	if x != nil {
		return x.NextRedirectAppUrl
	}
	return ""
}

func (x *KakaoReadyResponse) GetNextRedirectMobileUrl() string {
	if x != nil {
		return x.NextRedirectMobileUrl
	}
	return ""
}

func (x *KakaoReadyResponse) GetNextRedirectPcUrl() string {
	if x != nil {
		return x.NextRedirectPcUrl
	}
	return ""
}

func (x *KakaoReadyResponse) GetAndroidAppScheme() string {
	if x != nil {
		return x.AndroidAppScheme
	}
	return ""
}

func (x *KakaoReadyResponse) GetIosAppScheme() string {
	if x != nil {
		return x.IosAppScheme
	}
	return ""
}

type KakaoApproveRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	PgToken       string                 `protobuf:"bytes,2,opt,name=pg_token,json=pgToken,proto3" json:"pg_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KakaoApproveRequest) Reset() {
	*x = KakaoApproveRequest{}
	mi := &file_payment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KakaoApproveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KakaoApproveRequest) ProtoMessage() {}

func (x *KakaoApproveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KakaoApproveRequest.ProtoReflect.Descriptor instead.
func (*KakaoApproveRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{2}
}

func (x *KakaoApproveRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *KakaoApproveRequest) GetPgToken() string {
	if x != nil {
		return x.PgToken
	}
	return ""
}

type KakaoApproveResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KakaoApproveResponse) Reset() {
	*x = KakaoApproveResponse{}
	mi := &file_payment_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KakaoApproveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KakaoApproveResponse) ProtoMessage() {}

func (x *KakaoApproveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KakaoApproveResponse.ProtoReflect.Descriptor instead.
func (*KakaoApproveResponse) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{3}
}

func (x *KakaoApproveResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type KakaoCancelRequest struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	OrderId               string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	CancelAmount          string                 `protobuf:"bytes,2,opt,name=cancel_amount,json=cancelAmount,proto3" json:"cancel_amount,omitempty"`
	CancelTaxFreeAmount   int64                  `protobuf:"varint,3,opt,name=cancel_tax_free_amount,json=cancelTaxFreeAmount,proto3" json:"cancel_tax_free_amount,omitempty"`
	CancelVatAmount       int64                  `protobuf:"varint,4,opt,name=cancel_vat_amount,json=cancelVatAmount,proto3" json:"cancel_vat_amount,omitempty"`
	CancelAvailableAmount int64                  `protobuf:"varint,5,opt,name=cancel_available_amount,json=cancelAvailableAmount,proto3" json:"cancel_available_amount,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *KakaoCancelRequest) Reset() {
	*x = KakaoCancelRequest{}
	mi := &file_payment_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KakaoCancelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KakaoCancelRequest) ProtoMessage() {}

func (x *KakaoCancelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KakaoCancelRequest.ProtoReflect.Descriptor instead.
func (*KakaoCancelRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{4}
}

func (x *KakaoCancelRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *KakaoCancelRequest) GetCancelAmount() string {
	if x != nil {
		return x.CancelAmount
	}
	return ""
}

func (x *KakaoCancelRequest) GetCancelTaxFreeAmount() int64 {
	if x != nil {
		return x.CancelTaxFreeAmount
	}
	return 0
}

func (x *KakaoCancelRequest) GetCancelVatAmount() int64 {
	if x != nil {
		return x.CancelVatAmount
	}
	return 0
}

func (x *KakaoCancelRequest) GetCancelAvailableAmount() int64 {
	if x != nil {
		return x.CancelAvailableAmount
	}
	return 0
}

type KakaoCancelResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       string                 `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KakaoCancelResponse) Reset() {
	*x = KakaoCancelResponse{}
	mi := &file_payment_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KakaoCancelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KakaoCancelResponse) ProtoMessage() {}

func (x *KakaoCancelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KakaoCancelResponse.ProtoReflect.Descriptor instead.
func (*KakaoCancelResponse) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{5}
}

func (x *KakaoCancelResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

var File_payment_proto protoreflect.FileDescriptor

const file_payment_proto_rawDesc = "" +
	"\n" +
	"\rpayment.proto\x12\"go.escape.ship.proto.paymentapi.v1\x1a\x1cgoogle/api/annotations.proto\x1a.protoc-gen-openapiv2/options/annotations.proto\"\xe9\x01\n" +
	"\x11KakaoReadyRequest\x12(\n" +
	"\x10partner_order_id\x18\x01 \x01(\tR\x0epartnerOrderId\x12&\n" +
	"\x0fpartner_user_id\x18\x02 \x01(\tR\rpartnerUserId\x12\x1b\n" +
	"\titem_name\x18\x03 \x01(\tR\bitemName\x12\x1a\n" +
	"\bquantity\x18\x04 \x01(\x05R\bquantity\x12!\n" +
	"\ftotal_amount\x18\x05 \x01(\x03R\vtotalAmount\x12&\n" +
	"\x0ftax_free_amount\x18\x06 \x01(\x03R\rtaxFreeAmount\"\x85\x02\n" +
	"\x12KakaoReadyResponse\x121\n" +
	"\x15next_redirect_app_url\x18\x01 \x01(\tR\x12nextRedirectAppUrl\x127\n" +
	"\x18next_redirect_mobile_url\x18\x02 \x01(\tR\x15nextRedirectMobileUrl\x12/\n" +
	"\x14next_redirect_pc_url\x18\x03 \x01(\tR\x11nextRedirectPcUrl\x12,\n" +
	"\x12android_app_scheme\x18\x04 \x01(\tR\x10androidAppScheme\x12$\n" +
	"\x0eios_app_scheme\x18\x05 \x01(\tR\fiosAppScheme\"K\n" +
	"\x13KakaoApproveRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\x12\x19\n" +
	"\bpg_token\x18\x02 \x01(\tR\apgToken\"1\n" +
	"\x14KakaoApproveResponse\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\"\xed\x01\n" +
	"\x12KakaoCancelRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId\x12#\n" +
	"\rcancel_amount\x18\x02 \x01(\tR\fcancelAmount\x123\n" +
	"\x16cancel_tax_free_amount\x18\x03 \x01(\x03R\x13cancelTaxFreeAmount\x12*\n" +
	"\x11cancel_vat_amount\x18\x04 \x01(\x03R\x0fcancelVatAmount\x126\n" +
	"\x17cancel_available_amount\x18\x05 \x01(\x03R\x15cancelAvailableAmount\"0\n" +
	"\x13KakaoCancelResponse\x12\x19\n" +
	"\border_id\x18\x01 \x01(\tR\aorderId2\x87\x06\n" +
	"\x0ePaymentService\x12\xef\x01\n" +
	"\n" +
	"KakaoReady\x125.go.escape.ship.proto.paymentapi.v1.KakaoReadyRequest\x1a6.go.escape.ship.proto.paymentapi.v1.KakaoReadyResponse\"r\x92AP\n" +
	"\x0eKakao Payments\x12\x18Ready payment with Kakao\x1a$Initiate payment process with Kakao.\x82\xd3\xe4\x93\x02\x19:\x01*\"\x14/payment/kakao/ready\x12\xfc\x01\n" +
	"\fKakaoApprove\x127.go.escape.ship.proto.paymentapi.v1.KakaoApproveRequest\x1a8.go.escape.ship.proto.paymentapi.v1.KakaoApproveResponse\"y\x92AU\n" +
	"\x0eKakao Payments\x12\x1aApprove payment with Kakao\x1a'Approve the payment process with Kakao.\x82\xd3\xe4\x93\x02\x1b:\x01*\"\x16/payment/kakao/approve\x12\x83\x02\n" +
	"\vKakaoCancel\x126.go.escape.ship.proto.paymentapi.v1.KakaoCancelRequest\x1a7.go.escape.ship.proto.paymentapi.v1.KakaoCancelResponse\"\x82\x01\x92A_\n" +
	"\x0eKakao Payments\x12\x19Cancel payment with Kakao\x1a2Cancel an ongoing or completed payment with Kakao.\x82\xd3\xe4\x93\x02\x1a:\x01*\"\x15/payment/kakao/cancelB-Z+github.com/escape-ship/paymentsrv/proto/genb\x06proto3"

var (
	file_payment_proto_rawDescOnce sync.Once
	file_payment_proto_rawDescData []byte
)

func file_payment_proto_rawDescGZIP() []byte {
	file_payment_proto_rawDescOnce.Do(func() {
		file_payment_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_payment_proto_rawDesc), len(file_payment_proto_rawDesc)))
	})
	return file_payment_proto_rawDescData
}

var file_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_payment_proto_goTypes = []any{
	(*KakaoReadyRequest)(nil),    // 0: go.escape.ship.proto.paymentapi.v1.KakaoReadyRequest
	(*KakaoReadyResponse)(nil),   // 1: go.escape.ship.proto.paymentapi.v1.KakaoReadyResponse
	(*KakaoApproveRequest)(nil),  // 2: go.escape.ship.proto.paymentapi.v1.KakaoApproveRequest
	(*KakaoApproveResponse)(nil), // 3: go.escape.ship.proto.paymentapi.v1.KakaoApproveResponse
	(*KakaoCancelRequest)(nil),   // 4: go.escape.ship.proto.paymentapi.v1.KakaoCancelRequest
	(*KakaoCancelResponse)(nil),  // 5: go.escape.ship.proto.paymentapi.v1.KakaoCancelResponse
}
var file_payment_proto_depIdxs = []int32{
	0, // 0: go.escape.ship.proto.paymentapi.v1.PaymentService.KakaoReady:input_type -> go.escape.ship.proto.paymentapi.v1.KakaoReadyRequest
	2, // 1: go.escape.ship.proto.paymentapi.v1.PaymentService.KakaoApprove:input_type -> go.escape.ship.proto.paymentapi.v1.KakaoApproveRequest
	4, // 2: go.escape.ship.proto.paymentapi.v1.PaymentService.KakaoCancel:input_type -> go.escape.ship.proto.paymentapi.v1.KakaoCancelRequest
	1, // 3: go.escape.ship.proto.paymentapi.v1.PaymentService.KakaoReady:output_type -> go.escape.ship.proto.paymentapi.v1.KakaoReadyResponse
	3, // 4: go.escape.ship.proto.paymentapi.v1.PaymentService.KakaoApprove:output_type -> go.escape.ship.proto.paymentapi.v1.KakaoApproveResponse
	5, // 5: go.escape.ship.proto.paymentapi.v1.PaymentService.KakaoCancel:output_type -> go.escape.ship.proto.paymentapi.v1.KakaoCancelResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_payment_proto_init() }
func file_payment_proto_init() {
	if File_payment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_payment_proto_rawDesc), len(file_payment_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_proto_goTypes,
		DependencyIndexes: file_payment_proto_depIdxs,
		MessageInfos:      file_payment_proto_msgTypes,
	}.Build()
	File_payment_proto = out.File
	file_payment_proto_goTypes = nil
	file_payment_proto_depIdxs = nil
}
