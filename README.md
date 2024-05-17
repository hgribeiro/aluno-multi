# CRUD in Go

This is a simple CRUD (Create, Read, Update, Delete) application implemented in Go.

## Prerequisites

- Go 1.16 or higher

## Getting Started

1. Clone the repository:

```shell
git clone https://github.com/your-username/your-repo.git
```

2. Change into the project directory:

```shell
cd your-repo
```

3. Build the application:

```shell
go build
```

4. Run the application:

```shell
./your-repo
```

## Usage

- Create a new item:

  ```shell
  curl -X POST -d '{"name":"New Item"}' http://localhost:8080/items
  ```

- Get all items:

  ```shell
  curl http://localhost:8080/items
  ```

- Get a specific item:

  ```shell
  curl http://localhost:8080/items/{id}
  ```

- Update an item:

  ```shell
  curl -X PUT -d '{"name":"Updated Item"}' http://localhost:8080/items/{id}
  ```

- Delete an item:

  ```shell
  curl -X DELETE http://localhost:8080/items/{id}
  ```

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
