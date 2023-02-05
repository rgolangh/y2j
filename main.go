package main

import (
	"encoding/json"
	"fmt"
	yaml "gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	data := make(map[string]interface{})
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("failed reading stdin %v", err)
	}

	err = yaml.Unmarshal(input, &data)
	if err != nil {
		log.Fatalf("cannot unmarshal data,"+
			" yaml is not well formatted\ninput:%v\nerr:\n%v", string(input), err)
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("failed marshaling json %v", err)
	}
	fmt.Println(string(marshal))
}
