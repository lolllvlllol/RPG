package workjson

import (
	"RPG/internal/calculate"
	"RPG/internal/model"
	"RPG/internal/update"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
)

func ApplyDailyReport(player *model.Player, strength model.StrengthReport, sleep model.SleepReport, programming model.ProgrammingReport, nutrition model.NutritionReport, discipline model.DisciplineReport) {
	// Подсчёт XP за день
	strengthXP := calculate.CalculateStrengthXP(strength)
	sleepXP := calculate.CalculateSleepXP(sleep)
	programmingXP := calculate.CalculateProgrammingXP(programming)
	nutritionXP := calculate.CalculateNutritionXP(nutrition)
	disciplineXP := calculate.CalculateDisciplineXP(discipline)

	// Начисление XP игроку и навыкам
	update.AddXPAll(strengthXP, &player.Skills[0], 1.1, player)
	update.AddXPAll(sleepXP, &player.Skills[1], 1.5, player)
	update.AddXPAll(programmingXP, &player.Skills[2], 1.3, player)
	update.AddXPAll(nutritionXP, &player.Skills[3], 1.4, player)
	update.AddXPAll(disciplineXP, &player.Skills[4], 1.5, player)

	// Обновление уровней навыков
	for i := range player.Skills {
		update.UpdateSkillLevel(&player.Skills[i])
	}

	// Обновление уровня игрока
	update.UpdatePlayerLevel(player)
	// Сохранение данных игрока в JSON
	SavePlayerJSON(*player)
}

func LoadPlayerJSON(player *model.Player) {
	data, err := os.ReadFile("save.json")
	if err != nil {
		fmt.Println(err)
		*player = model.Player{
			XP:             0,
			Level:          1,
			PotentialLevel: 1,
			BalanceLimit:   1,
			Skills: []model.Skill{
				{Name: "Сила", Level: 1, XP: 0, TotalXP: 0},
				{Name: "Сон", Level: 1, XP: 0, TotalXP: 0},
				{Name: "Программирование", Level: 1, XP: 0, TotalXP: 0},
				{Name: "Питание", Level: 1, XP: 0, TotalXP: 0},
				{Name: "Дисциплина", Level: 1, XP: 0, TotalXP: 0},
			},
		}
	} else {
		err = json.Unmarshal(data, player)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func SavePlayerJSON(player model.Player) {
	data, err := json.MarshalIndent(player, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := os.WriteFile("save.json", data, 0644); err != nil {
		fmt.Println(err)
		return
	}
}

func ReadNissionJSON() {
	var missions []model.Mission
	data, err := os.ReadFile("missions.json")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	err = json.Unmarshal(data, &missions)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	rand.Shuffle(len(missions), func(i, j int) {
		missions[i], missions[j] = missions[j], missions[i]
	})

	dailyMissions := missions[:5]

	for _, mission := range dailyMissions {
		fmt.Println("Миссия:", mission.Name)
		fmt.Println("Описание:", mission.Description)
		fmt.Println("Очки:", mission.Points, "MP")
		fmt.Println("Навык:", mission.Skill)
		fmt.Println("Опыт:", mission.XP, "XP")
		fmt.Println("----------------------")
	}
}
