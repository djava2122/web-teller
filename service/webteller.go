package service

import (
	jsoniter "github.com/json-iterator/go"

	"gitlab.pactindo.com/ebanking/common/micro"

	"gitlab.pactindo.com/ebanking/proto-common/gate"
)

var json = jsoniter.ConfigFastest

var (
	gateSvc gate.GateService
)

func Init() {
	gateSvc = gate.NewGateService("gate.shared", micro.Client())
}

type WebTellerHandler struct{}
