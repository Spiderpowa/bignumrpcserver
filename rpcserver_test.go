package main

import "testing"

import (
        "time"
        "net"
        "net/rpc"
        "net/rpc/jsonrpc"
)

type RPCClient struct {
    *rpc.Client
    listener net.Listener
    isClosed chan bool
}

func (client *RPCClient) Close(t *testing.T) {
    go func() {
        t.Log("Waiting for a brief time before closing current server...")
        time.Sleep(time.Millisecond*500)
        client.listener.Close()
    }()

    client.isClosed <- true
}

func initClient(t *testing.T) *RPCClient {
    listener := NewListener("tcp", ":1234")
    if listener == nil {
        t.Error("Error Listening")
    }
    isClosed := make(chan bool)
    go StartRPCServer(listener, isClosed)
    client, err := net.Dial("tcp", "localhost:1234")
    if err != nil {
        t.Error("Dialing:", err)
    }
    c := jsonrpc.NewClient(client)
    rpcClient := &RPCClient{c, listener, isClosed}
    return rpcClient
}

func makeTable(t *testing.T, client *RPCClient)  {
    numbers := [][]string {
        {"Zero", "0"},
        {"One", "1"},
        {"NegOne", "-1"},
        {"Two", "2"},
        {"Three", "3.0"},
        {"NegThree", "-3.0"},
        {"Four", "4"},
        {"Five.One", "5.1"},
        {"PI", "3.14"},
        {"OneAndLotsOfZero", "100000000000000000000000000000000"},
        {"LotsOfNine", "999999999999999999999999999999999999"},
        {"NegLotsOfNine", "-999999999999999999999999999999999999"},
    }
    var reply string
    for _, n := range numbers {
        err := client.Call("BigNumber.Create", n, &reply)
        if err != nil {
            t.Errorf("Setting error %q", n)
        }
    }
}

func TestCreate(t *testing.T) {
    client := initClient(t)
    defer client.Close(t)
    cases := []struct {
        in string
        out bool
    }{
        {"A", true},
        {"A", false},
        {"B", true},
        {"C", true},
        {"C", false},
    }
    var reply string
    for _, c := range cases {
        out := client.Call("BigNumber.Create", []string{c.in, "1"}, &reply) == nil
        if out != c.out {
            t.Errorf("Fail to create (%s) == %t, expect %t\n", c.in, out, c.out)
        }
    }
}

func TestAdd(t *testing.T) {
    client := initClient(t)
    defer client.Close(t)
    makeTable(t, client)
}