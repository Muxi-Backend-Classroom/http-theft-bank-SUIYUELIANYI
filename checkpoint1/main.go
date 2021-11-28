package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
)

func main(){
	url :="http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/code"
	method :="GET"
	client :=&http.Client{}
	request,err :=http.NewRequest(method,url,nil)
	if err !=nil{
		fmt.Println(err)
		return
	}
	request.Header.Add("code","250")
	response,err :=client.Do(request)
	if err != nil{
		fmt.Println(err)
		return
	}
	by, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(by))
}
