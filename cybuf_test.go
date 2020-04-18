package cybuf

import (
	"encoding/json"
	"testing"
)

var (
	cybufBytes = []byte(`
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
			Phone: 1010101
		}
		{
			Name: "Acm"
			Phone: 2333
		}
	]
}
`)
	jsonBytes = []byte(`
{
	"Name": "cybuf",
	"Age": 1,
	"Weight": 100.2,
	"School": {
		"Name": "Wuhan University",
		"Age": 120
	},
	"Friends": [
		{
			"Name": "Zerone",
			"Phone": 1010101
		},
		{
			"Name": "Acm",
			"Phone": 2333
		}
	]
}
`)
)

func TestCyBufUnmarshal(t *testing.T) {
	unmarshalMap := map[string]interface{}{}

	err := Unmarshal(cybufBytes, &unmarshalMap)
	if err != nil {
		t.Error(err)
	}
	t.Log(unmarshalMap)
}

func BenchmarkCyBufUnmarshal(b *testing.B) {
	b.ResetTimer()

	testMap := make(map[string]interface{})
	for i := 0; i < b.N; i++ {
		Unmarshal(cybufBytes, &testMap)
	}
}

func BenchmarkJsonUnmarshal(b *testing.B) {
	b.ResetTimer()

	testMap := make(map[string]interface{})
	for i := 0; i < b.N; i++ {
		json.Unmarshal(jsonBytes, &testMap)
	}
}
