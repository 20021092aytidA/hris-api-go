package requesthelper

import (
	"strings"
)

func ProcessQry(queryStr string) map[string]any {
	var qryMap = make(map[string]any)
	if queryStr == "" {
		return qryMap
	}

	qryArr := strings.Split(queryStr, "&")
	for _, qry := range qryArr {
		temp := strings.Split(qry, "=")
		qryMap[temp[0]] = temp[1]
	}

	return qryMap
}
