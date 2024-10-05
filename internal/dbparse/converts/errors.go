package converts

import (
	"fmt"

	"guthub.com/vloldik/dbml-gen/internal/dbparse/parseobj"
)

func ErrorUnknownSetting(setting parseobj.Setting, where string) error {
	return fmt.Errorf("%s setting is unknown for type %T", setting, where)
}
