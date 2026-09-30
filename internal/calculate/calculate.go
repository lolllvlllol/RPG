package calculate

import (
	"RPG/internal/model"
	"RPG/internal/update"
	"fmt"
)

// CalculateStrengthXP рассчитывает опыт за силовую тренировку.
// Награда зависит от общего объёма тренировки и примерного одноповторного максимума.
func CalculateStrengthXP(report model.StrengthReport) int {
	volume := report.Weight * float64(report.Reps) * float64(report.Sets)
	oneRepMax := report.Weight * (1.0 + float64(report.Reps)/30.0)

	xp := 40

	// Чем выше расчётный одноповторный максимум, тем больше бонус к опыту.
	if oneRepMax >= 80 {
		xp += 30
	} else if oneRepMax >= 60 {
		xp += 20
	} else if oneRepMax >= 40 {
		xp += 10
	}

	// Объём тренировки учитывает одновременно вес, количество повторений и подходов.
	if volume > 1200 {
		xp += 40
	} else if volume > 780 {
		xp += 30
	} else if volume > 360 {
		xp += 20
	} else {
		xp += 10
	}

	// Некорректные параметры тренировки не должны приносить опыт.
	if report.Weight <= 0 || report.Reps <= 0 || report.Sets <= 0 {
		return 0
	}

	return xp
}

// CalculateProgrammingXP рассчитывает опыт за занятие программированием.
// Учитываются продолжительность занятия, сложность задачи и факт её решения.
func CalculateProgrammingXP(report model.ProgrammingReport) int {
	xp := 20

	// Продолжительная практика увеличивает базовую награду.
	if report.CodeHours >= 3 {
		xp += 50
	} else if report.CodeHours == 2 {
		xp += 35
	} else if report.CodeHours == 1 {
		xp += 20
	}

	// Сложность задачи определяет дополнительный опыт.
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

	// Решённая задача даёт отдельный бонус независимо от её сложности.
	if report.TaskSolved {
		xp += 40
	}

	return xp
}

// CalculateSleepXP рассчитывает опыт за сон.
// Награда определяется относительно индивидуальной цели по продолжительности сна.
func CalculateSleepXP(report model.SleepReport) int {
	xp := 20

	if report.Hours >= report.TargetHours {
		xp += 80
	} else if report.Hours >= report.TargetHours-1 {
		xp += 50
	} else if report.Hours >= report.TargetHours-2 {
		xp += 25
	}

	return xp
}

// CalculateNutritionXP рассчитывает опыт за питание.
// Учитываются суточная калорийность и текущий вес пользователя.
func CalculateNutritionXP(report model.NutritionReport) int {
	xp := 20

	// Калорийность определяет основную часть награды за питание.
	if report.Calories >= 3000 {
		xp += 60
	} else if report.Calories >= 2500 {
		xp += 40
	} else if report.Calories >= 2000 {
		xp += 20
	}

	// Вес используется как дополнительный показатель прогресса.
	if report.Weight >= 80 {
		xp += 40
	} else if report.Weight >= 76 {
		xp += 30
	} else if report.Weight >= 72 {
		xp += 20
	}

	return xp
}

// CalculateDisciplineXP рассчитывает опыт за дисциплину.
// Итог зависит от оценки дня и продолжительности текущей серии успешных дней.
func CalculateDisciplineXP(report model.DisciplineReport) int {
	xp := 20

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

	// Длительная серия соблюдения режима даёт дополнительный опыт.
	if report.StreakDays >= 30 {
		xp += 80
	} else if report.StreakDays >= 14 {
		xp += 50
	} else if report.StreakDays >= 7 {
		xp += 30
	}

	// Опыт за отчёт не может стать отрицательным.
	if xp < 0 {
		xp = 0
	}

	return xp
}

// CalculateDailyMissions обрабатывает выполненные ежедневные миссии.
// Функция проверяет введённые ID, предотвращает повторный учёт одной миссии,
// начисляет опыт соответствующим навыкам и определяет результат дня по очкам.
func CalculateDailyMissions(missions []model.Mission, player *model.Player) {
	var count, number, point int
	var idnumber []int
	var alreadyUsed, missionFound bool
	var foundMission model.Mission

	fmt.Println("Сколько миссий выполнено?")
	fmt.Scan(&count)

	// За один день можно обработать не более пяти выданных миссий.
	if count >= 0 && count <= 5 {
		for len(idnumber) < count {
			fmt.Println("Введите ID миссии")
			fmt.Scan(&number)

			missionFound = false
			alreadyUsed = false

			// Проверяем, относится ли введённый ID к одной из сегодняшних миссий.
			for _, mission := range missions {
				if number == mission.ID {
					missionFound = true
					foundMission = mission
				}
			}

			if !missionFound {
				fmt.Println("ID такого нет")
			} else {
				// Одна и та же миссия не должна учитываться несколько раз.
				for _, v := range idnumber {
					if v == number {
						alreadyUsed = true
					}
				}
			}

			if alreadyUsed {
				fmt.Println("ID уже был введен")
			}

			// Начисляем награду только после успешной проверки ID.
			if missionFound && !alreadyUsed {
				idnumber = append(idnumber, number)

				// XP получает навык, к которому привязана выполненная миссия.
				for v := range player.Skills {
					if player.Skills[v].Name == foundMission.Skill {
						update.AddXPAll(
							foundMission.XP,
							&player.Skills[v],
							1.0,
							player,
						)

						update.UpdateSkillLevel(&player.Skills[v])
						point += foundMission.Points
					}
				}
			}
		}

		// Общий уровень игрока пересчитывается после обработки всех миссий.
		update.UpdatePlayerLevel(player)

		// Для успешного завершения дневного набора необходимо набрать 100 очков.
		if point >= 100 {
			fmt.Println("Миссии выполнены")
		} else {
			fmt.Println("Миссии провалены")
		}
	} else {
		fmt.Println("Максимум 5 миссий в день")
	}
}
