package main

import (
	"encoding/base64"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	f, err := os.Open("./pod.yaml")
	if err != nil {
		log.Fatalln(err)
		return
	}
	input, _ := ioutil.ReadAll(f)
	encodeString := base64.StdEncoding.EncodeToString(input)
	fmt.Println(encodeString)

	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		log.Fatalln(err)
		return
	}

	resultMap := make(map[string]interface{})
	if err := yaml.Unmarshal(decodeBytes, &resultMap); err != nil {
		// error handling
		log.Fatalln(err)

	}
	fmt.Println(resultMap)
}
