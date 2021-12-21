package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go Handle(conn)
	}

}

var data = make(map[string]string)

type Command struct {
	Fields []string
	Result chan string
}

func RedisServer(commands chan Command) {
	for cmd:=range commands{

	}

}

func Handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		//skip blank
		if len(fs) < 1 {
			continue
		}

		switch fs[0] {
		//GET VALUE FROM <KEY>
		case "GET":
			key := fs[1]
			value := data[key]
			fmt.Fprintf(conn, "%s\n", value)
			//SET <KEY> <VALUE>
		case "SET":
			if len(fs) != 3 {
				io.WriteString(conn, "INVALID COMMAND(need SET <key> <value> \n")
				continue
			}
			key := fs[1]
			value := fs[2]
			data[key] = value
			//DELETE <KEY>
		case "DEL":
			key := fs[1]
			delete(data, key)

		default:
			io.WriteString(conn, "INVALID COMMAND: "+fs[0])

		}
		fmt.Println(ln)
	}
}
