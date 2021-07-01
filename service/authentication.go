package service

import (
	"context"

	"gitlab.pactindo.com/ebanking/common/log"
	"gitlab.pactindo.com/ebanking/common/transport"
	"gitlab.pactindo.com/ebanking/common/trycatch"

	wtproto "gitlab.pactindo.com/ebanking/web-teller/proto"
)

func (h *WebTellerHandler) Authentication(ctx context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {

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
		gateMsg := transport.SendToGate("gate.shared", req.TxType, req.Params)
		if gateMsg.ResponseCode == "00" {

			userInfo := transport.SendToGate("gate.shared", "11", map[string]string{
				"name": id,
				"core": core,
			})

			if userInfo.ResponseCode == "00" {
				claims := new(Claims)
				claims.Core = core
				claims.TellerID = id
				claims.TellerPass = pass
				claims.CoCode = getData(gateMsg.Data, "coCode")
				claims.TillCoCode = getData(gateMsg.Data, "tillCoCode")
				claims.CompanyCode = getData(gateMsg.Data, "companyCode")
				claims.BeginBalance = getData(gateMsg.Data, "saldoAwalHari")
				claims.CurrentBalance = getData(gateMsg.Data, "saldoSekarang")

				token, err := createToken(claims)
				if err != nil {
					log.Errorf("generate token failed: %v", err)
					panic(err)
				}

				data := make(map[string]interface{})
				data["token"] = token
				data["tellerName"] = getData(gateMsg.Data, "userName")
				data["role"] = getData(userInfo.Data, "initApp")

				res.Response, _ = json.Marshal(successResp(data))
			}

		} else {
			res.Response, _ = json.Marshal(newResponse("02", "Invalid TellerID or Password"))
		}
	}

	return nil
}
