package service

import (
	"context"

	"gitlab.pactindo.com/ebanking/common/log"
	"gitlab.pactindo.com/ebanking/common/transport"
	"gitlab.pactindo.com/ebanking/common/trycatch"
	wtproto "gitlab.pactindo.com/ebanking/web-teller/proto"
	"gitlab.pactindo.com/ebanking/web-teller/repo"
)

func (h *WebTellerHandler) InquiryNomorRekening(ctx context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {
	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], string(res.Response))
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] error : %v", req.Headers["Request-ID"], e)
		res.Response, _ = json.Marshal(newResponse("99", "Internal Server Error"))
	})
	referenceNumber := req.Params["referenceNumber"]
	feature := req.Params["feature"]
	log.Infof("[%s] request: %s", req.Headers["Request-ID"], feature)
	log.Infof("[%s] request: %s", req.Headers["Request-ID"], referenceNumber)

	// jsonReq, _ := json.Marshal(req)
	// log.Infof("[%s] request: %v", req.Headers["Request-ID"], string(jsonReq))
	if feature == "receipt" {
		receipt := repo.Transaction.GetTransactionReceipt(referenceNumber)
		log.Infof("[%s] request: %s", req.Headers["Request-ID"], receipt)
		if receipt != nil {
			res.Response, _ = json.Marshal(successResp(receipt))
		} else {
			res.Response, _ = json.Marshal(newResponse("80", "Data Not Found"))
		}
	} else {
		srcAccount := req.Params["srcAccount"]
		req.Params["account"] = req.Params["srcAccount"]

		if srcAccount == "" {
			res.Response, _ = json.Marshal(newResponse("01", "Bad Request"))
		} else {
			gateMsg := transport.SendToGate("gate.shared", "01", req.Params)
			log.Infof("[%s] Info: %v", req.Params)
			if gateMsg.ResponseCode == "00" {
				res.Response, _ = json.Marshal(successResp(gateMsg.Data))
			} else {
				res.Response, _ = json.Marshal(newResponse(gateMsg.ResponseCode, "Data Not Found"))
			}
		}
	}

	return nil
}
