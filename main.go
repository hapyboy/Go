package main

import (
"fmt"
"time"
)

var c chan int
func main(){
	c =make(chan int)
	
	go out("90-1")
	go out("go-2")
	for i:=0;i<100;i++{
		<- c
		
	}
	
}
func out(s string){
	for i := 0; i<50; i++{
		c<-0
		fmt.Println(s)
		time.Sleep(1e8)
	}
}