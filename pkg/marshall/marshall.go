package marshall

import (
	"encoding/json"
)

/*
Marshall - set's the same json fields of obj 
*/
func Marshall(from interface{}, to interface{}) error {
	data, err := json.Marshal(from)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, to)
}