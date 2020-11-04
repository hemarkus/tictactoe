package tictactoe

type Player interface {
	RequestMove(board [3][3]TAG) (*Coordinate, error)
	GetTag() TAG
}

type GenericPlayer struct {
	tag TAG
}

func (gp *GenericPlayer) GetTag() TAG {
	return gp.tag
}
