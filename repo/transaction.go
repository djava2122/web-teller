package repo

import (
	"bytes"
	"encoding/json"
	"fmt"

	"gitlab.pactindo.com/ebanking/common/log"
	"gitlab.pactindo.com/ebanking/common/pg"
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
	ResponseCode      string  `json:"responseCode"`
	Receipt           string  `json:"receipt"`
}

type GetReceipt struct {
	Id          int                    `json:id`
	Receipt     map[string]interface{} `json:receipt`
	JumlahCetak int                    `json:jumlah_cetak`
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
}

type transaction struct{}

func (_ transaction) Update(trx UCetak) error {
	sql := `UPDATE t_transaction
	SET jumlah_cetak = $1
	WHERE id = $2;`

	ar, err := pg.DB.Exec(sql, trx.Cetak, trx.Id)
	log.Infof("[%s] Update Table: %v", ar)

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

	ar, err := pg.DB.Exec(sql, resp, sts, receipt, reff)
	log.Infof("[%s] Update Table: %v", ar)

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
				updated, updatedby, transaction_status, branch_code, response_code, feature_group_name, feature_group_code, receipt
			) values (
					$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24
			)`

	ar, err := pg.DB.Exec(sql,
		trx.ReferenceNumber, trx.FeatureId, trx.FeatureCode, trx.FeatureName, trx.ProductId, trx.ProductCode,
		trx.ProductName, trx.BillerName, trx.TransactionDate, trx.TransactionAmount, trx.Fee, trx.MerchantType, trx.CurrencyCode,
		trx.CustomerReference, trx.Created, trx.CreatedBy, trx.Updated, trx.UpdatedBy, trx.TransactionStatus, trx.BranchCode,
		trx.ResponseCode, trx.FeatureGroupName, trx.FeatureGroupCode, trx.Receipt)
	log.Infof("[%s] Insert Table: %v", ar)

	if err != nil {
		log.Errorf("OI OI ERROR :", err)
		return err
	}

	return nil
}

func (_ transaction) Filter(teller string) (result []MTransaction, err error) {
	query := bytes.NewBufferString("select feature_name,feature_code,feature_group_code,feature_group_name,transaction_date,transaction_amount,fee,transaction_status,reference_number,customer_reference,currency_code,createdby,branch_code from t_transaction ")
	if teller != "" {
		query.WriteString(fmt.Sprintf(" WHERE createdby = '%s'", teller))
	}

	query.WriteString(" ORDER BY created DESC")

	rows, err := pg.DB.Query(query.String())
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		datas := MTransaction{}
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
		)
		if err != nil {
			return nil, err
		}
		result = append(result, datas)
	}
	log.Infof("Result Report : %s", result)
	return
}

func (_ transaction) GetTrxCustom(teller, start, end string) (result []GetReceipt, err error) {
	query := bytes.NewBufferString("select id, receipt, jumlah_cetak from t_transaction")
	if teller != "" {
		query.WriteString(fmt.Sprintf(" WHERE (reference_number = '%s' or customer_reference = '%s') and (response_code = '00' or response_code = '06')", teller, teller))
	}
	if start != "All" {
		query.WriteString(fmt.Sprintf(" and  transaction_date between '%s' and '%s'", start, end))
	}
	query.WriteString(" ORDER BY created DESC")
	log.Infof("Query :", query.String())
	rows, err := pg.DB.Query(query.String())
	if err != nil {
		return nil, err
	}
	log.Infof("Test query:", rows)
	for rows.Next() {
		datas := GetReceipt{}
		var tampung string
		err := rows.Scan(
			&datas.Id,
			&tampung,
			&datas.JumlahCetak,
		)
		json.Unmarshal([]byte(tampung), &datas.Receipt)
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
	log.Infof("Select receip: %v", o)
	log.Infof("Select tampung: %v", tampung)

	if err == nil {
		return &o
	}

	return nil
}

var Transaction transaction
