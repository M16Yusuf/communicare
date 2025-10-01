package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/m16yusuf/communicare/internal/handlers"
	"github.com/m16yusuf/communicare/internal/middleware"
	"github.com/redis/go-redis/v9"

	docs "github.com/m16yusuf/communicare/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(db *pgxpool.Pool, rdb *redis.Client) *gin.Engine {
	// inisialization engine gin
	router := gin.Default()
	router.Use(middleware.CORSMiddleware)

	// swaggo configuration
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// make directori public accesible
	router.Static("/img", "public")

	// setup routing
	InitAuthRouter(router, db, rdb)
	InitUserRouter(router, db, rdb)

	router.NoRoute(handlers.NoRouteHandler)

	return router
}
