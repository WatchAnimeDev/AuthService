package interfaces

import "database/sql"

type UidMap struct {
	Id            int
	AnonUid       sql.NullString
	PocketbaseUid sql.NullString
}
