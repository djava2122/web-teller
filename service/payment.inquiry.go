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

	jsonReq, _ := json.Marshal(req)
	log.Infof("[%s] request: %v", req.Headers["Request-ID"], string(jsonReq))

	txType := req.Params["txType"]
	billerCode := req.Params["billerCode"]
	billerProductCode := req.Params["billerProductCode"]
	customerReference := req.Params["customerReference"]

	if txType == "" || billerCode == "" || billerProductCode == "" || customerReference == "" {
		res.Response, _ = json.Marshal(newResponse("01", "Bad Request"))
	} else {

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

		gateMsg := transport.SendToGate("gate.shared", req.TxType, map[string]string{
			"core": req.Params["core"],
			// "tellerID":          req.Params["tellerID"],
			// "tellerPass":        req.Params["tellerPass"],
			"txType":            txType,
			"billerId":          billerCode,
			"billerProductCode": billerProductCode,
			"customerId":        customerReference,
			"referenceNumber":   util.PadLeftZero(req.Headers["Request-ID"], 12),
			"termType":          "6010",
		})
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
			}

			res.Response, _ = json.Marshal(successResp(gateMsg.Data))
		} else {
			res.Response, _ = json.Marshal(newResponse(gateMsg.ResponseCode, gateMsg.Description))
		}
	}

	return nil
}
