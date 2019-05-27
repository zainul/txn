package usecaseerror

var OwnErrors map[string]string

const (
	NotFoundCode             = "E000000"
	InvalidToAccountNumber   = "E000001"
	InvalidFromAccountNumber = "E000002"
	InsufficientBalance      = "E000003"
	FailedToTransfer         = "E000004"
)

func init() {
	errs := make(map[string]string)

	errs[NotFoundCode] = "Unexpected error"
	errs[InvalidFromAccountNumber] = "Invalid from account number"
	errs[InvalidToAccountNumber] = "Invalid to account number"
	errs[InsufficientBalance] = "Insufficient balance"
	errs[FailedToTransfer] = "Failed to transfer from %s to %s"
	OwnErrors = errs
}

func GetCode(err string) string {
	if val, ok := OwnErrors[err]; ok {
		return val
	}

	return OwnErrors[NotFoundCode]
}
