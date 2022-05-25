# RESTFUL API using Gin

## Project Structure
```
├── app -> Router and Database
├── controller -> Controller
├── exception -> Exception Handler
├── helper -> Helper Common Function
├── model -> Application Model
│   └── domain Entity
│       └── dto Data Transfer Object
│   └── web Web API Response
├── repository -> Repository
└── service -> Service
├── .gitignore
├── README.MD
├── api.http Plain HTTP Request
├── go.mod
├── go.sum
├── main.go Main Function
├── openapi.json Open API Specs

```

## Running Application

1. Clone this project
2. Setup Database connection
3. Copy .env.example to .env and fill information with your own credentials
4. Run this command in console
`go run main.go`
5. Or create a json config if u are using vscode like this one
```
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/main.go"
        }
    ]
}
```
6. And then Run the app with `ctrl + F5` in vscode
7. Change dir to db/migrations and run this command
`goose mysql "user:password@/dbname?parseTime=true" up`
