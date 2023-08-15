package gouuid

import uuid "github.com/satori/go.uuid"

//goland:noinspection GoUnusedExportedFunction
func V4() string {
	return uuid.NewV4().String()
}
