# Go Hexagonal Architecture

This project is intended to be a practice about feature flags, module access and other stuff related to authorization in an application make using go

## Project Structure

This project is built on top of the [Gin Web Framework]() and it is structured as follows:

- Hexagonal Architecture
  - `database/`: Contains the database connection and the repository implementation
  - `pkg/`: Contains all the modules of the application
    - `auth/`: Contains the authentication module
      - `domain/`: Contains the domain entities
      - `application/`: Contains the application services
      - `data/` Contains the repository implementation
    - `feature/`: Contains the feature module
      - ...
    - `user/`: Contains the user module
      - ...

## Running the project

To run the project you need to have a postgres database running on your machine. You can use the following command to start a postgres database using docker:

```bash
docker compose up -d
```

After that, you can run the project using the following command:

```bash
# Install the dependencies
go mod tidy

# Run the project
air
```
