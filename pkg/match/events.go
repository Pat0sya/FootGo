package match

import (
	"fmt"
	"math/rand"
)

// Event представляет событие, которое может произойти во время матча

// GenerateRandomEvent генерирует случайное событие для симуляции матча
func (m *Match) GenerateRandomEvent(minute int) {
	eventTypes := []string{"Goal", "Foul", "Yellow Card", "Injury"}
	selectedEvent := eventTypes[rand.Intn(len(eventTypes))]
	team := m.Team1
	if rand.Intn(2) == 1 {
		team = m.Team2
	}
	score1 := m.Score1()
	score2 := m.Score2()
	switch selectedEvent {
	case "Goal":
		if team == m.Team1 {
			score1++ // Увеличиваем счет первой команды
		} else {
			score2++ // Увеличиваем счет второй команды
		}
		fmt.Printf("Гол! %s забивает! Счет: %d:%d\n", team.Name, score1, score2)
	case "Foul":
		fmt.Printf("Фол! Команда %s совершила фол.\n", team.Name)
	case "Yellow Card":
		fmt.Printf("Желтая карточка! Команда %s получает предупреждение.\n", team.Name)
	case "Injury":
		fmt.Printf("Травма! Игрок из команды %s получил травму.\n", team.Name)
	}
}
