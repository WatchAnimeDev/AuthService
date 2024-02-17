package routers

import (
	"auth-service/interfaces"
	"auth-service/mysql"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetUID(c *gin.Context) {
	db, err := mysql.CreatePool()

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	defer db.Close()

	query := `INSERT INTO uid_map(pocketbase_uid) VALUES (?)`
	insertResult, err := db.Exec(query, sql.NullString{})

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	lastInsertID, err := insertResult.LastInsertId()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	var uidmap interfaces.UidMap
	err = db.QueryRow("SELECT um.id,BIN_TO_UUID(um.anon_uid) AS anon_uid, um.pocketbase_uid FROM uid_map um WHERE um.id = ?", lastInsertID).Scan(&uidmap.Id, &uidmap.AnonUid, &uidmap.PocketbaseUid)

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"anonUid": uidmap.AnonUid.String,
		"id":      uidmap.Id,
		"pbUid":   uidmap.PocketbaseUid.String,
	})
}
