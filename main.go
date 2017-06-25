package main

import (
	"fmt"
	"net/http"
	"sync"
)

var pool = sync.Pool{
	New: func() interface{} {
		return "init"
	},
}

func get() interface{} {
	s := pool.Get().(string)
	cpS := s
	pool.Put(s)
	return cpS
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/pool", poolHandler)
	mux.HandleFunc("/pool/change", poolChangeHandler)
	http.ListenAndServe(":7777", mux)
}

func poolHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("----------------------------------")
	fmt.Println("/pool")
	fmt.Println(get().(string))
	fmt.Println("----------------------------------")
}

func poolChangeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("----------------------------------")
	fmt.Println("/pool/change")
	s := get().(string)
	s = "change handler"
	pool.Put(s)
	fmt.Println("-- changed --")
	fmt.Println("----------------------------------")
}
