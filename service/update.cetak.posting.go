package service

import (
	"context"
	"strconv"

	"gitlab.pactindo.com/ebanking/web-teller/repo"

	"gitlab.pactindo.com/ebanking/common/log"
	"gitlab.pactindo.com/ebanking/common/trycatch"

	wtproto "gitlab.pactindo.com/ebanking/web-teller/proto"
)

func (h *WebTellerHandler) UpdateCetak(_ context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {

	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], string(res.Response))
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] error : %v", req.Headers["Request-ID"], e)
		res.Response, _ = json.Marshal(newResponse("99", "Internal Server Error"))
	})

	jsonReq, _ := json.Marshal(req.Params)
	log.Infof("[%s] request: %v", req.Headers["Request-ID"], string(jsonReq))

	var params = map[string]string{
		"id":    req.Params["id"],
		"cetak": req.Params["cetak"],
	}

	trxData := UpdateCetakTransaksi(params)

	err := repo.Transaction.Update(trxData)
	if err != nil {
		res.Response, _ = json.Marshal(newResponse("01", "Error While Update Table"))
		log.Errorf("Error Update Transaction: %v", err)
	} else {
		res.Response, _ = json.Marshal(newResponse("00", "SUCCESS"))
	}
	return nil
}

func UpdateCetakTransaksi(params map[string]string) repo.UCetak {
	trx := repo.UCetak{}
	cetak, _ := strconv.Atoi(params["cetak"])
	trx.Id = params["id"]
	trx.Cetak = cetak
	//log.Infof("asasaaasas:", trx)
	return trx
}
