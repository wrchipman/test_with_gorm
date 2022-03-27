package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testwithgorm/pkg/mocks"

	"github.com/gorilla/mux"
)

func DeleteBook(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)

	id, _ :=strconv.Atoi(vars["id"])

	for index, book := range mocks.Books {
		if book.Id == id {
			mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)

			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}