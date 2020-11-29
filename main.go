package main

import (
	"fmt"
	"go-designpattern/builder"
)

func main() {
	builderPattern()
}

func builderPattern() {
	user1 := builder.New("gapgit@gmailgo.com", "123412341234", "", 0)        // required property
	user2 := builder.New("gapgit@gmailgo.com", "123412341234", "gapgit", 20) // extra property
	fmt.Println(user1) // &{gapgit@gmailgo.com 1234  0 {}}
	fmt.Println(user2) // &{gapgit@gmailgo.com 1234 gapgit 10 {}}
}
