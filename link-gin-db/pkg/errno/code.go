package errno

import "fmt"

var (
	_codes = map[int]struct{}{}
)

func New(e int) int {
	if e <= 0 {
		panic("business ecode must greater than zero")
	}
	return add(e)
}

func add(e int) int {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", e))
	}
	_codes[e] = struct{}{}
	return e
}

type Errno struct {
	Code int         `json:"Code"`
	Msg  string      `json:"Msg"`
	Data interface{} `json:"Data,omitempty"`
}

var (
	Success = &Errno{Code: New(1), Msg: "ok"}

	ErrParams     = &Errno{Code: New(1000), Msg: "param error"}
	ErrSystem     = &Errno{Code: New(1001), Msg: "Internal server error"}
	ErrDatabase   = &Errno{Code: New(1002), Msg: "database error"}
	ErrValidation = &Errno{Code: New(1003), Msg: "validation error"}
	ErrJson       = &Errno{Code: New(1004), Msg: "json error"}
	ErrBind       = &Errno{Code: New(1005), Msg: "bind json error"}
	ErrEncrypt    = &Errno{Code: New(1006), Msg: "encrypt password error"}
	ErrCompare    = &Errno{Code: New(1007), Msg: "Compare password error"}
	ErrSignToken  = &Errno{Code: New(1008), Msg: "sign token error"}
	ErrCheckToken = &Errno{Code: New(1009), Msg: "check token error"}

	ErrUserNotLogin = &Errno{Code: New(1020), Msg: "user not login error"}
	ErrUserNotFound = &Errno{Code: New(1021), Msg: "user not find error"}
	ErrUserHadFound = &Errno{Code: New(1022), Msg: "user had found error"}
)
