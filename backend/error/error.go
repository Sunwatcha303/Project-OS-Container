package error

import "errors"

type TemplateError interface {
	Code() string
	Error() string
	HttpStatusCode() int
	ThaiDescription() string
	EnglishDescription() string
	Response() (int, HttpErrorResponse)
}

type HttpError struct {
	error string
}

type HttpErrorResponse struct {
	Code               string `json:"code"`
	Error              string `json:"error"`
	ThaiDescription    string `json:"thai_description"`
	EnglishDescription string `json:"english_description"`
}

func New(error string) *HttpError {
	return &HttpError{
		error: error,
	}
}

func InitError(err error) error {
	if templateError, ok := err.(TemplateError); ok {
		return templateError
	} else {
		return errors.New(err.Error())
	}
}

func (e *HttpError) Code() string {
	if errorCode, ok := CodesError[e]; ok {
		return errorCode
	} else {
		return InternalServerError.Error()
	}
}

func (e *HttpError) Error() string {
	return e.error
}

func (e *HttpError) ThaiDescription() string {
	if thaiDescription, ok := ThaiDescription[e]; ok {
		return thaiDescription
	} else {
		return InternalServerError.ThaiDescription()
	}
}

func (e *HttpError) EnglishDescription() string {
	if englishDescription, ok := EnglishDescription[e]; ok {
		return englishDescription
	} else {
		return InternalServerError.EnglishDescription()
	}
}

func (e *HttpError) HttpStatusCode() int {
	if httpStatusCode, ok := HttpStatusCodes[e]; ok {
		return httpStatusCode
	} else {
		return InternalServerError.HttpStatusCode()
	}
}

func (e *HttpError) Response() (int, HttpErrorResponse) {
	if httpStatusCode, ok := HttpStatusCodes[e]; ok {
		return httpStatusCode, HttpErrorResponse{
			Code:               e.Code(),
			Error:              e.Error(),
			ThaiDescription:    e.ThaiDescription(),
			EnglishDescription: e.EnglishDescription(),
		}
	} else {
		return InternalServerError.HttpStatusCode(), HttpErrorResponse{
			Code:               InternalServerError.Code(),
			Error:              InternalServerError.Error(),
			ThaiDescription:    InternalServerError.ThaiDescription(),
			EnglishDescription: InternalServerError.EnglishDescription(),
		}
	}
}

func FindError(errorName string) (err error) {
	for err, _ := range EnglishDescription {
		if err.Error() == errorName {
			return err
		}
	}
	return InternalServerError
}

var (
	InternalServerError TemplateError = New("internal_server_error")
	ApiKeyError         TemplateError = New("api_key_error")
)

var HttpStatusCodes = map[error]int{
	InternalServerError: 500,
	ApiKeyError:         401,
}

var CodesError = map[error]string{
	InternalServerError: "1000",
	ApiKeyError:         "1001",
}

var ThaiDescription = map[error]string{
	InternalServerError: "เกิดข้อผิดพลาดที่เซิฟเวอร์",
	ApiKeyError:         "ไม่มีสิทธ์ในการร้องขอ",
}

var EnglishDescription = map[error]string{
	InternalServerError: "There was an error on the server,",
	ApiKeyError:         "Invalid credentials",
}

func GetErrorResponse(err error) (int, HttpErrorResponse) {
	var mapError bool
	var httpStatusCode int
	var errorCode string
	var thaiDescription string
	var englishDescription string

	if httpStatusCode, mapError = HttpStatusCodes[err]; mapError {
		if errorCode, mapError = CodesError[err]; mapError {
			if thaiDescription, mapError = ThaiDescription[err]; mapError {
				if englishDescription, mapError = EnglishDescription[err]; mapError {
					return httpStatusCode, HttpErrorResponse{
						Code:               errorCode,
						Error:              err.Error(),
						ThaiDescription:    thaiDescription,
						EnglishDescription: englishDescription,
					}
				}
			}
		}
	}
	return InternalServerError.HttpStatusCode(), HttpErrorResponse{
		Code:               InternalServerError.Code(),
		Error:              InternalServerError.Error(),
		ThaiDescription:    InternalServerError.ThaiDescription(),
		EnglishDescription: InternalServerError.EnglishDescription(),
	}
}
