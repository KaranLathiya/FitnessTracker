
# FitnessTracker

Fitness track website where a user can track their daily
exercises, meals, weight, and set of personal health goals.

Created using Go + cockroachdb.


## Deployment

Project also running on 

```bash
  https://fitnesstracker-k5h0.onrender.com
```


## Run Locally

Clone the project

```bash
  git clone https://github.com/KaranLathiya/FitnessTracker.git
```

Install Go version 1.21

```bash
  go install 1.21 
```

Install dependencies

```bash
  go mod tidy 
```
Set variables in .env file

Start the server

```bash
  go run server/server.go
```

# Routing

## For Signup 

To first time signup for new user  --POST

    http://localhost:8080/signup

## For Login

To login for user  --POST

    http://localhost:8080/login

## For UserProfile

To add/update user profile details for user --PUT

To fetch profile details for user --GET

    http://localhost:8080/user/profile

## For Change Password

To change password for user  --POST

    http://localhost:8080/user/change-password

## For Forgot Password

To get the otp in email for user  --POST

    http://localhost:8080/otp/request

To verify the otp for user  --POST

    http://localhost:8080/otp/verify
    
To set the new password after otp verification for user  --POST

    http://localhost:8080/forgot-password

## For Meal


To add meal details of user --POST

To update meal details of user --PUT

To delete meal details of user --DELETE

    http://localhost:8080/user/meal

## For Exercise


To add exercise details of user --POST

To update exercise details of user --PUT

To delete exercise details of user --DELETE

    http://localhost:8080/user/exercise

## For Weight

To add weight details of user --POST

To update weight details of user --PUT

To delete weight details of user --DELETE

    http://localhost:8080/user/weight

## For Water

To add water details of user --POST

To update water details of user --PUT

To delete water details of user --DELETE

    http://localhost:8080/user/water

## For All details by date 

To fetch monthly water details of user --GET
     
    http://localhost:8080/user/alldetails?date=2024-02-06
<!-- 
## For WaterIntake Monthly 

To fetch monthly water details of user --GET
     
    http://localhost:8080/user/water-intake-of-month -->

## For Weight Details Yearly 

To fetch yearly Weight Details of user --GET
     
    http://localhost:8080/user/yearly-weight-details?year=2024

## For  CaloriesBurned Details Yearly 

To fetch yearly CaloriesBurned details of user --GET
     
    http://localhost:8080/user/yearly-caloriesburned-details?year=2024