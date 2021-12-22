package handler

import (
	"net/http"
	"time"

	"github.com/JEONG-YUNHO01/test-jwt/db"
	"github.com/JEONG-YUNHO01/test-jwt/helper"
	"github.com/JEONG-YUNHO01/test-jwt/models"

	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}
	db := db.Connect()
	result := db.Find(&user, "email=?", user.Email)

	// 이미 이메일이 존재할 경우의 처리
	if result.RowsAffected != 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "existing email",
		})
	}

	// 비밀번호를 bycrypt 라이브러리로 해싱 처리
	hashpw, err := helper.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	user.Password = hashpw

	// 위의 두단계에서 err가 nil일 경우 DB에 유저를 생성
	if err := db.Create(&user); err.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed SignUp",
		})
	}

	// 모든 처리가 끝난 후 200, Success 메시지를 반환
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success",
	})
}

func SignIn(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}
	inputpw := user.Password

	db := db.Connect()
	result := db.Find(user, "email=?", user.Email)

	// 존재하지않는 아이디일 경우
	if result.RowsAffected == 0 {
		return echo.ErrBadRequest
	}

	res := helper.CheckPasswordHash(user.Password, inputpw)

	// 비밀번호 검증에 실패한 경우
	if !res {
		return echo.ErrUnauthorized
	}
	// 토큰 발행
	accessToken, err := helper.CreateJWT(user.Email)
	if err != nil {
		return echo.ErrInternalServerError
	}

	cookie := new(http.Cookie)
	cookie.Name = "access-token"
	cookie.Value = accessToken
	cookie.HttpOnly = true
	cookie.Expires = time.Now().Add(time.Hour * 24)

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login Success",
	})
}

// 테스트용 API
func MockData() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Mock Data를 생성한다.
		list := map[string]string{
			"1": "고양이",
			"2": "사자",
			"3": "호랑이",
		}
		return c.JSON(http.StatusOK, list)
	}
}
