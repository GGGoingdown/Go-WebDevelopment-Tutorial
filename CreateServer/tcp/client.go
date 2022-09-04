package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

func worker(number int, wg *sync.WaitGroup) {

	name := fmt.Sprintf("worker-%v", strconv.Itoa(number))
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalln("Connection failed", err)
	}

	_, err = conn.Write([]byte(name))
	if err != nil {
		log.Fatalf("Send failed, error : %v\n", err)
	}

	var buf [1024]byte
	n, err := conn.Read(buf[:])

	if err != nil {
		log.Fatalf("Read failed. error: %v\n", err)
	}

	fmt.Println(string(buf[:n]))
	wg.Done()
}

func main() {
	clientNumber := 50

	wg := &sync.WaitGroup{}
	wg.Add(clientNumber)
	fmt.Println("Start time", time.Now())
	for i := 1; i <= clientNumber; i++ {
		go worker(i, wg)
	}
	wg.Wait()
	fmt.Println("Stop time", time.Now())
	fmt.Println()
}
