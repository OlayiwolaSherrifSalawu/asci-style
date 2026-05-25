package api

type Consterror string

func (e Consterror) Error() string {
	return string(e)
}

const (
	INTERNAL_SERVER_ERROR = Consterror("An Internal Server Error Occured")
	FAILED_TO_LOAD_ENV    = Consterror("Failed to load env")
)
