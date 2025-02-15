# go-crud-postgresql
Implementing CRUD operations in Go using PostgreSQL as the database

To run this code, make sure to follow the below steps:

Install postgresql in your local system and login as the default user.

Command to login as root user: psql -U <username>
It will ask for the passowrd. Provide the <password>

Create a database with desired name using command:
CREATE DATABASE <database_name>;

User thi database by giving: \c <database_name>;

Create a table in this database: CREATE TABLE <tablename>(stockid SERIAL PRIMARY KEY, name VARCHAR(200) NOT NULL, price INT NOT NULL, company VARCHAR(250) NOT NULL);

Install required Go packages:
go get -u github.com/gorilla/mux
go get -u github.com/joho/godotenv
go get -u github.com/lib/pq

Create a .env file in root directory and add the following line from .env.example file. Make sure to edit the URL and add username, password and database name.

Enter commands from root directory:
go mod tidy
go build
go run .

Open postman and start your server as localhost on port 8080

Available Routes:

GET - /api/stock - List all stock details
GET - api/stock/{id} - GET stock by id
POST - api/createstock - Add a new stock
PUT - api/updatestock/{id} - Update a stock
DELETE - api/deletestock/{id} - DELETE a stock
