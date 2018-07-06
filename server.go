package main

import (
        "fmt"
        "log"
        "net"
        "net/rpc"
        "net/rpc/jsonrpc"
)

func StartRPCServer() {
    calc := CreateBigNumberHandler()
    server := rpc.NewServer()
    server.RegisterName("BigNumber", calc)
    listener, e := net.Listen("tcp", ":1234")
    if e != nil {
        log.Fatal("listen error:", e)
    }
    fmt.Println("RPC Listening")
    for {
        if conn, err := listener.Accept(); err != nil {
            log.Fatal("accept error: " + err.Error())
        } else {
            log.Printf("new connection established\n")
            go server.ServeCodec(jsonrpc.NewServerCodec(conn))
        }
    }
}
