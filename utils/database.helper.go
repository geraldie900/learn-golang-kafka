package utils

import (
	"fmt"
	"golang_kafka/config/database"
	"golang_kafka/config/globals"

	"github.com/arangodb/go-driver"
)

// Collection struct for define collection AQL
type Collection struct {
	CollectionVariable string
	CollectionName     string
	FilterData         []string
	ReturnData         []string
}

// ConnectCollection connect to Collection
func ConnectCollection(collectionName string) (driver.Collection, error) {
	col, err := (database.DB).Collection(globals.Ctx, collectionName)
	return col, err
}

// PasswordChecker validate password by input is same as in database or not
func PasswordChecker(key string, passwordInput string) bool {
	query := fmt.Sprintf("FOR u IN users FILTER u._key == \"%s\" RETURN u.password", key)
	cursor, err := database.DB.Query(globals.Ctx, query, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer cursor.Close()

	var password interface{}
	cursor.ReadDocument(globals.Ctx, &password)
	if fmt.Sprintf("%v", password) == passwordInput {
		return true
	}
	return false
}

// CursorCollection get cursor on result query
func CursorCollection(variableName string, collectionName string) (driver.Cursor, error) {
	query := fmt.Sprintf("%s %s", ForQuery(variableName, collectionName), ReturnQuery(variableName))
	// query := "FOR data IN @@collection RETURN data"
	// bindVars := map[string]interface{}{
	// 	"@collection": collectionName,
	// }
	// return database.DB.Query(globals.Ctx, query, bindVars)
	return database.DB.Query(globals.Ctx, query, nil)
}

// CurrentVersionAudit get current version of audit data
func CurrentVersionAudit(collectionName string, dataKey string) int {
	query := "FOR data IN @@collection FILTER data._key == @datakey RETURN data.audit_trail.current_version"
	bindVars := map[string]interface{}{
		"@collection": collectionName,
		"datakey":     dataKey,
	}
	cursor, err := database.DB.Query(globals.Ctx, query, bindVars)
	if err != nil {
		// err
	}
	defer cursor.Close()

	var result interface{}
	cursor.ReadDocument(globals.Ctx, &result)
	if result == nil {
		return 0
	}

	return int(result.(float64))
}
