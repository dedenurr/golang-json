package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)


func TestDecodeJSON(t *testing.T) {
	jsonString := `{"FirstName":"Zahra","MiddleName":"Assyifa","LastName":"Nurrahman","Age":6,"Married":false}`
	jsonBytes := []byte(jsonString)


	// tampung data dalam customer pointer
	customer := &Customer{}

	// proses decode dari json ke data golang
	err := json.Unmarshal(jsonBytes, &customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
	fmt.Println(customer.Married)
	


}

/* Noted
Decode => proses perubahan dari json ke data golang
gunakan Unmarshal untuk decode
*/