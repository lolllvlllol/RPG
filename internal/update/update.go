package update

import "RPG/internal/model"

// UpdateSkillLevel — переводит накопленный XP навыка в уровни
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

	needXP := 100 + skill.Level*50
	skill.NeedXP = needXP
	skill.LeftXP = needXP - skill.XP
	skill.Percent = skill.XP * 100 / needXP
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

// UpdatePlayerLevel — считает общий XP и уровень игрока
func UpdatePlayerLevel(player *model.Player) {
	var allLevel int

	for {
		needXP := 1000 + player.PotentialLevel*500
		if player.XP >= needXP {
			player.PotentialLevel += 1
			player.XP -= needXP
		} else {
			break
		}
	}

	for _, skill := range player.Skills {
		allLevel += skill.Level
	}

	player.BalanceLimit = allLevel / len(player.Skills)

	player.Level = min(player.BalanceLimit, player.PotentialLevel)

	//

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

func AddXPAll(xp int, skill *model.Skill, weight float64, player *model.Player) {
	skill.TotalXP += xp
	skill.XP += xp

	player.XP += int(float64(xp) * weight)
}
