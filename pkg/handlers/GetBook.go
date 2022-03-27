package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testwithgorm/pkg/mocks"

	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	
	id, _ := strconv.Atoi(vars["id"])
	
	for _, book := range mocks.Books {
		if book.Id == id {
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			break
		}
	} 
}