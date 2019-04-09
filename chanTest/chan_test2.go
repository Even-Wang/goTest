package chanTest

import (
	"time"
	"fmt"
)

func printnum(w chan bool)  {
	for i := 0; i <10; i++{
		time.Sleep(time.Millisecond)
		fmt.Printf("%d",i)
	}
	w <- true

}

func printstr(w chan bool)  {
	for i := 'A'; i <'A'+10; i++{
		time.Sleep(time.Millisecond)
		fmt.Printf("%c",i)
	}
	w <- true

}

func main()  {
	w1,w2 := make(chan bool),make(chan bool)
	go printnum(w1)
	go printstr(w2)
	<- w1          //bool类型通道实现同步
	<- w2
}