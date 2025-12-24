package example

import (
	"strings"

	"github.com/google/uuid"
)

// generateUuidWithoutDash 生成不带横杠的 UUID
func GenerateUuidWithoutDash() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
