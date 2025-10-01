package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/m16yusuf/communicare/internal/handlers"
	"github.com/m16yusuf/communicare/internal/middleware"
	"github.com/m16yusuf/communicare/internal/repositories"
	"github.com/redis/go-redis/v9"
)

func InitUserRouter(router *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	userRouter := router.Group("/users")
	userRepository := repositories.NewUserRepository(db, rdb)
	uh := handlers.NewUserHandler(userRepository)

	userRouter.POST("", middleware.VerifyToken(rdb), uh.UpdateProfile)
}
