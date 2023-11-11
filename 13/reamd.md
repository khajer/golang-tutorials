# run database with docker
docker run --name db_postgres -d -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 postgres 

docker exec -it db_postgres psql -U postgres
docker exec -it db_postgres psql -U postgres -d postgres

\l # list database 
CREATE DATABASE members;


\connect members

CREATE TABLE members (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL 
);
\dt

insert into members("id", "name", "email") values(2, 'itsara', 'khajerx@gmail.com');



go get github.com/lib/pq