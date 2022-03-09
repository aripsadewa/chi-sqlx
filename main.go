package main

import (
	"log"
	"net/http"
	"rest_api/app"
	"rest_api/controller"
	"rest_api/repository"
	"rest_api/service"
	"rest_api/utils"

	_ "rest_api/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	httpSwagger "github.com/swaggo/http-swagger"
)

// type M map[string]interface{}

// type MyClaims struct {
// 	jwt.StandardClaims
// 	Username string `json:"Username"`
// 	Email    string `json:"Email"`
// 	Group    string `json:"Group"`
// }

// var APPLICATION_NAME = "My Simple JWT App"
// var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
// var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
// var JWT_SIGNATURE_KEY = []byte(utils.EnvConfigs.SecretApp)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /api/v1

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
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	user := controller.NewUserController(userService)
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
		// r.Post("/login1", user.Login())
	})
	r.Route("/api/v1", func(r chi.Router) {
		app.NewCategoryRouter(r, cat)
		app.NewUserRouter(r, user)

	})
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3000/swagger/doc.json"), //The url pointing to API definition
	))
	log.Println("Running server at port " + utils.CekNilParameter(utils.EnvConfigs.AppPort, "3000"))

	err = http.ListenAndServe(":"+utils.CekNilParameter(utils.EnvConfigs.AppPort, "3000"), r)
	if err != nil {
		utils.InternalServerError(err)
	}

	log.Println("test...")
}
