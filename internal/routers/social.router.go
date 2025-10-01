package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/m16yusuf/communicare/internal/handlers"
	"github.com/m16yusuf/communicare/internal/middleware"
	"github.com/m16yusuf/communicare/internal/repositories"
	"github.com/redis/go-redis/v9"
)

func InitSocialRouter(router *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	socialRouter := router.Group("/social")
	socialRepository := repositories.NewSocialRepository(db, rdb)
	sh := handlers.NewSocialHandler(socialRepository)

	socialRouter.POST("/follow/:user_id", middleware.VerifyToken(rdb), sh.FollowAUser)
}
