package handlers

// import (
// 	// "ccg/feddis"
// 	"ccg/models"
// 	"ccg/repository"
// 	"fmt"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// // var fds feddis.Client = *feddis.GetClient()

// func GetGhoul(c echo.Context) error {
// 	// data, _ := repository.GetGhoul()

// 	// js, _ := json.Marshal(data)
// 	// fmt.Println(string(js))
// 	return c.String(http.StatusOK, "ну хз брат")
// }

// func AddGhoul(c echo.Context) error {

// 	var ghoul models.Ghoul

// 	c.Bind(&ghoul)

// 	fmt.Println(ghoul)

// 	repository.CreateGhoul(ghoul)

// 	return c.String(http.StatusOK, "иди нахуй")
// }

// func EditGhoul(c echo.Context) error {
// 	firstName := c.FormValue("firstName")
// 	lastName := c.FormValue("lastName")
// 	rank := c.FormValue("rank")

// 	ccg := models.CCG{
// 		FirstName: firstName,
// 		LastName:  lastName,
// 		Rank:      rank,
// 	}

// 	repository.EditCCG(ccg)

// 	return c.String(http.StatusOK, "всё норм бро")
// }

// func GetGhoulByName(name string) (models.Ghoul, error){
// 	return , nil
// }
