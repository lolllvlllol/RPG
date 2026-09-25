package iointerface

import (
	"RPG/internal/model"
	"fmt"
)

func InputInfo(s *model.StrengthReport, sl *model.SleepReport, p *model.ProgrammingReport, n *model.NutritionReport, d *model.DisciplineReport) {
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
	sl.TargetHours = 8

	fmt.Println("Введите данные по аспекту - Программирование")
	fmt.Println("Введите кол-во часов программирования")
	fmt.Scan(&p.CodeHours)
	fmt.Println("Введите сложность задачи (hard, medium, easy)")
	fmt.Scan(&p.TaskComplexity) // hard, medium, easy
	fmt.Println("Введите выполнение задачи (true, false)")
	fmt.Scan(&p.TaskSolved) // true, false

	fmt.Println("Введите данные по аспекту - Питание")
	fmt.Println("Введите ккал за день")
	fmt.Scan(&n.Calories)
	fmt.Println("Введите свой вес ")
	fmt.Scan(&n.Weight) // float

	fmt.Println("Введите данные по аспекту - Дисциплина")
	fmt.Println("Введите статус дня: good / normal / failed")
	fmt.Scan(&d.DayStatus)
	fmt.Println("Введите серию дней")
	fmt.Scan(&d.StreakDays)
}

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
