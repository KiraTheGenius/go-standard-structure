# go-standard-structure
Running the Server

To run the HTTP server:
```go
go run cmd/myApp/main.go
```
The server will start on http://localhost:8080 by default.

## Project Components
 ### cmd/
   + Packages that provide support for a specific program that is being built.
   + startup, shutdown and configuration.
     
> **Note:** A common convention is placing all commands in a repository into a `cmd` directory; while this isn’t strictly necessary in a repository that consists only of commands, it’s very useful in a mixed repository that has both commands and importable packages, as we will discuss next.

 ### internal/
   + Packages that provide support for the different programs the project owns.
   + CRUD, services or business logic.
  
  Contains the core server logic, including request handlers.
  ### internal/platform/
  + Packages that provide internal foundational support for the project.
  + database, authentication or marshaling.
  
### pkg/
  + Contains utility functions or packages that are intended to be used by other applications.
  + Unlike internal, this code is accessible to other projects.

### external/someapi
  + Integrations with external APIs.
  
### vendor
 + Contains project-specific dependencies.(third-party packages)
 + Managed by running go mod vendor.
  ```go 
    go mod vendor
  ```

## Refrences
+ https://go.dev/doc/modules/layout
+ https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html
+ https://github.com/golang-standards/project-layout
