package player

import (
	"fmt"

	"golang.org/x/exp/rand"
)

// Position представляет позицию игрока на поле
type Position string

const (
	Forward    Position = "Forward"
	Midfielder Position = "Midfielder"
	Defender   Position = "Defender"
	Goalkeeper Position = "Goalkeeper"
)

// Player представляет игрока в футбольной команде
type Player struct {
	Name         string
	Position     Position
	AttackSkill  int // Навыки атаки
	DefenseSkill int // Навыки защиты
	Goals        int // Забитые голы
	Fouls        int // Совершенные фолы
	YellowCards  int // Желтые карточки
	RedCards     int // Красные карточки
}

// NewPlayer создает нового игрока с заданными характеристиками
func NewPlayer(name string, position Position, attackSkill, defenseSkill int) *Player {
	return &Player{
		Name:         name,
		Position:     position,
		AttackSkill:  attackSkill,
		DefenseSkill: defenseSkill,
	}
}

// DisplayStats выводит статистику игрока
func (p *Player) DisplayStats() {
	fmt.Printf("%s (%s) - Голы: %d, Фолы: %d, Желтые карточки: %d, Красные карточки: %d\n",
		p.Name, p.Position, p.Goals, p.Fouls, p.YellowCards, p.RedCards)
}

// ScoreGoal увеличивает количество забитых голов игроком
func (p *Player) ScoreGoal() {
	p.Goals++
}

// CommitFoul увеличивает количество фолов и добавляет карточки
func (p *Player) CommitFoul() {
	p.Fouls++
	if p.YellowCards < 1 {
		p.YellowCards++
		fmt.Printf("%s получает желтую карточку.\n", p.Name)
	} else {
		p.RedCards++
		fmt.Printf("%s получает красную карточку и удаляется с поля.\n", p.Name)
	}
}

// GenerateRandomTeam создает случайную команду из 11 игроков
func GenerateRandomTeam(teamName string) []*Player {
	positions := []Position{Forward, Midfielder, Defender, Goalkeeper}
	players := []*Player{}

	for i := 0; i < 11; i++ {
		position := positions[i%len(positions)]
		player := NewPlayer(
			fmt.Sprintf("%s Player %d", teamName, i+1),
			position,
			rand.Intn(100), // Случайные навыки атаки
			rand.Intn(100), // Случайные навыки защиты
		)
		players = append(players, player)
	}

	return players
}
