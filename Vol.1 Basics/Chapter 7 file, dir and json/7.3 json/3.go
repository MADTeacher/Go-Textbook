package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Student struct {
	Name        string
	Age         uint8
	Course      uint8
	Single      bool
	Description []string
}

func main() {
	student := &Student{}

	data, err := ioutil.ReadFile("student.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, student)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", *student)
}
