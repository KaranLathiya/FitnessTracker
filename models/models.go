package models

type Users struct {
	Email        string  `json:"email" validate:"required,email" `
	FullName     string  `json:"fullName" validate:"required" `
	Age          *int    `json:"age" validate:"required,gt=0,lte=130" `
	Gender       *string `json:"gender" validate:"required" `
	Height       *int    `json:"height" validate:"required,gte=50,lte=300" `
	Weight       *int    `json:"weight" validate:"required,gte=2,lte=700" `
	HealthGoal   *string `json:"healthGoal" validate:"required" `
	ProfilePhoto *string `json:"profilePhoto" validate:"required" `
}

type UserSignup struct {
	Email    string `json:"email" validate:"required,email" `
	FullName string `json:"fullName" validate:"required" `
	Password string `json:"password" validate:"required" `
}

type UserLogin struct {
	Email    string `json:"email" validate:"required,email" `
	Password string `json:"password" validate:"required" `
}

type Exercise struct {
	ExerciseType   string `json:"exerciseType" validate:"required" `
	Duration       int    `json:"duration" validate:"required,gt=0,lte=1440"`
	CaloriesBurned int    `json:"caloriesBurned" validate:"required,gt=0,lte=20000"`
	Date           string `json:"date,omitempty"`
}

type Meal struct {
	MealType         string `json:"mealType" validate:"required" `
	Ingeredients     string `json:"ingredients" validate:"required" `
	CaloriesConsumed int    `json:"caloriesConsumed" validate:"required,gte=0,lte=20000" `
	Date             string `json:"date,omitempty"`
}

type Weight struct {
	DailyWeight int    `json:"dailyWeight" validate:"required,gte=2,lte=700" `
	Date        string `json:"date,omitempty"`
}

type Water struct {
	WaterIntake int    `json:"waterIntake" validate:"required,gt=0,lte=20" `
	Date        string `json:"date,omitempty"`
}

type Message struct {
	Code    int    `json:"code"  validate:"required"`
	Message string `json:"message"  validate:"required"`
}

type UserID struct {
	UserID string `json:"userId"  validate:"required" `
}
type ChangePassword struct {
	CurrentPassword string `json:"CurrentPassword"  validate:"required" `
	NewPassword     string `json:"newPassword"  validate:"required" `
}
type Date struct {
	Date string `json:"date"  validate:"required"`
}
type ForgotPasswordInput struct {
	Email     string `json:"email" validate:"required,email" `
	EventType string `json:"eventType" validate:"required" `
}

type ValidateOTP struct {
	Email     string `json:"email" validate:"required,email" `
	EventType string `json:"eventType" validate:"required" `
	OTP       string `json:"otp" validate:"required" `
}
type YearlyWeight struct {
	Month                int     `json:"month"  validate:"required"`
	AverageMonthlyWeight float32 `json:"averageMonthlyWeight" validate:"required"`
}
type Token struct {
	Token string `json:"token"  validate:"required"`
}
type YearlyCaloriesBurned struct {
	Month                        int     `json:"month"  validate:"required"`
	AverageMonthlyCaloriesBurned float32 `json:"averageMonthlyCaloriesBurned" validate:"required"`
}
type SetNewPaswordInput struct {
	Email       string `json:"email" validate:"required,email" `
	EventType   string `json:"eventType" validate:"required" `
	Token       string `json:"token"  validate:"required"`
	NewPassword string `json:"newPassword"  validate:"required"`
}
