package main

import (
	"net/http"
	"fmt"
	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/encrypt" // 加密解密模块
	"encoding/json"
	"strings"
	"io/ioutil"
)

func main() {
	client := &http.Client{}
	passport := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoicWlhbnJlbiIsImlhdCI6MTYzODI4MDQ4NSwibmJmIjoxNjM4MjgwNDg1fQ.Q4kTf6FrWVNjMrDe2eV_JdC5siD7KJtMQgzb4u4jM9o"
	infor := "c2VjcmV0X2tleTpNdXhpU3R1ZGlvMjAzMzA0LCBlcnJvcl9jb2RlOmZvciB7Z28gZnVuYygpe3RpbWUuU2xlZXAoMSp0aW1lLkhvdXIpfSgpfQ=="
	errorCode, err := encrypt.Base64Decode(infor) //base64解密
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(errorCode) //得到secret_key:MuxiStudio203304, error_code:for {go func(){time.Sleep(1*time.Hour)}()}
	secret_key := "MuxiStudio203304"
	error_code := "for {go func(){time.Sleep(1*time.Hour)}()}"

	//得出加密过的error_code
	data, err := encrypt.AESEncryptOutInBase64([]byte(error_code), []byte(secret_key))//用secret_key加密error_code
	
	type Body struct {
		Content string
	}
	body := Body{
		Content: string(data),
	}
	by, _ := json.Marshal((body))
	payload := strings.NewReader(string(by))
	req, _ := http.NewRequest("PUT", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate", payload)
	req.Header.Add("Code", "120")
	req.Header.Add("Passport", passport)
	res, _ := client.Do(req)
	//fmt.Println(res.Header)
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}
