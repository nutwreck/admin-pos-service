package schemes

type SchemeDatabaseError struct {
	Type string
	Code int
}

type ErrorResponse struct {
	StatusCode int         `json:"code"`
	Error      interface{} `json:"error"`
}

type UnathorizatedError struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Struktur data untuk membaca JSON Error
type ReadMsgErrorValidator struct {
	Results struct {
		Errors []map[string]struct {
			Message string `json:"message"`
			Value   string `json:"value"`
			Param   string `json:"param"`
			Tag     string `json:"tag"`
		} `json:"errors"`
	} `json:"results"`
}

// Struktur data untuk menyimpan hasil array list Error
type ResultMsgErrorValidator struct {
	Message string `json:"message"`
	Value   string `json:"value"`
	Param   string `json:"param"`
	Tag     string `json:"tag"`
}
