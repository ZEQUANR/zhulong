package logger

import "github.com/gin-gonic/gin"

func CreateLog(c *gin.Context, group ErrorGroup, who ErrorWho, errorType ErrorType, body ErrorBody, err error) {
	e := newError(group, who, errorType, body, err.Error())

	// middlewares.LogError(c, middlewares.LoggerTypeNormal, e)
	c.JSON(errorType.Code, gin.H{"message": e.Error()})
}
