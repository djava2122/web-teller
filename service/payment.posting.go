package service

import (
	"context"

	"github.com/valyala/fastjson"

	"gitlab.pactindo.com/backend-svc/common/log"
	"gitlab.pactindo.com/backend-svc/common/transport"
	"gitlab.pactindo.com/backend-svc/common/trycatch"
	"gitlab.pactindo.com/backend-svc/common/util"

	wtproto "gitlab.pactindo.com/ebanking/web-teller/proto"
)

func (h *WebTellerHandler) PaymentPosting(_ context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {

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

		inquiryData := req.Params["inquiryData"]
		if inquiryData == "" {
			res.Response, _ = json.Marshal(newResponse("02", "invalid inquiry data"))
			return nil
		}

		var inqDataObj *fastjson.Object
		inqData, err := fastjson.Parse(inquiryData)
		if err == nil {
			inqDataObj, err = inqData.Object()
		}
		if err != nil {
			log.Errorf("unable to parse inquiry data: %v", err)
			res.Response, _ = json.Marshal(newResponse("02", "invalid inquiry data"))
			return nil
		}

		var params = map[string]string{
			"core":              req.Params["core"],
			"tellerID":          req.Params["tellerID"],
			"tellerPass":        req.Params["tellerPass"],
			"txType":            txType,
			"billerId":          billerCode,
			"billerProductCode": billerProductCode,
			"customerId":        customerReference,
			"referenceNumber":   util.PadLeftZero(req.Headers["Request-ID"], 12),
			"termType":          "6010",
		}
		inqDataObj.Visit(func(key []byte, v *fastjson.Value) {
			if v.Type() == fastjson.TypeString {
				params[string(key)] = string(v.GetStringBytes())
			} else {
				params[string(key)] = v.String()
			}
		})

		gateMsg := transport.SendToGate("gate.shared", req.TxType, params)
		if gateMsg.ResponseCode == "00" {
			res.Response, _ = json.Marshal(successResp(gateMsg.Data))
		} else {
			res.Response, _ = json.Marshal(newResponse(gateMsg.ResponseCode, gateMsg.Description))
		}
	}

	return nil
}
