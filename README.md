package main

import (
	"fmt"
	"time"
)

func main(){
	a := make(chan int,10)
	d := make(chan int,10)
	go add(a,d)

	for{
		sysSelect(a,d,2500)
		fmt.Println("定时任务运行中")
	}

}

func add(a ,b chan int){
	for i:=0;i<10;i++{
		a <- i
		time.Sleep(time.Second)
		b <- i+10
		time.Sleep(time.Second)
	}
}

func sysSelect(a , d chan int,tt  int64){
	now := time.Now().UnixNano()/1e6
	for{
		select {
		case c :=<-a :
			fmt.Println("a读事件",c)
		case d :=<-d:
			fmt.Println("b写事件",d)
		default:
			if(time.Now().UnixNano()/1e6 - now  > tt){
				fmt.Println("结束周期启动定时任务")
				return
		}
		}
	}
}