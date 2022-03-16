package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
        response := map[string]string{
            "message": "Hello Docker!",
        }
        json.NewEncoder(rw).Encode(response)
    })

    log.Println("Server is running!")
    http.ListenAndServe(":4000", router)
}



// Understanding the go.mod File

// When you run commands with the go tool, the go.mod file is a very important part of the process. It’s the file that contains the name of the module and versions of other modules your own module depends on. It can also contain other directives, such as replace, which can be helpful for doing development on multiple modules at once.