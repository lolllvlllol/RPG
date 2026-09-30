package iointerface

import (
	"RPG/internal/model"
	"fmt"
)

// InputInfo запрашивает у пользователя данные для ежедневного отчёта
// и записывает введённые значения непосредственно в переданные структуры.
func InputInfo(
	s *model.StrengthReport,
	sl *model.SleepReport,
	p *model.ProgrammingReport,
	n *model.NutritionReport,
	d *model.DisciplineReport,
) {
	fmt.Println("Введите показатели:")

	fmt.Println("Введите данные по аспекту - Сила")
	fmt.Println("Введите вес тренажера")
	fmt.Scan(&s.Weight)
	fmt.Println("Введите повторения")
	fmt.Scan(&s.Reps)
	fmt.Println("Введите подходы")
	fmt.Scan(&s.Sets)

	fmt.Println("Введите данные по аспекту - Сон")
	fmt.Println("Введите кол-во часов сна")
	fmt.Scan(&sl.Hours)

	// Целевая продолжительность сна задаётся автоматически и не требует ввода.
	sl.TargetHours = 8

	fmt.Println("Введите данные по аспекту - Программирование")
	fmt.Println("Введите кол-во часов программирования")
	fmt.Scan(&p.CodeHours)
	fmt.Println("Введите сложность задачи (hard, medium, easy)")
	fmt.Scan(&p.TaskComplexity)
	fmt.Println("Введите выполнение задачи (true, false)")
	fmt.Scan(&p.TaskSolved)

	fmt.Println("Введите данные по аспекту - Питание")
	fmt.Println("Введите ккал за день")
	fmt.Scan(&n.Calories)
	fmt.Println("Введите свой вес ")
	fmt.Scan(&n.Weight)

	fmt.Println("Введите данные по аспекту - Дисциплина")
	fmt.Println("Введите статус дня: good / normal / failed")
	fmt.Scan(&d.DayStatus)
	fmt.Println("Введите серию дней")
	fmt.Scan(&d.StreakDays)
}

// ShowPlayer выводит текущее состояние игрока:
// общий уровень, прогресс и характеристики каждого навыка.
func ShowPlayer(player *model.Player) {
	fmt.Printf(
		"=== PLAYER ===\n"+
			"LVL        %d / 100\n"+
			"POTENTIAL  %d / 100\n"+
			"BALANCE    %d / 100\n\n"+
			"XP         %d / %d\n"+
			"BAR        %s %d%%\n\n",
		player.Level,
		player.PotentialLevel,
		player.BalanceLimit,
		player.XP,
		player.NeedXP,
		player.Bar,
		player.Percent,
	)

	fmt.Println("=== SKILLS ===")

	// Каждый навык выводится отдельно с его уровнем и текущим прогрессом.
	for _, skill := range player.Skills {
		fmt.Printf(
			"%-18s LVL %-2d | %-3d / %-3d | %-14s %d%%\n",
			skill.Name,
			skill.Level,
			skill.XP,
			skill.NeedXP,
			skill.Bar,
			skill.Percent,
		)
	}
}

// ShowMissions выводит переданный набор миссий и их игровые параметры.
// Функция только отображает миссии и не изменяет их состояние.
func ShowMissions(missions []model.Mission) {
	for _, mission := range missions {
		fmt.Println("Миссия:", mission.Name)
		fmt.Println("Описание:", mission.Description)
		fmt.Println("Очки:", mission.Points, "MP")
		fmt.Println("Навык:", mission.Skill)
		fmt.Println("Опыт:", mission.XP, "XP")
		fmt.Println("----------------------")
	}
}
