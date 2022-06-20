DROP TABLE IF EXISTS Users;
CREATE TABLE Users (
	id varchar(20) PRIMARY KEY,
	password bytea NOT NULL,
	addr varchar(50),
	phone_number char(11),
	email varchar(50)
);

DROP TABLE IF EXISTS Venders;
CREATE TABLE Venders (
	name varchar(20) PRIMARY KEY,
	password char(60) NOT NULL,
	addr varchar(50),
	phone_number char(11),
	email varchar(50),
	url varchar(500)
);

DROP TABLE IF EXISTS Categories;
CREATE TABLE Categories (
	prev varchar(50) UNIQUE,
	next varchar(50)
);

DROP TABLE IF EXISTS Items;
CREATE TABLE Items (
	id SERIAL PRIMARY Key,
	name varchar(500) UNIQUE NOT NULL,
	stock integer DEFAULT 0,
	price integer NOT NULL,
	content text,
	photo bytea,
	vender varchar(20) REFERENCES Venders(name) ON DELETE CASCADE,
	category varchar(50) REFERENCES Categories(prev) 
);

DROP TABLE IF EXISTS Orders;
CREATE TABLE Orders (
	order_id serial PRIMARY KEY,
	user_id varchar(50) REFERENCES Users(id) ON DELETE CASCADE,
	time timestamp DEFAULT CURRENT_TIMESTAMP,
	cart jsonb not null
);




-- JSONb로 수정