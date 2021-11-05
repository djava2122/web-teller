package service

import (
	"context"
	"database/sql"
	"fmt"

	"gitlab.pactindo.com/ebanking/common/log"
	"gitlab.pactindo.com/ebanking/common/trycatch"
	wtproto "gitlab.pactindo.com/ebanking/web-teller/proto"
	"gitlab.pactindo.com/ebanking/web-teller/repo"
)

func (h *WebTellerHandler) ReportMpnAll(_ context.Context, req *wtproto.APIREQ, res *wtproto.APIRES) error {
	defer func() {
		log.Infof("[%s] response: %v", req.Headers["Request-ID"], string(res.Response))
	}()
	defer trycatch.Catch(func(e trycatch.Exception) {
		log.Infof("[%s] error : %v", req.Headers["Request-ID"], e)
		res.Response, _ = json.Marshal(newResponse("99", "Internal Server Error"))
	})
	startDate := req.Params["startDate"]
	endDate := req.Params["endDate"]
	gateTabel, _ := InitReport(startDate, endDate)
	if gateTabel != nil {
		res.Response, _ = json.Marshal(successResp(gateTabel))
	} else {
		res.Response, _ = json.Marshal(newResponse("99", "Gate Table Error"))
	}

	return nil
}

func InitReport(startDate, endDate string) (result []map[string]interface{}, err error) {
	dsn := fmt.Sprintf("postgres://mgate:mgate2020@172.19.252.114/micro-gate?sslmode=disable")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Infof("unable to connect db, error : %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Infof("unable to connect db, error : %v", err)
	}
	sql := "select isomsg -> '102' as srcAccount,isomsg -> '18' as termType, isomsg -> '41' as termID, trrfnm, trrspc, concat(rtdate ,' ', isomsg -> '12') as trxDate, fields-> 'bookDate' as bookDate, fields -> 'ntb' as ntb , isomsg -> '11' as stan, fields -> 'customerId' as customerReference, fields -> 'npwp' as npwp, fields -> 'payerName'as payerName, fields -> 'payerAddress' as payerAddress, fields->'taxAccount' as taxAccount, fields -> 'depositTypeCode'as depositTypeCode,fields -> 'taxPeriod' as taxPeriod, fields -> 'skNumber' as skNumber,cast(isomsg -> '4' as float)/100 as amount, fields -> 'currencyCode' as currencyCode, fields -> 'ntpn' as ntpn, fields -> 'nop' as nop, fields -> 'payerName'as payerName,cast(isomsg -> '4' as float)/100 as amount, fields -> 'currencyCode' as currencyCode, fields -> 'kl' as kl, fields -> 'unitEselon' as unitEselon, fields -> 'satkerCode' as satkerCode, fields -> 'payerID' as payerId, fields -> 'payerName'as payerName, cast(isomsg -> '4' as float)/100 as amount, fields -> 'currencyCode' as currencyCode, fields -> 'documentType' as documentType, fields -> 'documentNumber' as documentNumber, fields -> 'documentDate' as documentDate, fields -> 'kppbcCode'as kppbcCode from mgate.t_transaction where rtdate between $1 and $2 and trftcd = 'MPN2' and (trrspc = '00' or trrspc = '06' or trrspc = '90' or trrspc = '92' or trrspc = 'AW') and mclass != 'R'"
	log.Infof("query : ", sql)
	rows, _ := db.Query(sql, startDate, endDate)
	log.Infof("Test : ", rows)
	for rows.Next() {
		dt := GetMgateStruct{}
		var rek string
		err = rows.Scan(
			&rek,
			&dt.TermType,
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
			&dt.PrayerName,
			&dt.Amount,
			&dt.CurrencyCode,
			&dt.Kl,
			&dt.UnitEselon,
			&dt.SatkerCode,
			&dt.PrayerId,
			&dt.PrayerName,
			&dt.Amount,
			&dt.CurrencyCode,
			&dt.DocumentType,
			&dt.DocumentNumber,
			&dt.DocumentDate,
			&dt.KppbcCode,
		)
		dt.FeatureCode = "404"
		dt.FeatureName = "MPN"
		status := ""
		var code string
		var name string
		var trxType string
		var src string
		var sumb string
		log.Infof("log quiry mgate 1 :", dt)
		if dt.TermType.String == "6010" {
			code, name, trxType, src, _ = repo.Transaction.GetBranch(dt.TxRefNumber)
			log.Infof("update table teller :", name)
			sumb = "TELLER"
			if name == "" {
				name = "1001 - Cabang Utama"
			}
		} else if dt.TermType.String == "6011" {
			code = "ID0011001"
			name = "1000 - ATM"
			src = rek
			sumb = "ATM"
		} else if dt.TermId.String == "K7020" {
			code = "ID0011001"
			name = "122 - SP2D"
			src = rek
			sumb = "SP2D"
		} else {
			src = rek
			code = "ID0011001"
			name = "1001 - Cabang Utama"
			sumb = dt.TermType.String
		}
		if src == "6000000000" || src == "1000000000" {
			src = "-"
		}
		var resp RespReceipt
		if dt.ResponseCode == "00" {
			status = "SUCCESS"
		} else {
			status = "PENDING"
		}
		resp.Id = 0
		resp.JumlahCetak = 0
		resp.Receipt = map[string]interface{}{
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
			"payerID":           dt.PrayerId.String,
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
			"channel":           sumb,
			"txStatus":          status,
		}
		log.Infof("Log Respon Mgate: ", resp)
		result = append(result, resp.Receipt)
	}
	db.Close()
	log.Infof("result : ", result)
	return result, err
}
