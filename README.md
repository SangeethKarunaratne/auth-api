# GoLang Micro Service Boilerplate

This project follows a well-organized folder structure inspired by Clean Architecture, promoting modularity, maintainability, and separation of concerns. Below is an overview of the directory layout and its purposes.

## Folder Structure

```plaintext
project/
|-- app/
|   |-- config/
|   |   |-- config.go
|   |   |-- parser.go
|   |   |-- reader.go
|   |-- container/
|   |   |-- container.go
|   |   |-- resolver.go
|   |   |-- resolve_adapters.go
|   |   |-- resolve_repositories.go
|   |   |-- resolve_services.go
|   |-- http/
|   |   |-- controllers/
|   |   |   |-- user_controller.go
|   |   |-- errors/
|   |   |   |-- errors.go
|   |   |-- middleware/
|   |   |   |-- auth_middleware.go
|   |   |-- request/
|   |   |   |-- login_request.go
|   |   |-- response/
|   |   |   |-- error_response.go
|   |   |-- routes/
|   |   |   |-- router.go
|-- config/
|   |-- app.yaml
|   |-- database.yaml
|-- domain/
|   |-- adapters/
|   |   |-- db_adapter_interface.go
|   |-- entities/
|   |   |-- user.go
|   |-- repositories/
|   |   |-- user_repository_interface.go
|   |-- services/
|   |   |-- notification_service_api_interface.go
|   |-- usecases/
|   |   |-- authentication_usecase.go
|   |   |-- user_usecase.go
|-- external/
|   |-- adapters/
|   |   |-- mysql_adapter.go
|   |-- errors/
|   |   |-- adapter_error.go
|   |-- repositories/
|   |   |-- user_repository.go
|   |-- services/
|   |   |-- notification_service_api.go
|-- tests/
|-- go.mod
|-- go.sum
```
# Clean Architecture
The project's folder structure adheres closely to the principles of Clean Architecture, as outlined by Robert C. Martin (Uncle Bob). Clean Architecture emphasizes separating concerns into layers, with business logic at the center and infrastructure details at the outer layers. This approach promotes testability, maintainability, and flexibility.

For more information on Clean Architecture, refer to The Clean Architecture blog post by Uncle Bob. https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture

# Usage

* Clone the repository: git clone https://github.com/your-username/your-project.git
* Navigate to the project directory: cd your-project
* Install dependencies: go mod tidy
* Run the application: go run main.go