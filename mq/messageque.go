package mq

import (
  "fmt"
  "time"
  "sync"
)
/*

subscribers - subscribes to a topic
Topic - maintains a queue of messages and subscribers
publisher - publishes messages to Topics
messageque - 
*/



type SubscriberList struct {
 subscribers []chan string
}


type messageque struct {
  Topics map[string]*SubscriberList
  mutex *sync.Mutex
}

func newMessageque() *messageque {
  topics := make(map[string]*SubscriberList)
  return &messageque{
    Topics: topics,
  }
}

func (mq *messageque) Publish(topic, message string) {
  fmt.Println("publishing : ", topic, message)
  sublist, ok := mq.Topics[topic]
  if ok {
    // subscribers found
    for _, subscriber := range sublist.subscribers {
        go func() {
            subscriber <- message
        }()
    }
  }
  
}

func (mq *messageque) Subscribe(topic string) {
  new_subscriber := make(chan string, 1)
  sublist, ok := mq.Topics[topic]
  if !ok {
    sublist = new(SubscriberList)
  }
  sublist.subscribers = append(sublist.subscribers, new_subscriber)

  mq.Topics[topic] = sublist
  fmt.Println("subscribing topic: ", topic)
  for {
    message := <- new_subscriber
    fmt.Println("got message: ", message)
  }
}

// global message que/ also can be on file/use linux message que/ pipe
var MessageQue *messageque = newMessageque()

type Publisher struct {}

func (p *Publisher) Publish(topic, message string) {
  MessageQue.Publish(topic, message)
}

type Subscriber struct {}

func (s *Subscriber) Subscribe(topic string) {
  MessageQue.Subscribe(topic)
}



func main() {
  // you can write to stdout for debugging purposes, e.g.
  go func() {
   // subscriber 1 
   var s1 Subscriber
   s1.Subscribe("Topic")
  }()
  go func() {
    // subsciber 2
    var s2 Subscriber
    s2.Subscribe("Topic 2")
  }()

    time.Sleep(1 * time.Second)
  go func() {
    // publisher
    var p Publisher 
    p.Publish("Topic", "this is a test message")
    p.Publish("Topic 2", "this is a test message 2")
  }()

  time.Sleep(10 * time.Second)
  fmt.Println("This is a debug message")
}
