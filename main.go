package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"database/sql"
	//"fmt"
	"log"
	//"net/http"
	//"strconv"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)
var db *sql.DB
var err error 
type User struct {
	Username string //`json:"Username"`
	pwd      string //`json:"pwd"`
}
type NumOne struct{
	Num1 float64 `json:"num1"`
}
//init database..
func initDB(){
	db, err = sql.Open("mysql", "erchizhang:123456@tcp(127.0.0.1)/Trail1")
	if err != nil {
		fmt.Println("righthere")
	}
}
func login(w http.ResponseWriter, r *http.Request){
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
func saveUser(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var(
		user User
	)
	decoder.Decode(&user)
	message := saveIn(user.Username, user.pwd)
	data,_:=json.Marshal(message)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers","*")
	w.Write(data)
	fmt.Println(user)
}

func saveIn(name string, pwd string) string {
	if (name=="")||(pwd==""){
		return "username and password cannot be empty"
	}else if (strings.ContainsAny(name,"%<>/\\")||strings.ContainsAny(pwd,"%<>/\\")){
		return "username and password caannot contain %<>/\\"
	}else if (len(name)>20)||(len(pwd)>20){
		return "username and password cannot exceed 20 characters long"
	}else{
		insert, err := db.Exec("insert into User(username, pwd, count) values(?, ?, 1)",name, pwd)
		//insert, err:= db.Exec("insert into finalTable(username, times) values("+"hehe"+",1)")
		if err != nil {
			fmt.Println(
				//insert
				insert, err)
		}
		return "注册成功！"
	}

}

func main() {
	//user := User{"abc", 123}
	initDB()
	http.HandleFunc("/", login)
	http.HandleFunc("/signup", saveUser)
	err:=http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	db.Close()
}
