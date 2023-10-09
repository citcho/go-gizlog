package main

import (
	"fmt"
	"log"

	"github.com/citcho/go-gizlog/internal/auth/application"
	auth_infra_user "github.com/citcho/go-gizlog/internal/auth/infrastructure/user"
	auth_interface_http "github.com/citcho/go-gizlog/internal/auth/interface/public/http"
	"github.com/citcho/go-gizlog/internal/common/auth"
	"github.com/citcho/go-gizlog/internal/common/clock"
	"github.com/citcho/go-gizlog/internal/common/config"
	"github.com/citcho/go-gizlog/internal/common/database"
	"github.com/citcho/go-gizlog/internal/common/server"
	report_app "github.com/citcho/go-gizlog/internal/report/application"
	"github.com/citcho/go-gizlog/internal/report/domain/report"
	report_infra "github.com/citcho/go-gizlog/internal/report/infrastructure/report"
	report_interface_http "github.com/citcho/go-gizlog/internal/report/interface/public/http"
	user_app "github.com/citcho/go-gizlog/internal/user/application"
	"github.com/citcho/go-gizlog/internal/user/domain/user"
	user_infra "github.com/citcho/go-gizlog/internal/user/infrastructure/user"
	user_interface_intraprocess "github.com/citcho/go-gizlog/internal/user/interface/private/intraprocess"
	user_interface_http "github.com/citcho/go-gizlog/internal/user/interface/public/http"
)

func main() {
	log.Println("Starting monolith service")

	cfg, err := config.NewDBConfig()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	db := database.NewDB(*cfg)

	jwter, err := auth.NewJWTer(clock.RealClocker{})
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	userRepository := user_infra.NewMySQLRepository(db)
	userService := user.NewUserService(userRepository)
	userUsecase := user_app.NewUserUsecase(userService, userRepository)
	userHttpController := user_interface_http.NewUserController(userUsecase)

	userIntraprocessController := user_interface_intraprocess.NewIntraprocessController(userRepository)
	authUsecase := application.NewAuthUsecase(
		auth_infra_user.NewIntraprocessService(userIntraprocessController), jwter,
	)
	authController := auth_interface_http.NewAuthController(authUsecase)

	reportRepository := report_infra.NewMySQLRepository(db)
	reportService := report.NewReportService(reportRepository)
	reportUsecase := report_app.NewReportUsecase(reportService, reportRepository)
	reportController := report_interface_http.NewReportController(reportUsecase)

	s := server.NewServer()

	user_interface_http.AddRoutes(s, *userHttpController)
	auth_interface_http.AddRoutes(s, *authController)
	report_interface_http.AddRoutes(s, jwter, *reportController)

	s.Logger.Fatal(s.Start(":" + fmt.Sprintf("%d", 8080)))
}
