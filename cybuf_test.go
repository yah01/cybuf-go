package cybuf

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type People struct {
	Name    string
	Age     int
	Weight  float64
	Live    bool
	Friends []People
	School  School
}

type School struct {
	Name string
	Age  int
}

var (
	people = People{
		Name:   "yah01",
		Age:    21,
		Weight: 100.2,
		Live:   true,
		Friends: []People{
			{
				Name:   "wmx",
				Age:    100,
				Weight: 200.5,
				Live:   false,
				Friends: []People{
					{
						Name:    "bytedance",
						Age:     8,
						Weight:  114514.5,
						Live:    true,
						Friends: nil,
						School: School{
							Name: "BDU",
							Age:  114514,
						},
					},
				},
				School: School{
					Name: "SHU",
					Age:  114514,
				},
			},
			{
				Name:    "bytedance",
				Age:     8,
				Weight:  114514.5,
				Live:    true,
				Friends: nil,
				School: School{
					Name: "BDU",
					Age:  114514,
				},
			},
		},
		School: School{
			Name: "Wuhan University",
			Age:  120,
		},
	}

	peopleMap        map[string]interface{}
	cybufBytes       []byte
	cybufIndentBytes []byte
	jsonBytes        []byte
	jsonIndentBytes  []byte
	err              error
)

func init() {
	if cybufBytes, err = Marshal(people); err != nil {
		panic(err)
	}
	if cybufIndentBytes, err = MarshalIndent(people); err != nil {
		panic(err)
	}
	if jsonBytes, err = json.Marshal(people); err != nil {
		panic(err)
	}
	if jsonIndentBytes, err = json.MarshalIndent(people, "", "\t"); err != nil {
		panic(err)
	}
	peopleMap = make(map[string]interface{})
	err = Unmarshal(cybufBytes, &peopleMap)
	if err != nil {
		panic(err)
	}

	fmt.Println(people, peopleMap, string(cybufBytes), string(cybufIndentBytes), string(jsonBytes), string(jsonIndentBytes))
}

func TestCyBufMarshalMap(t *testing.T) {
	_, err := Marshal(peopleMap)
	if err != nil {
		t.Error(err)
		return
	}
	//if !bytes.Equal(tBytes, cybufBytes) {
	//	t.Error(string(tBytes),string(cybufBytes))
	//}
}

func TestCyBufMarshalStruct(t *testing.T) {
	_, err := Marshal(people)
	if err != nil {
		t.Error(err)
		return
	}
	//if !bytes.Equal(_, cybufBytes) {
	//	t.Error(_)
	//}
}

func TestCyBufMarshalIndentMap(t *testing.T) {
	_, err := MarshalIndent(peopleMap)
	if err != nil {
		t.Error(err)
		return
	}
	//if !bytes.Equal(_, cybufIndentBytes) {
	//	t.Error(_)
	//}
}

func TestCyBufMarshalStructIndent(t *testing.T) {
	_, err := Marshal(people)
	if err != nil {
		t.Error(err)
		return
	}
	//if !bytes.Equal(_, cybufIndentBytes) {
	//	t.Error(_)
	//}
}

func TestCyBufUnmarshalMap(t *testing.T) {
	res := map[string]interface{}{}

	err = Unmarshal(cybufBytes, &res)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(res, peopleMap) {
		t.Error(res)
	}
}

func TestCyBufUnmarshalStruct(t *testing.T) {
	res := People{}

	err = Unmarshal(cybufBytes, &res)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(res, people) {
		t.Error(res, people)
	}
}

func BenchmarkCyBufMarshalMap(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := Marshal(peopleMap)
		if err != nil {
			b.Error(res, err)
			return
		}
	}
}

func BenchmarkJsonMarshalMap(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := json.Marshal(peopleMap)
		if err != nil {
			b.Error(res, err)
			return
		}
	}
}

func BenchmarkCyBufMarshalIndentMap(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := MarshalIndent(peopleMap)
		if err != nil {
			b.Error(res, err)
			return
		}
	}
}

func BenchmarkJsonMarshalIndentMap(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res, err := json.MarshalIndent(peopleMap, "", "\t")
		if err != nil {
			b.Error(res, err)
			return
		}
	}
}

func BenchmarkCyBufUnmarshalMap(b *testing.B) {
	res := make(map[string]interface{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = Unmarshal(cybufBytes, &res)
		if err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkJsonUnmarshalMap(b *testing.B) {
	res := make(map[string]interface{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = json.Unmarshal(jsonBytes, &res)
		if err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkCyBufUnmarshalStruct(b *testing.B) {
	people := People{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = Unmarshal(cybufBytes, &people)
		if err != nil {
			b.Error(err)
			return
		}
	}
}

func BenchmarkJsonUnmarshalStruct(b *testing.B) {
	people := People{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = json.Unmarshal(jsonBytes, &people)
		if err != nil {
			b.Error(err)
			return
		}
	}
}
