# Hexagonal Architecture Template
This is an API project template using Hexagonal Architecture in Go.

### What is the concept?

Architecture proposed in 2005 to:
* Separate concerns
    * Each component has a well-defined responsibility
* Focus on the business logic
    * With the separation of components and layers we have a better details about the application business logic
* Parallelize work
    * With the responsibility well defined, it is easier to parallelize the work
* Isolate tests
    * With less dependencies between components, writing quality tests is easier
* Infrastructure changes
    * With business logic properly separated from the external layer comunication, the changes are less painful, for example changing a database to another

### What it does?
It makes a project scalable, easy to maintain and productive when implemented.

**Layers:**

* CORE
    * In this architectural model, everything is built around the core. Think of the core as a hexagon structure, which is where the business rule is. Everything outside the hexagon is seen as the external world
    * When we are in the core planning phase, we must refrain from technical decisions (which database, relationship between tables, framework, REST/gRPC, etc)
    * This layer must be 100% agnostic to any technical decision and must be 100% focused on solving the problem so that the application is born
* ACTORS
    * They are “things” from the world outside the core, which can be people, databases or other applications. The actors are divided into two groups:
        * DRIVER: Trigger some interaction with the core of the application (e.g. people)
        * DRIVEN: They expect communication from the core (e.g. database, queue, external APIs)
* PORTS
    * They are interfaces defined in the core to communicate with the external world, and just like the actors, there are two types of ports
        * DRIVER: These are actions exposed to the eyes of the driver through functions and methods
        * DRIVEN: These are well-defined communication interfaces that must be implemented by driven actors
* ADAPTERS
    * These are components responsible for integrating the core with the external world. The role of these components is to translate what the driver actor wants into what the core understands, as well as what the core wants into what the driven actor understands
        * DRIVER: Translate requests into core service calls
        * DRIVEN: Translate the core request into what its actor understands, for example, SQL
* DEPENDENCY INJECTION
    * Used to not create dependency between the core and its adapters
    * Technique that consists of using interfaces as parameters of functions and methods instead of directly importing certain components
    * Technique that also facilitates the creation of tests

### Project Structure
```
├── cmd
│   └── api
│       └── main.go
├── internal
│   ├── core
│   │   ├── domain
│   │   │   └── message.go
│   │   ├── dto
│   │   │   ├── create_message.go
│   │   │   └── get_message.go
│   │   ├── ports
│   │   │   ├── message_repository.go
│   │   │   └── message_servicer.go
│   │   └── usecases
│   │       └── message
│   │           ├── message_service.go
│   │           └── message_service_test.go
│   ├── handlers
│   │   ├── message_handler.go
│   │   ├── message_handler_test.go
│   │   └── server.go
│   └── repositories
│       └── message
│           ├── memory_repository.go
│           └── memory_repository_test.go
├── pkg
│   ├── apperrors
│   │   └── apperrors.go
│   └── identifier
│       └── uuid_generator.go
└── test
    └── mocks
        ├── message_repository_mock.go
        ├── message_servicer_mock.go
        └── uuid_generator_mock.go
```