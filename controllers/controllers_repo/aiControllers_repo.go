package controller_repo

import "github.com/gin-gonic/gin"

type AiController interface {
	AiQuery() gin.HandlerFunc
}
