package main

import (
	"os"
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("content-type", "text/html")
    fmt.Fprintf(w, "<h2>Hello World from Distelli! You have a working Go application Deployment!</h2>")
}

func Average(xs []float64) float64 {
  total := float64(0)
  for _, x := range xs {
    total += x
  }
  return total / float64(len(xs))
}

func main() {
	port := os.Args[1]
	fmt.Println(port)

    http.HandleFunc("/", handler)
    http.ListenAndServe(":"+port, nil)
}
