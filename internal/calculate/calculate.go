package calculate

import (
	"RPG/internal/model"
	"fmt"
)

// CalculateStrengthXP — считает опыт за силовую тренировку
func CalculateStrengthXP(report model.StrengthReport) int {
	volume := report.Weight * float64(report.Reps) * float64(report.Sets)
	oneRepMax := report.Weight * (1.0 + float64(report.Reps)/30.0)

	xp := 40

	if oneRepMax >= 80 {
		xp += 30
	} else if oneRepMax >= 60 {
		xp += 20
	} else if oneRepMax >= 40 {
		xp += 10
	}

	if volume > 1200 {
		xp += 40
	} else if volume > 780 {
		xp += 30
	} else if volume > 360 {
		xp += 20
	} else {
		xp += 10
	}

	if report.Weight <= 0 || report.Reps <= 0 || report.Sets <= 0 {
		return 0
	}
	return xp
}

// CalculateProgrammingXP — считает опыт за программирование
func CalculateProgrammingXP(report model.ProgrammingReport) int {
	xp := 20 // базовый XP за факт занятия

	if report.CodeHours >= 3 {
		xp += 50
	} else if report.CodeHours == 2 {
		xp += 35
	} else if report.CodeHours == 1 {
		xp += 20
	}

	switch report.TaskComplexity {
	case "hard":
		xp += 60
	case "medium":
		xp += 40
	case "easy":
		xp += 20
	default:
		fmt.Println("Введите заданные параметры")
	}

	if report.TaskSolved {
		xp += 40
	}

	return xp
}

// CalculateSleepXP — считает опыт за сон
func CalculateSleepXP(report model.SleepReport) int {
	xp := 20 // базовый XP за запись сна

	if report.Hours >= report.TargetHours {
		xp += 80
	} else if report.Hours >= report.TargetHours-1 {
		xp += 50
	} else if report.Hours >= report.TargetHours-2 {
		xp += 25
	}

	return xp
}

// CalculateNutritionXP — считает опыт за питание
func CalculateNutritionXP(report model.NutritionReport) int {
	xp := 20 // базовый XP за запись питания

	if report.Calories >= 3000 {
		xp += 60
	} else if report.Calories >= 2500 {
		xp += 40
	} else if report.Calories >= 2000 {
		xp += 20
	}

	if report.Weight >= 80 {
		xp += 40
	} else if report.Weight >= 76 {
		xp += 30
	} else if report.Weight >= 72 {
		xp += 20
	}

	return xp
}

// CalculateDisciplineXP — считает опыт за дисциплину
func CalculateDisciplineXP(report model.DisciplineReport) int {
	xp := 20 // базовый XP за контроль дня

	switch report.DayStatus {
	case "good":
		xp += 80
	case "normal":
		xp += 30
	case "failed":
		xp -= 30
	default:
		fmt.Println("Введите заданные параметры")
	}

	if report.StreakDays >= 30 {
		xp += 80
	} else if report.StreakDays >= 14 {
		xp += 50
	} else if report.StreakDays >= 7 {
		xp += 30
	}

	if xp < 0 {
		xp = 0
	}

	return xp
}
