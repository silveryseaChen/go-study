package main

import (
	"fmt"
	"time"
)

func main() {

	//go Print("go go..")
	//Print("main..")
	/**运行结果 竟然不一样！？？？ 会出现go不足10次的情况**/

	/**通道 相当一个阻塞队列*/
	ch := make(chan int)
	//ch <- 10	error 无缓存区 必须放在异步中
	go cache(ch)
	x := <- ch
	fmt.Println(x)

	ch = make(chan int, 2)
	ch <- 10	//有缓冲区可以同步处理
	ch <- 20
	fmt.Println(<-ch, <-ch)
	//fmt.Println(<-ch, <-ch)	异常 无值时 不可获取

	ch <- 101
	ch <- 202
	close(ch)	//如果不加此行 该语句会报错
	for c := range ch{
		fmt.Println(c)
	}

}

func cache(ch chan int)  {
	ch <- 10
}

func Print(s string)  {

	for i:=0; i<10; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s, i)
	}

}

