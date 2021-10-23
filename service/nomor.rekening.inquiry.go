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
	TxRefNumber       string         `json:txRefNumber`
	TermId            sql.NullString `json:termId`
	ResponseCode      string         `json:responseCode`
	TxDate            string         `json:txDate`
	BookDate          sql.NullString `json:bookDate`
	Ntb               sql.NullString `json:ntb`
	Ntpn              sql.NullString `json:ntpn`
	Stan              sql.NullString `json:stan`
	CustomerReference sql.NullString `json:customerReference`
	Npwp              sql.NullString `json:npwp`
	PrayerName        sql.NullString `json:prayerName`
	PrayerAddress     sql.NullString `json:prayerAddress`
	Nop               sql.NullString `json:nop`
	TaxAccount        sql.NullString `json:taxAccount`
	DepositTypeCode   sql.NullString `json:depositTypeCode`
	TaxPeriod         sql.NullString `json:taxPeriod`
	SkNumber          sql.NullString `json:skNumber`
	Amount            sql.NullString `json:amount`
	CurrencyCode      sql.NullString `json:currencyCode`
	PrayerId          sql.NullString `json:prayerId`
	DocumentType      sql.NullString `json:documentType`
	DocumentNumber    sql.NullString `json:documentNumber`
	DocumentDate      sql.NullString `json:documentDate`
	KppbcCode         sql.NullString `json:kppbcCode`
	Kl                sql.NullString `json:kl`
	UnitEselon        sql.NullString `json:unitEselon`
	SatkerCode        sql.NullString `json:satkerCode`
	FeatureCode       string         `json:featureCode`
	FeatureName       string         `json:featurName`
}

type ReqMPN struct {
	TxRefNumber     string            `json:txRefNumber`
	ResponseCode    string            `json:responseCode`
	Datetime        sql.NullString    `json:dateTime`
	TermType        sql.NullString    `json:termType`
	TermId          sql.NullString    `json:termId`
	CustomerId      sql.NullString    `json:customerId`
	SrcAccount      sql.NullString    `json:srcAccount`
	Amount          sql.NullString    `json:amount`
	InqData         sql.NullString    `json:inqData`
	ReferenceNumber sql.NullString    `json:referenceNumber`
	Additional      map[string]string `json:additional`
}

func InitDb(req, startDate, endDate string) (result []interface{}, err error) {
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
		sql = "select fields -> 'termID' as termID, trrfnm, trrspc, concat(rtdate ,' ', isomsg -> '12') as trxDate,fields-> 'bookDate' as bookDate, fields -> 'ntb' as ntb , fields -> 'stan' as stan, fields -> 'customerId' as customerReference, fields -> 'npwp' as npwp, fields -> 'payerName'as payerName, fields -> 'payerAddress' as payerAddress, fields->'taxAccount' as taxAccount, fields -> 'depositTypeCode'as depositTypeCode,fields -> 'taxPeriod' as taxPeriod, fields -> 'skNumber' as skNumber,cast(isomsg -> '4' as float)/100 as amount, fields -> 'currencyCode' as currencyCode, fields -> 'ntpn' as ntpn, fields -> 'nop' as nop from mgate.t_transaction where rtdate between $1 and $2 and (trrfnm = $3 or fields -> 'customerId' = $4) and trftcd = 'MPN2' and (trrspc = '00' or trrspc = '06' or trrspc = '90' or trrspc = '92')"
	} else if jenis >= 7 {
		sql = "select fields -> 'termID' as termID, trrfnm, trrspc, concat(rtdate ,' ', isomsg -> '12') as trxDate,fields-> 'bookDate' as bookDate, fields -> 'ntb' as ntb , fields -> 'stan' as stan, fields -> 'customerId' as customerReference, fields -> 'payerName'as payerName,cast(isomsg -> '4' as float)/100 as amount, fields -> 'currencyCode' as currencyCode, fields -> 'kl' as kl, fields -> 'unitEselon' as unitEselon, fields -> 'satkerCode' as satkerCode,  fields -> 'ntpn' as ntpn from mgate.t_transaction where rtdate between $1 and $2 and (trrfnm = $3 or fields -> 'customerId' = $4)  and trftcd = 'MPN2' and (trrspc = '00' or trrspc = '06' or trrspc = '90' or trrspc = '92')"
	} else {
		sql = "select fields -> 'termID' as termID, trrfnm, trrspc, concat(rtdate ,' ', isomsg -> '12') as trxDate,fields-> 'bookDate' as bookDate, fields -> 'ntb' as ntb , fields -> 'stan' as stan, fields -> 'customerId' as customerReference, fields -> 'payerName'as payerName, cast(isomsg -> '4' as float)/100 as amount, fields -> 'currencyCode' as currencyCode, fields -> 'documentType' as documentType, fields -> 'documentNumber' as documentNumber, fields -> 'documentDate' as documentDate, fields -> 'kppbcCode'as kppbcCode, fields -> 'ntpn' as ntpn from mgate.t_transaction where rtdate between $1 and $2 and (trrfnm = $3 or fields -> 'customerId' = $4) and trftcd = 'MPN2' and (trrspc = '00' or trrspc = '06' or trrspc = '90' or trrspc = '92')"
	}
	log.Infof("query : ", sql)
	rows, _ := db.Query(sql, startDate, endDate, req, req)
	// if err != nil {
	// 	return nil, err
	// }
	log.Infof("Test : ", rows)
	for rows.Next() {
		dt := GetMgateStruct{}
		if jenis <= 3 {
			err = rows.Scan(
				&dt.TermId,
				&dt.TxRefNumber,
				&dt.ResponseCode,
				&dt.TxDate,
				&dt.BookDate,
				&dt.Ntb,
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
				&dt.Ntpn,
				&dt.Nop,
			)
		} else if jenis >= 7 {
			err = rows.Scan(
				&dt.TermId,
				&dt.TxRefNumber,
				&dt.ResponseCode,
				&dt.TxDate,
				&dt.BookDate,
				&dt.Ntb,
				&dt.Stan,
				&dt.CustomerReference,
				&dt.PrayerName,
				&dt.Amount,
				&dt.CurrencyCode,
				&dt.Kl,
				&dt.UnitEselon,
				&dt.SatkerCode,
				&dt.Ntpn,
			)
		} else {
			err = rows.Scan(
				&dt.TermId,
				&dt.TxRefNumber,
				&dt.ResponseCode,
				&dt.TxDate,
				&dt.BookDate,
				&dt.Ntb,
				&dt.Stan,
				&dt.CustomerReference,
				&dt.PrayerName,
				&dt.Amount,
				&dt.CurrencyCode,
				&dt.DocumentType,
				&dt.DocumentNumber,
				&dt.DocumentDate,
				&dt.KppbcCode,
				&dt.Ntpn,
			)
		}
		dt.FeatureCode = "404"
		dt.FeatureName = "MPN"
		status := ""
		var code string
		var name string
		var trxType string
		var src string
		if dt.TermId.String == "WTELLER" || dt.TermId.String == "KWTELLER" {
			code, name, trxType, src, _ = repo.Transaction.GetBranch(dt.TxRefNumber)
			log.Infof("update table teller :", name)
		} else {
			code = "ID0011001"
			name = "1001 - Cabang Utama"
		}
		resp := make(map[string]interface{})
		if dt.ResponseCode == "00" {
			status = "SUCCESS"
		} else if dt.ResponseCode == "99" || dt.ResponseCode == "19" {
			status = "FAILED"
		} else {
			sql = "select trrfnm, trrspc, fields -> 'dateTime' as dateTime, fields -> 'inqData' as inqData, fields -> 'termID' as termID, fields -> 'termType' as termType, cast(isomsg -> '4' as integer)/100 as amount, fields -> 'additional' as additional, fields -> 'srcAccount' as srcAccount from mgate.t_transaction where (rtdate between $1 and $2 ) and (trrfnm = $3 or fields -> 'customerId' = $4) and trftcd = 'MPN2' and (trrspc = '06'  or trrspc = '90'  or trrspc = '92')"
			rows, _ = db.Query(sql, startDate, endDate, req, req)
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
				dt.TxDate = Rdt.Datetime.String
				status = "PENDING"
				json.Unmarshal([]byte(tamp), &Rdt.Additional)
				reqParams := map[string]string{
					"txType":             "07",
					"dateTime":           time.Now().Format("20060102150405"),
					"termType":           Rdt.TermType.String,
					"termId":             Rdt.TermId.String,
					"customerId":         req,
					"srcAccount":         Rdt.SrcAccount.String,
					"amount":             Rdt.Amount.String,
					"description":        "",
					"inqData":            Rdt.InqData.String,
					"referenceNumber":    util.RandomNumber(12),
					"orgDateTime":        Rdt.Datetime.String,
					"orgReferenceNumber": Rdt.TxRefNumber,
				}
				log.Infof("ddd ", reqParams)
				gateMsg := transport.SendToGate("gate.shared", "27", reqParams)
				log.Infof("Gate send response:  ", gateMsg)

				if gateMsg.ResponseCode == "00" {
					status = "SUCCESS"
				} else if gateMsg.ResponseCode == "AW" {
					dt.ResponseCode = Rdt.ResponseCode
					status = "Reinquiry Gagal"
				} else if gateMsg.ResponseCode == "99" || dt.ResponseCode == "19" {
					status = "FAILED"
				} else {
					status = "PENDING"
				}
				log.Infof("Reinquiry Mgate rc timeout :", gateMsg)
				if gateMsg.ResponseCode == "00" {
					resp["Receipt"] = gateMsg.Data
					gateMsg.Data["featureCode"] = dt.FeatureCode
					gateMsg.Data["featureName"] = dt.FeatureName
					gateMsg.Data["txRefNumber"] = dt.TxRefNumber
					gateMsg.Data["txDate"] = reqParams["orgDateTime"]
					gateMsg.Data["customerReference"] = reqParams["customerId"]
					gateMsg.Data["txStatus"] = status
					receipt, _ := json.Marshal(gateMsg.Data)
					// sqlMgate := `update mgate.t_transaction set trrspc = $1 where trrfnm = $2;`
					// db.Exec(sqlMgate, "00", dt.TxRefNumber)
					// db.Close()
					if Rdt.TermId.String == "WTELLER" || Rdt.TermId.String == "KWTELLER" {
						updateTabel := repo.Transaction.UpdateMpn(gateMsg.ResponseCode, status, string(receipt), dt.TxRefNumber)
						log.Infof("update table teller :", updateTabel)
					}
				} else {
					resp["Receipt"] = map[string]string{
						"txRefNumber":       dt.TxRefNumber,
						"responseCode":      dt.ResponseCode,
						"txDate":            dt.TxDate,
						"bookDate":          dt.BookDate.String,
						"ntb":               dt.Ntb.String,
						"ntpn":              dt.Ntpn.String,
						"stan":              dt.Stan.String,
						"customerReference": dt.CustomerReference.String,
						"npwp":              dt.Npwp.String,
						"payerName":         dt.PrayerName.String,
						"payerAddress":      dt.PrayerAddress.String,
						"nop":               dt.Nop.String,
						"taxAccount":        dt.TaxAccount.String,
						"depositTypeCode":   dt.DepositTypeCode.String,
						"taxPeriod":         dt.TaxPeriod.String,
						"skNumber":          dt.SkNumber.String,
						"amount":            dt.Amount.String,
						"currencyCode":      dt.CurrencyCode.String,
						"payerId":           dt.PrayerId.String,
						"documentType":      dt.DocumentType.String,
						"documentNumber":    dt.DocumentNumber.String,
						"documentDate":      dt.DocumentDate.String,
						"kppbcCode":         dt.KppbcCode.String,
						"kl":                dt.Kl.String,
						"unitEselon":        dt.UnitEselon.String,
						"satkerCode":        dt.SatkerCode.String,
						"featureCode":       dt.FeatureCode,
						"featureName":       dt.FeatureName,
						"branchCode":        code,
						"branchName":        name,
						"transactionType":   trxType,
						"srcAccount":        src,
						"txStatus":          status,
					}
				}
				result = append(result, resp)
				log.Infof("result : ", result...)
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
			"bookDate":          dt.BookDate.String,
			"ntb":               dt.Ntb.String,
			"ntpn":              dt.Ntpn.String,
			"stan":              dt.Stan.String,
			"customerReference": dt.CustomerReference.String,
			"npwp":              dt.Npwp.String,
			"payerName":         dt.PrayerName.String,
			"payerAddress":      dt.PrayerAddress.String,
			"nop":               dt.Nop.String,
			"taxAccount":        dt.TaxAccount.String,
			"depositTypeCode":   dt.DepositTypeCode.String,
			"taxPeriod":         dt.TaxPeriod.String,
			"skNumber":          dt.SkNumber.String,
			"amount":            dt.Amount.String,
			"currencyCode":      dt.CurrencyCode.String,
			"payerId":           dt.PrayerId.String,
			"documentType":      dt.DocumentType.String,
			"documentNumber":    dt.DocumentNumber.String,
			"documentDate":      dt.DocumentDate.String,
			"kppbcCode":         dt.KppbcCode.String,
			"kl":                dt.Kl.String,
			"unitEselon":        dt.UnitEselon.String,
			"satkerCode":        dt.SatkerCode.String,
			"featureCode":       dt.FeatureCode,
			"featureName":       dt.FeatureName,
			"branchCode":        code,
			"branchName":        name,
			"transactionType":   trxType,
			"srcAccount":        src,
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
	startDate := req.Params["startDate"]
	endDate := req.Params["endDate"]

	if startDate == "All" {
		startDate = time.Now().Format("2006-01-02")
	}
	if endDate == "All" {
		endDate = time.Now().Format("2006-01-02")
	}
	log.Infof("[%s] request: %s", req.Headers["Request-ID"], feature)
	log.Infof("[%s] request: %s", req.Headers["Request-ID"], referenceNumber)
	log.Infof("[%s] request: %s", req.Headers["Request-ID"], startDate)
	log.Infof("[%s] request: %s", req.Headers["Request-ID"], endDate)

	// jsonReq, _ := json.Marshal(req)
	// log.Infof("[%s] request: %v", req.Headers["Request-ID"], string(jsonReq))
	if feature == "404" {
		resp, _ := InitDb(referenceNumber, startDate, endDate)
		res.Response, _ = json.Marshal(successResp(resp))
	} else if feature == "receipt" {
		receipt, _ := repo.Transaction.GetTrxCustom(referenceNumber, req.Params["startDate"], endDate)
		log.Infof("[%s] request: %s", req.Headers["Request-ID"], receipt)
		if receipt != nil {
			res.Response, _ = json.Marshal(successResp(receipt))
		} else {
			res.Response, _ = json.Marshal(newResponse("80", "Data Not Found"))
		}
	} else if feature == "Reinquiry" {
		params := map[string]string{
			"type":            endDate,
			"referenceNumber": req.Params["referenceNumber"],
		}
		gateMsg := transport.SendToGate("gate.shared", "B003", params)
		log.Infof("[%s] request: %s", req.Headers["Request-ID"], gateMsg)
		if gateMsg.ResponseCode == "00" {
			gateMsg.Data["responseCode"] = gateMsg.ResponseCode
			gateMsg.Data["message"] = gateMsg.Description
			gateMsg.Data["featureCode"] = req.Params["startDate"]
			gateMsg.Data["txStatus"] = gateMsg.Description
			gateMsg.Data["txRefNumber"] = req.Params["referenceNumber"]
			gateMsg.Data["txDate"] = time.Now().Format("2006-01-02 15:04:05")
			gateMsg.Data["featureName"] = req.Params["featureName"]
			if gateMsg.Data["featureCode"] == "202" {
				gateMsg.Data["amount"] = gateMsg.Data["rpPayment"]
				gateMsg.Data["customerReference"] = gateMsg.Data["customerId"]
			} else if gateMsg.Data["featureCode"] == "319" || gateMsg.Data["featureCode"] == "303" {
				gateMsg.Data["amount"] = gateMsg.Data["totalAmount"]
				gateMsg.Data["customerReference"] = gateMsg.Data["customerId"]
			}
			var receipt repo.GetReceipt
			receipt.Id = 0
			receipt.JumlahCetak = 0
			receipt.Receipt = gateMsg.Data
			recp, _ := json.Marshal(gateMsg.Data)
			updateTabel := repo.Transaction.UpdateMpn(gateMsg.ResponseCode, gateMsg.Description, string(recp), req.Params["referenceNumber"])
			log.Infof("update table teller :", updateTabel)
			res.Response, _ = json.Marshal(successResp(receipt))
		} else {
			receipt, _ := repo.Transaction.GetTrxCustom(req.Params["referenceNumber"], "All", time.Now().Format("2006-01-02"))
			log.Infof("[%s] request: %s", req.Headers["Request-ID"], receipt)
			sts := "FAILED"
			if gateMsg.ResponseCode == "06" {
				receipt[0].Receipt["responseCode"] = "06"
				receipt[0].Receipt["txStatus"] = "PENDING"
				sts = "PENDING"
			} else {
				gateMsg.ResponseCode = "19"
				receipt[0].Receipt["responseCode"] = "91"
				receipt[0].Receipt["txStatus"] = "FAILED"
				sts = "FAILED"
			}
			recp, _ := json.Marshal(receipt[0].Receipt)
			updateTabel := repo.Transaction.UpdateMpn(gateMsg.ResponseCode, sts, string(recp), req.Params["referenceNumber"])
			log.Infof("update table teller :", updateTabel)
			if receipt != nil {
				res.Response, _ = json.Marshal(respCekStatus(receipt, gateMsg.ResponseCode, gateMsg.Description))
			} else {
				res.Response, _ = json.Marshal(newResponse("80", "Data Not Found"))
			}
		}
	} else {
		srcAccount := req.Params["srcAccount"]
		req.Params["account"] = req.Params["srcAccount"]

		if srcAccount == "" {
			res.Response, _ = json.Marshal(newResponse("01", "Bad Request"))
		} else {
			gateMsg := transport.SendToGate("gate.shared", "01", req.Params)
			if gateMsg.Data["customerName"] == "" {
				gateMsg.ResponseCode = "80"
				gateMsg.Description = "Data Not Found"
			}
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
