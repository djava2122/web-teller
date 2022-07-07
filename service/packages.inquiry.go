package service

import (
	"context"
	"git.pactindo.com/ebanking/common/log"
	"git.pactindo.com/ebanking/common/transport"
	"git.pactindo.com/ebanking/common/trycatch"
	wtproto "git.pactindo.com/ebanking/web-teller/proto"
	"strconv"
)

func (h *WebTellerHandler) InquiryPackages(ctx context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {
	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], string(res.Response))
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] error : %v", req.Headers["Request-ID"], e)
		res.Response, _ = json.Marshal(newResponse("99", "Internal Server Error"))
	})

	gateMsg := transport.SendToGate("gate.shared", "AJ002", map[string]string{
		"customerId":  req.Params["customerId"],
		"productCode": req.Params["productCode"],
		"srcAccType":  "00",
		"termId":      "WTELLER",
		"termType":    "6010",
	})

	if gateMsg.ResponseCode == "00" {
		price, _ := strconv.ParseFloat(gateMsg.Data["voucherAmount"].(string), 64)
		gateMsg.Data["voucherAmount"] = strconv.FormatFloat(price/100, 'f', 0, 64)
		res.Response, _ = json.Marshal(successResp(gateMsg.Data))
	} else {
		res.Response, _ = json.Marshal(newResponse(gateMsg.ResponseCode, gateMsg.Description))
	}

	return nil
}
