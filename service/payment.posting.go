package service

import (
	"context"
	"strconv"
	"time"

	"gitlab.pactindo.com/ebanking/web-teller/repo"

	"github.com/valyala/fastjson"

	"gitlab.pactindo.com/ebanking/common/log"
	"gitlab.pactindo.com/ebanking/common/transport"
	"gitlab.pactindo.com/ebanking/common/trycatch"
	"gitlab.pactindo.com/ebanking/common/util"

	wtproto "gitlab.pactindo.com/ebanking/web-teller/proto"
)

func (h *WebTellerHandler) PaymentPosting(_ context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {

	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], string(res.Response))
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] error : %v", req.Headers["Request-ID"], e)
		res.Response, _ = json.Marshal(newResponse("99", "Internal Server Error"))
	})

	jsonReq, _ := json.Marshal(req.Params)
	log.Infof("[%s] request: %v", req.Headers["Request-ID"], string(jsonReq))

	txType := req.Params["txType"]
	billerCode := req.Params["billerCode"]
	billerProductCode := req.Params["billerProductCode"]
	customerReference := req.Params["customerReference"]
	featureName := req.Params["featureName"]
	inquiryData := req.Params["inquiryData"]

	var inqDataObj *fastjson.Object

	if featureName != "MPN" {
		if txType == "" || billerCode == "" || billerProductCode == "" || customerReference == "" {
			res.Response, _ = json.Marshal(newResponse("01", "Bad Request"))
		}
	}

	txDate := time.Now()

	if featureName != "MPN" || featureName != "INSTITUSI" {
		if inquiryData == "" {
			res.Response, _ = json.Marshal(newResponse("02", "invalid inquiry data"))
			return nil
		}

		inqData, err := fastjson.Parse(inquiryData)
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
		"amount":            req.Params["amount"],
		"fee":               req.Params["fee"],
		"txType":            txType,
		"billerId":          billerCode,
		"billerProductCode": billerProductCode,
		"customerId":        customerReference,
		"inqData":           inquiryData,
		"referenceNumber":   util.RandomNumber(12),
		"termType":          "6010",
		"termId":            "WTELLER",
	}

	if featureName != "MPN" {
		inqDataObj.Visit(func(key []byte, v *fastjson.Value) {
			if v.Type() == fastjson.TypeString {
				params[string(key)] = string(v.GetStringBytes())
			} else {
				params[string(key)] = v.String()
			}
		})
	}

	gateMsg := transport.SendToGate("gate.shared", req.TxType, params)
	if gateMsg.ResponseCode == "00" {
		gateMsg.Data["txDate"] = txDate.Format("20060102 15:04:05")
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

	return nil
}

func BuildDataTransaction(data map[string]string, params map[string]string, status string, code string) repo.MTransaction {
	trx := repo.MTransaction{}
	trx.ReferenceNumber = params["referenceNumber"]
	trx.FeatureId, _ = strconv.Atoi(data["featureId"])
	trx.FeatureCode, _ = strconv.Atoi(data["featureCode"])
	trx.FeatureName = params["featureName"]
	trx.FeatureGroupCode = data["featureGroupCode"]
	trx.FeatureGroupName = data["featureGroupName"]
	trx.ProductId, _ = strconv.Atoi(data["billerProductId"])
	trx.ProductCode = data["billerProductCode"]
	trx.ProductName = data["billerCode"]
	trx.BillerName = data["billerProductCode"]
	trx.CustomerReference = params["customerId"]
	trx.TransactionDate = time.Now().Format("20060102 15:04:05")
	trx.TransactionAmount, _ = strconv.ParseFloat(params["amount"], 64)
	trx.CurrencyCode = "IDR"
	trx.MerchantType = "6010"
	trx.Created = time.Now().Format("2006-01-02 15:04:05.000")
	trx.CreatedBy = data["tellerID"]
	trx.Updated = time.Now().Format("2006-01-02 15:04:05.000")
	trx.UpdatedBy = data["tellerID"]
	trx.TransactionStatus = status
	trx.BranchCode = data["branchCode"]
	trx.ResponseCode = code
	return trx
}
