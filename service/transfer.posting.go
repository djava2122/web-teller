package service

import (
	"context"
	"git.pactindo.com/ebanking/common/log"
	"git.pactindo.com/ebanking/common/transport"
	"git.pactindo.com/ebanking/common/trycatch"
	"git.pactindo.com/ebanking/common/util"
	wtproto "git.pactindo.com/ebanking/web-teller/proto"
	"git.pactindo.com/ebanking/web-teller/repo"
)

func (h *WebTellerHandler) TransferPosting(_ context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {
	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], string(res.Response))
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] error : %v", req.Headers["Request-ID"], e)
		res.Response, _ = json.Marshal(newResponse("99", "Internal Server Error"))
	})

	jsonReq, _ := json.Marshal(req.Params)
	log.Infof("[%s] request: %v", req.Headers["Request-ID"], string(jsonReq))

	var srcAccount string
	switch req.Params["core"] {
	case "K":
		srcAccount = "1000000000"
	case "S":
		srcAccount = "6000000000"
	default:
		srcAccount = "1000000000"
	}

	params := map[string]string{
		"tellerID":        req.Params["tellerID"],
		"tellerPass":      req.Params["tellerPass"],
		"txType":          req.Params["txType"],
		"amount":          req.Params["amount"],
		"fee":             req.Params["fee"],
		"srcAccount":      srcAccount,
		"destAccount":     req.Params["destAccount"],
		"referenceNumber": util.RandomNumber(12),
		"termType":        "6010",
		"termId":          "WTELLER",
	}

	gateMsg := transport.SendToGate("gate.shared", req.TxType, params)

	if gateMsg.ResponseCode == "00" {
		gateMsg.Data = make(map[string]interface{})
		gateMsg.Data["txDate"] = FormattedTime(req.Params["txDate"], "20060102 15:04:05")
		gateMsg.Data["txRefNumber"] = params["referenceNumber"]
		gateMsg.Data["txStatus"] = "SUCCESS"

		res.Response, _ = json.Marshal(successResp(gateMsg.Data))

		trxData := BuildDataTransaction(req.Params, params, "SUCCESS", gateMsg.ResponseCode)

		err := repo.Transaction.Save(trxData)
		if err != nil {
			log.Errorf("error save transaction: %v", err)
		}
	} else {
		gateMsg.Data = make(map[string]interface{})
		gateMsg.Data["txDate"] = FormattedTime(req.Params["txDate"], "20060102 15:04:05")
		gateMsg.Data["txRefNumber"] = params["referenceNumber"]
		gateMsg.Data["txStatus"] = "FAILED"

		res.Response, _ = json.Marshal(newResponseWithData(gateMsg.ResponseCode, ParseRC(gateMsg.ResponseCode), gateMsg.Data))

		trxData := BuildDataTransaction(req.Params, params, "FAILED", gateMsg.ResponseCode)

		err := repo.Transaction.Save(trxData)
		if err != nil {
			log.Errorf("error save transaction: %v", err)
		}
	}

	return nil
}
