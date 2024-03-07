package models

// Users example
type Users struct {
	Email        string   `json:"email" validate:"required,email" `
	FullName     string   `json:"fullName" validate:"required" `
	Age          *int     `json:"age" validate:"required,gt=0,lte=130" `
	Gender       *string  `json:"gender" validate:"required,oneof=male female other" `
	Height       *float32 `json:"height" validate:"required,gte=50,lte=300" `
	Weight       *float32 `json:"weight" validate:"required,gte=2,lte=700" `
	HealthGoal   *string  `json:"healthGoal" validate:"required,oneof=weight_loss weight_gain muscle_building maintain_body" `
	ProfilePhoto *string  `json:"profilePhoto"  `
}

// UserSignup example
type UserSignup struct {
	Email    string `json:"email" validate:"required,email" `
	FullName string `json:"fullName" validate:"required" `
	Password string `json:"password" validate:"required" `
}

// UserLogin example
type UserLogin struct {
	Email    string `json:"email" validate:"required,email" `
	Password string `json:"password" validate:"required" `
}

// Exercise example
type Exercise struct {
	ExerciseType   string  `json:"exerciseType" validate:"required,oneof=weight_lifting walking running gym yoga" `
	Duration       int     `json:"duration" validate:"required,gt=0,lte=1440"`
	CaloriesBurned float32 `json:"caloriesBurned" validate:"required,gt=0,lte=20000"`
	Date           string  `json:"date,omitempty"` 
}

// Meal example
type Meal struct {
	MealType         string  `json:"mealType" validate:"required,oneof=breakfast lunch snacks dinner" `
	Ingeredients     string  `json:"ingredients" validate:"required" `
	CaloriesConsumed float32 `json:"caloriesConsumed" validate:"required,gte=0,lte=20000" `
	Date             string  `json:"date,omitempty"`
}

// MealType example
type MealType struct {
	MealType string `json:"mealType" validate:"required,oneof=breakfast lunch snacks dinner" `
}

// ExerciseType example
type ExerciseType struct {
	ExerciseType string `json:"exerciseType" validate:"required,oneof=weight_lifting walking running gym yoga" `
}

// Weight example
type Weight struct {
	DailyWeight float32 `json:"dailyWeight" validate:"required,gte=2,lte=700" `
	Date        string  `json:"date,omitempty"`
}

// Water example
type Water struct {
	WaterIntake float32 `json:"waterIntake" validate:"required,gt=0,lte=20" `
	Date        string  `json:"date,omitempty"`
}

// Message example
type Message struct {
	Code    int    `json:"code"  validate:"required"`
	Message string `json:"message"  validate:"required"`
}

// UserID example
type UserID struct {
	UserID string `json:"userId"  validate:"required" `
}

// ChangePassword example
type ChangePassword struct {
	CurrentPassword string `json:"currentPassword"  validate:"required" `
	NewPassword     string `json:"newPassword"  validate:"required" `
}

// Date example
type Date struct {
	Date string `json:"date"  validate:"required"`
}

// RequestOTP example
type RequestOTP struct {
	Email     string `json:"email" validate:"required,email" `
	EventType string `json:"eventType" validate:"required,oneof=forgot_password" `
}

// VerifyOTP example
type VerifyOTP struct {
	Email     string `json:"email" validate:"required,email" `
	EventType string `json:"eventType" validate:"required" `
	OTP       string `json:"otp" validate:"required" `
}

// YearlyWeight example
type YearlyWeight struct {
	Month                int     `json:"month"  validate:"required"`
	AverageMonthlyWeight float32 `json:"averageMonthlyWeight" validate:"required"`
}

// Token example
type Token struct {
	Token string `json:"token"  validate:"required"`
}

// YearlyCaloriesBurned example
type YearlyCaloriesBurned struct {
	Month                        int     `json:"month"  validate:"required"`
	AverageMonthlyCaloriesBurned float32 `json:"averageMonthlyCaloriesBurned" validate:"required"`
}

// SetNewPaswordInput example
type SetNewPaswordInput struct {
	Email       string `json:"email" validate:"required,email" `
	EventType   string `json:"eventType" validate:"required,oneof=forgot_password" `
	Token       string `json:"token"  validate:"required"`
	NewPassword string `json:"newPassword"  validate:"required"`
}

type AllDetails struct{
	MealDetails []Meal `json:"mealDetails" `
	ExerciseDetails []Exercise `json:"exerciseDetails" `
	WaterDetails []Meal `json:"waterDetails" `
	WeightDetails []Exercise `json:"weightDetails" `
}
