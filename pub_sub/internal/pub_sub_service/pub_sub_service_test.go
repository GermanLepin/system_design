package pub_sub_service

import (
	"fmt"
	"testing"
)

func Test_PubSub(t *testing.T) {
	t.Run("succeeded pub sub test", func(t *testing.T) {
		agent := NewAgent()
		defer agent.Close()

		//	r := require.New(t)

		sub1 := agent.Subscribe("topic1")
		sub2 := agent.Subscribe("topic2")

		go agent.Publish("topic1", "hello there!")
		go agent.Publish("topic1", "are you alright?")
		go agent.Publish("topic2", "check this out")

		go fmt.Println(<-sub1)
		go fmt.Println(<-sub2)

		//go r.Equal(<-sub1, "message 1")
		//go r.Equal(<-sub2, "message 2")
	})
}
