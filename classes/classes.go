package main

import (
	"fmt"
	"time"
)

type Message struct {
	text   string
	sender string
	sent   time.Time
}

func main() {
	var message Message = Message{"asdf", "me", time.Now()}
	fmt.Println(message)
	var pointer *Message = &message
	fmt.Println(pointer.text) //implicit dereference
	var (
		//marr = make([]Message, 12)
		//marrp *[]Message = &marr
		mmap map[int]Message = make(map[int]Message) //map[key]value
	)
	mmap[0] = message
	fmt.Println(mmap)
	mmapliteral := map[int]Message{
		0: Message{
			"hi", "we", time.Now(),
		},
	}
	fmt.Printf("%T, %v\n", mmapliteral, mmapliteral)
	var message2 Message = Message{"asdf", "me", time.Time{}}
	fmt.Println(message2)
	message2.send()
	fmt.Println(message2)
	//var a time.Time
	//a.send()
}
func (m *Message) send() {
	m.sent = time.Now()
}
