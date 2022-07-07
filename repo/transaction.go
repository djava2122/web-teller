package repo

import (
	"bytes"
	"encoding/json"
	"fmt"

	"git.pactindo.com/ebanking/common/log"
	"git.pactindo.com/ebanking/common/pg"
)

type MTransaction struct {
	ID                int32   `json:"id"`
	ReferenceNumber   string  `json:"referenceNumber"`
	FeatureId         int     `json:"featureId"`
	FeatureCode       int     `json:"featureCode"`
	FeatureName       string  `json:"featureName"`
	FeatureGroupCode  string  `json:"featureGroupCode"`
	FeatureGroupName  string  `json:"featureGroupName"`
	ProductId         int     `json:"productId"`
	ProductCode       string  `json:"productCode"`
	ProductName       string  `json:"productName"`
	TransactionDate   string  `json:"transactionDate"`
	CurrencyCode      string  `json:"currencyCode"`
	TransactionAmount float64 `json:"transactionAmount"`
	Fee               float64 `json:"fee"`
	CustomerReference string  `json:"customerReference"`
	BillerName        string  `json:"billerName"`
	MerchantType      string  `json:"merchanType"`
	TransactionStatus string  `json:"transactionStatus"`
	Created           string  `json:"created"`
	CreatedBy         string  `json:"createdBy"`
	Updated           string  `json:"updated"`
	UpdatedBy         string  `json:"updatedBy"`
	BranchCode        string  `json:"branchCode"`
	BranchName        string  `json:"branchName"`
	TransactionType   string  `json:transactionType`
	SrcAccount        string  `json:srcAccount`
	ResponseCode      string  `json:"responseCode"`
	Receipt           string  `json:"receipt"`
}

type GetReceipt struct {
	Id              int                    `json:id`
	BranchCode      string                 `json:"branchCode"`
	BranchName      string                 `json:"branchName"`
	TransactionType string                 `json:transactionType`
	SrcAccount      string                 `json:srcAccount`
	Receipt         map[string]interface{} `json:receipt`
	JumlahCetak     int                    `json:jumlah_cetak`
	ResponseCode    string                 `json:responseCode`
}
type Mbranch struct {
	BranchCode      string `json:"branchCode"`
	BranchName      string `json:"branchName"`
	TransactionType string `json:transactionType`
	SrcAccount      string `json:srcAccount`
	JumlahCetak     int    `json:jumlah_cetak`
}
type UCetak struct {
	Id    string `json:"id"`
	Cetak int    `json:"cetak"`
}
type Filter struct {
	FeatureCode string `json:"featureCode"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
}

type TransactionReport struct {
	FeatureName       string  `json:"featureName"`
	FeatureCode       string  `json:"featureCode"`
	FeatureGroupCode  string  `json:"featureGroupCode"`
	FeatureGroupName  string  `json:"featureGroupName"`
	TransactionDate   string  `json:"transactionDate"`
	TransactionAmount float64 `json:"transactionAmount"`
	Fee               float64 `json:"fee"`
	TransactionStatus string  `json:"transactionStatus"`
	ReferenceNumber   string  `json:"referenceNumber"`
	CustomerReference string  `json:"customerReference"`
	CurrencyCode      string  `json:"currencyCode"`
	CreatedBy         string  `json:"createdBy"`
	BranchCode        string  `json:"branchCode"`
	BranchName        string  `json:"branchName"`
	TransactionType   string  `json:transactionType`
	SrcAccount        string  `json:srcAccount`
	ResponseCode      string  `json:"responseCode"`
}

type transaction struct{}

func (_ transaction) Update(trx UCetak) error {
	sql := `UPDATE t_transaction
	SET jumlah_cetak = $1
	WHERE id = $2;`

	_, err := pg.DB.Exec(sql, trx.Cetak, trx.Id)
	//log.Infof("[%s] Update Table: %v", ar)

	if err != nil {
		log.Errorf("OI OI ERROR :", err)
		return err
	}

	return nil
}
func (_ transaction) UpdateMpn(resp, sts, receipt, reff string) error {
	sql := `UPDATE t_transaction
	SET response_code = $1, transaction_status = $2,receipt = $3
	WHERE reference_number = $4;`

	_, err := pg.DB.Exec(sql, resp, sts, receipt, reff)
	//log.Infof("[%s] Update Table: %v", ar)

	if err != nil {
		log.Errorf("OI OI ERROR :", err)
		return err
	}

	return nil
}
func (_ transaction) Save(trx MTransaction) error {
	sql := `insert into t_transaction (
				reference_number, feature_id, feature_code, feature_name, product_id, product_code, product_name,
				biller_name, transaction_date, transaction_amount, fee, merchant_type, currency_code, customer_reference, created, createdby, 
				updated, updatedby, transaction_status, branch_code, response_code, feature_group_name, feature_group_code, src_account, trx_type, branch_name, receipt
			) values (
					$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27
			)`

	_, err := pg.DB.Exec(sql,
		trx.ReferenceNumber, trx.FeatureId, trx.FeatureCode, trx.FeatureName, trx.ProductId, trx.ProductCode,
		trx.ProductName, trx.BillerName, trx.TransactionDate, trx.TransactionAmount, trx.Fee, trx.MerchantType, trx.CurrencyCode,
		trx.CustomerReference, trx.Created, trx.CreatedBy, trx.Updated, trx.UpdatedBy, trx.TransactionStatus, trx.BranchCode,
		trx.ResponseCode, trx.FeatureGroupName, trx.FeatureGroupCode, trx.SrcAccount, trx.TransactionType, trx.BranchName, trx.Receipt)
	//log.Infof("[%s] Insert Table: %v", ar)

	if err != nil {
		log.Errorf("OI OI ERROR :", err)
		return err
	}

	return nil
}

func (_ transaction) Filter(teller string, cabang string) (result []TransactionReport, err error) {
	query := bytes.NewBufferString("select feature_name,feature_code,feature_group_code,feature_group_name,transaction_date,transaction_amount,fee,transaction_status,reference_number,customer_reference,currency_code,createdby,branch_code,branch_name, trx_type, src_account, response_code from t_transaction ")
	if teller != "" && cabang == "" {
		query.WriteString(fmt.Sprintf(" WHERE createdby = '%s'", teller))
	} else if cabang != "" {
		query.WriteString(fmt.Sprintf(" WHERE branch_code = '%s'", cabang))
	}

	query.WriteString(" ORDER BY created DESC")

	rows, err := pg.DB.Query(query.String())
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		datas := TransactionReport{}
		err := rows.Scan(
			&datas.FeatureName,
			&datas.FeatureCode,
			&datas.FeatureGroupCode,
			&datas.FeatureGroupName,
			&datas.TransactionDate,
			&datas.TransactionAmount,
			&datas.Fee,
			&datas.TransactionStatus,
			&datas.ReferenceNumber,
			&datas.CustomerReference,
			&datas.CurrencyCode,
			&datas.CreatedBy,
			&datas.BranchCode,
			&datas.BranchName,
			&datas.TransactionType,
			&datas.SrcAccount,
			&datas.ResponseCode,
		)
		if datas.TransactionType == "TUNAI" {
			datas.SrcAccount = "-"
		}
		if err != nil {
			return nil, err
		}
		result = append(result, datas)
	}
	//log.Infof("Result Report : %s", result)
	return
}

func (_ transaction) GetTrxCustom(teller, start, end string) (result []GetReceipt, err error) {
	query := bytes.NewBufferString("select id, branch_code, branch_name, trx_type, src_account, receipt, jumlah_cetak, response_code, createdby from t_transaction")
	if teller != "" {
		query.WriteString(fmt.Sprintf(" WHERE (reference_number = '%s' or customer_reference = '%s') and (response_code = '00' or response_code = '06')", teller, teller))
	}
	if start != "All" {
		tglAwal := start + " 00:00:00"
		tglAkhir := end + " 23:59:59"
		query.WriteString(fmt.Sprintf(" and  transaction_date between '%s' and '%s'", tglAwal, tglAkhir))
	}
	query.WriteString(" ORDER BY created DESC")
	//log.Infof("Query :", query.String())
	rows, err := pg.DB.Query(query.String())
	if err != nil {
		return nil, err
	}
	//log.Infof("Test query:", rows)
	for rows.Next() {
		datas := GetReceipt{}
		var tampung string
		var createdBy string
		err := rows.Scan(
			&datas.Id,
			&datas.BranchCode,
			&datas.BranchName,
			&datas.TransactionType,
			&datas.SrcAccount,
			&tampung,
			&datas.JumlahCetak,
			&datas.ResponseCode,
			&createdBy,
		)
		json.Unmarshal([]byte(tampung), &datas.Receipt)
		datas.Receipt["branchName"] = datas.BranchName
		datas.Receipt["transactionType"] = datas.TransactionType
		datas.Receipt["srcAccount"] = datas.SrcAccount
		datas.Receipt["responseCode"] = datas.ResponseCode
		datas.Receipt["createdBy"] = createdBy
		if err != nil {
			return nil, err
		}
		result = append(result, datas)
	}

	return
}

func (_ transaction) GetTransactionReceipt(reffNumber string) (result interface{}) {
	sql := "receipt from t_transaction where reference_number = $1"

	var o interface{}
	var tampung string
	err := pg.DB.QueryRow(sql, reffNumber).Scan(&tampung)
	json.Unmarshal([]byte(tampung), &o)
	//log.Infof("Select receip: %v", o)
	//log.Infof("Select tampung: %v", tampung)

	if err == nil {
		return &o
	}

	return nil
}
func (_ transaction) GetBranch(reffNumber string) (code, name, trxType, src string, err error) {
	sql := "select branch_code,branch_name, trx_type, src_account from t_transaction where reference_number = $1"

	err = pg.DB.QueryRow(sql, reffNumber).Scan(&code, &name, &trxType, &src)
	//log.Infof("Select receip: ", err)

	if err == nil {
		return code, name, trxType, src, err
	}

	return code, name, trxType, src, err
}

func (_ transaction) FindTransaction(custRef string) (string, string, error) {
	var customerRef, trxStatus string
	stmt, err := pg.DB.Prepare("SELECT customer_reference, transaction_status FROM t_transaction WHERE customer_reference=$1")
	if err != nil {
		return "", "", err
	}
	defer stmt.Close()
	err = stmt.QueryRow(custRef).Scan(&customerRef, &trxStatus)
	if err != nil {
		return "", "", err
	}
	return customerRef, trxStatus, nil
}

func (_ transaction) TransactionReport(filter string) (result []TransactionReport, err error) {
	queryGetTransaction := "select feature_name,feature_code,feature_group_code,feature_group_name,transaction_date,transaction_amount,fee,transaction_status,reference_number,customer_reference,currency_code,createdby,branch_code,branch_name, trx_type, src_account, response_code from t_transaction "
	rows, err := pg.DB.Query(fmt.Sprintf("%v %v", queryGetTransaction, filter))
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		datas := TransactionReport{}
		err := rows.Scan(
			&datas.FeatureName,
			&datas.FeatureCode,
			&datas.FeatureGroupCode,
			&datas.FeatureGroupName,
			&datas.TransactionDate,
			&datas.TransactionAmount,
			&datas.Fee,
			&datas.TransactionStatus,
			&datas.ReferenceNumber,
			&datas.CustomerReference,
			&datas.CurrencyCode,
			&datas.CreatedBy,
			&datas.BranchCode,
			&datas.BranchName,
			&datas.TransactionType,
			&datas.SrcAccount,
			&datas.ResponseCode,
		)
		if datas.TransactionType == "TUNAI" {
			datas.SrcAccount = "-"
		}
		if err != nil {
			return nil, err
		}
		result = append(result, datas)
	}
	return
}

func (_ transaction) FilterTransaction(filter string, page, pageSize int) (result []TransactionReport, rowCount int, err error) {
	queryGetTransaction := "select feature_name,feature_code,feature_group_code,feature_group_name,transaction_date,transaction_amount,fee,transaction_status,reference_number,customer_reference,currency_code,createdby,branch_code,branch_name, trx_type, src_account, response_code from t_transaction "
	rows, err := pg.DB.Query(fmt.Sprintf("%v %v LIMIT %v OFFSET %v", queryGetTransaction, filter, pageSize, (page-1)*pageSize))
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	queryCountData := "SELECT COUNT(id) FROM t_transaction "
	err = pg.DB.QueryRow(fmt.Sprintf("%v %v", queryCountData, filter)).Scan(&rowCount)
	log.Infof(fmt.Sprintf("%v %v", queryCountData, filter))
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		datas := TransactionReport{}
		err := rows.Scan(
			&datas.FeatureName,
			&datas.FeatureCode,
			&datas.FeatureGroupCode,
			&datas.FeatureGroupName,
			&datas.TransactionDate,
			&datas.TransactionAmount,
			&datas.Fee,
			&datas.TransactionStatus,
			&datas.ReferenceNumber,
			&datas.CustomerReference,
			&datas.CurrencyCode,
			&datas.CreatedBy,
			&datas.BranchCode,
			&datas.BranchName,
			&datas.TransactionType,
			&datas.SrcAccount,
			&datas.ResponseCode,
		)
		if datas.TransactionType == "TUNAI" {
			datas.SrcAccount = "-"
		}
		if err != nil {
			return nil, 0, err
		}
		result = append(result, datas)
	}
	return
}

var Transaction transaction
