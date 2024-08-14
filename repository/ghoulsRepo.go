package repository

import (
	"ccg/models"
	"ccg/storage"
	"ccg/usecases"
	"context"

	"github.com/labstack/echo/v4"
)

type GhoulRepository interface {
	CreateGhoul(ctx echo.Context) error
	GetGhoul(ctx echo.Context, id int) error
	GetAllGhouls(ctx echo.Context) error
}

func CreateGhoul(ghoul models.Ghoul) error {
	docId := ghoul.FirstName + "_" + ghoul.LastName

	ref := storage.DB.Collection("ghouls").Doc(docId)

	data := usecases.StructToInterface(ghoul)

	ref.Create(context.Background(), data)

	return nil
}

func EditGhoul(ghoul models.Ghoul) error {
	docId := ghoul.FirstName + "_" + ghoul.LastName

	ref := storage.DB.Collection("ghouls").Doc(docId)

	data := usecases.StructToInterface(ghoul)

	ref.Set(context.Background(), data)

	return nil
}

func GetGhoul() ([]models.Ghoul, error) {
	var ghouls []models.Ghoul

	docs, err := storage.DB.Collection("ghouls").Documents(context.Background()).GetAll()
	if err != nil {
		return nil, err
	}
	for _, doc := range docs {
		var ghoul models.Ghoul
		if err := doc.DataTo(&ghoul); err != nil {
			return nil, err
		}
		ghouls = append(ghouls, ghoul)
	}
	return ghouls, nil
}
