package file

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
)

// func Read(path string) ([]byte, error) {
// 	// var buf bytes.Buffer
// 	return ioutil.ReadFile(path)
// }

func ReadAll(path string) ([]byte, error) {
	var buf bytes.Buffer

	err := open(path, &buf)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(&buf)
}

func ReadByReader(path string) (io.Reader, error) {
	var buf bytes.Buffer

	err := open(path, &buf)
	if err != nil {
		return &buf, err
	}

	return &buf, nil
}

func open(path string, buf *bytes.Buffer) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = io.Copy(buf, file)
	if err != nil {
		return err
	}

	return err
}
