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

func init() {
	MarshalSep = ' '
}

func TestCyBufMarshal(t *testing.T) {
	marshalBytes, err = Marshal(marshalMap)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("\n" + string(marshalBytes) + "\n")
	t.Log("CyBuf Marshal length:", len(marshalBytes), "\n")
	marshalBytes, _ = json.Marshal(marshalMap)
	t.Log("JSON Marshal length:", len(marshalBytes), "\n")
}

func TestCyBufMarshalIndent(t *testing.T) {

	marshalBytes, err = MarshalIndent(marshalMap)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("\n" + string(marshalBytes) + "\n")
	t.Log("CyBuf MarshalIndent length:", len(marshalBytes), "\n")
	marshalBytes, _ = json.MarshalIndent(marshalMap, "", "\t")
	t.Log("JSON MarshalIndent length:", len(marshalBytes), "\n")
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
		marshalBytes, err = Marshal(marshalMap)
		if err != nil {
			b.Fatal(err)
			return
		}
	}
}

func BenchmarkCyBufMarshalIndent(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		marshalBytes, err = MarshalIndent(marshalMap)
		if err != nil {
			b.Fatal(err)
			return
		}
	}
}

func BenchmarkJsonMarshal(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		marshalBytes, err = json.Marshal(marshalMap)
		if err != nil {
			b.Fatal(err)
			return
		}
	}
}

func BenchmarkJsonMarshalIndent(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		marshalBytes, err = json.MarshalIndent(marshalMap, "", "\t")
		if err != nil {
			b.Fatal(err)
			return
		}
	}
}

func BenchmarkCyBufUnmarshal(b *testing.B) {
	b.ResetTimer()

	testMap := make(map[string]interface{})
	for i := 0; i < b.N; i++ {
		err = Unmarshal(cybufBytes, &testMap)
		if err != nil {
			b.Fatal(err)
			return
		}
	}
}

func BenchmarkJsonUnmarshal(b *testing.B) {
	b.ResetTimer()

	testMap := make(map[string]interface{})
	for i := 0; i < b.N; i++ {
		err = json.Unmarshal(jsonBytes, &testMap)
		if err != nil {
			b.Fatal(err)
			return
		}
	}
}
