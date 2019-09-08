package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rhyuen/golang_mongo_faas/types"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update endpoint")
	cap := types.Quote{"Captain America", "I can do this all day."}
	json.NewEncoder(w).Encode(cap)
}
