package app

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type YamlReader struct {
}

func (reader *YamlReader) Read(filename *string) map[string]interface{} {
	var b, err = ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatalf("read file error: %v", err)
		return nil
	}
	return convertByteToMap(b)
}

func convertByteToMap(data []byte) map[string]interface{} {
	m := make(map[string]interface{})

	err := yaml.Unmarshal(data, &m)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil
	}

	return m
}
