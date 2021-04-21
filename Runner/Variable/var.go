package Variable

import "333exec/Runner/Types"

type Var struct {
	Name  string
	Value interface{}
	Type  Types.Type
}
