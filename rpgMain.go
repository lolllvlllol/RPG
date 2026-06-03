package main

import "fmt"

// Player — главный герой / общий профиль
type Player struct {
	XP             int // XP внутри текущего потенциального уровня
	Level          int // текущий уровень с учетом баланса
	PotentialLevel int // уровень, который игрок заслужил по XP
	BalanceLimit   int // ограничение уровня по среднему уровню навыков

	Percent int    // процент заполнения XP-бара
	NeedXP  int    // сколько XP нужно до следующего уровня
	LeftXP  int    // сколько XP осталось до следующего уровня
	Bar     string // визуальная шкала XP

	Skills []Skill // список навыков игрока
}

// Skill — один аспект жизни: сила, сон, питание и т.д.
type Skill struct {
	Name    string
	Level   int
	XP      int // XP внутри текущего уровня навыка
	TotalXP int // общий XP навыка за всё время, пока оставляем на будущее

	Percent int
	NeedXP  int
	LeftXP  int
	Bar     string
}

// StrengthReport — отчёт по силовой тренировке
type StrengthReport struct {
	Weight float64 // рабочий вес
	Reps   int     // повторения
	Sets   int     // подходы
}

// SleepReport — отчёт по сну
type SleepReport struct {
	Hours       int // сколько часов спал
	TargetHours int // сколько часов нужно
}

// ProgrammingReport — отчёт по программированию
type ProgrammingReport struct {
	CodeHours      int
	TaskComplexity string // easy / medium / hard
	TaskSolved     bool
}

// NutritionReport — отчёт по питанию
type NutritionReport struct {
	Calories int
	Weight   float64
}

// DisciplineReport — отчёт по дисциплине
type DisciplineReport struct {
	DailyCompleted bool
	DailyFailed    bool
	StreakDays     int
}

func main() {
	player := Player{
		XP:             0,
		Level:          1,
		PotentialLevel: 1,
		BalanceLimit:   1,
		Skills: []Skill{
			{Name: "Сила", Level: 1, XP: 0, TotalXP: 0},
			{Name: "Сон", Level: 1, XP: 0, TotalXP: 0},
			{Name: "Программирование", Level: 1, XP: 0, TotalXP: 0},
			{Name: "Питание", Level: 1, XP: 0, TotalXP: 0},
			{Name: "Дисциплина", Level: 1, XP: 0, TotalXP: 0},
		},
	}

	strengthReport := StrengthReport{}
	sleepReport := SleepReport{}
	programmingReport := ProgrammingReport{}
	nutritionReport := NutritionReport{}
	disciplineReport := DisciplineReport{}

	InputInfo(&strengthReport, &sleepReport, &programmingReport, &nutritionReport, &disciplineReport)

	// Подсчёт XP за день
	strengthXP := CalculateStrengthXP(strengthReport)
	sleepXP := CalculateSleepXP(sleepReport)
	programmingXP := CalculateProgrammingXP(programmingReport)
	nutritionXP := CalculateNutritionXP(nutritionReport)
	disciplineXP := CalculateDisciplineXP(disciplineReport)

	// Начисление XP игроку и навыкам
	AddXPAll(strengthXP, &player.Skills[0], 1.1, &player)
	AddXPAll(sleepXP, &player.Skills[1], 1.5, &player)
	AddXPAll(programmingXP, &player.Skills[2], 1.3, &player)
	AddXPAll(nutritionXP, &player.Skills[3], 1.4, &player)
	AddXPAll(disciplineXP, &player.Skills[4], 1.5, &player)

	// Обновление уровней навыков
	for i := range player.Skills {
		UpdateSkillLevel(&player.Skills[i])
	}

	// Обновление уровня игрока
	UpdatePlayerLevel(&player)

	ShowPlayer(&player)
}

// CalculateStrengthXP — считает опыт за силовую тренировку
func CalculateStrengthXP(report StrengthReport) int {
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
func CalculateProgrammingXP(report ProgrammingReport) int {
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
	}

	if report.TaskSolved {
		xp += 40
	}

	return xp
}

// CalculateSleepXP — считает опыт за сон
func CalculateSleepXP(report SleepReport) int {
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
func CalculateNutritionXP(report NutritionReport) int {
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
func CalculateDisciplineXP(report DisciplineReport) int {
	xp := 20 // базовый XP за контроль дня

	if report.DailyCompleted {
		xp += 80
	}

	if report.DailyFailed {
		xp -= 30
	}

	if report.StreakDays >= 30 {
		xp += 80
	} else if report.StreakDays >= 14 {
		xp += 50
	} else if report.StreakDays >= 7 {
		xp += 30
	}

	if xp < 0 { //???
		xp = 0
	}

	return xp
}

// UpdateSkillLevel — переводит накопленный XP навыка в уровни
func UpdateSkillLevel(skill *Skill) {
	for {
		needXP := 100 + skill.Level*50
		if skill.XP >= needXP {
			skill.Level += 1
			skill.XP -= needXP
		} else {
			break
		}
	}

	//

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
func UpdatePlayerLevel(player *Player) {
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

func AddXPAll(xp int, skill *Skill, weight float64, player *Player) {
	skill.TotalXP += xp
	skill.XP += xp

	player.XP += int(float64(xp) * weight)
}

func InputInfo(s *StrengthReport, sl *SleepReport, p *ProgrammingReport, n *NutritionReport, d *DisciplineReport) {
	fmt.Println("Введите показатели:")
	fmt.Println("Введите данные по аспекту - Сила")
	fmt.Scan(&s.Weight)
	fmt.Scan(&s.Reps)
	fmt.Scan(&s.Sets)

	fmt.Println("Введите данные по аспекту - Сон")
	fmt.Scan(&sl.Hours)
	sl.TargetHours = 8

	fmt.Println("Введите данные по аспекту - Программирование")
	fmt.Scan(&p.CodeHours)
	fmt.Scan(&p.TaskComplexity) // hard, medium, easy
	fmt.Scan(&p.TaskSolved)     // true, false

	fmt.Println("Введите данные по аспекту - Питание")
	fmt.Scan(&n.Calories)
	fmt.Scan(&n.Weight) // float

	fmt.Println("Введите данные по аспекту - Дисциплина")
	fmt.Scan(&d.DailyCompleted) // true, false
	fmt.Scan(&d.DailyFailed)    // true, false
	fmt.Scan(&d.StreakDays)
}

func ShowPlayer(player *Player) {
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
