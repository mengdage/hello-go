package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"runtime"
	"strings"
)

const (
	close string = "CLOSE"
)

func printRequest(conn net.Conn, user string) {
	scanner := bufio.NewScanner(conn)
	requestJS := false
	i := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Received from %s: %s\n", user, line)
		if i == 0 {
			if strings.Contains(strings.Fields(line)[1], ".js") {
				requestJS = true
			}
		}

		if len(line) == 0 {
			fmt.Print("-------- CRLF ----------")
			break
		}
		i++
	}

	if requestJS {
		sendScript(conn)
	} else {
		sendResponse(conn)
	}
}

func sendResponse(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Document</title>
	</head>
	<body>
		<h1>Hello World</h1>
		<script src="app.js"></script>
	</body>
	</html>
	`

	fmt.Fprint(conn, "HTTP1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %v\r\n", len(body))
	fmt.Fprint(conn, "text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func sendScript(conn net.Conn) {
	st := `
	const div = document.createElement('div')
	div.textContent = 'hello, js'
	document.body.appendChild(div)
	console.log('hello')
	`

	fmt.Fprint(conn, "HTTP1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %v\r\n", len(st))
	fmt.Fprint(conn, "application/javascript\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, st)
}

func handler(conn net.Conn) {
	fmt.Println("Gorountine#:", runtime.NumGoroutine())
	user := conn.RemoteAddr()
	fmt.Println("User:", user)

	printRequest(conn, fmt.Sprint(user))
}

func main() {
	ln, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(ln.Addr())

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handler(conn)
	}
}
