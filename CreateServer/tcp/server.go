package main

import (
	"bufio"
	"log"
	"net"
	"fmt"
	"time"
	"errors"
)

func process(conn net.Conn){

	defer conn.Close()
	reader := bufio.NewReader(conn)
	var buf [128]byte
	n, err := reader.Read(buf[:])
	if err != nil{
		log.Fatalf("read from conn failed: %v\n", err)
	}

	recv := string(buf[:n])
	fmt.Printf("Receive : %v with %v length at %v\n", recv, n, time.Now())

	_, err = conn.Write([]byte("ok"))
	if err != nil{
		log.Fatalf("wrtie from conn failed: %v\n", err)
	}
}


func main(){
	listen, err := net.Listen("tcp", ":8000")
	defer listen.Close() // Close server 
	if err != nil{
		panic(errors.New("Listen failed"))
	}
	
	for {
		conn, err := listen.Accept()
		if err != nil{
			panic(errors.New("Conn failed"))
		}
		go process(conn)
	}
}