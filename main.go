package main

import (
	"log"
	"net/http"
	"rest_api/app"
	"rest_api/config"
	"rest_api/controller"
	"rest_api/repository"
	"rest_api/service"
	"rest_api/utils"

	_ "rest_api/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

// var (
// 	db          *gorm.DB                          = config.SetupDatabaseConnection()
// 	catRepo     repository.CategoryGormRepository = repository.NewCategoryGormRepository(db)
// 	catService  service.CategoryService           = service.NewCategoryService(catRepo)
// 	cat         controller.CategoryController     = controller.NewCategoryController(catService, validator.New())
// 	userRepo    repository.UserRepository         = repository.NewUserRepository(db)
// 	userService service.UserService               = service.NewUserService(userRepo)
// 	user        controller.UserController         = controller.NewUserController(userService)
// 	err         error
// )

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
	db := config.SetupDatabaseConnection()
	defer config.CloseDatabaseConnection(db)
	catRepo := repository.NewCategoryRepository(db)
	catService := service.NewCategoryService(catRepo)
	cat := controller.NewCategoryController(catService, validator.New())
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	user := controller.NewUserController(userService)

	// var db *gorm.DB = config.SetupDatabaseConnection()
	// fmt.Printf(" %+v", db)
	// db, err := sqlx.Connect("mysql", "root:@(localhost:3306)/rest-pzn")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// db.SetMaxOpenConns(200)
	// db.SetMaxIdleConns(10)

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
		// 	// r.Post("/login1", user.Login())
	})
	r.Route("/api/v1", func(r chi.Router) {
		app.NewCategoryRouter(r, cat)
		app.NewUserRouter(r, user)

	})
	// r.Get("/swagger/*", httpSwagger.Handler(
	// 	httpSwagger.URL("http://localhost:3000/swagger/doc.json"), //The url pointing to API definition
	// ))
	log.Println("Running server at port " + utils.CekNilParameter(utils.EnvConfigs.AppPort, "3000"))

	err := http.ListenAndServe(":"+utils.CekNilParameter(utils.EnvConfigs.AppPort, "3000"), r)
	if err != nil {
		utils.InternalServerError(err)
	}

	log.Println("test...")
}
