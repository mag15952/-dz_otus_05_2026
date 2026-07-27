package models

type Game struct {
	boardCount   int
	moves        []MoveStruct
	isFinished   bool
	isAuto       bool
	autoMoveChan chan MoveStruct
}

func NewGame(boardCount int) *Game {
	return &Game{
		boardCount:   boardCount,
		moves:        make([]MoveStruct, 0),
		isFinished:   false,
		isAuto:       false,
		autoMoveChan: make(chan MoveStruct, 100),
	}
}

func SetGameBoardcount(g Game, boardCount int) {
	g.boardCount = boardCount

}

func (g *Game) GetBoardCount() int {
	return g.boardCount
}
