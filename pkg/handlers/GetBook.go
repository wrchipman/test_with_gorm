package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testwithgorm/pkg/models"

	"github.com/gorilla/mux"
)

func (h handler) GetBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	
	id, _ := strconv.Atoi(vars["id"])
	
	var book []models.Book

	if result := h.DB.Find(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}