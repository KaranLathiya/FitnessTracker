package models

type Users struct {
	UserID       int    `json:"userId"`
	Email        string `json:"email"  `
	FullName     string `json:"fullName" validate:"required" `
	Age          int    `json:"age" validate:"required,gte=0,lte=130" `
	Gender       string `json:"gender" validate:"required" `
	Height       int    `json:"height" validate:"required,gte=15,lte=300" `
	Weight       int    `json:"weight" validate:"required,gte=0,lte=700" `
	HealthGoal   string `json:"healthGoal" validate:"required" `
	ProfilePhoto string `json:"profilePhoto" validate:"required" `
}

type UserSignup struct {
	UserID   int    `json:"userId"`
	Email    string `json:"email" validate:"required,email" `
	FullName string `json:"fullName"  `
	Password string `json:"password" validate:"required" `
}

type Exercise struct {
	UserID         int    `json:"userId"`
	ExerciseType   string `json:"exerciseType" validate:"required" `
	Duration       int    `json:"duration" validate:"required" `
	CaloriesBurned int    `json:"caloriesBurned" validate:"required" `
	Date           string `json:"date"`
}
type Meal struct {
	UserID           int    `json:"userId"`
	MealType         string `json:"mealType" validate:"required" `
	Ingeredients     string `json:"ingredients" validate:"required" `
	CaloriesConsumed int    `json:"caloriesConsumed" validate:"required" `
	Date             string `json:"date"`
}

type Weight struct {
	UserID      int    `json:"userId"`
	DailyWeight string `json:"dailyWeight" validate:"required" `
	Date        string `json:"date"`
}

type Water struct {
	UserID      int    `json:"userId"`
	WaterIntake int    `json:"waterIntake" validate:"required" `
	Date        string `json:"date"`
}

type MyError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// BreakFast        string `json:"BreakFast"`
// 	Launch           string `json:"launch"`
// 	Snacks           string `json:"snacks"`
// 	Dinner           string `json:"dinner"`
// 	WaterConsumption int    `json:"waterConsumption"`
// 	CaloriesTaken    int    `json:"caloriesTaken"`
