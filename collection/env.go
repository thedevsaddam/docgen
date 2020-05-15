package collection

import (
	"encoding/json"
	"io"
)

//Environment represents environment variables
type Environment struct {
	ID                   json.RawMessage `json:"id"`
	Name                 string          `json:"name"`
	Values               []Field         `json:"values"`
	PostmanVariableScope json.RawMessage `json:"_postman_variable_scope"`
	PostmanExportedAt    json.RawMessage `json:"_postman_exported_at"`
	PostmanExportedUsing json.RawMessage `json:"_postman_exported_using"`
}

// Open open an environment file
func (e *Environment) Open(rdr io.Reader) error {
	dcr := json.NewDecoder(rdr)
	if err := dcr.Decode(&e); err != nil {
		return err
	}
	for i := len(e.Values) - 1; i >= 0; i-- {
		if !e.Values[i].Enabled {
			e.Values = append(e.Values[:i], e.Values[i+1:]...)
		}
	}
	return nil
}

func (e *Environment) upsertVariable(f Field) {
	var update bool
	for i, v := range e.Values {
		if v.Key == f.Key {
			e.Values[i] = f
			update = true
		}
	}
	if !update {
		e.Values = append(e.Values, f)
	}
}

// SetCollectionVariables set collection variables to the env // if env provided then the values will be replaced
func (e *Environment) SetCollectionVariables(vv []Field) {
	for _, v := range vv {
		if !v.Disabled {
			e.upsertVariable(v)
		}
	}
}
