package service

import (
	"context"
	"gitlab.pactindo.com/ebanking/common/log"
	"gitlab.pactindo.com/ebanking/common/transport"
	"gitlab.pactindo.com/ebanking/common/trycatch"
	wtproto "gitlab.pactindo.com/ebanking/web-teller/proto"
)

func (h *WebTellerHandler) CashTellerInquiry(ctx context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {
	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], string(res.Response))
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] error : %v", req.Headers["Request-ID"], e)
		res.Response, _ = json.Marshal(newResponse("99", "Internal Server Error"))
	})

	jsonReq, _ := json.Marshal(req)
	log.Infof("[%s] request: %v", req.Headers["Request-ID"], string(jsonReq))

	core := req.Params["core"]
	id := req.Params["tellerID"]
	pass := req.Params["tellerPass"]

	if core == "" || id == "" || pass == "" {
		res.Response, _ = json.Marshal(newResponse("01", "Bad Request"))
	} else if core != "K" && core != "S" {
		res.Response, _ = json.Marshal(newResponse("01", "Bad Request"))
	} else {
		gateMsg := transport.SendToGate("gate.shared", "10", req.Params)
		if gateMsg.ResponseCode == "00" {

			//userInfo := transport.SendToGate("gate.shared", "11", map[string]string{
			//	"name": id,
			//	"core": core,
			//})

			//if userInfo.ResponseCode == "00" {
				data := make(map[string]interface{})
				data["tellerName"] = getData(gateMsg.Data, "userName")
				data["role"] = ParseRoleTeller(getData(gateMsg.Data, "kdSPV1"))
				data["branchCode"] = getData(gateMsg.Data, "companyCode")
				data["beginBalance"] = getData(gateMsg.Data, "saldoAwalHari")
				data["CurrentBalance"] = getData(gateMsg.Data, "saldoSekarang")

				res.Response, _ = json.Marshal(successResp(data))
			//}

		} else {
			res.Response, _ = json.Marshal(newResponse("02", "Invalid TellerID or Password"))
		}
	}

	return nil
}
