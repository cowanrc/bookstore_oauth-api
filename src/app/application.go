package app

import (
	"bookstore_oauth-api/src/domain/access_token"
	"bookstore_oauth-api/src/httpPkg"
	"bookstore_oauth-api/src/repository/db"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := httpPkg.NewHandler(access_token.NewService(db.New()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run(":8080")
}
