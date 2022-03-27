package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testwithgorm/pkg/models"

	"github.com/gorilla/mux"
)

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)

	id, _ :=strconv.Atoi(vars["id"])

	var book models.Book
	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	h.DB.Delete(&book)

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Deleted")
}