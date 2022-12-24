package json

import (
	"encoding/json"

	"github.com/wawakakakyakya/configloader/file"
)

func Load(path string, jc interface{}) error {
	raw, err := file.ReadAll(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(raw, &jc)
}
