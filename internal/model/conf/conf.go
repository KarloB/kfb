package conf

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

// Configuration app configuration
type Configuration struct {
	Key     string `json:"key"`     // secret key
	DB      string `json:"db"`      // database ID
	Timeout int    `json:"timeout"` // timeout in seconds
}

// Get get configuration
func Get(path string) (*Configuration, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buff := bytes.NewBuffer(nil)
	_, err = io.Copy(buff, f)
	if err != nil {
		return nil, err
	}

	conf := &Configuration{}
	err = json.Unmarshal(buff.Bytes(), conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
