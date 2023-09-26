package schemes

type Responses struct {
	StatusCode int         `json:"code" example:"200"`
	Message    string      `json:"message" example:"status OK"`
	Data       interface{} `json:"data"`
}

type ResponsesPagination struct {
	StatusCode int         `json:"code" example:"200"`
	Message    string      `json:"message" example:"status OK"`
	Page       int         `json:"page" example:"1"`
	PerPage    int         `json:"per_page" example:"10"`
	TotalPage  int         `json:"total_page" example:"5"`
	TotalData  int         `json:"total_data" example:"50"`
	Data       interface{} `json:"data"`
}

type Responses201Example struct {
	StatusCode int         `json:"code" example:"201"`
	Message    string      `json:"message" example:"Create Successfully"`
	Data       interface{} `json:"data"`
}

type Responses400Example struct {
	StatusCode int         `json:"code" example:"400"`
	Message    string      `json:"message" example:"status bad request"`
	Data       interface{} `json:"data"`
}

type Responses401Example struct {
	StatusCode int    `json:"code" example:"401"`
	Message    string `json:"message" example:"Authorization header is required"`
}

type Responses403Example struct {
	StatusCode int         `json:"code" example:"403"`
	Message    string      `json:"message" example:"status forbidden"`
	Data       interface{} `json:"data"`
}

type Responses404Example struct {
	StatusCode int         `json:"code" example:"404"`
	Message    string      `json:"message" example:"status not found"`
	Data       interface{} `json:"data"`
}

type Responses409Example struct {
	StatusCode int         `json:"code" example:"409"`
	Message    string      `json:"message" example:"status conflict data"`
	Data       interface{} `json:"data"`
}

type Responses500Example struct {
	StatusCode int         `json:"code" example:"500"`
	Message    string      `json:"message" example:"status internal error"`
	Data       interface{} `json:"data"`
}
