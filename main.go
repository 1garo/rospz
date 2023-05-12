package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type User struct {
    Age int `json:"age"`
    Name string `json:"name"`
    N   int `json:"count"`
}

type Error struct {
    Message string `json:"message"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/json")

    var u User
    u.Name = r.URL.Query().Get("name")

    age, err := strconv.Atoi(r.URL.Query().Get("age")); 
    if err != nil {
        var e Error
        
        w.WriteHeader(http.StatusBadRequest)
        e.Message = "failed to parse `age` query param"

        err = json.NewEncoder(w).Encode(e)
        if err != nil {
            log.Fatal(err)
        }
        return
    }
    u.Age = age

    err = json.NewEncoder(w).Encode(u)
    if err != nil {
        log.Fatal(err.Error())

    }
}

func main() {
    http.HandleFunc("/user", GetUser)
    log.Fatal(http.ListenAndServe(":8001", nil))
}

