package helper

import (
	"fmt"

	"github.com/kodinggo/user-service-gb1/internal/config"
)

// prepare connection string
// charset=utf8mb4 agar dapat menyimpan semua karakter unicode
// parseTime=true agar dapat diparsing dari timestamp ke tipe time.Time
func GetConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		config.GetDbUser(),
		config.GetDbPassword(),
		config.GetDbHost(),
		config.GetDbPort(),
		config.GetDbName(),
	)
}
