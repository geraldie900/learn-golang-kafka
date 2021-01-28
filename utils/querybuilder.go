package utils

import (
	"fmt"
	"strings"
)

// ForQuery 'for <alias> in <collectionName>' query
func ForQuery(variableName string, collectionName string) string {
	query := fmt.Sprintf("FOR %s IN %s", variableName, collectionName)
	return query
}

// FilterQuery 'FILTER <variableName>.<expression>'
func FilterQuery(variableName string, expression string, dataSearch string) string {
	query := fmt.Sprintf("FILTER %s.%s == %s", variableName, expression, dataSearch)
	return query
}

// ReturnQuery 'RETURN <variableName>' OR 'RETURN {<dataReturn> : <variableName>.<dataReturn>}'
func ReturnQuery(variableName string, dataReturn ...string) string {
	query := fmt.Sprintf("RETURN %s", variableName)
	if len(dataReturn) != 0 {
		var dataReturnMerge []string
		for i, data := range dataReturn {
			dataReturnMerge[i] = fmt.Sprintf("%s : %s.%s", data, variableName, data)
		}
		query = fmt.Sprintf("RETURN { %s }", strings.Join(dataReturnMerge, ", "))
	}
	return query
}
