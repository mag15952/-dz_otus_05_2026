package models

type move interface {
	SetMove(int, string, bool)
}

type MoveStruct struct {
	Movecount int
	Move      string
	History   []string
	Cancel    bool
	player    int
}

func NewMove() MoveStruct {
	return MoveStruct{
		Movecount: 0,
		Move:      "",
		Cancel:    false,
		player:    0,
	}
}

func (m *MoveStruct) GetMove() string {
	return m.Move
}

func (m *MoveStruct) SetMove(movecount int,
	move string, cancel bool,
	player int) {

	m.Movecount = movecount
	m.Move = move
	m.Cancel = cancel
	m.player = player

	m.History = append(m.History, move)

}
