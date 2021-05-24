// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: web-teller.proto

package proto

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
		name = "proto"
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

// Server API for WebTeller service

type WebTellerHandler interface {
	Authentication(context.Context, *APIREQ, *APIRES) error
	SessionValidate(context.Context, *APIREQ, *APIRES) error
	PaymentInquiry(context.Context, *APIREQ, *APIRES) error
	PaymentPosting(context.Context, *APIREQ, *APIRES) error
}

func RegisterWebTellerHandler(s server.Server, hdlr WebTellerHandler, opts ...server.HandlerOption) error {
	type webTeller interface {
		Authentication(ctx context.Context, in *APIREQ, out *APIRES) error
		SessionValidate(ctx context.Context, in *APIREQ, out *APIRES) error
		PaymentInquiry(ctx context.Context, in *APIREQ, out *APIRES) error
		PaymentPosting(ctx context.Context, in *APIREQ, out *APIRES) error
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
