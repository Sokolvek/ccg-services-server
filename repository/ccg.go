package repository

import (
	"ccg/models"
	"ccg/storage"
	"context"

	"github.com/labstack/echo/v4"
)

type CCGRepository interface {
	CreateCCG(ctx echo.Context) error
	GetCGG(ctx echo.Context, id int) error
	GetAllCCG(ctx echo.Context)
}

func CreateCCG(ccg models.CCG) error {
	docId := ccg.FirstName + "_" + ccg.LastName

	ref := storage.DB.Collection("ccgs").Doc(docId)
	data := map[string]interface{}{
		"FirstName": ccg.FirstName,
		"LastName":  ccg.LastName,
		"Rank":      ccg.Rank,
	}
	ref.Create(context.Background(), data)
	return nil
}

func GetCCG() ([]models.CCG, error) {
	var ccgs []models.CCG

	docs, err := storage.DB.Collection("ccgs").Documents(context.Background()).GetAll()
	if err != nil {
		return nil, err
	}
	for _, doc := range docs {
		var ccg models.CCG
		if err := doc.DataTo(&ccg); err != nil {
			return nil, err
		}
		ccgs = append(ccgs, ccg)
	}
	return ccgs, nil
}
