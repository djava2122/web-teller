package util

import (
	"encoding/json"
	"time"
)

const (
	DefaultFormatDate = "20060102 15:04:05"
	SCHEDULED         = "SCHEDULED"
)

func ParseIntToDate(tmInt int64) time.Time {
	return time.Unix(tmInt, 0)
}

func ConvertMap(m map[string]interface{}) map[string]string {
	m2 := make(map[string]string)
	for k, value := range m {
		switch v := value.(type) {
		case []interface{}:
			bj, _ := json.Marshal(v)
			m2[k] = string(bj)
		case string:
			m2[k] = v
		case map[string]interface{}:
			bj, _ := json.Marshal(v)
			m2[k] = string(bj)
		default:
			bj, _ := json.Marshal(v)
			m2[k] = string(bj)
		}
	}
	return m2
}
