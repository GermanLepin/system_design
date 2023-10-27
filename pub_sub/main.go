package main

import (
	"fmt"
	"sync"
)

// Pub/sub, short for “publish-subscribe,” is a messaging pattern that decouples systems
// and communications between them. As indicated in the name, a pub/sub system has two
// types of entities: Publishers and Subscribers. Publishers are entities that produce
// messages and publish them to a specific topic, whereas Subscribers are entities that
// receive messages by subscribing to a specific topic. The pub/sub pattern is often
// used to decouple different parts of a system in a cloud system. For example, a message
// queue or message bus can act as a pub/sub system, allowing different parts of a system
// to communicate without being directly connected to each other. This makes it easier to
// manage complex systems and improve scalability and reliability.

func main() {
	agent := newAgent()

	sub1 := agent.Subscribe("topic1")
	sub4 := agent.Subscribe("topic1")

	sub2 := agent.Subscribe("topic2")
	sub5 := agent.Subscribe("topic2")

	sub3 := agent.Subscribe("topic3")
	sub6 := agent.Subscribe("topic3")
	sub7 := agent.Subscribe("topic3")
	sub8 := agent.Subscribe("topic3")

	agent.wg.Add(11)
	go agent.Publish("topic1", "message 1")
	go agent.Publish("topic2", "message 2")
	go agent.Publish("topic3", "message 3")

	go agent.Read(sub1)
	go agent.Read(sub2)
	go agent.Read(sub3)
	go agent.Read(sub4)
	go agent.Read(sub5)
	go agent.Read(sub6)
	go agent.Read(sub7)
	go agent.Read(sub8)

	agent.wg.Wait()
}

func (a *Agent) Subscribe(topic string) chan string {
	ch := make(chan string)
	a.subs[topic] = append(a.subs[topic], ch)

	return ch
}

func (a *Agent) Publish(topic string, msg string) {
	defer a.wg.Done()

	for _, ch := range a.subs[topic] {
		ch <- msg
		close(ch)
	}
}

func (a *Agent) Read(ch chan string) {
	defer a.wg.Done()

	fmt.Println(<-ch)
}

type Agent struct {
	wg   sync.WaitGroup
	subs map[string][]chan string
}

func newAgent() *Agent {
	return &Agent{
		subs: make(map[string][]chan string),
	}
}
