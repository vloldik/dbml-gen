package genutil

import (
	"strings"
)

func GormTagsFromList(tags ...string) map[string]string {
	return map[string]string{"gorm": strings.Join(tags, ";")}
}
