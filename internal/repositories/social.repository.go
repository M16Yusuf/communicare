package repositories

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type SocialRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewSocialRepository(db *pgxpool.Pool, rdb *redis.Client) *SocialRepository {
	return &SocialRepository{db: db, rdb: rdb}
}

// follow a user
func (sr *SocialRepository) FollowUser(cntxt context.Context, user_following, user_follower string) error {
	sql := `INSERT INTO follows (user_id, follower_id) VALUES($1, $2)`
	values := []any{user_following, user_follower}
	cmd, err := sr.db.Exec(cntxt, sql, values...)
	if err != nil {
		log.Println("Failed execute query follow\nCause:", err)
		return err
	}
	if cmd.RowsAffected() == 0 {
		log.Println("no row effected when insert users maybe failed?")
		return errors.New("failed follow user")
	}

	// return error nil is success
	return nil
}
