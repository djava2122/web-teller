package service

import (
	"context"
	"git.pactindo.com/ebanking/common/log"
	"git.pactindo.com/ebanking/common/transport"
	"git.pactindo.com/ebanking/common/trycatch"
	"git.pactindo.com/ebanking/common/util"
	pfee "git.pactindo.com/ebanking/proto-common/fee"
	wtproto "git.pactindo.com/ebanking/web-teller/proto"
	"strconv"
)

func (h *WebTellerHandler) TransferInquiry(ctx context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {
	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], string(res.Response))
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] error : %v", req.Headers["Request-ID"], e)
		res.Response, _ = json.Marshal(newResponse("99", "Internal Server Error"))
	})

	jsonReq, _ := json.Marshal(req.Params)
	log.Infof("[%s] request: %v", req.Headers["Request-ID"], string(jsonReq))

	var fee int
	var reqGetFee = pfee.ReqFee{FeatureCode: req.Params["featureCode"], RequestId: req.Headers["Request-ID"]}
	rFee, err := feeSvc.GetFeatureFee(ctx, &reqGetFee)
	if err != nil {
		panic(err)
	}

	if rFee.Rc != "00" {
		res.Response, _ = json.Marshal(newResponse("02", "invalid fee"))
		return nil
	}

	fee = int(rFee.Fee.Charge)

	gateMsg := transport.SendToGate("gate.shared", req.TxType, map[string]string{
		"core": req.Params["core"],
		// "tellerID":          req.Params["tellerID"],
		// "tellerPass":        req.Params["tellerPass"],
		"txType":          req.Params["txType"],
		"destAccount":     req.Params["destAccount"],
		"referenceNumber": util.PadLeftZero(req.Headers["Request-ID"], 12),
		"termType":        "6010",
		"termId":          "WTELLER",
	})

	if gateMsg.ResponseCode == "00" {

		var amount float64 = 0
		var err error
		if v, ok := gateMsg.Data["amount"]; ok && v != "" {
			amount, err = strconv.ParseFloat(v.(string), 64)
			if err != nil {
				panic("error parsing amount [billAmount]")
			}
		}

		total := amount + float64(fee)

		gateMsg.Data["txFee"] = strconv.Itoa(fee)
		gateMsg.Data["total"] = strconv.FormatFloat(total, 'f', 0, 64)

		res.Response, _ = json.Marshal(successResp(gateMsg.Data))

	} else {
		res.Response, _ = json.Marshal(newResponse(gateMsg.ResponseCode, gateMsg.Description))
	}

	return nil
}
