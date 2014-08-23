package impltests

import (
	. "github.com/puffinframework/pubsub"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSubscribe(t *testing.T, pb PubSub) {
	countTopic1 := 0
	countTopic2 := 0
	countTopic3 := 0

	sub1, err1 := pb.Subscribe("topic1", func(data []byte) {
		countTopic1++
		assert.Equal(t, "value1", string(data))
	})
	assert.NotNil(t, sub1)
	assert.Nil(t, err1)

	sub2, err2 := pb.Subscribe("topic2", func(data []byte) {
		countTopic2++
		assert.Equal(t, "value2", string(data))
	})
	assert.NotNil(t, sub2)
	assert.Nil(t, err2)

	sub3, err3 := pb.Subscribe("topic3", func(data []byte) { countTopic3++ })
	assert.NotNil(t, sub3)
	assert.Nil(t, err3)

	assert.Nil(t, pb.Publish("topic1", []byte("value1")))
	assert.Nil(t, pb.Publish("topic2", []byte("value2")))

	time.Sleep(10 * time.Millisecond)

	assert.Equal(t, 1, countTopic1)
	assert.Equal(t, 1, countTopic2)
	assert.Equal(t, 0, countTopic3)
}

func TestUnsubscribe(t *testing.T, pb PubSub) {
	countTopic1 := 0
	countTopic2 := 0

	sub1, err1 := pb.Subscribe("topic1", func(data []byte) { countTopic1++ })
	assert.NotNil(t, sub1)
	assert.Nil(t, err1)

	sub2, err2 := pb.Subscribe("topic2", func(data []byte) { countTopic2++ })
	assert.NotNil(t, sub2)
	assert.Nil(t, err2)

	assert.Nil(t, pb.Publish("topic1", nil))
	assert.Nil(t, pb.Publish("topic2", nil))

	time.Sleep(10 * time.Millisecond)

	assert.Nil(t, sub1.Unsubscribe())

	assert.Nil(t, pb.Publish("topic1", nil))
	assert.Nil(t, pb.Publish("topic2", nil))

	time.Sleep(10 * time.Millisecond)

	assert.Equal(t, 1, countTopic1)
	assert.Equal(t, 2, countTopic2)

	assert.Nil(t, sub1.Unsubscribe())
}

func TestSubscribeSync(t *testing.T, pb PubSub) {
	countTopic1 := 0
	countTopic2 := 0

	sub1, err1 := pb.SubscribeSync("topicX", func(data []byte) (result []byte, err error) {
		countTopic1++
		assert.Equal(t, "request1", string(data))
		return []byte("result1"), nil
	})
	assert.NotNil(t, sub1)
	assert.Nil(t, err1)

	sub2, err2 := pb.Subscribe("topicX", func(data []byte) { countTopic2++ })
	assert.NotNil(t, sub2)
	assert.Nil(t, err2)

	result1, err1 := pb.PublishSync("topicX", []byte("request1"), 10*time.Millisecond)
	assert.Equal(t, "result1", string(result1))

	time.Sleep(10 * time.Millisecond)

	assert.Equal(t, 1, countTopic1)
	assert.Equal(t, 1, countTopic2)
}
