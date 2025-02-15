# go-crud-postgresql
Implementing CRUD operations in Go using PostgreSQL as the database<br>

To run this code, make sure to follow the below steps:<br>

Install postgresql in your local system and login as the default user.<br>

Command to login as root user: psql -U <username><br>
It will ask for the password. Provide the <password><br>

Create a database with desired name using command:<br>
CREATE DATABASE <database_name>;<br>

Use this database by giving: \c <database_name>;<br>

Create a table in this database: CREATE TABLE <tablename>(stockid SERIAL PRIMARY KEY, name VARCHAR(200) NOT NULL, price INT NOT NULL, company VARCHAR(250) NOT NULL);<br>

Install required Go packages:<br>
go get -u github.com/gorilla/mux<br>
go get -u github.com/joho/godotenv<br>
go get -u github.com/lib/pq<br>

Create a .env file in root directory and add the following line from .env.example file. Make sure to edit the URL and add <username>, <password> and <database_name>.<br>

Enter commands from root directory:<br>
go mod tidy<br>
go build<br>
go run .<br>

Open postman and start your server as localhost on port 8080<br>

Available Routes:<br>

GET - /api/stock - List all stock details<br>
GET - api/stock/{id} - GET stock by id<br>
POST - api/createstock - Add a new stock<br>
PUT - api/updatestock/{id} - Update a stock<br>
DELETE - api/deletestock/{id} - DELETE a stock
