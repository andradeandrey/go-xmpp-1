package main

import (
  "fmt"
  "xmpp"
)

func main() {
  client, err := xmpp.NewClient("talk.google.com", "andrade.andrey@gmail.com", "xxx")

  if err != nil {
    fmt.Printf("Failed due to: %s\n", err)
  } else {
    client.OnAny(func(msg string) {
      log(msg)
    })

    client.OnMessage(func(msg xmpp.Message) {
      log("Got a message from %s to %s: %s\n", msg.From(), msg.To(), msg.Body())
    })

    client.SendChat("nilton.kummer@gmail.com", "Hello world!")

    client.Loop()
  }
}

func log(format string, args ...interface{}) {
  fmt.Printf("MAIN LOG: " + format + "\n", args...)
}
