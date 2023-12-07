package pub_sub_service

import (
	"fmt"
	"sync"
)

func (a *agent) Subscribe(topic string) <-chan string {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.closed {
		return nil
	}

	ch := make(chan string)
	a.subs[topic] = append(a.subs[topic], ch)
	return ch
}

func (a *agent) Publish(topic string, msg string) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.closed {
		return
	}

	for _, ch := range a.subs[topic] {
		ch <- msg
	}
}

func (a *agent) Close() {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.closed {
		return
	}
	a.closed = true

	for _, ch := range a.subs {
		for _, sub := range ch {
			close(sub)
		}
	}
}

func (a *agent) Read(ch <-chan string) {
	//defer a.Wg.Done()

	for value := range ch {
		fmt.Println(value)
	}
}

type agent struct {
	subs map[string][]chan string

	closed bool
	Wg     sync.WaitGroup
	mu     sync.RWMutex
}

func NewAgent() *agent {
	return &agent{
		subs: make(map[string][]chan string),
	}
}
