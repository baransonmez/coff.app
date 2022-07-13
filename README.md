# Coff.app - Clean Architecture

This project is a sample implementation for demonstrating the power of Clean Architecture written in GO.

## Motivation

My purpose in creating this project is to try a clean architecture implementation in Go language. Thanks to Clean
Architecture, I aimed to develop a more manageable application by separating business rules and infrastructure
dependencies.

## Structure

While planning this repository, my main priority was to separate between business rules and infrastructure codes.
Codes under Business package include domain, use-cases and output adapters.
Domain and use-cases are
written in a way that does not depend on output adapters.
The reason why output adapters are not moved to a different directory is that I
think it is more useful to use this structure when feature-based development.

Use-cases communicate with output adapters using interfaces. They are not directly dependent on output adapters, so I
have separated business rules from infrastructure codes.

```bash
.
.
├── business
│   ├── common
│   │   ├── command.go
│   │   └── entity.go
│   └── core
│       ├── coffee
│       │   ├── commands.go
│       │   ├── domain.go
│       │   ├── output_adapters
│       │   │   └── persistence
│       │   │       ├── db.go
│       │   │       ├── in_mem.go
│       │   │       └── models.go
│       │   ├── ports.go
│       │   └── usecases.go
.       .   .
.       .   .
```

I paid attention that the model used by the output adapters and the model I used in the domain were separate models. In
this way, the model I used in the domain became independent of the database used and its restrictions. With this
approach, I got ahead of database oriented design and made sure to think about business rules first.

On the other hand, I designed the use-cases to be accessible from outside the package. It was developed to take command
type as a parameter so that it doesn't care about any details about where it is used. In this way, it has been ensured
that no changes are made in the functions according to the place of use, that is, it is not affected by any
infrastructure restrictions.

```
├── app
│   ├── cli
│   │   └── main.go
│   └── web
│       ├── coffee
│       │   ├── input_adapters
│       │   │   └── handlers
│       │   │       └── api.go
│       │   └── main.go
│       ├── recipe
│       │   ├── input_adapters
│       │   │   └── handlers
│       │   │       └── api.go
│       │   └── main.go
│       └── user
│           ├── input_adapters
│           │   ├── grpc
│           │   │   ├── pb
│           │   │   │   ├── user.pb.go
│           │   │   │   └── user_grpc.pb.go
│           │   │   ├── protos
│           │   │   │   └── user.proto
│           │   │   └── server.go
│           │   └── handlers
│           │       └── api.go
│           └── main.go
.           .
.           .
```

The app package corresponds to incoming ports in hexagonal architecture. We can think of it as the place where use-cases are
triggered for use. At this point, for example, if we are using REST, we should convert the incoming request to the
relevant command according to the use-case to be used and create a response from the result of the use-case
in this layer.

## Credits

[Alistair Cockburn - Hexagonal Architecture](http://alistair.cockburn.us/Hexagonal+architecture)

[Robert C. Martin - Clean Architecture: A Craftsman’s Guide to Software Structure and Design](https://www.oreilly.com/library/view/clean-architecture-a/9780134494272/)

[Ardanlabs - Service Repo](https://github.com/ardanlabs/service)
