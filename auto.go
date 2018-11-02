package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "net/http"
    "net/url"

)
func config() (string,string,string) {
	data, err := ioutil.ReadFile("config.txt")
    if err != nil {
        fmt.Println("File reading error", err)
    }
    //fmt.Println(string(data))
    config := string(data)
    a := strings.Split(config," ")[0]
    b := strings.Split(config," ")[1]
    cc := strings.Split(config," ")[2]
    var c string
    if cc == "1" {
    	c = "%E7%A7%BB%E5%8A%A8"
    } else if cc == "2"{
    	c = "%E7%94%B5%E4%BF%A1"
    } else {
    	c = "%E8%81%94%E9%80%9A"
    }
    return a,b,c
}
func tureurl() string {
    resp, _ := http.Get("http://10.10.9.4")
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

	t1 := string(body)
	t2 := strings.Split(t1,"?")[1]
	cc := strings.Split(t2,"'")[0]

	//fmt.Println(resp.Text(),cc)
	return cc
}

func httpPostForm(a,b,c,cc string) () {

    resp, _ := http.PostForm("http://10.10.9.4/eportal/InterFace.do?method=login",

        url.Values{"userId": {a}, "password": {b}, "service" : {c}, "queryString" : {cc}, "operatorPwd": {""}, "operatorUserId": {""}, "validcode": {""}, "passwordEncrypt": {"false"}})
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))

 

}

func main() {
	a,b,c := config()
	cc := tureurl()
    //fmt.Println(a,b,c,cc)
    //getnet(a,b,c,cc)
    httpPostForm(a,b,c,cc)


}