package json

import (
	"encoding/json"
)

func MustMarshal(v interface{}) []byte {
	rv, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return rv
}
