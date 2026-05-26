package core

type Consterror string

func (e Consterror) Error() string {
	return string(e)
}

const (
	INVALID_CHAR = Consterror("Invalid Chaacter")
)
