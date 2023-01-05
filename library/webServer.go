package library

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Response) {
	fmt.Fprintln(w, "apa kabar!")
}
