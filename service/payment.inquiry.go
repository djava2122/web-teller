package service

import (
	"context"

	"gitlab.pactindo.com/backend-svc/common/log"
	"gitlab.pactindo.com/backend-svc/common/transport"
	"gitlab.pactindo.com/backend-svc/common/trycatch"
	"gitlab.pactindo.com/backend-svc/common/util"

	wtproto "gitlab.pactindo.com/ebanking/web-teller/proto"
)

func (h *WebTellerHandler) PaymentInquiry(_ context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {

	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], string(res.Response))
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] error : %v", req.Headers["Request-ID"], e)
		res.Response, _ = json.Marshal(newResponse("99", "Internal Server Error"))
	})

	jsonReq, _ := json.Marshal(req)
	log.Infof("[%s] request: %v", req.Headers["Request-ID"], string(jsonReq))

	txType := req.Params["txType"]
	billerCode := req.Params["billerCode"]
	billerProductCode := req.Params["billerProductCode"]
	customerReference := req.Params["customerReference"]

	if txType == "" || billerCode == "" || billerProductCode == "" || customerReference == "" {
		res.Response, _ = json.Marshal(newResponse("01", "Bad Request"))
	} else {
		// req.Params["referenceNumber"] = util.PadLeftZero(req.Headers["Request-ID"], 12)

		gateMsg := transport.SendToGate("gate.shared", req.TxType, map[string]string{
			"core": req.Params["core"],
			// "tellerID":          req.Params["tellerID"],
			// "tellerPass":        req.Params["tellerPass"],
			"txType":            txType,
			"billerId":          billerCode,
			"billerProductCode": billerProductCode,
			"customerId":        customerReference,
			"referenceNumber":   util.PadLeftZero(req.Headers["Request-ID"], 12),
		})
		if gateMsg.ResponseCode == "00" {

			switch txType {
			case "LB": // local biller
				gateMsg.Data["inquiryData"] = map[string]interface{}{
					"amount":    gateMsg.Data["amount"],
					"refnum":    gateMsg.Data["refnum"],
					"totalBill": gateMsg.Data["totalBill"],
					"rpTag":     gateMsg.Data["rpTag"],
					"rpFee":     gateMsg.Data["rpFee"],
				}
			}

			res.Response, _ = json.Marshal(successResp(gateMsg.Data))
		} else {
			res.Response, _ = json.Marshal(newResponse(gateMsg.ResponseCode, gateMsg.Description))
		}
	}

	return nil
}
