package handler

import (
	"net/http"
	"fmt"
)


func GlobalStatusHandler(w http.ResponseWriter, _ *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Hello, your mahasan bot dev is up 🚀")
}
