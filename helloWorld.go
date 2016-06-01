package main

import (
  "encoding/json"
  "net/http"
  "log"
)

type User struct{
  FirstName string
  LastName string
}

type GreetUser struct{
  Body string
}

func HiHandler( w http.ResponseWriter, r *http.Request) {
  var u User
  // checks the request body
  if r.Body == nil {
      http.Error(w, "Please send a request body", 400)
      return
  }
  // checks json
  err := json.NewDecoder(r.Body).Decode(&u)
  if err != nil {
      http.Error(w, err.Error(), 400)
      return
  }
  // checks fields in json 
  if u.FirstName == "" || u.LastName == "" {
    http.Error(w, "Please fill both fields", 400)
    return
  }
  // sends back json
  json.NewEncoder(w).Encode(GreetUser{"Hi " + u.FirstName + " " + u.LastName})

}

func main() {
  http.HandleFunc( "/hi", HiHandler )
  log.Fatal(http.ListenAndServe(":8080", nil))
}
