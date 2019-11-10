package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/leejh3224/golang-starter/config"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})

	port := config.LoadConfig().GetInt("port")

	fmt.Println("🚀 server started on " + strconv.Itoa(port))

	err := http.ListenAndServe(":"+strconv.Itoa(port), r)

	if err != nil {
		log.Fatal("listen and serve!", err)
	}
}
