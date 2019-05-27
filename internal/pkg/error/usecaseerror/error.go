package usecaseerror

var OwnErrors map[string]string

const (
	NotFoundCode = "E000000"
)

func NewError() {
	errs := make(map[string]string)

	errs[NotFoundCode] = "00000"
	OwnErrors = errs
}

func GetCode(err string) string {
	if val, ok := OwnErrors[err]; ok {
		return val
	}

	return NotFoundCode
}
