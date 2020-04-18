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

	marshalMap = map[string]interface{}{
		"Name": "yah01",
		"Age":  21,
		"Live": true,
		"School": map[string]interface{}{
			"Name": "Wuhan University",
			"Age":  120,
		},
		"Wallet": []float64{1.0, 10.0, 100.0},
	}

	marshalBytes []byte
	err          error
)

func TestCyBufMarshal(t *testing.T) {

	marshalBytes, err = Marshal(marshalMap)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("\n" + string(marshalBytes))
}

func TestCyBufUnmarshal(t *testing.T) {
	unmarshalMap := map[string]interface{}{}

	err = Unmarshal(cybufBytes, &unmarshalMap)
	if err != nil {
		t.Error(err)
	}
	t.Log(unmarshalMap)
}

func BenchmarkCyBufMarshal(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		marshalBytes, _ = Marshal(marshalMap)
	}
}

func BenchmarkJsonMarshal(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		marshalBytes, _ = json.Marshal(marshalMap)
	}
}

func BenchmarkJsonMarshalIndent(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		marshalBytes, _ = json.MarshalIndent(marshalMap, "", "\t")
	}
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
