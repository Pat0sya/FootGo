package team

import (
	"Footballsim/pkg/player"
	"fmt"
)

// Team представляет футбольную команду
type Team struct {
	Name    string
	Players []*player.Player
	Wins    int
	Losses  int
	Draws   int
}

// NewTeam создает новую команду с заданным именем и случайными игроками
func NewTeam(name string) *Team {
	return &Team{
		Name:    name,
		Players: player.GenerateRandomTeam(name),
	}
}

// AddPlayer добавляет игрока в команду
func (t *Team) AddPlayer(p *player.Player) {
	t.Players = append(t.Players, p)
}

// DisplayTeamStats выводит статистику команды и каждого игрока
func (t *Team) DisplayTeamStats() {
	fmt.Printf("|---------|Команда: %s|---------|\n", t.Name)
	fmt.Printf("|Победы: %d|  Поражения: %d   | Ничьи: %d|\n", t.Wins, t.Losses, t.Draws)
	fmt.Println("Игроки:")
	for _, player := range t.Players {
		player.DisplayStats()
	}
}

// RecordWin увеличивает счетчик побед команды
func (t *Team) RecordWin() {
	t.Wins++
}

// RecordLoss увеличивает счетчик поражений команды
func (t *Team) RecordLoss() {
	t.Losses++
}

// RecordDraw увеличивает счетчик ничьих команды
func (t *Team) RecordDraw() {
	t.Draws++
}
