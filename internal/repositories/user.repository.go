package repositories

import (
	"context"
	"log"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/m16yusuf/communicare/internal/models"
	"github.com/redis/go-redis/v9"
)

type UserRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewUserRepository(db *pgxpool.Pool, rdb *redis.Client) *UserRepository {
	return &UserRepository{db: db, rdb: rdb}
}

// update profile
func (ur *UserRepository) UpdateProfileUser(cntxt context.Context, userProfile models.Profile) (models.ProfileResponse, error) {
	sql := `UPDATE users SET updated_at=CURRENT_TIMESTAMP `

	values := []any{}
	if userProfile.Fullname != nil {
		idx := strconv.Itoa(len(values) + 1)
		sql += ", fullname=$" + idx + ""
		values = append(values, userProfile.Fullname)
	}
	if userProfile.Avatar != nil {
		idx := strconv.Itoa(len(values) + 1)
		sql += ", avatar=$" + idx + ""
		values = append(values, userProfile.Avatar)
	}
	if userProfile.Bio != nil {
		idx := strconv.Itoa(len(values) + 1)
		sql += ", bio=$" + idx + ""
		values = append(values, userProfile.Bio)
	}

	idx := strconv.Itoa(len(values) + 1)
	sql += " WHERE id=$" + idx + " RETURNING id, fullname, avatar, bio"
	values = append(values, userProfile.Id)

	var newProfile models.ProfileResponse
	if err := ur.db.QueryRow(cntxt, sql, values...).Scan(&newProfile.Id, &newProfile.Fullname, &newProfile.Avatar, &newProfile.Bio); err != nil {
		log.Println("scan Error. ", err.Error())
		return models.ProfileResponse{}, err
	}

	// return error nil if not error
	return newProfile, nil
}
