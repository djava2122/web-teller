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
	core := req.Params["core"]
	if err := json.Unmarshal([]byte(req.Params["bulk"]), &newBulkPayment); err != nil {
		log.Infof("Error Unmarshal : ", err.Error())
		return err
	}

	for _, val := range newBulkPayment {

		var inqDataObj *fastjson.Object
		ftBol := false
		if val.FeatureCode == "404" || val.FeatureCode == "103" || val.FeatureCode == "315" || val.FeatureGroupCode == "002" || val.FeatureCode == "301" || val.FeatureCode == "311" {
			ftBol = false
		} else {
			ftBol = true
		}
		log.Infof("log feature: ", val.FeatureCode, val.FeatureGroupCode)
		if ftBol == true {
			if val.InquiryData == "" {
				res.Response, _ = json.Marshal(newResponse("02", "invalid inquiry data"))
				return nil
			}
			if val.FeatureCode != "306" {
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
			txDate := time.Now()
			params = map[string]string{
				"core":              req.Params["core"],
				"tellerID":          req.Params["tellerID"],
				"tellerPass":        req.Params["tellerPass"],
				"amount":            val.Amount,
				"branchCode":        req.Params["branchCode"],
				"txAmount":          val.TxAmount,
				"fee":               val.Fee,
				"txType":            val.Txtype,
				"billerCode":        val.BillerCode,
				"billerId":          val.BillerCode,
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
				"dateTime":          txDate.Format("20060102150405"),
			}

			if val.FeatureCode == "315" {
				params["amount"] = params["txAmount"]
				params["rpTag"] = params["txAmount"]
			}
			if val.FeatureCode == "404" && core == "S" {
				params = map[string]string{
					"tellerID":        req.Params["tellerID"],
					"tellerPass":      req.Params["tellerPass"],
					"amount":          val.Amount,
					"txType":          val.Txtype,
					"srcAccount":      val.BillerProductCode,
					"customerId":      val.CustomerReference,
					"inqData":         val.InquiryData,
					"referenceNumber": util.RandomNumber(12),
					"termType":        "6010",
					"termId":          "KWTELLER",
					"dateTime":        txDate.Format("20060102150405"),
				}
			}
			if val.FeatureCode == "404" {
				if val.BillerProductCode == "" {
					if core == "S" {
						val.BillerProductCode = "6000000000"
					} else {
						val.BillerProductCode = "1000000000"
					}
				}
				params["srcAccount"] = val.BillerProductCode
				branch := req.Params["branchCode"]
				substring := branch[3:9]
				params["branchCode"] = substring
			}

			if val.FeatureCode == "302" {
				params["formatType"] = "501"
			} else if val.FeatureCode != "302" {
				params["formatType"] = "506"
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
			if val.FeatureCode == "318" {
				params["norangka"] = val.CustomerReference2
			}
			if val.FeatureCode == "201" || val.FeatureCode == "203" || val.FeatureCode == "311" || val.FeatureCode == "301" || val.FeatureCode == "306" {
				tag, _ := strconv.Atoi(val.Amount)
				fee, _ := strconv.Atoi(val.Fee)
				total := strconv.Itoa(tag + fee)
				params["amount"] = total
			}
			log.Infof("Request-ID:[%s] param Send to Gate: %v", req.Headers["Request-ID"], params)
			if val.FeatureCode == "404" {
				gateMsg = transport.SendToGate("gate.shared", "26", params)
			} else {
				gateMsg = transport.SendToGate("gate.shared", "12", params)
			}

			if val.FeatureCode == "302" || val.FeatureCode == "304" ||
				val.FeatureCode == "315" || val.FeatureCode == "312" ||
				val.FeatureCode == "318" || val.FeatureCode == "314" ||
				val.FeatureCode == "316" || val.FeatureCode == "317" {
				params["fee"] = params["rpFeeStruk"]
			} else {
				params["fee"] = val.Fee
			}
			params["billerProductCode"] = val.BillerProductCode
		}
		var tampungData map[string]string
		stan := params["referenceNumber"]
		json.Unmarshal([]byte(val.PaymentOptions), &tampungData)
		tampungData["featureCode"] = val.FeatureCode
		tampungData["featureName"] = val.FeatureName
		tampungData["txRefNumber"] = params["referenceNumber"]
		tampungData["txDate"] = params["dateTime"]
		tampungData["ntb"] = params["referenceNumber"]
		tampungData["stan"] = stan[0:6]

		log.Infof("['%s'] Response Gate: %v", req.Headers["Request-ID"], gateMsg)
		if gateMsg.ResponseCode == "00" {
			if val.FeatureCode == "103" {
				gateMsg.Data = make(map[string]interface{})
				gateMsg.Data["txDate"] = params["dateTime"]
				gateMsg.Data["customerReference"] = val.CustomerReference
				gateMsg.Data["txRefNumber"] = params["referenceNumber"]
				gateMsg.Data["amount"] = val.Amount
				gateMsg.Data["fee"] = params["fee"]
				gateMsg.Data["accountName"] = val.BillerId
				gateMsg.Data["featureCode"] = val.FeatureCode
				gateMsg.Data["featureName"] = val.FeatureName
				gateMsg.Data["responseCode"] = gateMsg.ResponseCode
				gateMsg.Data["txStatus"] = "SUCCESS"
			} else {
				if val.FeatureCode == "201" || val.FeatureCode == "203" || val.FeatureCode == "306" || val.FeatureCode == "311" || val.FeatureCode == "301" {
					gateMsg.Data["amount"] = val.Amount
					gateMsg.Data["productName"] = val.ProductName
				}
				if val.FeatureCode == "404" && params["srcAccount"] != "" {
					gateMsg.Data["srcAccount"] = val.BillerProductCode
				}
				gateMsg.Data["txDate"] = params["dateTime"]
				gateMsg.Data["customerReference"] = val.CustomerReference
				gateMsg.Data["featureName"] = val.FeatureName
				gateMsg.Data["fee"] = params["fee"]
				gateMsg.Data["featureCode"] = val.FeatureCode
				gateMsg.Data["txRefNumber"] = params["referenceNumber"]
				gateMsg.Data["responseCode"] = gateMsg.ResponseCode
				gateMsg.Data["txStatus"] = "SUCCESS"

			}
			dataReceipt, _ := json.Marshal(gateMsg.Data)
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
			var sts string
			bookDate := ""
			if val.FeatureCode == "404" && gateMsg.ResponseCode != "19" && gateMsg.ResponseCode != "99" {
				var rec map[string]interface{}
				json.Unmarshal([]byte(val.PaymentOptions), &rec)
				d := time.Now()
				t := d.Format("1504")
				tInt, _ := strconv.Atoi(t)
				if tInt > 1500 {
					rec["bookDate"] = d.AddDate(0, 0, 1).Format("20060102")
					bookDate = d.AddDate(0, 0, 1).Format("20060102")
				} else {
					rec["bookDate"] = d.Format("20060102")
					bookDate = d.Format("20060102")
				}
				gateMsg.Data = rec
			}
			if gateMsg.ResponseCode == "19" {
				gateMsg.Data["txStatus"] = "FAILED"
				sts = "FAILED"
			} else if gateMsg.ResponseCode == "06" {
				gateMsg.Data["txStatus"] = "PENDING - Silahkan Cetak BPN Sementara"
				sts = "PENDING"
			} else {
				gateMsg.Data["txStatus"] = "FAILED"
				sts = "FAILED"
				// gateMsg.Data["txStatus"] = gateMsg.Description
			}
			tampungData["customerReference"] = val.CustomerReference
			tampungData["txStatus"] = sts
			tampungData["bookDate"] = bookDate
			tampungData["responseCode"] = gateMsg.ResponseCode
			tampungData["ntpn"] = ""
			data, _ := json.Marshal(tampungData)

			gateMsg.Data["txDate"] = params["dateTime"]
			gateMsg.Data["customerReference"] = val.CustomerReference
			gateMsg.Data["featureName"] = val.FeatureName
			gateMsg.Data["featureCode"] = val.FeatureCode
			gateMsg.Data["txRefNumber"] = params["referenceNumber"]
			gateMsg.Data["responseCode"] = gateMsg.ResponseCode

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
			params["receipt"] = string(data)
			params["txStatus"] = sts
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
