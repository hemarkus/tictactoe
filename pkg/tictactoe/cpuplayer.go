package tictactoe

type CPUPlayer struct {
	GenericPlayer
}

func NewCPUPlayer(tag TAG) *CPUPlayer {
	return &CPUPlayer{GenericPlayer{tag: tag}}
}

func (cpu *CPUPlayer) RequestMove(board [3][3]TAG) (*Coordinate, error) {
	candidates := map[*Coordinate]int{}
	for _, l := range lanes {
		myTag := 0
		otherTag := 0
		laneCandidates := []*Coordinate{}
		for _, c := range l {
			switch board[c.X][c.Y] {
			case cpu.tag:
				myTag++
			case NO:
				laneCandidates = append(laneCandidates, c)
			default:
				otherTag++
			}
		}
		if len(laneCandidates) == 0 {
			// remove lane
			continue
		}
		// weight win and rescue
		if myTag == 2 {
			for i := 0; i < 16; i++ {
				laneCandidates = append(laneCandidates, laneCandidates[0])
			}
		}
		if otherTag == 2 {
			for i := 0; i < 8; i++ {
				laneCandidates = append(laneCandidates, laneCandidates[0])
			}
		}
		for _, lc := range laneCandidates {
			candidates[lc] = candidates[lc] + 1
		}
	}
	var result *Coordinate
	count := 0
	for k, v := range candidates {
		if v > count {
			result = k
			count = v
		}
	}
	if result == nil {
		return nil, GameOverErr
	}
	return result, nil
}
