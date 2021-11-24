package main

import (
	"fmt"
	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/httptool"
	//"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/encrypt" 
	//"net/http"
)
/*type Request struct {
	Content string `json:"content"`
	ExtraInfo string `json:"extra_info"`
   } //请求游戏服务端的底层格式
   type HttpRequest struct {
	Req *http.Request
	Body *Request // 上述的 request 格式
	BodyType int
   } // ⽤户的请求类型, ⽤户⽤这个结构体就可以了
   type TextInfo struct {
	Text string `json:"text"`
	ExtraInfo string `json:"extra_info"`
   } // 服务端响应的底层格式
   type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data TextInfo `json:"data"`
   } // 响应内容
   type HttpResponse struct {
	Body Response `json:"body"`
	Raw string
	raw *http.Response
   } // ⽤户的响应类型，⽤户⽤这个结构体就可以了*/


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

	fmt.Println(req)
	//req.AddHeader("Code","123")
	resp, err := req.SendRequest()
	// 调⽤例 (request是*HttpRequest类型变量)
	resp.ShowHeader()
	//fmt.Println()
	// write your code below
	// ...
	value, err := resp.GetHeader("Passport")
	req.AddHeader("Passport",value[0])
    resp1, err := req.SendRequest()
	resp1.ShowBody()
}
