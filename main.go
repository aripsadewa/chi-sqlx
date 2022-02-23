package main

import (
	"log"
	"net/http"
	"os"
	"rest_api/controller"
	"rest_api/repository"
	"rest_api/service"
	"rest_api/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	appConfig := utils.AppConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error on loading .enc file")
	}
	appConfig.AppPort = getEnv("APP_PORT", "3000")
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
	r.Post("/save", cat.Create())
	r.Get("/{id}", cat.FindById())
	r.Put("/{id}", cat.Update())
	r.Delete("/{id}", cat.Delete())
	r.Get("/cat", cat.FindAll())
	// r.With(5).Get()

	log.Println("Running server at port " + appConfig.AppPort)

	err = http.ListenAndServe(":"+appConfig.AppPort, r)
	if err != nil {
		utils.InternalServerError(err)
	}

	log.Println("test...")
}
