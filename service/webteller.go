package service

import (
	jsoniter "github.com/json-iterator/go"

	"github.com/micro/go-micro/client"

	"gitlab.pactindo.com/ebanking/proto-common/fee"
)

var json = jsoniter.ConfigFastest

var (
	feeSvc fee.FeeService
)

func Init(mclient client.Client) {
	feeSvc = fee.NewFeeService("fee.shared", mclient)
}

type WebTellerHandler struct{}
