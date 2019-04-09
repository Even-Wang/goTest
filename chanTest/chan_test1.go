package chanTest

import (
	"fmt"
	"time"
)

func thrower(c chan int){
	for i := 0; i<5; i++{
		c <- i
		fmt.Println("throw >>",i)
	}
}


func catcher(c chan int) {
	for i :=0; i<5; i++ {
		msg := <-c
		fmt.Println("catch <<",msg)
	}
}


func main(){
	c := make(chan int,3)  //容量为3缓冲通道
	go thrower(c)
	go catcher(c)
	time.Sleep(1000 * time.Millisecond)
}