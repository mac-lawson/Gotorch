package security

import (
	"fmt"
	"net/http"
)

func OpenWebhook(port uint64) error {
	http.HandleFunc("/tensor", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
		fmt.Println(r.URL.Query().Get("id"))
	})
}
