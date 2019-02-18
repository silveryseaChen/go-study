package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	testDefer()

	testHttp()
	
}


func testDefer()  {
	/** defer 后的语句 将在函数结束（return）之前执行 */
	defer fmt.Println("this is the defer line")
	fmt.Println("do something...")
}

func testHttp(){
	http.HandleFunc("/go", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(" a request received")
		io.WriteString(writer, "hello world")
	})
	http.ListenAndServe(":8080", nil)
}
