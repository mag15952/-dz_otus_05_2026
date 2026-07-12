package move

type move interface {
	SetMove(int, string, bool)
}

type MoveStruct struct {
	Movecount int
	Move      string
	History   []string
	Cancel    bool
}

func NewMove() MoveStruct {
	return MoveStruct{
		Movecount: 0,
		Move:      "",
		Cancel:    false,
	}
}

func (m *MoveStruct) SetMove(movecount int, move string, cancel bool) {

	m.Movecount = movecount
	m.Move = move
	m.Cancel = cancel

	m.History = append(m.History, move)

}
