package core

type Consterror string

func (e Consterror) Error() string {
	return string(e)
}

const (
	INVALID_CHAR     = Consterror("Invalid Chaacter")
	Banner_Not_Found = Consterror("Banner Not Found")
)
