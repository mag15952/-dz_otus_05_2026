package models

type Game struct {
	boardCount int
	moves      []MoveStruct
	board      Chessboard
	players    Players
}

func NewGame(boardCount int) *Game {
	return &Game{
		boardCount: boardCount,
		moves:      make([]MoveStruct, 0),
		board:      Chessboard{},
		players:    Players{},
		//autoMoveChan: make(chan MoveStruct, 100),
	}
}

func (g *Game) SetGameBoardcount(boardCount int) {
	g.boardCount = boardCount
}

func (g *Game) GetGameBoardcount() int {
	return g.boardCount
}

func (g *Game) SetGameBoardArr(arr [][]string) {
	g.board.SetAllBoard(arr)

}

func (g *Game) SetGameBoard(board *Chessboard) {
	g.board = *board

}

func (g *Game) GetGameBoard() *Chessboard {
	return &g.board
}

func (g *Game) SetGamePlayers(players *Players) {
	g.players = *players

}

func (g *Game) GetGamePlayers() Players {
	return g.players

}

func (g *Game) SetGameMoves(move MoveStruct) {
	g.moves = append(g.moves, move)

}

func (g *Game) GetGameMoves() []MoveStruct {
	return g.moves

}
