package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	passport := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoicWlhbnJlbiIsImlhdCI6MTYzODI4MDQ4NSwibmJmIjoxNjM4MjgwNDg1fQ.Q4kTf6FrWVNjMrDe2eV_JdC5siD7KJtMQgzb4u4jM9o"
	req, _ := http.NewRequest("GET", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/secret_key", nil)
	req.Header.Add("Code", "200")
	req.Header.Add("Passport", passport)
	res, _ := client.Do(req)
	//fmt.Println(res.Header)
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}
