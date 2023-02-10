package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// json to map
func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"P001","name":"MacBook Pro 01","price":20000000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes,&result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])

}

// map to json
func TestMapEncode(t *testing.T)  {
	product := map[string]interface{}{
		"id":"P001",
		"name":"MacBook Pro 01",
		"price":20000000,
	}
	bytes,_ := json.Marshal(product)

	fmt.Println(string(bytes))
}