package match

import (
	"Footballsim/pkg/player"
	"Footballsim/pkg/team"
	"fmt"
	"math/rand"
)

// Event представляет событие, которое может произойти во время матча
type Event struct {
	Type   string
	Minute int
	Team   *team.Team
	Player *player.Player
}

// Match представляет футбольный матч
type Match struct {
	Team1      *team.Team
	Team2      *team.Team
	Duration   int
	IsFinished bool
	Events     []Event // Список событий матча
}

// NewMatch создает новый матч между двумя командами с заданной продолжительностью
func NewMatch(team1Name, team2Name string, duration int) *Match {
	return &Match{
		Team1:    team.NewTeam(team1Name),
		Team2:    team.NewTeam(team2Name),
		Duration: duration,
	}
}

func (m *Match) Start() {
	fmt.Printf("Матч начался: %s против %s\n", m.Team1.Name, m.Team2.Name)
	rand.New(rand.NewSource(12121212121212))

	for minute := 1; minute <= m.Duration; minute++ {
		m.simulateMinute(minute)
		if m.IsFinished {
			break
		}
	}

	score1 := m.Score1()
	score2 := m.Score2()
	fmt.Printf("Матч завершен! Результат: %s %d:%d %s\n", m.Team1.Name, score1, score2, m.Team2.Name)

	if score1 > score2 {
		m.Team1.RecordWin()
		m.Team2.RecordLoss()
	} else if score1 < score2 {
		m.Team2.RecordWin()
		m.Team1.RecordLoss()
	} else {
		m.Team1.RecordDraw()
		m.Team2.RecordDraw()
	}
}

func (m *Match) Score1() int {
	score := 0
	for _, player := range m.Team1.Players {
		score += player.Goals
	}
	return score
}

// Score2 возвращает общий счет команды 2
func (m *Match) Score2() int {
	score := 0
	for _, player := range m.Team2.Players {
		score += player.Goals
	}
	return score
}

func (m *Match) simulateMinute(minute int) {
	fmt.Printf("Минута %d: ", minute)

	if rand.Intn(100) < 5 {
		scoringTeam := m.Team1
		if rand.Intn(2) == 1 {
			scoringTeam = m.Team2
		}
		player := scoringTeam.Players[rand.Intn(len(scoringTeam.Players))]
		player.ScoreGoal()
		fmt.Printf("Гол! %s забивает. Игрок: %s. Счет: %d:%d\n", scoringTeam.Name, player.Name, m.Score1(), m.Score2())

		// Добавляем событие гола
		m.Events = append(m.Events, Event{Type: "Goal", Minute: minute, Team: scoringTeam, Player: player})
	} else {
		// Генерация случайного события (фол, желтая карточка, травма)
		eventTypes := []string{"Foul", "Yellow Card", "Injury"}
		if rand.Intn(100) < 10 { // 10% шанс на событие
			eventType := eventTypes[rand.Intn(len(eventTypes))]
			player := m.Team1.Players[rand.Intn(len(m.Team1.Players))]
			if rand.Intn(2) == 1 {
				player = m.Team2.Players[rand.Intn(len(m.Team2.Players))]
			}
			fmt.Printf("%s совершил событие: %s\n", player.Name, eventType)
			m.Events = append(m.Events, Event{Type: eventType, Minute: minute, Team: m.Team1, Player: player})
		} else {
			fmt.Println("Ничего не произошло.")
		}
	}

	if minute == m.Duration {
		m.IsFinished = true
	}
}
