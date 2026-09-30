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

// ApplyDailyReport обрабатывает ежедневный отчёт игрока.
// Функция рассчитывает XP по каждому направлению, начисляет награды,
// обновляет уровни навыков и игрока, после чего сохраняет новый прогресс.
func ApplyDailyReport(
	player *model.Player,
	strength model.StrengthReport,
	sleep model.SleepReport,
	programming model.ProgrammingReport,
	nutrition model.NutritionReport,
	discipline model.DisciplineReport,
) {
	// Каждый тип активности использует собственные правила расчёта XP.
	strengthXP := calculate.CalculateStrengthXP(strength)
	sleepXP := calculate.CalculateSleepXP(sleep)
	programmingXP := calculate.CalculateProgrammingXP(programming)
	nutritionXP := calculate.CalculateNutritionXP(nutrition)
	disciplineXP := calculate.CalculateDisciplineXP(discipline)

	// XP начисляется соответствующему навыку и одновременно
	// учитывается в общем прогрессе игрока с заданным коэффициентом.
	update.AddXPAll(strengthXP, &player.Skills[0], 1.1, player)
	update.AddXPAll(sleepXP, &player.Skills[1], 1.5, player)
	update.AddXPAll(programmingXP, &player.Skills[2], 1.3, player)
	update.AddXPAll(nutritionXP, &player.Skills[3], 1.4, player)
	update.AddXPAll(disciplineXP, &player.Skills[4], 1.5, player)

	// После начисления XP пересчитываем уровень и прогресс каждого навыка.
	for i := range player.Skills {
		update.UpdateSkillLevel(&player.Skills[i])
	}

	// Общий уровень пересчитывается после обновления всех навыков.
	update.UpdatePlayerLevel(player)

	// Обновлённое состояние игрока сохраняется между запусками программы.
	SavePlayerJSON(*player)
}

// LoadPlayerJSON загружает состояние игрока из save.json.
// Если сохранение недоступно, создаётся новый игрок с начальными параметрами.
func LoadPlayerJSON(player *model.Player) {
	data, err := os.ReadFile("save.json")

	if err != nil {
		fmt.Println(err)

		// Начальное состояние используется при первом запуске,
		// когда сохранённого профиля игрока ещё нет.
		*player = model.Player{
			XP:             0,
			Level:          1,
			PotentialLevel: 1,
			BalanceLimit:   1,
			Skills: []model.Skill{
				{Name: "Strength", Level: 1, XP: 0, TotalXP: 0},
				{Name: "Sleep", Level: 1, XP: 0, TotalXP: 0},
				{Name: "Programming", Level: 1, XP: 0, TotalXP: 0},
				{Name: "Nutrition", Level: 1, XP: 0, TotalXP: 0},
				{Name: "Discipline", Level: 1, XP: 0, TotalXP: 0},
			},
		}
	} else {
		// Преобразуем сохранённый JSON обратно в структуру Player.
		// Указатель позволяет заполнить исходную переменную игрока.
		err = json.Unmarshal(data, player)

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

// LoadDailyStateJSON загружает сохранённое состояние ежедневных миссий.
// Возвращает true только в случае успешного чтения и разбора daily.json.
func LoadDailyStateJSON(dailyState *model.DailyState) bool {
	data, err := os.ReadFile("daily.json")

	if err != nil {
		fmt.Println(err)
		return false
	} else {
		// Данные из JSON записываются непосредственно
		// в переданный DailyState через указатель.
		err = json.Unmarshal(data, dailyState)

		if err != nil {
			fmt.Println(err)
			return false
		}
	}

	return true
}

// SavePlayerJSON сохраняет текущее состояние игрока в save.json.
// Структура преобразуется в JSON перед записью на диск.
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

// SaveDailyStateJSON сохраняет состояние текущего дня в daily.json.
// Благодаря этому выбранные миссии и флаг Completed сохраняются
// после завершения программы.
func SaveDailyStateJSON(dailyState model.DailyState) {
	data, err := json.MarshalIndent(dailyState, "", "	")

	if err != nil {
		fmt.Println(err)
		return
	}

	if err := os.WriteFile("daily.json", data, 0644); err != nil {
		fmt.Println(err)
		return
	}
}

// ReadMissionJSON загружает общий список миссий из missions.json
// и случайным образом выбирает пять миссий для нового дня.
func ReadMissionJSON() []model.Mission {
	var missions []model.Mission

	data, err := os.ReadFile("missions.json")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return nil
	}

	// Преобразуем JSON-массив в срез структур Mission.
	err = json.Unmarshal(data, &missions)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return nil
	}

	// Перемешивание позволяет получать новый случайный набор
	// при формировании миссий следующего дня.
	rand.Shuffle(len(missions), func(i, j int) {
		missions[i], missions[j] = missions[j], missions[i]
	})

	// После перемешивания первые пять элементов становятся миссиями дня.
	dailyMissions := missions[:5]

	return dailyMissions
}
