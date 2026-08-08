package repository2

import (
	"fmt"
	"main/dz7/internal/models"
	"sync"
)

var (
	games     []*models.Game
	muGames   sync.RWMutex
	muPlayers sync.RWMutex
	muMoves   sync.RWMutex
	muBoards  sync.RWMutex
)

func init() {
	games = make([]*models.Game, 0)

}

// Repository интерфейс для работы с хранилищем
type Repo interface {
	SaveObject(entity interface{}, i int)
	GetPlayers(int) models.Players
	GetGame(int) *models.Game
	GetMoves(int) []models.MoveStruct
	GetBoard(int) *models.Chessboard
	Clear()
}

type repository struct{}

func NewRepo() Repo {
	return &repository{}
}

func (r *repository) SaveObject(obj interface{}, i int) {

	muGames.Lock()
	defer muGames.Unlock()

	switch v := obj.(type) {
	case *models.Players:

		games[i].SetGamePlayers(obj.(*models.Players))

	case *models.Game:

		games = append(games, obj.(*models.Game))

	case models.MoveStruct:

		games[i].SetGameMoves(obj.(models.MoveStruct))

	case *models.Chessboard:

		games[i].SetGameBoard(obj.(*models.Chessboard))

	default:
		fmt.Printf("Неизвестный тип: %T (значение: %v)\n", v, obj)
	}
}

func (r *repository) GetPlayers(i int) models.Players {
	muPlayers.RLock()
	defer muPlayers.RUnlock()
	return games[i].GetGamePlayers()
}

func (r *repository) GetGame(i int) *models.Game {
	muGames.RLock()
	defer muGames.RUnlock()
	return games[i]
}

func (r *repository) GetMoves(i int) []models.MoveStruct {
	muMoves.RLock()
	defer muMoves.RUnlock()
	return games[i].GetGameMoves()
}

func (r *repository) GetBoard(i int) *models.Chessboard {
	muBoards.RLock()
	defer muBoards.RUnlock()
	return games[i].GetGameBoard()
}

func (r *repository) Clear() {

	muGames.Lock()
	defer muGames.Unlock()

	games = make([]*models.Game, 0)

}
