package main

import (
	"encoding/json"
	"fmt"
	"github.com/loudbund/go-request/request_v1"
	"log"
	"net/http"
	"strings"
	"time"
)

func init() {
}

// 启动一个server来接收参数
func server() {
	// 处理get
	http.HandleFunc("/get", func(writer http.ResponseWriter, request *http.Request) {
		query := request.URL.Query()
		fmt.Println("This is server, GET:", query)
		_, _ = writer.Write([]byte(`{"code":0}`))
	})

	// 处理post form
	http.HandleFunc("/post_form", func(writer http.ResponseWriter, request *http.Request) {
		_ = request.ParseForm()
		fmt.Println("This is server, POST FORM:", request.Form)
		_, _ = writer.Write([]byte(`{"code":0}`))
	})

	// 处理post json
	http.HandleFunc("/post_json", func(writer http.ResponseWriter, request *http.Request) {
		decoder := json.NewDecoder(request.Body)
		var params map[string]interface{}
		_ = decoder.Decode(&params)
		fmt.Println("This is server, POST JSON:", params)
		_, _ = writer.Write([]byte(`{"code":0}`))
	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func main() {
	// 启动一个http服务
	go server()
	time.Sleep(time.Second)

	// 示例1、Get
	fmt.Println(strings.Repeat("=", 30), "start get ")
	code, body, err := request_v1.Get("http://127.0.0.1:3000/get?userid=12345")
	fmt.Println("client get result:", code, body, err)

	// 示例2、PostForm
	fmt.Println(strings.Repeat("=", 30), "start PostForm ")
	code, body, err = request_v1.PostForm("http://127.0.0.1:3000/post_form", map[string]string{
		"userName": "haha",
	})
	fmt.Println("client post_form result:", code, body, err)

	// 示例3、PostJson
	fmt.Println(strings.Repeat("=", 30), "start PostJson ")
	code, body, err = request_v1.PostJson("http://127.0.0.1:3000/post_json?userid=12345", map[string]interface{}{
		"name": "wawa",
		"data": map[string]interface{}{
			"age": 123,
		},
	})
	fmt.Println("client post_json result:", code, body, err)

}
