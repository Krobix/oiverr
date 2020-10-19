package oliverr

/* This is the types "enum" */
const (
	IntT = iota
	StrT
	FloatT
	FptrT //function pointer
	DictT //dictionary type; basically a map
)

/*Object represents all objects in general */
type Object interface {
	/*Conversion methods*/
	getStr() string
	getInt() int64
	getFloat() float64
	getFptr() int64
	getDict() map[Object]Object //Will not work for all types, probably only for dict type itself
}

/* Individual type structs */

type StringObject struct {
	value string
}

type IntObject struct {
	value int64
}

type FloatObject struct {
	value float64
}

