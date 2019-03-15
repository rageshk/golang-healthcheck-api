package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)

type ApplicationData struct {
   ApplicationVersion string `json:"version"`
   Description string `json:"description"`
   Lastcommitsha string `json:"lastcommitsha"`  
}



func handleRequests() {
    http.HandleFunc("/healthcheck", returnHealthCheckData)
    log.Fatal(http.ListenAndServe(":8082", nil))
}

func returnHealthCheckData(w http.ResponseWriter, r *http.Request){

    //Hardcoding the response as of now.
    //TODO - integrate with Github API to retrive updated application status
    //response,err := http.Get("https://api.github.com/repos/rageshk/health_check_api/statuses/53c7d37603b5125c701a384e34ac75b2856427cb")
    appData := ApplicationData{ApplicationVersion: "1.0", Description: "pre-interview technical test", Lastcommitsha: "abc57858585"}

    fmt.Println("Endpoint Hit: returnHealthCheckData")

    json.NewEncoder(w).Encode(appData)
}


func main() {
    handleRequests()
}