package main

import (
  "context"
  "fmt"
  "net"
  "os"
  "time"
)

func main() {
  params := os.Args[1:]
  if len(params) < 3 {
    fmt.Println("usage: knock host port1 port2 portN")
    return
  }
  host := params[:1]
  ports := params[1:]
  var d net.Dialer
  for _, port := range ports {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
    defer cancel()
    _, _ = d.DialContext(ctx, "tcp", fmt.Sprintf("%s:%v", host, port))
    fmt.Println("knock on port", port)
  }
}
