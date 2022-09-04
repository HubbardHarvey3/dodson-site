package main

import (
  "net/http"
)

func main() {


http.Handle("/", http.FileServer(http.Dir("../dodson-site/dist/dodson-site")))

http.ListenAndServe("127.0.0.1:3000", nil) 


}
