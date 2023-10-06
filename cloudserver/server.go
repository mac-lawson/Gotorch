package cloudserver

import (
	"fmt"
	"net/http"
)

func tensorAPI() {
	http.HandleFunc("/tensor", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
		fmt.Println(r.URL.Query().Get("id"))
	})
}

func RunAPIServer() {

	tensorAPI()

	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
