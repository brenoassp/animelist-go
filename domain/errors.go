package domain

type DomainErr struct {
	error
	Code  string
	Title string
	Data  map[string]interface{}
}

func NewNotFoundErr(title string, data map[string]interface{}) error {
	return DomainErr{Code: "NotFoundErr", Title: title, Data: data}
}

func NewBadRequestErr(title string, data map[string]interface{}) error {
	return DomainErr{Code: "BadRequestErr", Title: title, Data: data}
}

func NewDatabaseErr(title string, data map[string]interface{}) error {
	return DomainErr{Code: "DatabaseErr", Title: title, Data: data}
}
