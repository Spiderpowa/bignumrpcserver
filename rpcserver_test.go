package main

import "testing"

import (
	"math/big"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

type _RPCClient struct {
	*rpc.Client
	listener net.Listener
	isClosed chan bool
}

func (client *_RPCClient) close(t *testing.T) {
	go func() {
		go func() {
			time.Sleep(time.Millisecond * 500)
			client.listener.Close()
		}()
		client.isClosed <- true
	}()
}

func initClient(t *testing.T, address string) *_RPCClient {
	listener := NewListener("tcp", address)
	if listener == nil {
		t.Error("Error Listening")
	}
	isClosed := make(chan bool)
	go StartRPCServer(listener, isClosed)
	client, err := net.Dial("tcp", address)
	if err != nil {
		t.Error("Dialing:", err)
	}
	c := jsonrpc.NewClient(client)
	rpcClient := &_RPCClient{c, listener, isClosed}
	return rpcClient
}

func makeTable(t *testing.T, client *_RPCClient) {
	numbers := [][]string{
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
		{"LotsOfNine", "99999999999999999999999999999999"},
		{"NegLotsOfNine", "-99999999999999999999999999999999"},
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
	client := initClient(t, ":1234")
	defer client.close(t)
	cases := []struct {
		in  string
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

func TestUpdate(t *testing.T) {
	client := initClient(t, ":1230")
	defer client.close(t)
	makeTable(t, client)
	cases := []struct {
		name, val string
		out       bool
	}{
		{"One", "10", true},
		{"Two", "20", true},
		{"NotExists", "5", false},
	}
	var reply string
	for _, c := range cases {
		out := client.Call("BigNumber.Update", []string{c.name, c.val}, &reply) == nil
		if out != c.out {
			t.Errorf("Fail to update (%s) == %t, expect %t\n", c.name, out, c.out)
		}
	}

	casesVerify := []struct {
		x, y, ans string
	}{
		{"One", "Two", "30"},
		{"One", "One", "20"},
		{"Two", "Two", "40"},
	}
	for _, c := range casesVerify {
		err := client.Call("BigNumber.Add", []string{c.x, c.y}, &reply)
		if err != nil {
			t.Errorf("Add (%q %q) Error:%q", c.x, c.y, err)
			continue
		}
		ans, _ := new(big.Float).SetString(c.ans)
		if reply != ans.String() {
			t.Errorf("Update-Add %q %q == %q, expect %q", c.x, c.y, reply, ans.String())
		}
	}
}

func TestDelete(t *testing.T) {
	client := initClient(t, ":1231")
	defer client.close(t)
	makeTable(t, client)
	cases := []struct {
		name string
		out  bool
	}{
		{"One", true},
		{"Two", true},
		{"NotExists", false},
	}
	var reply string
	for _, c := range cases {
		out := client.Call("BigNumber.Delete", c.name, &reply) == nil
		if out != c.out {
			t.Errorf("Fail to delete (%s) == %t, expect %t\n", c.name, out, c.out)
		}
		if !out {
			continue
		}
		suc := client.Call("BigNumber.Update", []string{c.name, "1"}, &reply) != nil
		if !suc {
			t.Errorf("Fail to delete (%s), still found in server\n", c.name)
		}
	}
}

func TestAdd(t *testing.T) {
	client := initClient(t, ":1235")
	defer client.close(t)
	makeTable(t, client)

	cases := []struct {
		x, y, ans string
	}{
		{"One", "One", "2.0"},
		{"One", "NegOne", "0"},
		{"7", "8", "15"},
		{"Five.One", "10", "15.1"},
		{"One", "LotsOfNine", "100000000000000000000000000000000"},
	}
	var reply string
	for _, c := range cases {
		err := client.Call("BigNumber.Add", []string{c.x, c.y}, &reply)
		if err != nil {
			t.Errorf("Add (%q %q) Error:%q", c.x, c.y, err)
			continue
		}
		ans, _ := new(big.Float).SetString(c.ans)
		if reply != ans.String() {
			t.Errorf("Add %q %q == %q, expect %q", c.x, c.y, reply, ans.String())
		}
	}
}

func TestSubtract(t *testing.T) {
	client := initClient(t, ":1236")
	defer client.close(t)
	makeTable(t, client)

	cases := []struct {
		x, y, ans string
	}{
		{"One", "One", "0"},
		{"One", "NegOne", "2"},
		{"Zero", "One", "-1"},
		{"7", "8", "-1"},
		{"Five.One", "10", "-4.9"},
		{"OneAndLotsOfZero", "One", "99999999999999999999999999999999"},
		{"One", "NegLotsOfNine", "100000000000000000000000000000000"},
	}
	var reply string
	for _, c := range cases {
		err := client.Call("BigNumber.Subtract", []string{c.x, c.y}, &reply)
		if err != nil {
			t.Errorf("Sub (%q %q) Error:%q", c.x, c.y, err)
			continue
		}
		ans, _ := new(big.Float).SetString(c.ans)
		if reply != ans.String() {
			t.Errorf("Sub %q %q == %q, expect %q", c.x, c.y, reply, ans.String())
		}
	}
}

func TestMultiply(t *testing.T) {
	client := initClient(t, ":1237")
	defer client.close(t)
	makeTable(t, client)

	cases := []struct {
		x, y, ans string
	}{
		{"One", "One", "1"},
		{"One", "NegOne", "-1"},
		{"Zero", "One", "0"},
		{"NegThree", "NegOne", "3"},
		{"7", "8", "56"},
		{"Five.One", "10", "51"},
		{"OneAndLotsOfZero", "One", "100000000000000000000000000000000"},
		{"OneAndLotsOfZero", "OneAndLotsOfZero", "10000000000000000000000000000000000000000000000000000000000000000"},
	}
	var reply string
	for _, c := range cases {
		err := client.Call("BigNumber.Multiply", []string{c.x, c.y}, &reply)
		if err != nil {
			t.Errorf("Mul (%q %q) Error:%q", c.x, c.y, err)
			continue
		}
		ans, _ := new(big.Float).SetString(c.ans)
		if reply != ans.String() {
			t.Errorf("Mul %q %q == %q, expect %q", c.x, c.y, reply, ans.String())
		}
	}
}

func TestDivision(t *testing.T) {
	client := initClient(t, ":1238")
	defer client.close(t)
	makeTable(t, client)

	cases := []struct {
		x, y, ans string
	}{
		{"One", "One", "1"},
		{"One", "NegOne", "-1"},
		{"Zero", "One", "0"},
		{"NegThree", "NegOne", "3"},
		{"8", "2", "4"},
		{"Five.One", "Three", "1.7"},
		{"OneAndLotsOfZero", "One", "100000000000000000000000000000000"},
		{"OneAndLotsOfZero", "OneAndLotsOfZero", "1"},
	}
	var reply string
	for _, c := range cases {
		err := client.Call("BigNumber.Division", []string{c.x, c.y}, &reply)
		if err != nil {
			t.Errorf("Div (%q %q) Error:%q", c.x, c.y, err)
			continue
		}
		ans, _ := new(big.Float).SetString(c.ans)
		if reply != ans.String() {
			t.Errorf("Div %q %q == %q, expect %q", c.x, c.y, reply, ans.String())
		}
	}
}
