// Raw socket connection with no validation and string quoting logic.
// Refer to protocol description:
// http://questdb.io/docs/reference/api/ilp/overview

package configs

import (
	"fmt"
	"log"
	"net"
)

func QestdbConn() *net.TCPConn {
	host := "127.0.0.1:9009"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", host)
	checkErr(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkErr(err)
	fmt.Println("Connected to QuestDB")
	return conn
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
var QestDb *net.TCPConn = QestdbConn()
