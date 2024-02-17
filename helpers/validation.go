package helpers

import "github.com/gin-gonic/gin"

func ValidateUidMapRequest(c *gin.Context) []string {
	errorList := []string{}
	if c.Query("anonuid") == "" {
		errorList = append(errorList, "anonuid is required")
	}
	if c.Query("pocketbaseuid") == "" {
		errorList = append(errorList, "pocketbaseuid is required")
	}
	return errorList
}
