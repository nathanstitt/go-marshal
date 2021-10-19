package marshal

import (
	"encoding/json"
)

func MarshalJSONB(mp *map[string]interface{}) ([]byte, error) {
	bt, err := json.Marshal(mp)
	if err != nil {
		return []byte("null"), err
	}
	return bt, nil
}

func UnmarshalJSONB(b []byte, mp *map[string]interface{}) error {
	err := json.Unmarshal(b, mp)
	return err
}
