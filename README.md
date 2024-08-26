# RSSAGG
This is a simple CRUD (Create, Read, Update, Delete) application built using Go (Golang). The project is designed to help you understand the basics of building RESTful APIs and managing data with Go. It includes features to create, read, update, and delete resources using a RESTful API.

# How to run 
- clone the project
- ```go mod tidy```
- ```go get -u```
- ```go mod vendor```
- make a .env file and populate it with ``` PORT=<portNumber>
DB_URL=postgres://postgres:@localhost:5432/<dbName>?sslmode=disable```
- ```sql/schema ```
- ```goose postgres postgres://postgres:@localhost:5432/<dbName> up```
- ```cd ../../```
- ```sqlc generate```
- ```go build && ./rssagg```
