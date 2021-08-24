// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: web-teller.proto

package wtproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for WebTeller service

type WebTellerService interface {
	Authentication(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error)
	SessionValidate(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error)
	PaymentInquiry(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error)
	PaymentPosting(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error)
	BulkPaymentPosting(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error)
	TransferInquiry(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error)
	TransferPosting(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error)
	TransactionReport(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error)
	CashTellerInquiry(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error)
	InquiryNomorRekening(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error)
	UpdateCetak(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error)
}

type webTellerService struct {
	c    client.Client
	name string
}

func NewWebTellerService(name string, c client.Client) WebTellerService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "wtproto"
	}
	return &webTellerService{
		c:    c,
		name: name,
	}
}

func (c *webTellerService) Authentication(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error) {
	req := c.c.NewRequest(c.name, "WebTeller.Authentication", in)
	out := new(APIRES)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webTellerService) SessionValidate(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error) {
	req := c.c.NewRequest(c.name, "WebTeller.SessionValidate", in)
	out := new(APIRES)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webTellerService) PaymentInquiry(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error) {
	req := c.c.NewRequest(c.name, "WebTeller.PaymentInquiry", in)
	out := new(APIRES)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webTellerService) PaymentPosting(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error) {
	req := c.c.NewRequest(c.name, "WebTeller.PaymentPosting", in)
	out := new(APIRES)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webTellerService) BulkPaymentPosting(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error) {
	req := c.c.NewRequest(c.name, "WebTeller.BulkPaymentPosting", in)
	out := new(APIRES)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webTellerService) TransferInquiry(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error) {
	req := c.c.NewRequest(c.name, "WebTeller.TransferInquiry", in)
	out := new(APIRES)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webTellerService) TransferPosting(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error) {
	req := c.c.NewRequest(c.name, "WebTeller.TransferPosting", in)
	out := new(APIRES)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webTellerService) TransactionReport(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error) {
	req := c.c.NewRequest(c.name, "WebTeller.TransactionReport", in)
	out := new(APIRES)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webTellerService) CashTellerInquiry(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error) {
	req := c.c.NewRequest(c.name, "WebTeller.CashTellerInquiry", in)
	out := new(APIRES)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webTellerService) InquiryNomorRekening(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error) {
	req := c.c.NewRequest(c.name, "WebTeller.InquiryNomorRekening", in)
	out := new(APIRES)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webTellerService) UpdateCetak(ctx context.Context, in *APIREQ, opts ...client.CallOption) (*APIRES, error) {
	req := c.c.NewRequest(c.name, "WebTeller.UpdateCetak", in)
	out := new(APIRES)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WebTeller service

type WebTellerHandler interface {
	Authentication(context.Context, *APIREQ, *APIRES) error
	SessionValidate(context.Context, *APIREQ, *APIRES) error
	PaymentInquiry(context.Context, *APIREQ, *APIRES) error
	PaymentPosting(context.Context, *APIREQ, *APIRES) error
	BulkPaymentPosting(context.Context, *APIREQ, *APIRES) error
	TransferInquiry(context.Context, *APIREQ, *APIRES) error
	TransferPosting(context.Context, *APIREQ, *APIRES) error
	TransactionReport(context.Context, *APIREQ, *APIRES) error
	CashTellerInquiry(context.Context, *APIREQ, *APIRES) error
	InquiryNomorRekening(context.Context, *APIREQ, *APIRES) error
	UpdateCetak(context.Context, *APIREQ, *APIRES) error
}

func RegisterWebTellerHandler(s server.Server, hdlr WebTellerHandler, opts ...server.HandlerOption) error {
	type webTeller interface {
		Authentication(ctx context.Context, in *APIREQ, out *APIRES) error
		SessionValidate(ctx context.Context, in *APIREQ, out *APIRES) error
		PaymentInquiry(ctx context.Context, in *APIREQ, out *APIRES) error
		PaymentPosting(ctx context.Context, in *APIREQ, out *APIRES) error
		BulkPaymentPosting(ctx context.Context, in *APIREQ, out *APIRES) error
		TransferInquiry(ctx context.Context, in *APIREQ, out *APIRES) error
		TransferPosting(ctx context.Context, in *APIREQ, out *APIRES) error
		TransactionReport(ctx context.Context, in *APIREQ, out *APIRES) error
		CashTellerInquiry(ctx context.Context, in *APIREQ, out *APIRES) error
		InquiryNomorRekening(ctx context.Context, in *APIREQ, out *APIRES) error
		UpdateCetak(ctx context.Context, in *APIREQ, out *APIRES) error
	}
	type WebTeller struct {
		webTeller
	}
	h := &webTellerHandler{hdlr}
	return s.Handle(s.NewHandler(&WebTeller{h}, opts...))
}

type webTellerHandler struct {
	WebTellerHandler
}

func (h *webTellerHandler) Authentication(ctx context.Context, in *APIREQ, out *APIRES) error {
	return h.WebTellerHandler.Authentication(ctx, in, out)
}

func (h *webTellerHandler) SessionValidate(ctx context.Context, in *APIREQ, out *APIRES) error {
	return h.WebTellerHandler.SessionValidate(ctx, in, out)
}

func (h *webTellerHandler) PaymentInquiry(ctx context.Context, in *APIREQ, out *APIRES) error {
	return h.WebTellerHandler.PaymentInquiry(ctx, in, out)
}

func (h *webTellerHandler) PaymentPosting(ctx context.Context, in *APIREQ, out *APIRES) error {
	return h.WebTellerHandler.PaymentPosting(ctx, in, out)
}

func (h *webTellerHandler) BulkPaymentPosting(ctx context.Context, in *APIREQ, out *APIRES) error {
	return h.WebTellerHandler.BulkPaymentPosting(ctx, in, out)
}

func (h *webTellerHandler) TransferInquiry(ctx context.Context, in *APIREQ, out *APIRES) error {
	return h.WebTellerHandler.TransferInquiry(ctx, in, out)
}

func (h *webTellerHandler) TransferPosting(ctx context.Context, in *APIREQ, out *APIRES) error {
	return h.WebTellerHandler.TransferPosting(ctx, in, out)
}

func (h *webTellerHandler) TransactionReport(ctx context.Context, in *APIREQ, out *APIRES) error {
	return h.WebTellerHandler.TransactionReport(ctx, in, out)
}

func (h *webTellerHandler) CashTellerInquiry(ctx context.Context, in *APIREQ, out *APIRES) error {
	return h.WebTellerHandler.CashTellerInquiry(ctx, in, out)
}

func (h *webTellerHandler) InquiryNomorRekening(ctx context.Context, in *APIREQ, out *APIRES) error {
	return h.WebTellerHandler.InquiryNomorRekening(ctx, in, out)
}

func (h *webTellerHandler) UpdateCetak(ctx context.Context, in *APIREQ, out *APIRES) error {
	return h.WebTellerHandler.UpdateCetak(ctx, in, out)
}
