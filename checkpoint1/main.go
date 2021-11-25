package main

import (
	"fmt"
	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/httptool"
)

func main() {
	req, err := httptool.NewRequest(
		httptool.GETMETHOD,
		"http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/code",
		"",
		httptool.DEFAULT, // 这里可能不是 DEFAULT，自己去翻阅文档
	)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(req)
	//resp, err := req.SendRequest()
	//resp.ShowHeader()
	//value, err := resp.GetHeader("Passport")
	//req.AddHeader("Passport",value[0])
    resp1, err := req.SendRequest()
	resp1.ShowBody()
}
