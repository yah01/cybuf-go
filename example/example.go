package main

import (
	"fmt"
	"github.com/yah01/cybuf-go"
)

func main() {
	cybufBytes := []byte(`{
Name : "cybuf"
Age : 1
Weight : 100.2
Fuck : {
Name : "you"
Age : 99
}
Emails: [{
	Prefix: "kagaminehuan"
	Name: "live.com"
}{
	Prefix: "cenyang"
	Name: "bytedance.com"
}]}`)

	testMap := map[string]interface{}{}

	err := cybuf.Unmarshal(cybufBytes, &testMap)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", testMap)
}
