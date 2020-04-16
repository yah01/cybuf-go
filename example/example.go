package main

import (
	"fmt"
	"github.com/yah01/cybuf-go"
)

func main() {
	cybufBytes := []byte(`Name : "cybuf"
Age : 1
Weight : 100.2`)

	testMap := map[string]interface{}{}

	err := cybuf.Unmarshal(cybufBytes, &testMap)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", testMap)
}
