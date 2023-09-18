package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mis-hashemi/pgcontroller"
	"github.com/mis-hashemi/pgcontroller/example/example01/entity"
	"github.com/mis-hashemi/pgcontroller/example/example01/repository"
	"github.com/mis-hashemi/pgcontroller/example/example01/service/userservice"
	"github.com/mis-hashemi/pgcontroller/migrator"
	"github.com/mis-hashemi/pgcontroller/page"
	"github.com/mis-hashemi/pgcontroller/sort"
	requestparameter "github.com/mis-hashemi/request-parameter"
	"github.com/mis-hashemi/request-parameter/query"
)

var migrateFlag = flag.String("migrate", "", "Run migration up or down")

func main() {
	flag.Parse()
	controller := pgcontroller.NewPgController(pgcontroller.Config{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "123",
		DBName:   "ex01",
	}, false)

	err := controller.Generate()
	if err != nil {
		panic(err)
	}
	err = controller.Init()
	if err != nil {
		panic(err)
	}
	mgr := migrator.New(controller.GetDataContext(), "./repository/migrations")

	migrateOperation(*migrateFlag, mgr)

	e := echo.New()
	// Define the query parameter information.
	paramInfoMap := map[string]query.RequestParameter{
		entity.FieldFirstName: {
			Definition: query.NewQueryDefinition(entity.FieldFirstName, query.GetAllStringQueryOperator(), query.DataTypeString),
			Optional:   true,
		},
		entity.FieldLastName: {
			Definition: query.NewQueryDefinition(entity.FieldFirstName, query.GetAllStringQueryOperator(), query.DataTypeString),
			Optional:   true,
		},
		entity.FieldPhoneNumber: {
			Definition: query.NewQueryDefinition(entity.FieldPhoneNumber, query.GetAllStringQueryOperator(), query.DataTypeString),
			Optional:   true,
		},
	}
	userRepo := repository.NewUserRepository(controller)
	userSrv := userservice.New(userRepo)
	e.GET("/users", func(c echo.Context) error {

		queryInfo, err := requestparameter.ParseEchoQueryString(c, paramInfoMap)
		if err != nil {
			return c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
		}
		sorts := sort.ParseEchoQueryParamSort(c)
		page, _ := page.ParseEchoQueryParamPagination(c)
		res, err := userSrv.GetAll(c.Request().Context(), queryInfo, sorts, page)
		if err != nil {
			return c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err.Error()))
		}
		return c.JSON(http.StatusOK, res)
	})

	e.Start(":8080")
}

func migrateOperation(flag string, mg migrator.Migrator) {
	switch flag {
	case "up":
		mg.Up()
	case "down":
		mg.Down()
	case "":
	default:
		log.Println("flag value is invalid, this flag only accepts the following values: up, down.")
	}
}
