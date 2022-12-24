package yml

import (
	"github.com/wawakakakyakya/configloader/file"
	yaml "gopkg.in/yaml.v2"
)

func Load(path string, yc interface{}) error {
	f, err := file.ReadAll(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(f, yc)
}
