package main

import (
	"fmt"
	"github.com/yah01/cybuf-go"
)

func main() {
	cybufBytes := []byte(`
{
	Name: "cybuf"
	Age: 1
	Weight: 100.2
	School: {
	  Name: "Wuhan University"
	  Age: 120
	}
	Friends: [{
		Name: "Zerone"
		Phone: 01010101
	}
	{
		Name: "Acm"
		Phone: 2333
	}]
}
`)

	testMap := map[string]interface{}{}

	err := cybuf.Unmarshal(cybufBytes, &testMap)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", testMap)
}
