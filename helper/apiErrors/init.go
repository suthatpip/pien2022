package apiErrors

type apiError struct {
	Id      string
	Message string
	Code    int
	Status  int
	Detail  string
}

func (e *apiError) Error() string {
	return e.Message
}

func NewError() *apiError {
	return &apiError{}
}

// Use for test, parse to ApiError...
type ApiError apiError

// Use for Error API
var ApiErrors []apiError

func init() {

	ApiErrors = append(ApiErrors, apiErrors...)
}

func cloneError(e *apiError) *apiError {
	newError := *e
	return &newError
}

// Use for Error API
func findErrorById(errorId string) *apiError {
	for index := range ApiErrors {
		if ApiErrors[index].Id == errorId {
			return cloneError(&ApiErrors[index])
		}
	}
	return nil
}

func ThrowError(errorId string) *apiError {
	if err := findErrorById(errorId); err != nil {
		return err
	}
	panic("Error To Throw Not Defined")
}

func ParseError(err error) *apiError {
	if parseError, ok := err.(*apiError); ok {
		return parseError
	}
	return nil
}
