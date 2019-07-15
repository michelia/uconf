package uconf

import (
	"io/ioutil"
	"log"

	"github.com/ghodss/yaml"
)

func Parse(conf interface{}, fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal("ReadFile error: ", err)
	}
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		log.Fatal("yaml Unmarshal error: ", err)
	}
}
