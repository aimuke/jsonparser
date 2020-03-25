package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("****************** start *******************")
	err := parser("data.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("******************  end  *******************")
}

func parser(p string) error {
	bs, err := ioutil.ReadFile(p)
	if err != nil {
		return err
	}

	var data interface{}
	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}

	bs, err = json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("target.json", bs, 0644)
	return err
}
