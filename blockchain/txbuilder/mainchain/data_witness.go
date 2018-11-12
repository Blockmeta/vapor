package mainchain

import (
	"encoding/json"

	chainjson "github.com/bytom/encoding/json"
)

// DataWitness used sign transaction
type DataWitness chainjson.HexBytes

func (dw DataWitness) Materialize(args *[][]byte) error {
	*args = append(*args, dw)
	return nil
}

// MarshalJSON marshal DataWitness
func (dw DataWitness) MarshalJSON() ([]byte, error) {
	x := struct {
		Type  string             `json:"type"`
		Value chainjson.HexBytes `json:"value"`
	}{
		Type:  "data",
		Value: chainjson.HexBytes(dw),
	}
	return json.Marshal(x)
}
