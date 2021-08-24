package service

import (
	"context"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fastjson"
	"gitlab.pactindo.com/ebanking/web-teller/model"
	"gitlab.pactindo.com/ebanking/web-teller/repo"

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

	newDataTrx := make([]map[string]interface{}, 0)
	newBulkPayment := make([]model.BulkPayment, 0)
	gateMsg := transport.GateMsg{}
	params := map[string]string{}

	if err := json.Unmarshal([]byte(req.Params["bulk"]), &newBulkPayment); err != nil {
		log.Infof("Error Unmarshal : ", err.Error())
		return err
	}

	for _, val := range newBulkPayment {

		var inqDataObj *fastjson.Object
		ftBol := false
		if val.FeatureCode == "404" || val.FeatureCode == "103" || val.FeatureCode == "315" {
			ftBol = false
		} else {
			ftBol = true
		}
		if ftBol == true {
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

			log.Infof("inquiry data: %v", inqDataObj)
		}

		if val.FeatureCode == "103" {

			var srcAccount string
			switch req.Params["core"] {
			case "K":
				srcAccount = "1000000000"
			case "S":
				srcAccount = "6000000000"
			default:
				srcAccount = "1000000000"
			}
			dest := val.FeatureName + " " + val.CustomerReference
			params = map[string]string{
				"tellerID":        req.Params["tellerID"],
				"tellerPass":      req.Params["tellerPass"],
				"txType":          val.Txtype,
				"amount":          val.Amount,
				"txAmount":        val.TxAmount,
				"fee":             val.Fee,
				"srcAccount":      srcAccount,
				"destAccount":     val.CustomerReference,
				"destAccountName": val.BillerId,
				"vaType":          val.BillerCode,
				"description":     dest,
				"referenceNumber": util.RandomNumber(12),
				"termType":        "6010",
				"termId":          "WTELLER",
			}

			gateMsg = transport.SendToGate("gate.shared", "07", params)
		} else {
			if val.FeatureCode == "315" {
				if val.PaymentOptions != "" {
					additional := make([]map[string][]interface{}, 0)
					addData := make(map[string][]interface{})
					js := jsoniter.ConfigCompatibleWithStandardLibrary
					arrAdd := make([]paymentOption, 0)
					mumbill := 0
					js.UnmarshalFromString(val.PaymentOptions, &arrAdd)
					for _, v := range arrAdd {
						arryData := make([]interface{}, 2)
						mumbill++
						numofbill, _ := strconv.Atoi(v.NumOfBill)
						arryData[0] = v.Amount
						arryData[1] = numofbill
						addData[v.Bill] = arryData
					}
					additional = append(additional, addData)
					jAdd, err := js.Marshal(additional)
					if err != nil {
						log.Error(err.Error())
					}
					val.PaymentOptions = string(jAdd)
				}
			}

			params = map[string]string{
				"core":              req.Params["core"],
				"tellerID":          req.Params["tellerID"],
				"tellerPass":        req.Params["tellerPass"],
				"amount":            val.Amount,
				"txAmount":          val.TxAmount,
				"fee":               val.Fee,
				"txType":            val.Txtype,
				"billerCode":        val.BillerCode,
				"billerProductCode": val.BillerProductCode,
				"customerId":        val.CustomerReference,
				"inqData":           val.InquiryData,
				"additional":        val.PaymentOptions,
				"referenceNumber":   util.RandomNumber(12),
				"refnum":            val.Refnum,
				"rpFee":             val.RpFee,
				"rpFeeStruk":        val.RpFeeStruk,
				"rpTag":             val.RpTag,
				"totalBill":         val.TotalBill,
				"termType":          "6010",
				"termId":            "WTELLER",
			}
			if val.FeatureCode == "315" {
				params["amount"] = params["txAmount"]
				params["rpTag"] = params["txAmount"]
			}
			log.Infof("asasaaasas:", params)
			if val.FeatureName != "MPN" {
				inqDataObj.Visit(func(key []byte, v *fastjson.Value) {
					if v.Type() == fastjson.TypeString {
						params[string(key)] = string(v.GetStringBytes())
					} else {
						params[string(key)] = v.String()
					}
				})
			}
			if val.FeatureCode == "404" {
				params["srcAccount"] = val.BillerProductCode
				params["termId"] = "SWTELLER"
			}
			log.Infof("[%s] param: %v", req.Headers["Request-ID"], params)

			gateMsg = transport.SendToGate("gate.shared", "12", params)
		}

		if gateMsg.ResponseCode == "00" {
			if val.FeatureCode == "103" {
				gateMsg.Data = make(map[string]interface{})
				gateMsg.Data["txDate"] = FormattedTime(req.Params["txDate"], "20060102 15:04:05")
				gateMsg.Data["customerReference"] = val.CustomerReference
				gateMsg.Data["txRefNumber"] = params["referenceNumber"]
				gateMsg.Data["amount"] = val.Amount
				gateMsg.Data["accountName"] = val.BillerId
				gateMsg.Data["featureCode"] = val.FeatureCode
				gateMsg.Data["featureName"] = val.FeatureName
				gateMsg.Data["responseCode"] = gateMsg.ResponseCode
				gateMsg.Data["txStatus"] = "SUCCESS"
			} else {
				if val.FeatureCode == "404" && params["srcAccount"] != "" {
					gateMsg.Data["srcAccount"] = val.BillerProductCode
				}
				gateMsg.Data["txDate"] = time.Now().Format("20060102 15:04:05")
				gateMsg.Data["customerReference"] = val.CustomerReference
				gateMsg.Data["featureName"] = val.FeatureName
				gateMsg.Data["featureCode"] = val.FeatureCode
				gateMsg.Data["txRefNumber"] = params["referenceNumber"]
				gateMsg.Data["responseCode"] = gateMsg.ResponseCode
				gateMsg.Data["txStatus"] = "SUCCESS"

			}
			dataReceipt, _ := json.Marshal(gateMsg.Data["receipt"])
			log.Infof("Data test Response:", dataReceipt)
			newDataTrx = append(newDataTrx, gateMsg.Data)
			//res.Response, _ = json.Marshal(successResp(gateMsg.Data))
			params["featureGroupCode"] = val.FeatureGroupCode
			params["featureGroupName"] = val.FeatureGroupName
			params["customerId"] = val.CustomerReference
			params["featureName"] = val.FeatureName
			params["featureCode"] = val.FeatureCode
			params["productCode"] = val.BillerCode
			params["featureId"] = val.FeatureId
			params["branchCode"] = req.Params["branchCode"]
			params["receipt"] = string(dataReceipt)
			params["txStatus"] = "SUCCESS"

		} else {
			gateMsg.Data = make(map[string]interface{})
			gateMsg.Data["txDate"] = FormattedTime(req.Params["txDate"], "20060102 15:04:05")
			gateMsg.Data["customerReference"] = val.CustomerReference
			gateMsg.Data["featureName"] = val.FeatureName
			gateMsg.Data["featureCode"] = val.FeatureCode
			gateMsg.Data["txRefNumber"] = params["referenceNumber"]
			gateMsg.Data["responseCode"] = gateMsg.ResponseCode
			gateMsg.Data["txStatus"] = "FAILED"

			newDataTrx = append(newDataTrx, gateMsg.Data)

			//res.Response, _ = json.Marshal(newResponseWithData(gateMsg.ResponseCode, ParseRC(gateMsg.ResponseCode), gateMsg.Data))
			params["featureGroupCode"] = val.FeatureGroupCode
			params["featureGroupName"] = val.FeatureGroupName
			params["customerId"] = val.CustomerReference
			params["featureName"] = val.FeatureName
			params["featureCode"] = val.FeatureCode
			params["productCode"] = val.BillerCode
			params["featureId"] = val.FeatureId
			params["branchCode"] = req.Params["branchCode"]
			params["txStatus"] = "FAILED"

		}
		trxData := BuildDataTransaction(req.Params, params, params["txStatus"], gateMsg.ResponseCode)

		err := repo.Transaction.Save(trxData)
		if err != nil {
			log.Errorf("error save transaction: %v", err)
		}
		res.Response, _ = json.Marshal(successResp(newDataTrx))
	}
	return nil
}

type paymentOption struct {
	Bill      string  `json:"bill"`
	Index     string  `json:"index"`
	Amount    float64 `json:"amount"`
	NumOfBill string  `json:"numOfBill"`
}
