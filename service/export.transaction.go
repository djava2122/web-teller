package service

import (
	"encoding/csv"
	"git.pactindo.com/ebanking/common/log"
	"git.pactindo.com/ebanking/common/micro"
	"git.pactindo.com/ebanking/common/util"
	"git.pactindo.com/ebanking/web-teller/repo"
	"os"
	"strconv"
)

func ExportTransaction(data []repo.TransactionReport) (bool, string, error) {
	conf := micro.GetConfig()
	rand := util.RandomNumber(4)
	nameFiles := "file-" + rand + ".csv"
	file, err := os.Create(conf["FILE_LOCATION"] + nameFiles)
	if err != nil {
		log.Errorf("Cannot create file", err)
	}

	defer file.Close()

	w := csv.NewWriter(file)

	defer w.Flush()

	// Using Write
	_ = w.Write([]string{
		"Feature Code",
		"Feature Name",
		"Feature Group Code",
		"Feature Group Name",
		"Transaction Date",
		"Amount",
		"Fee",
		"Transaction Status",
		"Ref Number",
		"Customer Ref",
		"Currency Code",
		"Created By",
		"Branch Code",
		"Branch Name",
		"Transaction Type",
		"Source Account",
		"Response Code",
	})
	for _, v := range data {
		row := []string{
			v.FeatureCode,
			v.FeatureName,
			v.FeatureGroupCode,
			v.FeatureGroupName,
			v.TransactionDate,
			strconv.FormatFloat(v.TransactionAmount, 'f', -1, 64),
			strconv.FormatFloat(v.Fee, 'f', -1, 64),
			v.TransactionStatus,
			v.ReferenceNumber,
			v.CustomerReference,
			v.CurrencyCode,
			v.CreatedBy,
			v.BranchCode,
			v.BranchName,
			v.TransactionType,
			v.SrcAccount,
			v.ResponseCode,
		}
		if err := w.Write(row); err != nil {
			log.Infof("Cannot write to file", err)
			//return false, err
		}
	}
	return true, nameFiles, nil
}
