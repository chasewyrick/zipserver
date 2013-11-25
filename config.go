package zip_server

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var DefaultConfigFname = "zip_server.json"

type Config struct {
	PrivateKeyPath string
	ClientEmail string
	Bucket string
	ExtractPrefix string
	MaxFileSize int
	MaxNumFiles int
}

var defaultConfig = Config{
	MaxFileSize: 1024 * 1024,
	MaxNumFiles: 100,
	ExtractPrefix: "html5",
}

func LoadConfig(fname string) *Config {
	jsonBlob, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatal(err)
	}

	config := defaultConfig
	err = json.Unmarshal(jsonBlob, &config)

	if err != nil {
		log.Fatal("Failed parsing config: " + fname + ": " + err.Error())
	}

	return &config
}


func (c *Config) Dump() string {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}