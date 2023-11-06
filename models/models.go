package models

type Users struct {
	Age          int    `json:"age" validate:"required,gte=0,lte=130" `
	Gender       string `json:"gender" validate:"required" `
	Height       int    `json:"height" validate:"required,gte=50,lte=300" `
	Weight       int    `json:"weight" validate:"required,gte=2,lte=700" `
	HealthGoal   string `json:"healthGoal" validate:"required" `
	ProfilePhoto string `json:"profilePhoto" validate:"required" `
}

type UserSignup struct {
	Email    string `json:"email" validate:"required,email" `
	FullName string `json:"fullName" `
	Password string `json:"password" validate:"required" `
}

type Exercise struct {
	ExerciseType   string `json:"exerciseType" validate:"required" `
	Duration       int    `json:"duration" `
	CaloriesBurned int    `json:"caloriesBurned"`
	Date           string `json:"date"`
}
type Meal struct {
	MealType         string `json:"mealType" validate:"required" `
	Ingeredients     string `json:"ingredients" `
	CaloriesConsumed int    `json:"caloriesConsumed" `
	Date             string `json:"date"`
}

type Weight struct {
	DailyWeight string `json:"dailyWeight" validate:"required" `
	Date        string `json:"date"`
}

type Water struct {
	WaterIntake int    `json:"waterIntake" validate:"required" `
	Date        string `json:"date"`
}

type Message struct {
	Code    int    `json:"code"  validate:"required"`
	Message string `json:"message"  validate:"required"`
}

type UserID struct {
	UserID string `json:"userId"  validate:"required" `
}

// BreakFast        string `json:"BreakFast"`
// 	Launch           string `json:"launch"`
// 	Snacks           string `json:"snacks"`
// 	Dinner           string `json:"dinner"`
// 	WaterConsumption int    `json:"waterConsumption"`
// 	CaloriesTaken    int    `json:"caloriesTaken"`