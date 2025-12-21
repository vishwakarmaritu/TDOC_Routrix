package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/lugnitdgp/TDOC_Routrix/internal/core"
)

func startDummyBackend(port string) {
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello from backend %s", port)
		})

		log.Printf("[DUMMY] starting backend on :%s\n", port)
		if err := http.ListenAndServe(":"+port, mux); err != nil {
			log.Printf("[DUMMY %s] crashed: %v", port, err)
		}
	}()
}

func main() {
	//Start
	startDummyBackend("9001")
	startDummyBackend("9002")	
	startDummyBackend("9003")	

//backend pool
pool := core.NewServerPool()
pool.AddServer(&core.Backend{Address: "localhost:9001", Alive:true})
pool.AddServer(&core.Backend{Address: "localhost:9002", Alive:true})
pool.AddServer(&core.Backend{Address: "localhost:9003", Alive:true})
log.Println("added 3 servers")
}