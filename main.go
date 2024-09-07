package main

import (
	"embed"
	"fmt"
	"invoice/config"
	"invoice/pkg/db"
	"net/http"
	"os"
	"time"

	orgHandler "invoice/internal/organisation/delivery/http"
	orgRepo "invoice/internal/organisation/repo"

	outletHandler "invoice/internal/outlet/delivery/http"
	outletRepo "invoice/internal/outlet/repo"

	userHandler "invoice/internal/user/delivery/http"
	userRepo "invoice/internal/user/repo"

	customerHandler "invoice/internal/customer/delivery/http"
	customerRepo "invoice/internal/customer/repo"

	bankAccountHandler "invoice/internal/bank.account/delivery/http"
	bankAccountRepo "invoice/internal/bank.account/repo"

	categoryHandler "invoice/internal/category/delivery/http"
	categoryRepo "invoice/internal/category/repo"

	itemHandler "invoice/internal/item/delivery/http"
	itemRepo "invoice/internal/item/repo"

	currenyHandler "invoice/internal/currency/delivery/http"
	currencyRepo "invoice/internal/currency/repo"

	variantHandler "invoice/internal/variant/delivery/http"
	variantRepo "invoice/internal/variant/repo"

	noteHandler "invoice/internal/note/delivery/http"
	noteRepo "invoice/internal/note/repo"

	inventoryHandler "invoice/internal/inventory/delivery/http"
	inventoryRepo "invoice/internal/inventory/repo"

	invoiceHandler "invoice/internal/invoice/delivery/http"
	invoiceRepo "invoice/internal/invoice/repo"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

//go:embed migrations
var fs embed.FS

func main() {

	logger, _ := zap.NewProduction()
	defer logger.Sync()

	config := config.LoadConfig()

	db, db_err := db.NewDB(config, fs)

	if db_err != nil {
		logger.Error("Error connecting to database", zap.Error(db_err))
		os.Exit(1)
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowWildcard:    true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1", "::1"})

	api := router.Group("/api/v1")

	{
		orgRepo := orgRepo.NewRepo(logger, db)
		orgHandler := orgHandler.NewHandler(api, orgRepo, logger)
		orgHandler.Route()
	}

	{
		outletRepo := outletRepo.NewRepo(logger, db)
		outletHandler := outletHandler.NewHandler(api, outletRepo, logger)
		outletHandler.Route()
	}

	{
		userRepo := userRepo.NewRepo(logger, db)
		userHandler := userHandler.NewHandler(api, userRepo, logger)
		userHandler.Route()
	}

	{
		customerRepo := customerRepo.NewRepo(logger, db)
		customerHandler := customerHandler.NewHandler(api, customerRepo, logger)
		customerHandler.Route()
	}

	{
		bankAccountRepo := bankAccountRepo.NewRepo(logger, db)
		bankAccountHandler := bankAccountHandler.NewHandler(api, bankAccountRepo, logger)
		bankAccountHandler.Route()
	}

	{
		categoryRepo := categoryRepo.NewRepo(logger, db)
		categoryHandler := categoryHandler.NewHandler(api, categoryRepo, logger)
		categoryHandler.Route()
	}

	{
		itemRepo := itemRepo.NewRepo(logger, db)
		itemHandler := itemHandler.NewHandler(api, itemRepo, logger)
		itemHandler.Route()
	}

	{
		currencyRepo := currencyRepo.NewRepo(logger, db)
		currencyHandler := currenyHandler.NewHandler(api, currencyRepo, logger)
		currencyHandler.Route()

		variantRepo := variantRepo.NewRepo(logger, db, currencyRepo)
		variantHandler := variantHandler.NewHandler(api, variantRepo, logger)
		variantHandler.Route()

		noteRepo := noteRepo.NewRepo(logger, db)
		noteHandler := noteHandler.NewHandler(api, noteRepo, logger)
		noteHandler.Route()

		inventoryRepo := inventoryRepo.NewRepo(logger, db, variantRepo)
		inventoryHandler := inventoryHandler.NewHandler(api, inventoryRepo, logger)
		inventoryHandler.Route()

		invoiceBoqRepo := invoiceRepo.BoqNewRepo(logger, db, inventoryRepo, noteRepo)
		invoiceRepo := invoiceRepo.NewRepo(logger, db, inventoryRepo, noteRepo)
		invoiceHandler := invoiceHandler.NewHandler(api, invoiceRepo, logger, invoiceBoqRepo)
		invoiceHandler.Route()
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, map[string]any{
			"status":    "error",
			"message":   "Resource not found",
			"reference": "404",
		})
	})

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8081"
	}
	router.Run(":" + PORT)
	fmt.Println("Application Server Running on: ", PORT)
}
