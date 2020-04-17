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
	Friends: [
		{
			Name: "Zerone"
			Phone: 01010101
		}
		{
			Name: "Acm"
			Phone: 2333
		}
	]
}
`)
	unmarshalMap := map[string]interface{}{}

	err := cybuf.Unmarshal(cybufBytes, &unmarshalMap)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", unmarshalMap)

	marshalMap := map[string]interface{}{
		"First":  1,
		"Second": 2,
		"Half":   0.5,
	}

	cybufBytes, err = cybuf.Marshal(marshalMap)
	fmt.Println(string(cybufBytes))
}
