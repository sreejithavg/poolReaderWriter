package main

import (

	"fmt"
	"sync"
	"time"
)

type queue struct {
	Timestamp string `json:"timestamp"`
	Data personJson `json:"data"`
}
type personJson struct {
	name string
	address string
	metaData []MetaData
	age int
}
type MetaData struct {
	key   string
	value interface{}
}

var personalDetails  = []personJson{
	{
	name:     "dgfhg",
	address:  "sdhtbh",
	metaData: []MetaData{{
		key:   "education ",
		value: "btech",
	}},
	age:      23,
},
	{
		name:     "djau",
		address:  "cvxd",
		metaData: []MetaData{{
			key:   "education ",
			value: "+2",
		}},
		age:      20,
	},
}
func main()  {
	var q =make(chan queue)
	var wg sync.WaitGroup
	wg.Add(2)
	go writer(q,&wg)
	time.Sleep(time.Millisecond*200)
	go reader(q,&wg)
	wg.Wait()
}
func writer(q chan queue,wg *sync.WaitGroup){
	var queue queue
	defer wg.Done()
	t:=time.Now()
	queue.Timestamp=t.String()
	for _,v:=range personalDetails{
		queue.Data=v
		q<-queue
	}
}

func reader(q chan queue,wg *sync.WaitGroup)  {

	for {
		select {
		case msg,ok:=<-q :
			if ok{
				fmt.Println(msg)
			}else{
				time.Sleep(time.Second*1)
				fmt.Println("channel is closed")
			}

		default:
			fmt.Println("empty queue")
			goto DONE
		}
	}
	DONE:wg.Done()

}
