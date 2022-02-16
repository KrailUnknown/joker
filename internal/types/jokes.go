package types

type JokeType struct {
	currentType string
}

// type JokeType interface {
// 	SetToDadType()
// 	SetToChuckType()
// 	GetType() string
// }

const DAD_JOKE_TYPE = "dad"
const CHUCK_JOKE_TYPE = "chuck"

func (jk *JokeType) SetToDadType() {
	jk.currentType = DAD_JOKE_TYPE
}

func (jk *JokeType) SetToChuckType() {
	jk.currentType = CHUCK_JOKE_TYPE
}

func (jk *JokeType) GetType() string {
	return jk.currentType
}
