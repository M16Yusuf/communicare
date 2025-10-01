package repositories

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/m16yusuf/communicare/internal/models"
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

// get list post from followed user
func (sr *SocialRepository) GetPost(cntxt context.Context, user_id string) ([]models.PostDetail, error) {

	sql := `SELECT p.id AS post_id, p.caption, p.photo, p.created_at, u.id AS user_id, u.fullname
		FROM posts p
		INNER JOIN users u ON p.user_id = u.id
		INNER JOIN follows f ON p.user_id = f.user_id 
		WHERE f.follower_id =  $1
		ORDER BY p.created_at DESC; `

	rows, err := sr.db.Query(cntxt, sql, user_id)
	if err != nil {
		log.Println("failed execute query \ncause: ", err)
		return []models.PostDetail{}, err
	}
	defer rows.Close()

	// processsing data rows into slice
	var posts []models.PostDetail

	for rows.Next() {
		var post models.PostDetail
		if err := rows.Scan(&post.Post_id, &post.Caption, &post.Photo, &post.Created_at, &post.User_id, &post.Fullname); err != nil {
			log.Println("Scan Error, ", err.Error())
			return []models.PostDetail{}, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (sr *SocialRepository) LikeAndOrCommentPost(cntxt context.Context, userID string, body models.InteracTionPost) error {

	if body.Is_like {
		sqlLike := `INSERT INTO likes ( post_id, user_id) VALUES( $1, $2 )`
		values := []any{body.Post_id, userID}
		cmd, err := sr.db.Exec(cntxt, sqlLike, values...)
		if err != nil {
			log.Println("Failed execute query like post\nCause:", err)
			return err
		}
		if cmd.RowsAffected() == 0 {
			log.Println("no row effected when insert like post maybe failed?")
			return errors.New("failed like post")
		}
	}

	if body.Comment != nil {
		sqlComment := `INSERT INTO comments ( post_id, user_id, comment ) VALUES ( $1, $2, $3 )`
		values := []any{body.Post_id, userID, body.Comment}

		cmd, err := sr.db.Exec(cntxt, sqlComment, values...)
		if err != nil {
			log.Println("Failed execute query comment post\nCause:", err)
			return err
		}
		if cmd.RowsAffected() == 0 {
			log.Println("no row effected when insert comment maybe failed?")
			return errors.New("failed comment post")
		}
	}

	// return error nil is success
	return nil
}
