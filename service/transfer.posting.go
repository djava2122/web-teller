package service

import (
	"context"
	"gitlab.pactindo.com/ebanking/common/log"
	"gitlab.pactindo.com/ebanking/common/transport"
	"gitlab.pactindo.com/ebanking/common/trycatch"
	"gitlab.pactindo.com/ebanking/common/util"
	wtproto "gitlab.pactindo.com/ebanking/web-teller/proto"
	"gitlab.pactindo.com/ebanking/web-teller/repo"
)

func (h *WebTellerHandler) TransferPosting(_ context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {
	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], string(res.Response))
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] error : %v", req.Headers["Request-ID"], e)
		res.Response, _ = json.Marshal(newResponse("99", "Internal Server Error"))
	})

	params := map[string]string{
		"core": 			 req.Params["core"],
		"tellerID":          req.Params["tellerID"],
		"tellerPass":        req.Params["tellerPass"],
		"txType":            req.Params["txType"],
		"amount": 			 req.Params["amount"],
		"fee":               req.Params["fee"],
		"destAccount":       req.Params["destAccount"],
		"referenceNumber":   util.PadLeftZero(req.Headers["Request-ID"], 12),
		"termType":          "6010",
		"termId": 			 "WTELLER",
	}

	gateMsg := transport.SendToGate("gate.shared", req.TxType, params)

	if gateMsg.ResponseCode == "00" {
		//gateMsg.Data["txDate"] = txDate.Format("20060102 15:04:05")
		//gateMsg.Data["txRefNumber"] = params["referenceNumber"]

		res.Response, _ = json.Marshal(successResp(gateMsg.Data))

		trxData := BuildDataTransaction(req.Params, params, "SUCCESS", gateMsg.ResponseCode)

		err := repo.Transaction.Save(trxData)
		if err != nil {
			log.Errorf("error save transaction: %v", err)
		}
	} else {
		res.Response, _ = json.Marshal(newResponse(gateMsg.ResponseCode, gateMsg.Description))

		trxData := BuildDataTransaction(req.Params, params, "FAILED", gateMsg.ResponseCode)

		err := repo.Transaction.Save(trxData)
		if err != nil {
			log.Errorf("error save transaction: %v", err)
		}
	}

	return nil
}
