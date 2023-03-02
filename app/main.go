package main

import (
    "fmt"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "welcome")
    fmt.Println("Endpont Hit: homePage")
}

func handleRequest() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":3010", nil))
}

func main(){
    handleRequest();
}