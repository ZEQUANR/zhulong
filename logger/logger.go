package logger

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateLog(c *gin.Context, who ErrorWho, errorCause ErrorAction, effect ErrorBody, err error) {
	e := fmt.Errorf(fmt.Sprintf("%s|%s|%s|[%s]", who, errorCause, effect, err.Error()))

	// middlewares.LogError(c, middlewares.LoggerTypeNormal, e)
	c.JSON(200, gin.H{"message": e.Error()})
}
