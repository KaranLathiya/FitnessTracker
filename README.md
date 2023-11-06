
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

## For Meal

To fetch meal details of user --GET

To add meal details of user --POST

To update meal details of user --PUT

To delete meal details of user --DELETE

    http://localhost:8080/user/meal

## For Exercise

To fetch exercise details of user --GET

To add exercise details of user --POST

To update exercise details of user --PUT

To delete exercise details of user --DELETE

    http://localhost:8080/user/exercise

## For Weight

To fetch weight details of user --GET

To add weight details of user --POST

To update weight details of user --PUT

To delete weight details of user --DELETE

    http://localhost:8080/user/weight

## For Water

To fetch water details of user --GET

To add water details of user --POST

To update water details of user --PUT

To delete water details of user --DELETE

    http://localhost:8080/user/water