package repositories

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/m16yusuf/communicare/internal/models"
	"github.com/m16yusuf/communicare/internal/utils"
	"github.com/redis/go-redis/v9"
)

type AuthRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewAuthRepository(db *pgxpool.Pool, rdb *redis.Client) *AuthRepository {
	return &AuthRepository{
		db:  db,
		rdb: rdb,
	}
}

// register new user
func (ar *AuthRepository) CreateUser(cntxt context.Context, email, password string) error {
	sql := `INSERT INTO users (email, password) VALUES ($1, $2)`
	values := []any{email, password}
	cmd, err := ar.db.Exec(cntxt, sql, values...)
	if err != nil {
		log.Println("Failed execute query register\nCause:", err)
		return err
	}
	if cmd.RowsAffected() == 0 {
		log.Println("no row effected when insert users maybe failed?")
		return errors.New("failed register")
	}
	// return nil if not error
	return nil
}

// login registered user
func (ar *AuthRepository) LoginUser(cntxt context.Context, email string) (models.User, error) {
	sql := `SELECT id, email, password FROM users WHERE email=$1`
	var user models.User
	if err := ar.db.QueryRow(cntxt, sql, email).Scan(&user.User_id, &user.Email, &user.Password); err != nil {
		if err == pgx.ErrNoRows {
			return models.User{}, errors.New("user not found")
		}
		log.Println("Internal Server Error.\nCause: ", err.Error())
		return models.User{}, err
	}
	// return user and error nil if success
	return user, nil
}

// BlacklistToken: blacklist user token (logout)
func (ar *AuthRepository) BlacklistToken(c context.Context, token string) error {
	// use utils.BlackListTokenRedish for logout token
	if err := utils.BlackListTokenRedish(c, *ar.rdb, token); err != nil {
		log.Println("failed blacklist token, ", err)
		return err
	}
	// is success return nil
	return nil
}
