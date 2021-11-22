package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
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

		//TODO: INI BUAT INQURY KE TABLE TRANSACTION
		//if val.FeatureCode == "404" || val.FeatureCode == "303" {
		//	customerRef, trxStatus, err := repo.Transaction.FindTransaction(val.CustomerReference)
		//	if err != nil {
		//		return err
		//	}
		//	if strings.EqualFold(val.CustomerReference, customerRef) {
		//		if trxStatus == "SUCCESS" || trxStatus == "PENDING" {
		//			res.Response, _ = json.Marshal(newResponse("99", "Invalid Duplicate Transaction"))
		//			return nil
		//		}
		//	}
		//}

		var inqDataObj *fastjson.Object
		ftBol := false
		if val.FeatureCode == "404" || val.FeatureCode == "103" || val.FeatureCode == "315" || val.FeatureGroupCode == "002" || val.FeatureCode == "301" || val.FeatureCode == "311" || val.FeatureCode == "319" || val.FeatureCode == "303" {
			ftBol = false
		} else {
			ftBol = true
		}
		//log.Infof("log feature: ", val.FeatureCode, val.FeatureGroupCode)
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
				//log.Infof("inquiry data: %v", inqDataObj)
			}
		}
		var srcAccount string
		if val.SrcAccount == "" {
			switch req.Params["core"] {
			case "K":
				srcAccount = "1000000000"
			case "S":
				srcAccount = "6000000000"
			default:
				srcAccount = "1000000000"
			}
		} else {
			srcAccount = strings.TrimSpace(val.SrcAccount)
		}

		if val.FeatureCode == "103" {
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
				"branchName":        req.Params["branchName"],
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
				"srcAccount":        srcAccount,
				"termType":          "6010",
				"termId":            "WTELLER",
				"dateTime":          txDate.Format("20060102150405"),
			}
			if val.FeatureCode == "315" {
				params["amount"] = params["txAmount"]
				params["rpTag"] = params["txAmount"]
			}
			if val.FeatureCode == "303" || val.FeatureCode == "319" {
				descr := ""
				if val.Txtype == "03" {
					descr = "PLN PAYMENT"
				} else if val.Txtype == "04" {
					descr = "NTAGLIST"
				} else {
					descr = "BPJS Kesehatan"
				}
				params = map[string]string{
					"core":              req.Params["core"],
					"tellerID":          req.Params["tellerID"],
					"tellerPass":        req.Params["tellerPass"],
					"amount":            val.Amount,
					"branchCode":        req.Params["branchCode"],
					"branchName":        req.Params["branchName"],
					"txAmount":          val.TxAmount,
					"fee":               val.Fee,
					"billerCode":        val.BillerCode,
					"billerId":          val.BillerCode,
					"billerProductCode": val.BillerProductCode,
					"txType":            val.Txtype,
					"customerId":        val.CustomerReference,
					"type":              val.Txtype,
					"productCode":       val.BillerProductCode,
					"description":       descr,
					"srcAccount":        srcAccount,
					"srcAccType":        "00",
					"inqData":           val.InquiryData,
					"referenceNumber":   util.RandomNumber(12),
					"termType":          "6010",
					"termId":            "WTELLER",
					"dateTime":          txDate.Format("20060102150405"),
				}
			}
			if val.FeatureCode == "404" && core == "S" {
				params["termId"] = "KWTELLER"

			}
			if val.FeatureCode == "404" {
				params["srcAccount"] = srcAccount
				branch := req.Params["branchCode"]
				substring := branch[3:9]
				params["branchCode"] = substring
			}

			if val.FeatureCode == "302" {
				params["formatType"] = "501"
			} else if val.FeatureCode != "302" {
				params["formatType"] = "506"
			}
			if val.FeatureCode == "202" {
				params["unsold"] = val.Unsold
				params["token"] = val.Token
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
			if val.FeatureCode == "305" {
				params["norangka"] = val.CustomerReference2
			}
			if params["fee"] == "0" {
				params["fee"] = ""
			}
			//log.Infof("Request-ID:[%s] param Send to Gate: %v", req.Headers["Request-ID"], params)
			if val.FeatureCode == "404" {
				gateMsg = transport.SendToGate("gate.shared", "26", params)
			} else {
				gateMsg = transport.SendToGate("gate.shared", "12", params)
			}

			if val.FeatureCode == "302" || val.FeatureCode == "304" ||
				val.FeatureCode == "315" || val.FeatureCode == "312" ||
				val.FeatureCode == "305" || val.FeatureCode == "314" ||
				val.FeatureCode == "316" || val.FeatureCode == "317" {
				params["fee"] = params["rpFeeStruk"]
			} else {
				params["fee"] = val.Fee
			}
			params["billerProductCode"] = val.BillerProductCode
		}
		var tampungData map[string]interface{}
		stan := params["referenceNumber"]
		json.Unmarshal([]byte(val.PaymentOptions), &tampungData)
		tampungData["featureCode"] = val.FeatureCode
		tampungData["featureName"] = val.FeatureName
		tampungData["txRefNumber"] = params["referenceNumber"]
		tampungData["txDate"] = params["dateTime"]
		tampungData["ntb"] = params["referenceNumber"]
		tampungData["stan"] = stan[0:6]

		//log.Infof("['%s'] Response Gate: %v", req.Headers["Request-ID"], gateMsg)
		if gateMsg.ResponseCode == "00" {
			if val.FeatureCode == "319" || val.FeatureCode == "303" {
				if val.Txtype == "03" {
					var tagihan []interface{}
					var ok bool
					if tagihan, ok = gateMsg.Data["tagihan"].([]interface{}); !ok {
						log.InfoS("error parsing tagihan")
						return nil
					}
					if tagihan == nil || len(tagihan) == 0 {
						log.InfoS("tagihan nil")
						return nil
					}
					var denda float64 = 0
					var periodToBePaid1 int
					var periodToBePaid2 int
					var dueDate string
					var slaAwal string
					var slaAkhir string
					var shaAwal string
					var shaAkhir string
					var periodPaid string
					for index, tag := range tagihan {
						//log.Infof("tagihan : ", tag)
						mapTag := tag.(map[string]interface{})
						if len(tagihan) == 1 {
							slaAwal = mapTag["slalwbp1"].(string)
							shaAkhir = mapTag["sahlwbp1"].(string)
							periodToBePaid1, _ = strconv.Atoi(mapTag["periodBillToBePaid"].(string))
							denda, _ = strconv.ParseFloat(mapTag["penaltyToBePaid"].(string), 64)
							dueDate = mapTag["dueDate"].(string)
						} else {

							if index == 0 {
								periodToBePaid1, _ = strconv.Atoi(mapTag["periodBillToBePaid"].(string))
								slaAwal = mapTag["slalwbp1"].(string)
								shaAwal = mapTag["sahlwbp1"].(string)
							}
							if index == len(tagihan)-1 {
								periodToBePaid2, _ = strconv.Atoi(mapTag["periodBillToBePaid"].(string))
								slaAkhir = mapTag["slalwbp1"].(string)
								shaAkhir = mapTag["sahlwbp1"].(string)
							}
							if index != 0 {
								periodPaid = periodPaid + "," + mapTag["periodBillToBePaid"].(string)
								dueDate = dueDate + "," + mapTag["dueDate"].(string)
							} else {
								periodPaid = mapTag["periodBillToBePaid"].(string)
								dueDate = mapTag["dueDate"].(string)
							}
							dendaPaid, _ := strconv.ParseFloat(mapTag["penaltyToBePaid"].(string), 64)
							denda = denda + dendaPaid
						}
					}
					if len(tagihan) == 1 {
						gateMsg.Data["standMeter"] = slaAwal + "-" + shaAkhir
						gateMsg.Data["periodPaid"] = periodToBePaid1
						gateMsg.Data["penaltyToBePaid"] = denda
						gateMsg.Data["dueDate"] = dueDate
					} else {
						if periodToBePaid1 < periodToBePaid2 {
							gateMsg.Data["standMeter"] = slaAwal + "-" + shaAkhir
						} else {
							gateMsg.Data["standMeter"] = slaAkhir + "-" + shaAwal
						}
						gateMsg.Data["penaltyToBePaid"] = denda
						gateMsg.Data["periodPaid"] = periodPaid
						gateMsg.Data["dueDate"] = dueDate
					}
				}
			}
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
				gateMsg.Data["message"] = gateMsg.Description
				gateMsg.Data["txStatus"] = "SUCCESS"
			} else {
				if val.FeatureCode == "201" || val.FeatureCode == "203" || val.FeatureCode == "306" || val.FeatureCode == "311" || val.FeatureCode == "301" || val.FeatureCode == "303" || val.FeatureCode == "319" || val.FeatureCode == "202" {
					gateMsg.Data["amount"] = val.Amount
					gateMsg.Data["txAmount"] = val.Amount
					gateMsg.Data["productName"] = val.ProductName
					if val.FeatureCode == "319" && gateMsg.Data["numbBill"] == "" {
						gateMsg.Data["numbBill"] = tampungData["numbBill"]
					}
				}
				if val.FeatureCode == "404" && params["srcAccount"] != "" {
					gateMsg.Data["srcAccount"] = srcAccount
				}
				gateMsg.Data["txDate"] = params["dateTime"]
				gateMsg.Data["customerReference"] = val.CustomerReference
				gateMsg.Data["featureName"] = val.FeatureName
				gateMsg.Data["fee"] = params["fee"]
				gateMsg.Data["featureCode"] = val.FeatureCode
				gateMsg.Data["transactionType"] = val.Txtype
				gateMsg.Data["txRefNumber"] = params["referenceNumber"]
				gateMsg.Data["responseCode"] = gateMsg.ResponseCode
				gateMsg.Data["message"] = gateMsg.Description
				gateMsg.Data["txStatus"] = "SUCCESS"

			}
			if val.TransactionType == "NON-TUNAI" {
				gateMsg.Data["srcAccount"] = val.SrcAccount
				gateMsg.Data["srcAccountName"] = val.SrcAccountName
			}
			gateMsg.Data["transactionType"] = val.TransactionType
			gateMsg.Data["txType"] = val.Txtype
			dataReceipt, _ := json.Marshal(gateMsg.Data)
			//log.Infof("Data test Response:", dataReceipt)
			newDataTrx = append(newDataTrx, gateMsg.Data)
			// res.Response, _ = json.Marshal(successResp(gateMsg.Data))
			params["featureGroupCode"] = val.FeatureGroupCode
			params["featureGroupName"] = val.FeatureGroupName
			params["customerId"] = val.CustomerReference
			params["featureName"] = val.FeatureName
			params["featureCode"] = val.FeatureCode
			params["productCode"] = val.BillerCode
			params["featureId"] = val.FeatureId
			params["transactionType"] = val.TransactionType
			params["branchCode"] = req.Params["branchCode"]

			params["receipt"] = string(dataReceipt)
			params["txStatus"] = "SUCCESS"
		} else {
			var sts string
			bookDate := ""
			if gateMsg.ResponseCode == "06" && val.FeatureCode == "404" {
				bookDate = fmt.Sprintf("%s", gateMsg.Data["bookDate"])
			}
			if val.FeatureCode == "404" && gateMsg.ResponseCode != "19" && gateMsg.ResponseCode != "99" && bookDate == "" {
				// gateMsg.Data = make(map[string]interface{})
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

				//log.Infof("Data Tampung: ", gateMsg.Data)
			}
			mpnPending := false
			if val.FeatureCode == "404" {
				gateMsg.Data = tampungData
				if gateMsg.ResponseCode == "06" || gateMsg.ResponseCode == "90" || gateMsg.ResponseCode == "92" {
					mpnPending = true
				} else {
					mpnPending = false
				}
			}
			if gateMsg.ResponseCode == "19" {
				var rec map[string]interface{}
				json.Unmarshal([]byte(val.PaymentOptions), &rec)
				gateMsg.Data = rec
				gateMsg.Data["txStatus"] = "FAILED"
				sts = "FAILED"
			} else if gateMsg.ResponseCode == "06" || gateMsg.ResponseCode == "05" || mpnPending == true {
				var rec map[string]interface{}
				json.Unmarshal([]byte(val.PaymentOptions), &rec)
				gateMsg.Data = rec
				gateMsg.ResponseCode = "06"
				gateMsg.Data["amount"] = val.Amount
				gateMsg.Data["productName"] = val.ProductName
				if mpnPending == true {
					gateMsg.Data["txStatus"] = "PENDING"
					gateMsg.Description = "Silahkan Cetak BPN Sementara"
				} else if val.FeatureCode == "202" {
					gateMsg.Data["txStatus"] = "PENDING"
					gateMsg.Description = "Silahkan Advice Manual pada menu cetak ulang"
				} else {
					gateMsg.Data["txStatus"] = "PENDING"
				}
				sts = "PENDING"
			} else {
				var rec map[string]interface{}
				json.Unmarshal([]byte(val.PaymentOptions), &rec)
				gateMsg.Data = rec
				gateMsg.Data["txStatus"] = "FAILED"
				sts = "FAILED"
			}

			//log.Infof("data gateData: ", gateMsg.Data)
			tampungData["customerReference"] = val.CustomerReference
			tampungData["txStatus"] = sts
			tampungData["bookDate"] = bookDate
			tampungData["responseCode"] = gateMsg.ResponseCode
			tampungData["ntpn"] = ""
			if val.FeatureCode == "319" || val.FeatureCode == "303" {
				tampungData["amount"] = val.Amount
			}

			data, _ := json.Marshal(tampungData)
			if val.FeatureCode == "404" && gateMsg.ResponseCode == "M4" {
				gateMsg.Data = tampungData
				//log.Infof("gate Date M4 :", gateMsg)
				gateMsg.Data["txDate"] = params["dateTime"]
				gateMsg.Data["customerReference"] = val.CustomerReference
				gateMsg.Data["featureName"] = val.FeatureName
				gateMsg.Data["fee"] = params["fee"]
				gateMsg.Data["featureCode"] = val.FeatureCode
				gateMsg.Data["txRefNumber"] = params["referenceNumber"]
				gateMsg.Data["responseCode"] = gateMsg.ResponseCode
				gateMsg.Data["message"] = gateMsg.Description
				gateMsg.Data["txStatus"] = "Failed"
			}
			gateMsg.Data["transactionType"] = val.TransactionType
			gateMsg.Data["txType"] = val.Txtype
			gateMsg.Data["txDate"] = params["dateTime"]
			gateMsg.Data["customerReference"] = val.CustomerReference
			gateMsg.Data["featureName"] = val.FeatureName
			gateMsg.Data["featureCode"] = val.FeatureCode
			gateMsg.Data["bookDate"] = bookDate
			gateMsg.Data["txRefNumber"] = params["referenceNumber"]
			gateMsg.Data["message"] = gateMsg.Description
			gateMsg.Data["responseCode"] = gateMsg.ResponseCode
			if val.FeatureCode == "202" || val.FeatureCode == "303" {
				data, _ = json.Marshal(gateMsg.Data)
			}
			newDataTrx = append(newDataTrx, gateMsg.Data)
			brN := req.Params["branchName"]
			if brN == "" {
				brN = "1001 - Cabang Utama"
			}
			if val.TransactionType == "" {
				val.TransactionType = "TUNAI"
			}
			//res.Response, _ = json.Marshal(newResponseWithData(gateMsg.ResponseCode, ParseRC(gateMsg.ResponseCode), gateMsg.Data))
			params["featureGroupCode"] = val.FeatureGroupCode
			params["featureGroupName"] = val.FeatureGroupName
			params["customerId"] = val.CustomerReference
			params["featureName"] = val.FeatureName
			params["featureCode"] = val.FeatureCode
			params["productCode"] = val.BillerCode
			params["featureId"] = val.FeatureId
			params["branchCode"] = req.Params["branchCode"]
			params["branchName"] = brN
			params["transactionType"] = val.TransactionType
			params["srcAccount"] = srcAccount
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
