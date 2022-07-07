package service

import (
	"context"
	"git.pactindo.com/ebanking/common/log"
	"git.pactindo.com/ebanking/common/transport"
	"git.pactindo.com/ebanking/common/trycatch"
	wtproto "git.pactindo.com/ebanking/web-teller/proto"
	"git.pactindo.com/ebanking/web-teller/repo"
	"strings"
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
		//log.Infof("Log Gate Auth: ", gateMsg.Data)
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
				claims.branchName = getData(gateMsg.Data, "branchName")
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
				//data["role"] = ParseRoleTeller(strings.TrimSpace(getData(userInfo.Data, "initApp")))
				//data["branchCode"] = getData(userInfo.Data, "companyCode")
				//data["role"] = ParseRoleTeller(strings.TrimSpace(getData(gateMsg.Data, "kdSPV1")))
				data["role"] = ParseRoles(strings.TrimSpace(getData(gateMsg.Data, "kdSPV1")))
				data["branchCode"] = ParseBranchCode(getData(gateMsg.Data, "companyCode"))
				data["branchName"] = getData(gateMsg.Data, "branchName")
				data["beginBalance"] = getData(gateMsg.Data, "saldoAwalHari")
				data["CurrentBalance"] = getData(gateMsg.Data, "saldoSekarang")

				if data["role"] == "A" {
					res.Response, _ = json.Marshal(newResponse("99", "Invalid Role Teller"))
					return nil
				}

				res.Response, _ = json.Marshal(successResp(data))
			}

		} else {
			res.Response, _ = json.Marshal(newResponse("02", "Invalid TellerID or Password"))
		}
	}

	return nil
}

func ParseRoles(roles string) string {
	_, role, err := repo.Transaction.FindRole(roles)
	if err != nil {
		panic(err)
	}
	if role == "" {
		return "A"
	}
	return role
}
