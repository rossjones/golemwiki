package main

import (
	"encoding/json"
	"io/ioutil"
)

type JsonConfiguration struct {
	Server   map[string]string
	Wiki     map[string]string
	Database map[string]string
}

func (l *JsonConfiguration) LoadFrom(path string) (err error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &l)
	if err != nil {
		return err
	}

	return nil
}
