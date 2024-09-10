# go-assigment

The purpose of this project is to solve a training assigment in go.
The details of this assignment can be found in [data_api.md](data_api.md)

## Setup

In order to set up the application, a database has to be created and linked with the code. Currently, the application used an MYSQL database.

## Configuration

After creating the database, there are a few changes in the code that are needed to properly configure it.

In the [docker](docker-compose.yml) file, the envinronment and volumes have to be changed to match your own database, such as the name of the database and the password. Moreover, this change also has to happen in initialization function in [data_structures.go](data_structures.go), where the database connection is created.

## Running the app
```console
go run .
```

Simply runnign this command in the terminal will start the server.
Interacting with the features in the application is done through the endpoints created. Their description and use cases are located in [data_api.md](data_api.md)