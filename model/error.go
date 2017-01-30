package model

type apiErrorCode int

type ApiError struct {
	Message		string 		`json:"message,omitempty"`
	Code		apiErrorCode	`json:"code,omitempty"`
}


const (
	ProductNotExistsOrSold = apiErrorCode(1)
)

var codeForMessages = map[apiErrorCode]string {
	ProductNotExistsOrSold : "The product doesn't exist or already has been sold",
}

func ApiErrorForReason(reason apiErrorCode) ApiError {
	return ApiError{codeForMessages[reason], reason}
}