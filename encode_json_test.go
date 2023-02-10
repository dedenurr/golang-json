package golang_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

//proses konversi tipe data di golang kedalam bentuk json(encode)
func logJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

//test konversi tipe data
func TestEncode(t *testing.T) {
	logJson("eko")
	logJson(1)
	logJson(true)
	logJson([]string{"Eko","Kurniawan","Khannedy"})
}

/* Noted
encode => proses perubahan dari data golang ke json 
gunakan Marshal untuk encode
*/