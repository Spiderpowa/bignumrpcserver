package main

func main() {
    listener := NewListener("tcp", ":1234")
    if listener != nil {
        StartRPCServer(listener, nil)
    }
}
