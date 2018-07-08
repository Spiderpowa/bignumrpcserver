package main

import "testing"
import (
	"bytes"
	"log"
	"strings"
	"time"
)

func TestNewListener(t *testing.T) {
	l1 := NewListener("tcp", ":5678")
	l2 := NewListener("tcp", ":5678")
	if l2 != nil {
		t.Errorf("Expect nil when listen on duplicated port\n")
	}
	l1.Close()
}

func TestStartRPCServer(t *testing.T) {
	listener := NewListener("tcp", ":5555")
	isClosed := make(chan bool)
	go StartRPCServer(listener, isClosed)
	go func() {
		time.Sleep(time.Millisecond * 300)
		listener.Close()
	}()
	isClosed <- true
	var buf bytes.Buffer
	log.SetOutput(&buf)
	go StartRPCServer(listener, nil)
	time.Sleep(time.Millisecond * 300)
	if !strings.Contains(buf.String(), "accept error") {
		t.Errorf("Expect to have 'accept error', but get %s\n", buf.String())
	}
}
