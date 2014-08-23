package pubsub

import (
	"time"
)

type PubSub interface {
	Subscribe(topic string, callback Callback) (Subscription, error)
	SubscribeSync(topic string, callback CallbackSync) (Subscription, error)
	Publish(topic string, data []byte) error
	PublishSync(topic string, data []byte, timeout time.Duration) ([]byte, error)
	Close()
}

type Subscription interface {
	Unsubscribe() error
}

type Callback func(data []byte)

type CallbackSync func(data []byte) ([]byte, error)
