package service

import (
	"context"

	"gitlab.pactindo.com/ebanking/common/log"
	"gitlab.pactindo.com/ebanking/common/trycatch"

	wtproto "gitlab.pactindo.com/ebanking/web-teller/proto"
)

func (h *WebTellerHandler) SessionValidate(ctx context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {

	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], string(res.Response))
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] error : %v", req.Headers["Request-ID"], e)
		res.Response, _ = json.Marshal(newResponse("99", "Internal Server Error"))
	})

	jsonReq, _ := json.Marshal(req)
	log.Infof("[%s] request: %v", req.Headers["Request-ID"], string(jsonReq))

	token, e := req.Headers["Authorization"]
	if !e {
		res.Response, _ = json.Marshal(newResponse("SE", "invalid session"))
		return nil
	}

	// -- get claim data by token jwt
	claim, err := getClaims(token)
	if err != nil {
		res.Response, _ = json.Marshal(newResponse("SE", "invalid session"))
	} else {
		res.Response, _ = json.Marshal(response2{
			Code:    "00",
			Message: "Success",
			Data:    claim,
		})
	}

	return nil
}
