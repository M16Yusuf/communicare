package models

import (
	"mime/multipart"
	"time"
)

type User struct {
	User_id  string `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type Profile struct {
	Id        string    `db:"id"`
	Fullname  *string   `db:"fullname"`
	Avatar    *string   `db:"avatar"`
	Bio       *string   `db:"bio"`
	CreatedAt time.Time `db:"created_at"`
}

type ProfileRequest struct {
	Fullname *string               `json:"fullname" form:"fullname"`
	Avatar   *multipart.FileHeader `form:"profile_picture"`
	Bio      *string               `json:"bio" form:"bio"`
}

type ProfileResponse struct {
	Id        string    `json:"user_id"`
	Fullname  string    `json:"fullname"`
	Avatar    *string   `json:"avatar"`
	Bio       *string   `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
}
