package json

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func Encode(data interface{}) *bytes.Buffer {
	b := new(bytes.Buffer)

	json.NewEncoder(b).Encode(data)

	return b
}

func Decode(data []byte, v interface{}) error {
	if data == nil {
		return fmt.Errorf("Malformed request")
	}

	if err := json.NewDecoder(bytes.NewReader(data)).Decode(v); err != nil {
		return err
	}

	return nil
}
