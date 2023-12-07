package main

import (
	"fmt"
	"system_design/pub_sub/internal/pub_sub_service"
)

func main() {
	agent := pub_sub_service.NewAgent()
	defer agent.Close()

	sub1 := agent.Subscribe("topic1")
	sub2 := agent.Subscribe("topic1")

	sub3 := agent.Subscribe("topic2")
	sub4 := agent.Subscribe("topic2")
	sub5 := agent.Subscribe("topic2")

	sub6 := agent.Subscribe("topic3")
	sub7 := agent.Subscribe("topic3")
	sub8 := agent.Subscribe("topic3")
	sub9 := agent.Subscribe("topic3")

	go agent.Publish("topic1", "hello there!")
	go agent.Publish("topic1", "are you alright?")
	go agent.Publish("topic2", "check this out")
	go agent.Publish("topic2", "Wow! It is awesome!")
	go agent.Publish("topic3", "Appreciate it")

	go agent.Read(sub1)
	go agent.Read(sub2)
	go agent.Read(sub3)
	go agent.Read(sub4)
	go agent.Read(sub5)
	go agent.Read(sub6)
	go fmt.Println(<-sub7)
	go fmt.Println(<-sub8)
	go fmt.Println(<-sub9)
}
