package service

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"gitlab.pactindo.com/ebanking/common/log"
	"gitlab.pactindo.com/ebanking/common/transport"
	"gitlab.pactindo.com/ebanking/common/trycatch"
	"gitlab.pactindo.com/ebanking/common/util"
	wtproto "gitlab.pactindo.com/ebanking/web-teller/proto"
	"gitlab.pactindo.com/ebanking/web-teller/repo"
)

type GetMgateStruct struct {
	TxRefNumber       string `json:txRefNumber`
	ResponseCode      string `json:responseCode`
	TxDate            string `json:txDate`
	BookDate          string `json:bookDate`
	Ntb               string `json:ntb`
	Ntpn              string `json:ntpn`
	Stan              string `json:stan`
	CustomerReference string `json:customerReference`
	Npwp              string `json:npwp`
	PrayerName        string `json:prayerName`
	PrayerAddress     string `json:prayerAddress`
	Nop               string `json:nop`
	TaxAccount        string `json:taxAccount`
	DepositTypeCode   string `json:depositTypeCode`
	TaxPeriod         string `json:taxPeriod`
	SkNumber          string `json:skNumber`
	Amount            string `json:amount`
	CurrencyCode      string `json:currencyCode`
	PrayerId          string `json:prayerId`
	DocumentType      string `json:documentType`
	DocumentNumber    string `json:documentNumber`
	DocumentDate      string `json:documentDate`
	KppbcCode         string `json:kppbcCode`
	Kl                string `json:kl`
	UnitEselon        string `json:unitEselon`
	SatkerCode        string `json:satkerCode`
	FeatureCode       string `json:featureCode`
	FeatureName       string `json:featurName`
}

type ReqMPN struct {
	TxRefNumber     string            `json:txRefNumber`
	ResponseCode    string            `json:responseCode`
	Datetime        string            `json:dateTime`
	TermType        string            `json:termType`
	TermId          string            `json:termId`
	CustomerId      string            `json:customerId`
	SrcAccount      string            `json:srcAccount`
	Amount          string            `json:amount`
	InqData         string            `json:inqData`
	ReferenceNumber string            `json:referenceNumber`
	Additional      map[string]string `json:additional`
}

func InitDb(req string) (result []interface{}, err error) {
	dsn := fmt.Sprintf("postgres://mgate:mgate2020@172.19.252.114/micro-gate?sslmode=disable")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Infof("unable to connect db, error : %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Infof("unable to connect db, error : %v", err)
	}
	// sql := "select fields -> 'additional' from mgate.t_transaction where trrfnm = $1 or fields -> 'customerId' = $2 and trftcd = 'MPN2' and trrspc = '00'"
	sql := ""
	jenis, _ := strconv.Atoi(req[0:1])
	if jenis <= 3 {
		sql = "select trrfnm, trrspc, fields -> 'trxDate' as trxDate,fields-> 'bookDate' as bookDate, fields -> 'ntb' as ntb , fields -> 'ntpn' as ntpn, fields -> 'stan' as stan, fields -> 'customerId' as customerReference, fields -> 'npwp' as npwp, fields -> 'payerName'as payerName, fields -> 'payerAddress' as payerAddress, fields->'taxAccount' as taxAccount, fields -> 'depositTypeCode'as depositTypeCode,fields -> 'taxPeriod' as taxPeriod, fields -> 'skNumber' as skNumber,fields -> 'amount' as amount, fields -> 'currencyCode' as currencyCode, fields -> 'nop' as nop from mgate.t_transaction where trrfnm = $1 or fields -> 'customerId' = $2 and trftcd = 'MPN2' and (trrspc = '00' or trrspc = '06' or trrspc = '45') and mclass <> 'R'"
	} else if jenis >= 7 {
		sql = "select trrfnm, trrspc, fields -> 'trxDate' as trxDate,fields-> 'bookDate' as bookDate, fields -> 'ntb' as ntb , fields -> 'ntpn' as ntpn, fields -> 'stan' as stan, fields -> 'customerId' as customerReference, fields -> 'payerName'as payerName,fields -> 'amount'as amount, fields -> 'currencyCode' as currencyCode, fields -> 'kl' as kl, fields -> 'unitEselon' as unitEselon, fields -> 'satkerCode' as satkerCode from mgate.t_transaction where trrfnm = $1 or fields -> 'customerId' = $2 and trftcd = 'MPN2' and (trrspc = '00' or trrspc = '06' or trrspc = '45') and mclass <> 'R'"
	} else {
		sql = "select trrfnm, trrspc, fields -> 'trxDate' as trxDate,fields-> 'bookDate' as bookDate, fields -> 'ntb' as ntb , fields -> 'ntpn' as ntpn, fields -> 'stan' as stan, fields -> 'customerId' as customerReference, fields -> 'payerName'as payerName, fields -> 'amount'as amount, fields -> 'currencyCode' as currencyCode, fields -> 'documentType' as documentType, fields -> 'documentNumber' as documentNumber, fields -> 'documentDate' as documentDate, fields -> 'kppbcCode'as kppbcCode from mgate.t_transaction where trrfnm = $1 or fields -> 'customerId' = $2 and trftcd = 'MPN2' and (trrspc = '00' or trrspc = '06' or trrspc = '45') and mclass <> 'R'"
	}
	rows, _ := db.Query(sql, req, req)
	// if err != nil {
	// 	return nil, err
	// }
	log.Infof("Test : ", rows)
	for rows.Next() {
		dt := GetMgateStruct{}
		if jenis <= 3 {
			err = rows.Scan(
				&dt.TxRefNumber,
				&dt.ResponseCode,
				&dt.TxDate,
				&dt.BookDate,
				&dt.Ntb,
				&dt.Ntpn,
				&dt.Stan,
				&dt.CustomerReference,
				&dt.Npwp,
				&dt.PrayerName,
				&dt.PrayerAddress,
				&dt.TaxAccount,
				&dt.DepositTypeCode,
				&dt.TaxPeriod,
				&dt.SkNumber,
				&dt.Amount,
				&dt.CurrencyCode,
				&dt.Nop,
			)
		} else if jenis >= 7 {
			err = rows.Scan(
				&dt.TxRefNumber,
				&dt.ResponseCode,
				&dt.TxDate,
				&dt.BookDate,
				&dt.Ntb,
				&dt.Ntpn,
				&dt.Stan,
				&dt.CustomerReference,
				&dt.PrayerName,
				&dt.Amount,
				&dt.CurrencyCode,
				&dt.Kl,
				&dt.UnitEselon,
				&dt.SatkerCode,
			)
		} else {
			err = rows.Scan(
				&dt.TxRefNumber,
				&dt.ResponseCode,
				&dt.TxDate,
				&dt.BookDate,
				&dt.Ntb,
				&dt.Ntpn,
				&dt.Stan,
				&dt.CustomerReference,
				&dt.PrayerName,
				&dt.Amount,
				&dt.CurrencyCode,
				&dt.DocumentType,
				&dt.DocumentNumber,
				&dt.DocumentDate,
				&dt.KppbcCode,
			)
		}
		dt.FeatureCode = "404"
		dt.FeatureName = "MPN"
		status := ""
		resp := make(map[string]interface{})
		if dt.ResponseCode == "00" {
			status = "SUCCESS"
		} else if dt.ResponseCode == "99" || dt.ResponseCode == "19" {
			status = "FAILED"
		} else {
			sql = "select trrfnm, trrspc, fields -> 'dateTime' as dateTime, fields -> 'inqData' as inqData, fields -> 'termID' as termID, fields -> 'termType' as termType, fields -> 'amount' as amount, fields -> 'additional' as additional, fields -> 'srcAccount' as srcAccount from mgate.t_transaction where trrfnm = $1 or fields -> 'customerId' = $2 and trftcd = 'MPN2' and 	(trrspc = '06' or trrspc = '45')"
			rows, _ = db.Query(sql, req, req)
			for rows.Next() {
				Rdt := ReqMPN{}
				var tamp string
				err = rows.Scan(
					&Rdt.TxRefNumber,
					&Rdt.ResponseCode,
					&Rdt.Datetime,
					&Rdt.InqData,
					&Rdt.TermId,
					&Rdt.TermType,
					&Rdt.Amount,
					&tamp,
					&Rdt.SrcAccount,
				)
				dt.Amount = Rdt.Amount
				dt.TxRefNumber = Rdt.TxRefNumber
				dt.TxDate = Rdt.Datetime
				status = "PENDING"
				json.Unmarshal([]byte(tamp), &Rdt.Additional)
				reqParams := map[string]string{
					"txType":             "07",
					"dateTime":           time.Now().Format("20060102150405"),
					"termType":           Rdt.TermType,
					"termId":             Rdt.TermId,
					"customerId":         req,
					"srcAccount":         Rdt.SrcAccount,
					"amount":             Rdt.Amount,
					"description":        "",
					"inqData":            Rdt.InqData,
					"referenceNumber":    util.RandomNumber(12),
					"orgDateTime":        Rdt.Datetime,
					"orgReferenceNumber": Rdt.TxRefNumber,
				}
				log.Infof("ddd ", reqParams)
				gateMsg := transport.SendToGate("gate.shared", "27", reqParams)
				if gateMsg.ResponseCode == "00" {
					status = "SUCCESS"
				} else if gateMsg.ResponseCode == "99" || dt.ResponseCode == "19" {
					status = "FAILED"
				} else {
					status = "PENDING"
				}
				log.Infof("termId :", gateMsg.Data["ntpn"])
				log.Infof("Reinquiry Mgate rc timeout :", gateMsg)
				resp["Receipt"] = gateMsg.Data
				gateMsg.Data["featureCode"] = dt.FeatureCode
				gateMsg.Data["featureName"] = dt.FeatureName
				gateMsg.Data["txRefNumber"] = dt.TxRefNumber
				gateMsg.Data["txDate"] = reqParams["orgDateTime"]
				gateMsg.Data["customerReference"] = reqParams["customerId"]
				gateMsg.Data["txStatus"] = status
				receipt, _ := json.Marshal(gateMsg.Data)
				sqlMgate := `update mgate.t_transaction set trrspc = $1 where trrfnm = $2;`
				_, err := db.Exec(sqlMgate, "00", dt.TxRefNumber)
				db.Close()
				if Rdt.TermId == "WTELLER" || Rdt.TermId == "KWTELLER" {
					updateTabel := repo.Transaction.UpdateMpn(gateMsg.ResponseCode, status, string(receipt), dt.TxRefNumber)
					log.Infof("update table teller :", updateTabel)
				}
				result = append(result, resp)
				return result, err
			}

		}
		// params := map[string]string{
		// 	"featureCode":       "404",
		// 	"txDate":            dt.TxDate,
		// 	"bookDate":          dt.BookDate,
		// 	"ntpn":              dt.Ntpn,
		// 	"customerReference": dt.CustomerReference,
		// }

		resp["Id"] = 0
		resp["JumlahCetak"] = 0
		resp["Receipt"] = map[string]string{
			"txRefNumber":       dt.TxRefNumber,
			"responseCode":      dt.ResponseCode,
			"txDate":            dt.TxDate,
			"bookDate":          dt.BookDate,
			"ntb":               dt.Ntb,
			"ntpn":              dt.Ntpn,
			"stan":              dt.Stan,
			"customerReference": dt.CustomerReference,
			"npwp":              dt.Npwp,
			"payerName":         dt.PrayerName,
			"payerAddress":      dt.PrayerAddress,
			"nop":               dt.Nop,
			"taxAccount":        dt.TaxAccount,
			"depositTypeCode":   dt.DepositTypeCode,
			"taxPeriod":         dt.TaxPeriod,
			"skNumber":          dt.SkNumber,
			"amount":            dt.Amount,
			"currencyCode":      dt.CurrencyCode,
			"payerId":           dt.PrayerId,
			"documentType":      dt.DocumentType,
			"documentNumber":    dt.DocumentNumber,
			"documentDate":      dt.DocumentDate,
			"kppbcCode":         dt.KppbcCode,
			"kl":                dt.Kl,
			"unitEselon":        dt.UnitEselon,
			"satkerCode":        dt.SatkerCode,
			"featureCode":       dt.FeatureCode,
			"featureName":       dt.FeatureName,
			"txStatus":          status,
		}
		log.Infof("Log Respon Mgate: ", dt)
		log.Infof("Log Respon Mgate: ", resp)
		result = append(result, resp)
		if err != nil {
			return result, err
		}
	}
	db.Close()
	// var data interface{}
	// var trrspc string
	// tampung := hstore.Hstore{}
	// err = db.QueryRow(sql, req, req).Scan(&trrspc, &tampung)
	// json.Unmarshal([]byte(tampung), &data)
	// log.Infof("Koneksi Mgate: %s, res :%v", data, err)
	// if err == nil {
	// 	return nil
	// }
	return result, err
}
func (h *WebTellerHandler) InquiryNomorRekening(ctx context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {
	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], string(res.Response))
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] error : %v", req.Headers["Request-ID"], e)
		res.Response, _ = json.Marshal(newResponse("99", "Internal Server Error"))
	})
	referenceNumber := req.Params["referenceNumber"]
	feature := req.Params["feature"]
	log.Infof("[%s] request: %s", req.Headers["Request-ID"], feature)
	log.Infof("[%s] request: %s", req.Headers["Request-ID"], referenceNumber)

	// jsonReq, _ := json.Marshal(req)
	// log.Infof("[%s] request: %v", req.Headers["Request-ID"], string(jsonReq))
	if feature == "404" {
		resp, _ := InitDb(referenceNumber)
		res.Response, _ = json.Marshal(successResp(resp))
	} else if feature == "receipt" {
		receipt, _ := repo.Transaction.GetTrxCustom(referenceNumber)
		log.Infof("[%s] request: %s", req.Headers["Request-ID"], receipt)
		if receipt != nil {
			res.Response, _ = json.Marshal(successResp(receipt))
		} else {
			res.Response, _ = json.Marshal(newResponse("80", "Data Not Found"))
		}
	} else {
		srcAccount := req.Params["srcAccount"]
		req.Params["account"] = req.Params["srcAccount"]

		if srcAccount == "" {
			res.Response, _ = json.Marshal(newResponse("01", "Bad Request"))
		} else {
			gateMsg := transport.SendToGate("gate.shared", "01", req.Params)
			log.Infof("[%s] Info: %v", req.Params)
			if gateMsg.ResponseCode == "00" {
				res.Response, _ = json.Marshal(successResp(gateMsg.Data))
			} else {
				res.Response, _ = json.Marshal(newResponse(gateMsg.ResponseCode, "Data Not Found"))
			}
		}
	}

	return nil
}

func (h *WebTellerHandler) ReInquiryMPN(ctx context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {
	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], string(res.Response))
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] error : %v", req.Headers["Request-ID"], e)
		res.Response, _ = json.Marshal(newResponse("99", "Internal Server Error"))
	})
	params := map[string]string{
		"tellerID":        req.Params["tellerID"],
		"tellerPass":      req.Params["tellerPass"],
		"amount":          req.Params["amount"],
		"txType":          req.Params["txType"],
		"srcAccount":      req.Params["srcAccount"],
		"customerId":      req.Params["customerId"],
		"inqData":         req.Params["inqData"],
		"referenceNumber": req.Params["referenceNumber"],
		"termType":        "6010",
		"termId":          "WTELLER",
	}
	gateMsg := transport.SendToGate("gate.shared", "69", params)

	gateMsg.Data["featureName"] = req.Params["featureName"]
	gateMsg.Data["featureCode"] = req.Params["featureCode"]
	gateMsg.Data["txRefNumber"] = req.Params["referenceNumber"]
	gateMsg.Data["responseCode"] = gateMsg.ResponseCode
	gateMsg.Data["txStatus"] = "SUCCESS"

	dataReceipt, _ := json.Marshal(gateMsg.Data)
	log.Infof("Data test Response:", dataReceipt)
	res.Response, _ = json.Marshal(successResp(gateMsg))
	return nil
}
