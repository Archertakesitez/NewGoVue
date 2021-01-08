package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type User struct {
	Username string
	Count    int
}
type NumOne struct{
	Num1 float64 `json:"num1"`
}
func test(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
		fmt.Println(decoder)
		var(
			numsData NumOne
		)
		decoder.Decode(&numsData)
		fmt.Println(numsData)
		numsData.Num1++
		
		data, _ := json.Marshal(numsData.Num1)
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers","*")
		w.Write(data)
}
func main() {
	//user := User{"abc", 123}
	http.HandleFunc("/", test)
	http.HandleFunc("/signup", test)
	http.ListenAndServe(":8080", nil)
}
