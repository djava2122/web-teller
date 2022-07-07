package service

import (
	"bytes"
	"context"
	"git.pactindo.com/ebanking/common/log"
	"git.pactindo.com/ebanking/common/micro"
	"git.pactindo.com/ebanking/common/transport"
	"git.pactindo.com/ebanking/common/trycatch"
	wtproto "git.pactindo.com/ebanking/web-teller/proto"
	"net/http"
)

func (h *WebTellerHandler) InquirySiskohat(ctx context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {

	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], string(res.Response))
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] error : %v", req.Headers["Request-ID"], e)
		res.Response, _ = json.Marshal(newResponse("99", "Internal Server Error"))
	})

	jsonReq, _ := json.Marshal(req.Params)
	log.Infof("[%s] request: %v", req.Headers["Request-ID"], string(jsonReq))

	gateMsg := transport.SendToGate("gate.shared", "29", map[string]string{
		"account": req.Params["account"],
	})

	if gateMsg.ResponseCode == "00" {
		res.Response, _ = json.Marshal(successResp(gateMsg.Data))
	} else {
		res.Response, _ = json.Marshal(newResponse(gateMsg.ResponseCode, gateMsg.Description))
	}
	return nil
}

func (h *WebTellerHandler) LoginSiskopatuh(ctx context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {
	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], string(res.Response))
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] error : %v", req.Headers["Request-ID"], e)
		res.Response, _ = json.Marshal(newResponse("99", "Internal Server Error"))
	})

	conf := micro.GetConfig()
	var client = &http.Client{}
	var data ResSiskopatuh

	postData := map[string]interface{}{
		"userid":   conf["USER_SIKOPATUH"],
		"password": conf["PASS_SIKOPATUH"],
	}

	reqBody, _ := json.Marshal(postData)

	request, _ := http.NewRequest("POST", conf["URL_SIKOPATUH"], bytes.NewReader(reqBody))
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-key", "98erk34dj")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if response != nil {
		err := json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			panic(err)
		}
	}

	res.Response, _ = json.Marshal(newResponseWithData(data.RC, data.Message, map[string]interface{}{
		"token": data.Token,
	}))
	return nil
}
