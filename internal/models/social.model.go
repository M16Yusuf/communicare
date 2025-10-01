package models

import (
	"mime/multipart"
	"time"
)

type Post struct {
	Id        string    `db:"id"`
	UserId    string    `db:"user_id"`
	Caption   *string   `db:"caption"`
	Photo     *string   `db:"photo"`
	CreatedAt time.Time `db:"created_at"`
}

type PostRequest struct {
	Caption *string               `json:"caption" form:"caption"`
	Photo   *multipart.FileHeader `json:"photo" form:"photo"`
}

type PostDetail struct {
	Post_id    string    `db:"post_id" json:"post_id"`
	Caption    string    `db:"caption" json:"caption"`
	Photo      string    `db:"photo" json:"photo"`
	Created_at time.Time `db:"created_at" json:"created_at"`
	User_id    string    `db:"user_id" json:"user_id"`
	Fullname   string    `db:"fullname" json:"fullname"`
}

type InteracTionPost struct {
	Post_id string `db:"post_id" json:"post_id" binding:"required"`
	// User_id string  `db:"user_id" json:"user_id" binding:"required"`
	Is_like bool    `json:"is_like"`
	Comment *string `db:"comment" json:"comment"`
}
