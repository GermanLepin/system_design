package main

import (
	"fmt"
	"sync"
)

// Pub/sub, short for “publish-subscribe,” is a messaging pattern that decouples systems
// and communications between them. As indicated in the name, a pub/sub system has two
// types of entities: Publishers and Subscribers. Publishers are entities that produce
// messages and publish them to a specific topic, whereas Subscribers are entities that
// receive messages by subscribing to a specific topic.

func main() {
	agent := newAgent()

	sub1 := agent.subscribe("topic1")
	sub2 := agent.subscribe("topic1")

	sub3 := agent.subscribe("topic2")
	sub4 := agent.subscribe("topic2")

	sub5 := agent.subscribe("topic3")
	sub6 := agent.subscribe("topic3")
	sub7 := agent.subscribe("topic3")
	sub8 := agent.subscribe("topic3")

	agent.wg.Add(11)
	go agent.publish("topic1", "message 1")
	go agent.publish("topic2", "message 2")
	go agent.publish("topic3", "message 3")

	go agent.read(sub1)
	go agent.read(sub2)
	go agent.read(sub3)
	go agent.read(sub4)
	go agent.read(sub5)
	go agent.read(sub6)
	go agent.read(sub7)
	go agent.read(sub8)

	agent.wg.Wait()
	agent.close()
}

func (a *agent) subscribe(topic string) chan string {
	ch := make(chan string)
	a.subs[topic] = append(a.subs[topic], ch)

	return ch
}

func (a *agent) publish(topic string, msg string) {
	defer a.wg.Done()

	for _, ch := range a.subs[topic] {
		ch <- msg
	}
}

func (a *agent) read(ch chan string) {
	defer a.wg.Done()

	fmt.Println(<-ch)
}

func (a *agent) close() {
	for _, ch := range a.subs {
		for _, sub := range ch {
			close(sub)
		}
	}
}

type agent struct {
	wg   sync.WaitGroup
	subs map[string][]chan string
}

func newAgent() *agent {
	return &agent{
		subs: make(map[string][]chan string),
	}
}
