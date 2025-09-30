package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/m16yusuf/communicare/internal/models"
	"github.com/m16yusuf/communicare/internal/repositories"
	"github.com/m16yusuf/communicare/internal/utils"
	"github.com/m16yusuf/communicare/pkg"
)

type AuthHandler struct {
	ar *repositories.AuthRepository
}

func NewAuthHandler(ar *repositories.AuthRepository) *AuthHandler {
	return &AuthHandler{ar: ar}
}

func (a *AuthHandler) Login(ctx *gin.Context) {
	var body models.AuthRequest
	if err := ctx.ShouldBind(&body); err != nil {
		// check if failed binding bcs input not match with model require
		log.Println("error when binding \ncause", err)
		if strings.Contains(err.Error(), "required") {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
				Response: models.Response{
					IsSuccess: false,
					Code:      400,
				},
				Err: "Email or password cannot be empty",
			})
			return
		}
		// else binding error because server, log the error
		log.Println("Internal Server Error.\nCause: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Response: models.Response{
				IsSuccess: false,
				Code:      500,
			},
			Err: "internal server error",
		})
		return
	}

	// get user data
	user, err := a.ar.LoginUser(ctx.Request.Context(), body.Email)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
				Response: models.Response{
					IsSuccess: false,
					Code:      400,
				},
				Err: "Email or Password is incorrect",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Response: models.Response{
				IsSuccess: false,
				Code:      500,
			},
			Err: "internal server error",
		})
		return
	}

	// compare password
	// body.password => from http body / input user
	// user.Password => from query GetUserWithEmail
	hc := pkg.NewHashConfig()
	isMatched, err := hc.CompareHashAndPassword(body.Password, user.Password)
	if err != nil {
		log.Println("Internal Server Error.\nCause: ", err.Error())
		re := regexp.MustCompile("hash|crypto|argon2id|format")
		if re.Match([]byte(err.Error())) {
			log.Println("Error during Hashing")
		}
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Response: models.Response{
				IsSuccess: false,
				Code:      500,
			},
			Err: "internal server error",
		})
		return
	}
	// if not match sen https status as response
	if !isMatched {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Response: models.Response{
				IsSuccess: false,
				Code:      400,
			},
			Err: "Email or Password is incorrect",
		})
		return
	}
	// If match, generate jwt token and send as response
	claim := pkg.NewJWTClaims(user.User_id, "user")
	jwtToken, err := claim.GenToken()
	if err != nil {
		log.Println("Internal Server Error.\nCause: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Response: models.Response{
				IsSuccess: false,
				Code:      500,
			},
			Err: "internal server error",
		})
		return
	}
	// return token as response success
	ctx.JSON(http.StatusOK, models.LoginResponse{
		Response: models.Response{
			IsSuccess: true,
			Code:      http.StatusOK,
			Msg:       "login successfully",
		},
		Token: jwtToken,
	})
}

func (a *AuthHandler) Register(ctx *gin.Context) {
	var body models.AuthRequest

	// Binding data and show if there is error when binding data
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Response: models.Response{
				IsSuccess: false,
				Code:      500,
			},
			Err: "Failed binding data ...",
		})
		return
	}

	// validate register
	if err := utils.RegisterValidation(body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Response: models.Response{
				IsSuccess: false,
				Code:      400,
			},
			Err: err.Error(),
		})
		return
	} else {
		// hash new password??
		hc := pkg.NewHashConfig()
		hc.UseRecommended()
		hash, err := hc.GenHash(body.Password)
		if err != nil {
			log.Println("Failed hash new password ...")
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Response: models.Response{
					IsSuccess: false,
					Code:      500,
				},
				Err: err.Error(),
			})
			return
		}
		// create user to database
		if err := a.ar.CreateUser(ctx.Request.Context(), body.Email, hash); err != nil {
			re := regexp.MustCompile("duplicate|unique")
			if re.Match([]byte(err.Error())) {
				ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
					Response: models.Response{
						IsSuccess: false,
						Code:      http.StatusBadRequest,
					},
					Err: "Email already registered",
				})
				return
			}
			log.Println("internal server error", err)
			ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
				Response: models.Response{
					IsSuccess: false,
					Code:      http.StatusInternalServerError,
				},
				Err: "internal server error",
			})
			return
		}
		ctx.JSON(http.StatusOK, models.Response{
			IsSuccess: true,
			Code:      http.StatusOK,
			Msg:       "User registered successfully",
		})
	}
}

func (a *AuthHandler) Logout(ctx *gin.Context) {
	// get token user for logout
	bearerToken := ctx.GetHeader("Authorization")

	if err := a.ar.BlacklistToken(ctx.Request.Context(), bearerToken); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Response: models.Response{
				IsSuccess: false,
				Code:      500,
			},
			Err: err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, models.Response{
			IsSuccess: true,
			Code:      200,
			Msg:       "Logout successfully",
		})
	}
}
