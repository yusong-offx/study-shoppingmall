DROP TABLE IF EXISTS Items;
DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS Venders;
DROP TABLE IF EXISTS Categories;
DROP TABLE IF EXISTS Orders;

CREATE TABLE Users (
	id varchar(20) PRIMARY KEY,
	password bytea NOT NULL,
	addr varchar(50),
	phone_number char(11),
	email varchar(50)
);

CREATE TABLE Venders (
	id varchar(20) PRIMARY KEY,
	name varchar(20) UNIQUE,
	password char(60) NOT NULL,
	addr varchar(50),
	phone_number char(11),
	email varchar(50),
	url varchar(500)
);

-- 카테고리 이름 50자 이상 되면 안됨.
CREATE TABLE Categories (
	cur varchar(50) PRIMARY KEY,
	prev varchar (50)
);

CREATE TABLE Items (
	id SERIAL PRIMARY Key,
	name varchar(500) UNIQUE NOT NULL,
	stock integer DEFAULT 0,
	price integer NOT NULL,
	content text,
	photo bytea,
	vender varchar(20) REFERENCES Venders(name) ON DELETE CASCADE,
	category varchar(50) REFERENCES Categories(cur)
);


CREATE TABLE Orders (
	order_id serial PRIMARY KEY,
	user_id varchar(20) REFERENCES Users(id) ON DELETE CASCADE,
	time timestamp DEFAULT CURRENT_TIMESTAMP,
	cart jsonb not null
);