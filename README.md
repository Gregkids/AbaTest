# Technical Test - Backend IoT Developer

CRUD Design for User & Device Access Control API Using Golang with Fiber framework.

### Dependencies

Package used:
* github.com/gofiber/contrib/jwt
* github.com/gofiber/fiber/v2
* github.com/golang-jwt/jwt/v5

Database used: 
* PostgreSQL
* Schema of the database

![Screenshot 2025-07-07 093416](https://github.com/user-attachments/assets/8e0a08c0-831c-40f3-a41b-89a45f3c3840)

### Run and Test the system

* Run the golang program
```
go run main.go
```
* Hit the endpoint /login to get token
example:
input:
```
{
    "email": "useradmin@gmail.com",
    "password": "adminpass"
}
```
output:
```
{
    "code": 200,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTIxMTI1NzcsImlhdCI6MTc1MTg1MzM3Nywicm9sZSI6IkFkbWluIiwidXNlcklEIjoiQURNMDEifQ.K8C5r6U2rQE-ComYMny17nw11-3KJk_Zbs8TiZw_vh8"
}
```
* Use the token as bearer token for authorization

