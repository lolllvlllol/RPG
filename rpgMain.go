package main

import (
	"RPG/internal/iointerface"
	"RPG/internal/model"
	"RPG/internal/workjson"
	"fmt"
)

func main() {
	var player model.Player
	workjson.LoadPlayerJSON(&player)
	for {
		var choice int

		fmt.Println("=== MENU ===")
		fmt.Println("1 — Ввести отчёт за день")
		fmt.Println("2 — Показать профиль")
		fmt.Println("3 — Выйти")
		fmt.Println("4 — Показать ежедневные миссии")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			strengthReport := model.StrengthReport{}
			sleepReport := model.SleepReport{}
			programmingReport := model.ProgrammingReport{}
			nutritionReport := model.NutritionReport{}
			disciplineReport := model.DisciplineReport{}

			iointerface.InputInfo(&strengthReport, &sleepReport, &programmingReport, &nutritionReport, &disciplineReport)
			workjson.ApplyDailyReport(&player, strengthReport, sleepReport, programmingReport, nutritionReport, disciplineReport)

		case 2:
			iointerface.ShowPlayer(&player)

		case 3:
			fmt.Println("Выход")
			return
		case 4:
			workjson.ReadNissionJSON()
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}
