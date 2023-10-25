package dal

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/cockroachdb/cockroach-go/crdb"
	_ "github.com/lib/pq"
)

var once sync.Once
var db *sql.DB

func Connect() (*sql.DB, error) {

	var err error
	once.Do(func() {
		connection_string := "postgresql://User:2a2dwrcFnaHmyS6I5mvE_A@solar-ape-6502.8nk.cockroachlabs.cloud:26257/fitnessdb?sslmode=verify-full"
		db, err = sql.Open("postgres", connection_string)
		if err != nil {
			fmt.Println("Database Connection err", err)
			return
		}

	})
	return db, err
}

func LogAndQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	fmt.Println(query)
	return db.Query(query, args...)
}

func MustExec(query string, args ...interface{}) (int64, error) {
	db := GetDB()
	result, err := db.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	RowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("RowsAffected Error", err)
		return 0,err
	}
	return RowsAffected,err

}
func InitDB(db *sql.DB) {
	// MustExec(db, "CREATE TABLE IF NOT EXISTS public.user_registration_details (   user_id INT8 NOT NULL DEFAULT unique_rowid(), email VARCHAR NOT NULL, password VARCHAR NOT NULL, CONSTRAINT user_registration_details_pk PRIMARY KEY (user_id ASC))")
	// MustExec(db, "CREATE TABLE IF NOT EXISTS public.user_profile_details ( user_id INT8 NOT NULL, age INT8 NOT NULL, gender VARCHAR NOT NULL, height INT8 NOT NULL, weight INT8 NOT NULL, health_goal VARCHAR NOT NULL, CONSTRAINT userdetails_pk PRIMARY KEY (user_id ASC), CONSTRAINT user_profile_details_fk FOREIGN KEY (user_id) REFERENCES public.user_registration_details(user_id) )")
	// MustExec(db, " CREATE TABLE IF NOT EXISTS public.exercise_details ( user_id INT8 NOT NULL, rowid INT8 NOT VISIBLE NOT NULL DEFAULT unique_rowid(), exercise_type VARCHAR NOT NULL, duration INT8 NOT NULL, calories_burned INT8 NOT NULL, date DATE NOT NULL, CONSTRAINT exercise_details_pkey PRIMARY KEY (rowid ASC), CONSTRAINT exercise_details_fk FOREIGN KEY (user_id) REFERENCES public.user_profile_details(user_id) );")
	// MustExec(db, "CREATE TABLE IF NOT EXISTS public.meal_details ( user_id INT8 NOT NULL, rowid INT8 NOT VISIBLE NOT NULL DEFAULT unique_rowid(), meal_type VARCHAR NOT NULL, ingredients VARCHAR NOT NULL, calories_consumed INT8 NOT NULL, date DATE NOT NULL, CONSTRAINT meal_details_pkey PRIMARY KEY (rowid ASC), CONSTRAINT meal_details_fk FOREIGN KEY (user_id) REFERENCES public.user_profile_details(user_id), CONSTRAINT check_meal_type CHECK (meal_type IN ('Breakfast':::STRING, 'Launch':::STRING, 'Snacks':::STRING, 'Dinner':::STRING)))")
}
func GetDB() *sql.DB {
	return db
}
