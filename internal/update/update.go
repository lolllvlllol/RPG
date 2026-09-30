package update

import (
	"RPG/internal/model"
	"time"
)

// UpdateSkillLevel пересчитывает уровень и отображаемый прогресс навыка.
// Накопленного XP может хватить сразу на несколько уровней,
// поэтому повышение выполняется до тех пор, пока опыта достаточно.
func UpdateSkillLevel(skill *model.Skill) {
	for {
		needXP := 100 + skill.Level*50

		if skill.XP >= needXP {
			skill.Level += 1
			skill.XP -= needXP
		} else {
			break
		}
	}

	// После повышения уровня пересчитываем показатели прогресса
	// относительно нового требования по XP.
	needXP := 100 + skill.Level*50
	skill.NeedXP = needXP
	skill.LeftXP = needXP - skill.XP
	skill.Percent = skill.XP * 100 / needXP

	// Шкала состоит из 10 сегментов, каждый из которых соответствует 10%.
	filled := skill.Percent / 10
	empty := 10 - filled

	skill.Bar = "["

	for range filled {
		skill.Bar += "▰"
	}

	for range empty {
		skill.Bar += "▱"
	}

	skill.Bar += "]"
}

// UpdatePlayerLevel пересчитывает потенциальный и фактический уровни игрока.
// Фактический уровень ограничивается средним уровнем всех навыков.
func UpdatePlayerLevel(player *model.Player) {
	var allLevel int

	// PotentialLevel определяется только накопленным общим XP игрока.
	// Цикл позволяет обработать повышение сразу на несколько уровней.
	for {
		needXP := 1000 + player.PotentialLevel*500

		if player.XP >= needXP {
			player.PotentialLevel += 1
			player.XP -= needXP
		} else {
			break
		}
	}

	// BalanceLimit не позволяет общему уровню сильно опережать развитие навыков.
	for _, skill := range player.Skills {
		allLevel += skill.Level
	}

	player.BalanceLimit = allLevel / len(player.Skills)

	// Фактический уровень ограничен меньшим из потенциального уровня
	// и уровня, доступного по балансу навыков.
	player.Level = min(player.BalanceLimit, player.PotentialLevel)

	// Пересчитываем отображаемый прогресс до следующего потенциального уровня.
	needXP := 1000 + player.PotentialLevel*500
	player.NeedXP = needXP
	player.LeftXP = needXP - player.XP
	player.Percent = player.XP * 100 / needXP

	filled := player.Percent / 10
	empty := 10 - filled

	player.Bar = "["

	for range filled {
		player.Bar += "▰"
	}

	for range empty {
		player.Bar += "▱"
	}

	player.Bar += "]"
}

// AddXPAll начисляет опыт одновременно выбранному навыку и игроку.
// weight определяет, какая доля XP навыка попадёт в общий прогресс игрока.
func AddXPAll(xp int, skill *model.Skill, weight float64, player *model.Player) {
	skill.TotalXP += xp
	skill.XP += xp

	player.XP += int(float64(xp) * weight)
}

// TodayDailyState создаёт состояние нового дня из уже выбранных миссий.
// При создании состояние считается ещё не завершённым.
func TodayDailyState(dailyMissions []model.Mission) model.DailyState {
	today := time.Now().Format("2006-01-02")

	dailyState := model.DailyState{
		Date:      today,
		Missions:  dailyMissions,
		Completed: false,
	}

	return dailyState
}

// IsDailyStateToday проверяет, относится ли сохранённое состояние
// ежедневных миссий к текущему календарному дню.
func IsDailyStateToday(dailyState model.DailyState) bool {
	today := time.Now().Format("2006-01-02")

	if dailyState.Date == today {
		return true
	} else {
		return false
	}
}
