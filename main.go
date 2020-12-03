package main

import (
	"fmt"
	"go-designpattern/builder"
	"go-designpattern/singleton"
	"go-designpattern/chain-resposibility"
	"sync"
)

func main() {
	//wg := &sync.WaitGroup{}
	//builderPattern()
	//singletonPattern(wg)
	//wg.Wait()
	chainPattern()
}

func builderPattern() {
	user1 := builder.New("gapgit@gmailgo.com", "123412341234", "", 0)        // required property
	user2 := builder.New("gapgit@gmailgo.com", "123412341234", "gapgit", 20) // extra property
	fmt.Println(user1) // &{gapgit@gmailgo.com 1234  0 {}}
	fmt.Println(user2) // &{gapgit@gmailgo.com 1234 gapgit 10 {}}
}

func singletonPattern(wg *sync.WaitGroup){
	channel := make(chan *singleton.ImageService)
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go singleton.GetInstance(channel, wg)
		fmt.Printf("%p \n", <- channel)
	}
}


func chainPattern(){
	s3 := chain_resposibility.NewS3()
	local := chain_resposibility.NewLocal()
	other := chain_resposibility.NewOther()
	s3.SetNext(local)
	local.SetNext(other)

	s3.Save(&chain_resposibility.Image{Size: 15000}) // Local save image
	local.Save(&chain_resposibility.Image{Size: 25000}) // Other save image
	other.Save(&chain_resposibility.Image{Size: 25000}) // Other save image
}