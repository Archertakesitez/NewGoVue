package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	//"fmt"
	"log"
	//"net/http"
	//"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	//"github.com/gin-gonic/gin"
)

var engine *xorm.Engine

//var db *sql.DB
var err error

type forLogin struct {
	Message string `json:"message"`
	Count   int    `json:"count"`
}

//PassUser ..
type PassUser struct {
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
}

// User dd
type User struct {
	UserID   int
	Username string
	Salt     string
	Count    int
	Pwd      string
}

//init database..
func initDB() {
	engine, err = xorm.NewEngine("mysql", "erchizhang:123456@tcp(127.0.0.1)/Trail1")
	//db, err = sql.Open("mysql", "erchizhang:123456@tcp(127.0.0.1)/Trail1")
	if err != nil {
		log.Fatal(err)
		return
	}
	err = engine.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = engine.Sync2(new(User))
	if err != nil {
		return
	}
	//engine.CreateTables(User{})

}
func login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	fmt.Println(decoder)
	var (
		user PassUser
	)
	decoder.Decode(&user)
	msg := saveLogIn(user.Username, user.Pwd)
	found := queryName(user.Username)
	loginPur := forLogin{
		Message: msg,
		Count:   found,
	}
	data, _ := json.Marshal(loginPur)

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Write(data)
	fmt.Println(user)
}
func saveUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	//fmt.Println(decoder)
	if err != nil {
		fmt.Println(err)
	}
	var (
		user PassUser
	)
	decoder.Decode(&user)
	fmt.Println(user)
	message := saveIn(user.Username, user.Pwd)
	data, _ := json.Marshal(message)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
	w.Write(data)
	fmt.Println(user)
}
func verifyPwd(name string, pwd string) bool {
	verified := false
	//var getPwd string
	user := &User{}
	//sqlStr := "select pwd from user where username=?;"
	//err = db.QueryRow(sqlStr, name).Scan(&getPwd)
	has, _ := engine.Where("username = ?", name).Get(user)
	if has {
		fmt.Printf("user:%v\n", user)
	}
	/*if err != nil {
		fmt.Println("verifyPwdError")
	} else {
		if getPwd == pwd {
			verified = true
		} else {
			verified = false
		}
	}*/
	if user.Pwd == pwd {
		verified = true
	} else {
		verified = false
	}

	return verified
}
func saveLogIn(name string, pwd string) string {
	found := queryName(name)
	verified := verifyPwd(name, pwd)
	str := ""
	//user := &User{}
	//engine.Where("username = ?", name).Get(&user)
	if (name == "") || (pwd == "") {
		str = "用户名或密码不能为空"
	} else {
		if found == -1 {
			str = "该用户名不存在"
		} else {
			if verified == true {
				str = "登录成功！跳转中..."
				user := &User{}
				user.Count = found + 1
				affected, err := engine.Update(user, &User{Username: name})
				if err != nil {
					fmt.Println(affected)
				}
				//db.Exec("update User set count=? where username=?", found+1, name)
			} else {
				str = "密码不正确"
			}
		}
	}
	return str
}
func queryName(name string) int {
	//var timeCount int
	//sqlStr := "select count from user where username=?;"
	//err = db.QueryRow(sqlStr, name).Scan(&timeCount)
	user := &User{}
	has, err := engine.Where("username=?", name).Get(user)
	switch {
	case has != true:
		return -1
	case err != nil:
		return -10
	default:
		return user.Count
	}
	/*switch {
	case err == sql.ErrNoRows:
		return -1
	case err != nil:
		return -10
	default:
		return timeCount
	}*/
}
func saveIn(name string, pwd string) string {
	str := ""
	timesR := queryName(name)
	if timesR == -1 { //no same name found
		if (name == "") || (pwd == "") {
			str = "username and password cannot be empty"
		} else if strings.ContainsAny(name, "%<>/\\") || strings.ContainsAny(pwd, "%<>/\\") {
			str = "username and password caannot contain %<>/\\"
		} else if (len(name) > 20) || (len(pwd) > 20) {
			str = "username and password cannot exceed 20 characters long"
		} else {
			user := &User{}
			user.Username = name
			user.Pwd = pwd
			user.Count = 1

			affected, _ := engine.Insert(user)
			fmt.Printf("%d records inserted, user.id:%d\n", affected, user.UserID)
			//insert, err := db.Exec("insert into User(username, pwd, count) values(?, ?, 1)", name, pwd)

			/*if err != nil {
				fmt.Println(insert)
			}*/
			str = "注册成功！"
		}

	} else {
		str = "该用户名已被占用"
	}
	return str
}

func main() {
	//var err error
	//engine, err = xorm.NewEngine("mysql", "root:123@/test?charset=utf8")

	//user := User{"abc", 123}
	initDB()
	http.HandleFunc("/", login)
	http.HandleFunc("/signup", saveUser)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	//db.Close()
}
