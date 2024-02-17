package routers

import (
	"auth-service/helpers"
	"auth-service/mysql"

	"github.com/gin-gonic/gin"
)

func MapUID(c *gin.Context) {

	errorList := helpers.ValidateUidMapRequest(c)

	if len(errorList) > 0 {
		c.JSON(400, gin.H{
			"errors": errorList,
		})
		return
	}

	db, err := mysql.CreatePool()

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	query := `UPDATE uid_map SET pocketbase_uid = ? WHERE anon_uid = UUID_TO_BIN(?);`
	_, err = db.Exec(query, c.Query("pocketbaseuid"), c.Query("anonuid"))

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	defer db.Close()
	c.JSON(200, gin.H{
		"message": "success",
	})

}
