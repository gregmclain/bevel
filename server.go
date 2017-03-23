package main

import (
    "fmt"
    "net"
    "os"
    "bufio"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "8080"
    CONN_TYPE = "tcp"
)

func ServerStart() {
	fmt.Println("Starting server")
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }

    defer l.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

    for {
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }

        go handleRequest(conn)
    }
}

func handleRequest(conn net.Conn) {
  request, err := bufio.NewReader(conn).ReadString('\n')
  if err != nil {
    fmt.Println("Error reading:", err.Error())
  }

  fmt.Println("Received: ", string(request))

  conn.Write([]byte("HTTP/1.1 200 OK\n\nHello, World!"))
  conn.Close()
}
