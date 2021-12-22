package main

import (
	"log"
	"os"

	"github.com/JEONG-YUNHO01/test-jwt/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	md "github.com/labstack/echo/v4/middleware"
)

func main() {

	// godotenv는 로컬 개발환경에서 .env를 통해 환경변수를 읽어올 때 쓰는 모듈이다.
	// 프로덕션 환경에서는 필요하지 않음.
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	e := echo.New()

	// 회원가입 API
	e.POST("/api/signup", handler.SignUp)

	// 로그인 API(현재는 테스트용)
	e.POST("/api/signin", handler.SignIn)

	// 목데이터로 테스트
	e.GET("/api/getlist", handler.MockData(), md.JWTWithConfig(md.JWTConfig{
		SigningKey:  []byte(os.Getenv("SECRET_KEY")),
		TokenLookup: "cookie:access-token",
	}))

	e.Logger.Fatal(e.Start(":1323"))
}
