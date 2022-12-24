package csv

import (
	"encoding/csv"

	"github.com/wawakakakyakya/configloader/file"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func Load(path string, useSJIS bool) ([][]string, error) {
	var lines [][]string
	f, err := file.ReadByReader(path)
	if err != nil {
		return lines, err
	}
	var r *csv.Reader
	if useSJIS {
		r = csv.NewReader(transform.NewReader(f, japanese.ShiftJIS.NewDecoder()))
	} else {
		r = csv.NewReader(f)
	}

	lines, err = r.ReadAll()
	if err != nil {
		return lines, err
	}

	return lines, nil
}
