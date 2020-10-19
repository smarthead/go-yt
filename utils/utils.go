package utils

import "strings"

func ConstructQuery(query string, fields []string) map[string]string {
	queryValues := make(map[string]string)

	if len(query) > 0 {
		queryValues["query"] = query
	}

	if fields != nil && len(fields) > 0 {
		queryValues["fields"] = strings.Join(fields, ",")
	}

	return queryValues
}
