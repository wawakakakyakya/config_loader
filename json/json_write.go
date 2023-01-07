package json

import (
	"encoding/json"

	"github.com/wawakakakyakya/configloader/file"
)

func Write(path string, jc interface{}) error {

	content, err := json.MarshalIndent(jc, "", " ")
	if err != nil {
		return err
	}
	return file.Write(path, content)
}
