package service

import (
	"context"
	"strconv"

	"gitlab.pactindo.com/ebanking/common/log"
	"gitlab.pactindo.com/ebanking/common/transport"
	"gitlab.pactindo.com/ebanking/common/trycatch"
	"gitlab.pactindo.com/ebanking/common/util"

	pfee "gitlab.pactindo.com/ebanking/proto-common/fee"
	wtproto "gitlab.pactindo.com/ebanking/web-teller/proto"
)

func (h *WebTellerHandler) PaymentInquiry(ctx context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {

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

	if req.Params["featureName"] != "MPN" {
		if txType == "" || billerCode == "" || billerProductCode == "" || customerReference == "" {
			res.Response, _ = json.Marshal(newResponse("01", "Bad Request"))
		}
	}

	var fee int
	switch txType {
	case "LB":
		// do nothing
	default:
		var reqGetFee = pfee.ReqFee{FeatureCode: req.Params["featureCode"], RequestId: req.Headers["Request-ID"]}
		rFee, err := feeSvc.GetFeatureFee(ctx, &reqGetFee)
		if err != nil {
			panic(err)
		}
		if rFee.Rc != "00" {
			res.Response, _ = json.Marshal(newResponse("02", "invalid fee"))
			return nil
		}
		fee = int(rFee.Fee.Charge)
	}
	log.Infof("[%s] request branch: %v", req.Headers["Request-ID"], req.Params["branchCode"])
	var substring string
	if req.Params["featureCode"] == "404" {
		branch := req.Params["branchCode"]
		substring = branch[3:9]
	} else {
		substring = req.Params["branchCode"]
	}
	numbBill := ""
	txFee := ""
	if req.Params["featureCode"] == "319" {
		numbBill = req.Params["numbBill"]
		txFee = req.Params["txFee"]
	}
	Params := map[string]string{
		"core":       req.Params["core"],
		"branchCode": substring,
		// "tellerID":          req.Params["tellerID"],
		// "tellerPass":        req.Params["tellerPass"],
		"txType":            txType,
		"numbBill":          numbBill,
		"fee":               txFee,
		"billerId":          billerCode,
		"billerProductCode": billerProductCode,
		"customerId":        customerReference,
		"referenceNumber":   util.RandomNumber(12),
		"termType":          "6010",
		"termId":            "WTELLER",
	}
	typeTamp := txType
	if req.Params["featureCode"] == "404" {
		typeTamp = "25"
	} else {
		typeTamp = req.TxType
	}
	if req.Params["featureCode"] == "318" {
		Params["norangka"] = req.Params["customerReference2"]
	}
	gateMsg := transport.SendToGate("gate.shared", typeTamp, Params)
	log.Infof("LOG Get :", gateMsg)
	if gateMsg.ResponseCode == "00" {

		switch txType {
		case "LB": // local biller
			gateMsg.Data["inquiryData"] = map[string]interface{}{
				"amount":    gateMsg.Data["amount"],
				"refnum":    gateMsg.Data["refnum"],
				"totalBill": gateMsg.Data["totalBill"],
				"rpTag":     gateMsg.Data["rpTag"],
				"rpFee":     gateMsg.Data["rpFee"],
			}
			gateMsg.Data["txFee"] = gateMsg.Data["rpFeeStruk"]
		case "01", "02", "03", "04": // PLN Prepaid, Telco Postpaid, PLN Postpaid, PLN Non Taglis
			var amount float64 = 0
			var err error
			if v, ok := gateMsg.Data["billAmount"]; ok && v != "" {
				amount, err = strconv.ParseFloat(v.(string), 64)
				if err != nil {
					panic("error parsing amount [billAmount]")
				}
			}

			total := amount + float64(fee)

			gateMsg.Data["txFee"] = strconv.Itoa(fee)
			gateMsg.Data["total"] = strconv.FormatFloat(total, 'f', 0, 64)

			gateMsg.Data["inquiryData"] = map[string]interface{}{
				"inqData": gateMsg.Data["inqData"],
				"amount":  amount,
				"fee":     fee,
			}
			delete(gateMsg.Data, "inqData")
		default:
			gateMsg.Data["txFee"] = strconv.Itoa(fee)
		}

		res.Response, _ = json.Marshal(successResp(gateMsg.Data))
	} else {
		res.Response, _ = json.Marshal(newResponse(gateMsg.ResponseCode, gateMsg.Description))
	}

	return nil
}
