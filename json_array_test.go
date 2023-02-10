package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Ade",
		MiddleName: "Hikmah",
		LastName:   "Patimah",
		Hobbies:    []string{"Traveling", "Shoping", "Eating"},
	}

	bytes,_ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T)  {
	jsonString := `{"FirstName":"Ade","MiddleName":"Hikmah","LastName":"Patimah","Age":0,"Married":false,"Hobbies":["Traveling","Shoping","Eating"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}

	err := json.Unmarshal(jsonBytes,customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)

}


func TestJSONArrayComplex(t *testing.T)  {
	customer := Customer{
		FirstName: "Dede",
		Addresses: []Address{
			{
				Street: "Jl Cinta",
				Country: "Zimbabue",
				PostalCode: "xxx",
			},
			{
				Street: "Jl Cinta 2",
				Country: "Zimbabue 2",
				PostalCode: "xxxzz",
			},
			
		},
	}
		bytes, err :=	json.Marshal(customer)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T)  {
	jsonString := `{"FirstName":"Dede","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jl Cinta","Country":"Zimbabue","PostalCode":"xxx"},{"Street":"Jl Cinta 2","Country":"Zimbabue 2","PostalCode":"xxxzz"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes,customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}	
func TestOnlyJSONArrayComplexDecode(t *testing.T)  {
	jsonString := `[{"Street":"Jl Cinta","Country":"Zimbabue","PostalCode":"xxx"},{"Street":"Jl Cinta 2","Country":"Zimbabue 2","PostalCode":"xxxzz"}]`
	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes,addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)

}	


func TestOnlyJSONArrayComplex(t *testing.T)  {
		addresses := []Address{
			{
				Street: "Jl Cinta",
				Country: "Zimbabue",
				PostalCode: "xxx",
			},
			{
				Street: "Jl Cinta 2",
				Country: "Zimbabue 2",
				PostalCode: "xxxzz",
			},
		}

		bytes, err :=	json.Marshal(addresses)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bytes))
}
