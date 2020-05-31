package strings

type Error int

var Errors map[Error]string

const (
	NO_ACTION Error = iota + 1
	PROC_IMPOS
)

func init() {
	Errors = make(map[Error]string)
	Errors[NO_ACTION] = "No such action!"
	Errors[PROC_IMPOS] = "String processing is impossible!"
}
