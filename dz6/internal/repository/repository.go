package repository

import (
	"fmt"
	"main/dz6/internal/models"
	"sync"
)

var (
	players []*models.Players
	games   []*models.Game
	moves   []models.MoveStruct
	boards  []*models.Chessboard
	mu      sync.RWMutex
)

func init() {
	players = make([]*models.Players, 0)
	games = make([]*models.Game, 0)
	moves = make([]models.MoveStruct, 0)
	boards = make([]*models.Chessboard, 0)
}

// Repository интерфейс для работы с хранилищем
type Repo interface {
	SaveObject(entity interface{}, i int)
	GetPlayers() []*models.Players
	GetGames() []*models.Game
	GetMoves() []models.MoveStruct
	GetBoards() []*models.Chessboard
	Clear()
}

type repository struct{}

func NewRepo() Repo {
	return &repository{}
}

func (r *repository) SaveObject(obj interface{}, i int) {
	mu.Lock()
	defer mu.Unlock()

	switch v := obj.(type) {
	case *models.Players:
		players = append(players, v)
		fmt.Printf("[Repo] Сохранен игрок: %s\n", v.GetPlayer(i))
	case *models.Game:
		games = append(games, v)
		//fmt.Printf("[Repo] Сохранена игра: %s\n", v.String())
	case models.MoveStruct:
		moves = append(moves, v)
		fmt.Printf("[Repo] Сохранен ход: %s\n", v.GetMove())
	case *models.Chessboard:
		boards = append(boards, v)
		fmt.Printf("[Repo] Сохранена доска %dx%d\n", v.GetLen(), v.GetLen())
	default:
		fmt.Printf("Неизвестный тип: %T (значение: %v)\n", v, obj)
	}
}

func (r *repository) GetPlayers() []*models.Players {
	mu.RLock()
	defer mu.RUnlock()
	return players
}

func (r *repository) GetGames() []*models.Game {
	mu.RLock()
	defer mu.RUnlock()
	return games
}

func (r *repository) GetMoves() []models.MoveStruct {
	mu.RLock()
	defer mu.RUnlock()
	return moves
}

func (r *repository) GetBoards() []*models.Chessboard {
	mu.RLock()
	defer mu.RUnlock()
	return boards
}

func (r *repository) Clear() {
	mu.Lock()
	defer mu.Unlock()
	players = make([]*models.Players, 0)
	games = make([]*models.Game, 0)
	moves = make([]models.MoveStruct, 0)
	boards = make([]*models.Chessboard, 0)
}

/*func SetObject(obj any) {

	switch obj_type := obj.(type) {
	case Chessbrd:
		obj_type.SetChessboard()
	case Player:
		obj_type.SetPlayers(2)
	case nil:
		fmt.Println("Получен nil")
	default:
		fmt.Printf("Неизвестный тип: %T (значение: %v)\n", obj_type, obj_type)
	}
}

func PrintObject(obj any, pnum int) {

	switch obj_type := obj.(type) {
	case Chessbrd:
		obj_type.PrintChessboard()
	case Player:
		obj_type.PrintPlayer(pnum)
	case nil:
		fmt.Println("Получен nil")
	default:
		fmt.Printf("Неизвестный тип: %T (значение: %v)\n", obj_type, obj)
	}
}*/
