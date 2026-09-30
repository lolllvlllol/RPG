package model

// Player хранит общее состояние игрока и прогресс его навыков.
type Player struct {
	XP             int // Опыт внутри текущего потенциального уровня.
	Level          int // Фактический уровень с учётом ограничения баланса.
	PotentialLevel int // Уровень, достигнутый на основе накопленного опыта.
	BalanceLimit   int // Максимальный доступный уровень с учётом развития навыков.

	Percent int    // Процент заполнения шкалы опыта.
	NeedXP  int    // Опыт, необходимый для перехода на следующий уровень.
	LeftXP  int    // Опыт, оставшийся до следующего уровня.
	Bar     string // Текстовое представление шкалы прогресса.

	Skills []Skill // Навыки игрока и их текущий прогресс.
}

// Skill описывает отдельный развиваемый навык игрока.
type Skill struct {
	Name    string
	Level   int
	XP      int // Опыт внутри текущего уровня навыка.
	TotalXP int // Общий опыт навыка за всё время.

	Percent int
	NeedXP  int
	LeftXP  int
	Bar     string
}

// StrengthReport содержит данные отчёта о силовой тренировке.
type StrengthReport struct {
	Weight float64 // Рабочий вес.
	Reps   int     // Количество повторений.
	Sets   int     // Количество подходов.
}

// SleepReport содержит данные отчёта о продолжительности сна.
type SleepReport struct {
	Hours       int // Фактическая продолжительность сна.
	TargetHours int // Целевая продолжительность сна.
}

// ProgrammingReport содержит данные отчёта о занятии программированием.
type ProgrammingReport struct {
	CodeHours      int
	TaskComplexity string // Допустимые значения: easy, medium, hard.
	TaskSolved     bool
}

// NutritionReport содержит данные отчёта о питании.
type NutritionReport struct {
	Calories int
	Weight   float64
}

// DisciplineReport содержит данные для оценки дисциплины за день.
type DisciplineReport struct {
	DayStatus  string // Допустимые значения: good, normal, failed.
	StreakDays int    // Продолжительность текущей серии дней.
}

// Mission описывает ежедневную миссию и награду за её выполнение.
type Mission struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Points      int    `json:"points"`
	Skill       string `json:"skill"`
	XP          int    `json:"xp"`
}

// DailyState хранит состояние ежедневных миссий.
// Оно сохраняется между запусками программы и обновляется при смене календарного дня.
type DailyState struct {
	Date      string    // Дата, для которой был сформирован набор миссий.
	Missions  []Mission // Миссии, выбранные на текущий день.
	Completed bool      // Показывает, сдавались ли миссии в этот день.
}
