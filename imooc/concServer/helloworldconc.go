package main

import (
	"fmt"
)

func printhelloworld(i int,ch chan string)  {
	 for{
		 ch<-fmt.Sprintf("hello world %d \n",i)
	 }
}

func main() {
    ch :=make(chan string)
	for i :=0;i<5 ;i++  {
		go printhelloworld(i,ch)
	}
	for   {
		msg := <-ch
		fmt.Println(msg)
	}

	//time.Sleep(time.Microsecond)
}
