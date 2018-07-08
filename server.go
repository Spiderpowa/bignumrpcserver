package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// NewListener will create listener and start listening
func NewListener(network, address string) net.Listener {
	listener, e := net.Listen(network, address)
	if e != nil {
		log.Println("listen error:", e)
		return nil
	}

	fmt.Println("RPC Listening")
	return listener
}

// StartRPCServer takes the listener and start aceepting connection from it
func StartRPCServer(listener net.Listener, isClosed chan bool) {
	calc := CreateBigNumberHandler()
	server := rpc.NewServer()
	server.RegisterName("BigNumber", calc)
	for {
		if conn, err := listener.Accept(); err != nil {
			select {
			case <-isClosed:
				return
			default:
				break
			}
			log.Println("accept error: " + err.Error())
			break
		} else {
			log.Printf("new connection established\n")
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}
	}
}
