package main

import (
	"log"
	"net/http"
	"rest_api/app"
	"rest_api/controller"
	"rest_api/repository"
	"rest_api/service"
	"rest_api/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	utils.InitiEnvConfigs()
	db, err := sqlx.Connect("mysql", "root:@(localhost:3306)/rest-pzn")
	if err != nil {
		log.Fatalln(err)
	}
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(10)

	catRepo := repository.NewCategoryRepository(db)
	catService := service.NewCategoryService(catRepo)
	cat := controller.NewCategoryController(catService, validator.New())
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Route("/v1", func(r chi.Router) {
		app.NewRouter(r, cat)
	})
	log.Println("Running server at port " + utils.CekNilParameter(utils.EnvConfigs.AppPort, "3000"))

	err = http.ListenAndServe(":"+utils.CekNilParameter(utils.EnvConfigs.AppPort, "3000"), r)
	if err != nil {
		utils.InternalServerError(err)
	}

	log.Println("test...")
}
