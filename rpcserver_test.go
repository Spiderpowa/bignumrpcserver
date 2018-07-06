package main

import "testing"

import (
        "net"
        "net/rpc"
        "net/rpc/jsonrpc"
)

func initClient(t *testing.T) *rpc.Client {
    go StartRPCServer()
    client, err := net.Dial("tcp", "localhost:1234")
    if err != nil {
        t.Error("Dialing:", err)
    }
    c := jsonrpc.NewClient(client)
    return c
}

func TestCreate(t *testing.T) {
    client := initClient(t)
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

