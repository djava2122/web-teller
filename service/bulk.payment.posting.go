package service

import (
	"context"
	"gitlab.pactindo.com/ebanking/web-teller/model"
	"gitlab.pactindo.com/ebanking/web-teller/repo"
	"time"

	"github.com/valyala/fastjson"

	"gitlab.pactindo.com/ebanking/common/log"
	"gitlab.pactindo.com/ebanking/common/transport"
	"gitlab.pactindo.com/ebanking/common/trycatch"
	"gitlab.pactindo.com/ebanking/common/util"

	wtproto "gitlab.pactindo.com/ebanking/web-teller/proto"
)

func (h *WebTellerHandler) BulkPaymentPosting(_ context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {

	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], string(res.Response))
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] error : %v", req.Headers["Request-ID"], e)
		res.Response, _ = json.Marshal(newResponse("99", "Internal Server Error"))
	})

	jsonReq, _ := json.Marshal(req.Params)
	log.Infof("[%s] request: %v", req.Headers["Request-ID"], string(jsonReq))

	newBulkPayment := make([]model.BulkPayment, 0)

	if err := json.Unmarshal([]byte(req.Params["bulk"]), &newBulkPayment); err != nil {
		log.Infof("Error Unmarshal : ", err.Error())
		return err
	}

	for _, val := range newBulkPayment {

		var inqDataObj *fastjson.Object

		if val.FeatureName != "MPN" || val.FeatureName != "INSTITUSI" {
			if val.InquiryData == "" {
				res.Response, _ = json.Marshal(newResponse("02", "invalid inquiry data"))
				return nil
			}

			inqData, err := fastjson.Parse(val.InquiryData)
			if err == nil {
				inqDataObj, err = inqData.Object()
			}
			if err != nil {
				log.Errorf("unable to parse inquiry data: %v", err)
				res.Response, _ = json.Marshal(newResponse("02", "invalid inquiry data"))
				return nil
			}
		}

		var params = map[string]string{
			"core":              req.Params["core"],
			"tellerID":          req.Params["tellerID"],
			"tellerPass":        req.Params["tellerPass"],
			"amount": 			 val.Amount,
			"fee":               val.Fee,
			"txType":            val.Txtype,
			"billerId":          val.BillerId,
			"billerProductCode": val.BillerProductCode,
			"customerId":        val.CustomerReference,
			"inqData": 			 val.CustomerReference,
			"referenceNumber":   util.RandomNumber(12),
			"termType":          "6010",
			"termId": 			 "WTELLER",
		}

		if val.FeatureName != "MPN" {
			inqDataObj.Visit(func(key []byte, v *fastjson.Value) {
				if v.Type() == fastjson.TypeString {
					params[string(key)] = string(v.GetStringBytes())
				} else {
					params[string(key)] = v.String()
				}
			})
		}

		gateMsg := transport.SendToGate("gate.shared", "12", params)
		if gateMsg.ResponseCode == "00" {
			gateMsg.Data["txDate"] = time.Now().Format("20060102 15:04:05")
			gateMsg.Data["txRefNumber"] = params["referenceNumber"]

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
	}
	return nil
}
