package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct{
	Id string `json:"id"`
	Name string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id: "P001",
		Name: "MacBook Pro 01",
		ImageUrl: "http://example.com/image.png",
	}
	bytes,_ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONDecode(t *testing.T)  {
	jsonString := `{"id":"P001","name":"MacBook Pro 01","image_url":"http://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	json.Unmarshal(jsonBytes,product)
	fmt.Println(product)

}