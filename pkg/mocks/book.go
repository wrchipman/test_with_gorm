package mocks

import "testwithgorm/pkg/models"



var Books = []models.Book{
	{
		Id: 1,
		Title: "Golang",
		Author: "Gopher",
		Desc: "A book for go",
	},
}