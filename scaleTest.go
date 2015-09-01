package main

import (
	//	"log"
	//	"net/http"
	//	_ "net/http/pprof"
	"network/ipv4/ipv4tps"
	"network/tcp"

	"os"
	"strconv"
	//"time"

	"github.com/hsheth2/logs"
)

func main() {
	numConn, _ := strconv.Atoi(os.Args[1])

	//	go func() {
	//		log.Println(http.ListenAndServe("localhost:6060", nil))
	//	}()

	s, err := tcp.New_Server_TCB()
	if err != nil {
		logs.Error.Println(err)
		return
	}

	err = s.BindListenWithQueueSize(49230, ipv4tps.IP_ALL, 10+numConn)
	if err != nil {
		logs.Error.Println(err)
		return
	}

	count := make(chan bool, numConn)
	done := make(chan bool)

	for i := 1; i <= numConn; i++ {
		conn, _, _, err := s.Accept()
		if err != nil {
			logs.Error.Println(err)
			return
		}
		//logs.Info.Println("Connection:", ip, port)

		go func(conn *tcp.TCB, count chan bool) {
			data, err := conn.Recv(10000)
			if err != nil {
				logs.Error.Println(err)
				return
			}

			logs.Info.Println("first 50 bytes of received data:", data[:50])

			//time.Sleep(500 * time.Millisecond)
			conn.Close()
			logs.Trace.Println("connection finished")

			count <- true
			if len(count) >= numConn {
				done <- true
			}
			logs.Info.Println("Chan len", len(count))
		}(conn, count)
		logs.Info.Println("Loop num", i)
	}
	logs.Info.Println("Exited loop")
	<-done
}
