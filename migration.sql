-- public.otp_details definition

-- Drop table

-- DROP TABLE public.otp_details;

CREATE TABLE public.otp_details (
	email VARCHAR NOT NULL,
	otp VARCHAR NOT NULL,
	event_type VARCHAR NOT NULL,
	expires_at TIMESTAMP NOT NULL,
	CONSTRAINT otp_details_pk PRIMARY KEY (email ASC, otp ASC),
	CONSTRAINT check_event_type CHECK (event_type IN ('forgot_password':::STRING,))
);


-- public.token_details definition

-- Drop table

-- DROP TABLE public.token_details;

CREATE TABLE public.token_details (
	email VARCHAR NOT NULL,
	event_type VARCHAR NOT NULL,
	token VARCHAR NOT NULL,
	CONSTRAINT token_details_pk PRIMARY KEY (email ASC, event_type ASC),
	CONSTRAINT check_event_type CHECK (event_type IN ('forgot_password':::STRING,))
);


-- public.user_details definition

-- Drop table

-- DROP TABLE public.user_details;

CREATE TABLE public.user_details (
	user_id VARCHAR NOT NULL DEFAULT unique_rowid(),
	email VARCHAR NOT NULL,
	password VARCHAR NOT NULL,
	fullname VARCHAR NOT NULL,
	age INT8 NULL,
	gender VARCHAR NULL,
	height FLOAT8 NULL,
	weight FLOAT8 NULL,
	health_goal VARCHAR NULL,
	profile_photo VARCHAR NULL,
	CONSTRAINT user_details_pk PRIMARY KEY (user_id ASC),
	UNIQUE INDEX user_details_un (email ASC),
	CONSTRAINT check_age CHECK ((age > 0:::INT8) AND (age < 130:::INT8)),
	CONSTRAINT check_height CHECK ((height > 50.0:::FLOAT8) AND (height < 300.0:::FLOAT8)),
	CONSTRAINT check_weight CHECK ((weight > 2.0:::FLOAT8) AND (weight < 700.0:::FLOAT8)),
	CONSTRAINT check_health_goal CHECK (health_goal IN ('weight_loss':::STRING, 'weight_gain':::STRING, 'muscle_building':::STRING, 'maintain_body':::STRING)),
	CONSTRAINT check_gender CHECK (gender IN ('male':::STRING, 'female':::STRING, 'other':::STRING))
);


-- public.exercise_details definition

-- Drop table

-- DROP TABLE public.exercise_details;

CREATE TABLE public.exercise_details (
	user_id VARCHAR NOT NULL,
	exercise_type VARCHAR NOT NULL,
	duration INT8 NOT NULL,
	calories_burned FLOAT8 NOT NULL,
	date DATE NOT NULL,
	CONSTRAINT exercise_details_pk PRIMARY KEY (user_id ASC, exercise_type ASC, date ASC),
	CONSTRAINT exercise_details_fk FOREIGN KEY (user_id) REFERENCES public.user_details(user_id),
	CONSTRAINT check_duration CHECK ((duration > 0:::INT8) AND (duration < 1440:::INT8)),
	CONSTRAINT check_calories_burned CHECK ((calories_burned > 0.0:::FLOAT8) AND (calories_burned < 20000.0:::FLOAT8)),
	CONSTRAINT exercise_details_check CHECK (exercise_type IN ('weight_lifting':::STRING, 'walking':::STRING, 'running':::STRING, 'gym':::STRING, 'yoga':::STRING))
);


-- public.meal_details definition

-- Drop table

-- DROP TABLE public.meal_details;

CREATE TABLE public.meal_details (
	user_id VARCHAR NOT NULL,
	meal_type VARCHAR NOT NULL,
	ingredients VARCHAR NOT NULL,
	calories_consumed FLOAT8 NOT NULL,
	date DATE NOT NULL,
	CONSTRAINT meal_details_pk PRIMARY KEY (user_id ASC, meal_type ASC, date ASC),
	CONSTRAINT meal_details_fk FOREIGN KEY (user_id) REFERENCES public.user_details(user_id),
	CONSTRAINT check_calories_consumed CHECK ((calories_consumed > 0.0:::FLOAT8) AND (calories_consumed < 20000.0:::FLOAT8)),
	CONSTRAINT meal_details_check CHECK (meal_type IN ('breakfast':::STRING, 'lunch':::STRING, 'snacks':::STRING, 'dinner':::STRING))
);


-- public.water_details definition

-- Drop table

-- DROP TABLE public.water_details;

CREATE TABLE public.water_details (
	user_id VARCHAR NOT NULL,
	water_intake FLOAT8 NOT NULL,
	date DATE NOT NULL,
	CONSTRAINT newtable_pk PRIMARY KEY (user_id ASC, date ASC),
	CONSTRAINT newtable_fk FOREIGN KEY (user_id) REFERENCES public.user_details(user_id),
	CONSTRAINT check_water_intake CHECK ((water_intake > 0.0:::FLOAT8) AND (water_intake < 20.0:::FLOAT8))
);


-- public.weight_details definition

-- Drop table

-- DROP TABLE public.weight_details;

CREATE TABLE public.weight_details (
	user_id VARCHAR NOT NULL,
	daily_weight FLOAT8 NOT NULL,
	date DATE NOT NULL,
	CONSTRAINT weight_details_pk PRIMARY KEY (user_id ASC, date ASC),
	CONSTRAINT weight_details_fk FOREIGN KEY (user_id) REFERENCES public.user_details(user_id),
	CONSTRAINT check_daily_weight CHECK ((daily_weight > 2.0:::FLOAT8) AND (daily_weight < 700.0:::FLOAT8))
);