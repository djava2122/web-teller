package repo

import (
	"bytes"
	"fmt"
	"gitlab.pactindo.com/ebanking/common/log"
	"gitlab.pactindo.com/ebanking/common/pg"
	"time"
)

type MTransaction struct {
	ID                int32   `json:"id"`
	ReferenceNumber   string  `json:"referenceNumber"`
	FeatureId         int     `json:"featureId"`
	FeatureCode       int     `json:"featureCode"`
	FeatureName       string  `json:"featureName"`
	ProductId         int     `json:"productId"`
	ProductCode       string  `json:"productCode"`
	ProductName       string  `json:"productName"`
	TransactionDate   string  `json:"transactionDate"`
	CurrencyCode      string  `json:"currencyCode"`
	TransactionAmount float64 `json:"transactionAmount"`
	Fee               int     `json:"fee"`
	CustomerReference string  `json:"customerReference"`
	BillerName        string  `json:"billerName"`
	MerchantType      string  `json:"merchanType"`
	TransactionStatus string  `json:"transactionStatus"`
	Created           string  `json:"created"`
	CreatedBy         string  `json:"createdBy"`
	Updated           string  `json:"updated"`
	UpdatedBy         string  `json:"updatedBy"`
}

type Filter struct {
	FeatureCode       string `json:"featureCode"`
	StartDate         string `json:"startDate"`
	EndDate           string `json:"endDate"`
}

type TransactionReport struct {
	TransactionDate   string  `json:"transactionDate"`
	TransactionAmount float64 `json:"transactionAmount"`
	TransactionStatus string  `json:"transactionStatus"`
	ReferenceNumber   string  `json:"referenceNumber"`
	CustomerReference string  `json:"customerReference"`
	CurrencyCode      string  `json:"currencyCode"`
}

type transaction struct {}

func (_ transaction) Save(trx MTransaction) error {
	sql := `insert into t_transaction (
				reference_number, feature_id, feature_code, feature_name, product_id, product_code, product_name,
				biller_name, transaction_date, transaction_amount, merchant_type, currency_code, customer_reference, created, createdby, 
				updated, updatedby, transaction_status
			) values (
					$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17
			)`

	_, err := pg.DB.Exec(sql,
		trx.ReferenceNumber, trx.FeatureId, trx.FeatureCode, trx.FeatureName, trx.ProductId, trx.ProductCode,
		trx.ProductName, trx.BillerName, trx.TransactionDate, trx.TransactionAmount, trx.MerchantType, trx.CurrencyCode,
		trx.CustomerReference, trx.Created, trx.CreatedBy, trx.Updated, trx.UpdatedBy, trx.TransactionStatus)

	if err != nil {
		log.Errorf("OI OI ERROR :", err)
		return err
	}

	return nil
}

func (_ transaction) Filter(featureCode, startDate, endDate string) (result []TransactionReport, err error) {
	query := bytes.NewBufferString("select transaction_date,transaction_amount,transaction_status,reference_number,customer_reference,currency_code from t_transaction ")
	if featureCode != "" {
		query.WriteString(fmt.Sprintf("WHERE feature_code='%s'", featureCode))
	}
	if startDate != "" {
		sDate, _ := time.Parse("20060102", startDate)
		query.WriteString(fmt.Sprintf(" AND created >= '%s'", sDate.Format("2006-01-02")))
	}
	if endDate != "" {
		eDate, _ := time.Parse("20060102", endDate)
		query.WriteString(fmt.Sprintf(" AND created < '%s'", eDate.Format("2006-01-02")))
	}

	query.WriteString(" ORDER BY created DESC")

	rows, err := pg.DB.Query(query.String())
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		datas := TransactionReport{}
		err := rows.Scan(
			&datas.TransactionDate,
			&datas.TransactionAmount,
			&datas.TransactionStatus,
			&datas.ReferenceNumber,
			&datas.CustomerReference,
			&datas.CurrencyCode,
			)
		if err != nil {
			return nil, err
		}
		result = append(result, datas)
	}

	return
}

var Transaction transaction
