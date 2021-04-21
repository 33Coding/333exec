package Code

import (
	"333exec/Runner/Variable"
	"errors"
	"fmt"
)

func LogPrintln(v Variable.Var) (string, error) {
	if v.Type.Name == "string" {
		return "log.Println(`" + fmt.Sprintf("%v", v.Value) + "`)", nil
	}
	return "", errors.New("invalid type: log.Println expects a string")
}
