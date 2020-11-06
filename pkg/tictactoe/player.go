package tictactoe

type Player interface {
	RequestMove(board [3][3]TAG) (*Coordinate, error)
	GetTag() TAG
	GetName() string
}

type GenericPlayer struct {
	tag  TAG
	name string
}

func (gp *GenericPlayer) GetTag() TAG {
	return gp.tag
}

func (gp *GenericPlayer) GetName() string {
	return gp.name
}
