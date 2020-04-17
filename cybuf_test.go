package cybuf

import (
	"encoding/json"
	"testing"
)

var (
	bytes = []byte(`
{
	Name: "cybuf"
	Age: 1
	Weight: 100.2
	Fuck: {
	Name: "you"
	Age: 99
	}
	Emails: [{
		Prefix: "kagaminehuan"
		Name: "live.com"
	}
	{
		Prefix: "cenyang"
		Name: "bytedance.com"
	}]
}
`)
	jsonBytes = []byte(`
{
	"Name": "cybuf",
	"Age": 1,
	"Weight": 100.2,
	"Fuck": {
		"Name": "you",
		"Age": 99
	},
	"Emails": [
		{
			"Prefix": "kagaminehuan",
			"Name": "live.com"
		},
		{
			"Prefix": "cenyang",
			"Name": "bytedance.com"
		}
	]
}
`)
)


func BenchmarkUnmarshal(b *testing.B) {
	b.ResetTimer()

	testMap := make(map[string]interface{})
	for i := 0; i < b.N; i++ {
		Unmarshal(bytes,&testMap)
	}
}

func BenchmarkJsonUnmarshal(b *testing.B) {
	b.ResetTimer()

	testMap := make(map[string]interface{})
	for i := 0; i < b.N; i++ {
		json.Unmarshal(jsonBytes,&testMap)
	}
}
