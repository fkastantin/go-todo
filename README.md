# TODO App - golang

This repo is created for learning purpose. The backend is developed using `GIN` framework using `golang`, `repository-service-controller` architecture and `dependency injection` principle are used, and `Docker` is used to host the API and mysql database.


## Model
Model is the layer where the data stored in the database is described, we have one entity `task`. `task.go` contains the struct with the name `Task` and the corresponding attributes. `TableName` method returns the name of the table is the database.

## Repository
Repository is the data access layer, it should provide all needed database operations (`CURD`). It uses `GORM` framework which is an ORM (object relational mapping) framework, it does the whole logic behind mapping between the database table and application objects.

We have one repository under `task_repositry.go` and it's powered with the CRUD operations for the `task` model.

## Service
Service layer is where the business logic. Since we do not have complicated business logic, it is more like a proxy for the data access layer.

## Controller
Controller is the responsible to return the appropriate response to the client.

## Dependency injection
Instances of the database pool, repository, service, controller, and route are created in the main function then they are injected into the objects that need them.
