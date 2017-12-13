package pgup

import (
	"encoding/json"

	"upper.io/db.v3"
)

type cond map[string]interface{}

// ParseQuery parses given string
// and returns parsed query and an error if any raised.
func ParseQuery(str string) (db.Cond, error) {
	var qry cond
	err := parseString(str, &qry)
	if err != nil || qry == nil {
		return nil, err
	}

	result := db.Cond{}
	for key, value := range qry {
		result[key] = value
	}
	return result, nil
}

// ParseOrder parses given string
// and returns parsed slice of order rules and an error if any raised.
func ParseOrder(str string) ([]interface{}, error) {
	var result []interface{}
	err := parseString(str, &result)
	return result, err
}

func parseString(str string, dest interface{}) error {
	if str == "" {
		return nil
	}
	return json.Unmarshal([]byte(str), dest)
}
