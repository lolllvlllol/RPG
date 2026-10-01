package main

import (
	"RPG/internal/calculate"
	"RPG/internal/iointerface"
	"RPG/internal/model"
	"RPG/internal/update"
	"RPG/internal/workjson"
	"fmt"
)

func main() {
	var player model.Player
	var dailyMissions []model.Mission

	// При запуске восстанавливаем сохранённое состояние игрока.
	// Если сохранения ещё нет, LoadPlayerJSON создаст начальный профиль.
	workjson.LoadPlayerJSON(&player)

	for {
		var choice int

		fmt.Println("=== MENU ===")
		fmt.Println("1 — Ввести отчёт за день")
		fmt.Println("2 — Показать профиль")
		fmt.Println("3 — Показать ежедневные миссии")
		fmt.Println("4 — Записать ежедневные миссии")
		fmt.Println("5 — Выйти")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Для каждого направления создаётся отдельный отчёт,
			// который затем заполняется пользовательским вводом.
			strengthReport := model.StrengthReport{}
			sleepReport := model.SleepReport{}
			programmingReport := model.ProgrammingReport{}
			nutritionReport := model.NutritionReport{}
			disciplineReport := model.DisciplineReport{}

			iointerface.InputInfo(
				&strengthReport,
				&sleepReport,
				&programmingReport,
				&nutritionReport,
				&disciplineReport,
			)

			// После заполнения отчётов рассчитываем XP,
			// обновляем прогресс игрока и сохраняем его состояние.
			workjson.ApplyDailyReport(
				&player,
				strengthReport,
				sleepReport,
				programmingReport,
				nutritionReport,
				disciplineReport,
			)

		case 2:
			// Выводим текущее состояние игрока без его изменения.
			iointerface.ShowPlayer(&player)

		case 3:
			var dailyState model.DailyState

			// Пытаемся восстановить ранее сохранённое состояние миссий.
			// loaded показывает только успешность загрузки,
			// а сами данные записываются в dailyState.
			loaded := workjson.LoadDailyStateJSON(&dailyState)

			if loaded {
				if update.IsDailyStateToday(dailyState) {
					// Если состояние относится к сегодняшнему дню,
					// используем уже сохранённый набор миссий.
					dailyMissions = dailyState.Missions
				} else {
					// При смене календарного дня формируем новый набор
					// и сохраняем новое состояние с Completed = false.
					dailyMissions = workjson.ReadMissionJSON()
					dailyState = update.TodayDailyState(dailyMissions)
					workjson.SaveDailyStateJSON(dailyState)
				}
			} else {
				// Если сохранённого состояния нет, создаём его впервые.
				dailyMissions = workjson.ReadMissionJSON()
				dailyState = update.TodayDailyState(dailyMissions)
				workjson.SaveDailyStateJSON(dailyState)
			}

			// Показываем уже выбранные миссии без повторной генерации.
			iointerface.ShowMissions(dailyMissions)

		case 4:
			var dailyState model.DailyState

			// Перед сдачей миссий загружаем состояние дня,
			// чтобы проверить, не были ли они уже сданы.
			loaded := workjson.LoadDailyStateJSON(&dailyState)

			if loaded {
				if update.IsDailyStateToday(dailyState) {
					if !dailyState.Completed {
						// Для начисления награды используются именно миссии,
						// сохранённые в состоянии текущего дня.
						calculate.CalculateDailyMissions(
							dailyState.Missions,
							&player,
						)

						// После первой сдачи блокируем повторное получение награды
						// и сохраняем обновлённое состояние дня.
						dailyState.Completed = true
						workjson.SaveDailyStateJSON(dailyState)
					} else {
						fmt.Println("Миссии уже сдавались")
					}
				} else {
					fmt.Println("Дата не соответствует. Загрузка новых миссий...")
					dailyMissions = workjson.ReadMissionJSON()
					dailyState = update.TodayDailyState(dailyMissions)
					workjson.SaveDailyStateJSON(dailyState)
				}
			} else {
				fmt.Println("Миссий нет. Загрузка новых миссий...")
				dailyMissions = workjson.ReadMissionJSON()
				dailyState = update.TodayDailyState(dailyMissions)
				workjson.SaveDailyStateJSON(dailyState)
			}

			// Сохраняем изменения XP и уровней игрока отдельно
			// от состояния ежедневных миссий.
			workjson.SavePlayerJSON(player)

		case 5:
			fmt.Println("Выход")
			return

		default:
			fmt.Println("Неизвестная команда")
		}
	}
}
