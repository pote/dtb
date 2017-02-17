package main

import(
	"fmt"
	"log"
	"math/rand"
  "net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", MESSAGES[rand.Intn(len(MESSAGES))])
	})

	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))
}
