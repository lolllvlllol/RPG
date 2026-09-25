package model

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
	DayStatus  string // good / normal / failed
	StreakDays int
}

// Mission - набор всяких дневных миссий в json созданных по структуре
type Mission struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Points      int    `json:"points"`
	Skill       string `json:"skill"`
	XP          int    `json:"xp"`
}
