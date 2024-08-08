# Person API

This is a REST API for managing person data, built using Go, MongoDB, and Clean Architecture.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
  - [Create a Person](#create-a-person)
  - [Update a Person](#update-a-person)
  - [Delete a Person](#delete-a-person)
  - [Get a Person by ID](#get-a-person-by-id)
  - [List Persons (with Pagination)](#list-persons-with-pagination)
- [Error Handling](#error-handling)
- [Tech Stack](#tech-stack)
- [Contributing](#contributing)
- [License](#license)

## Features

- Create, update, delete, and retrieve person data
- Pagination for listing persons
- Validation of person data
- Unique email constraint
- Customized error handling

## Installation

1. Clone the repository:

   ```shell
   git clone https://github.com/rpuglielli/person-api.git
   ```

2. Navigate to the project directory:

   ```shell
   cd person-api
   ```

3. Set up the environment:
   - Install Go (version 1.16 or higher)
   - Install MongoDB

4. Start the application:

   ```shell
   go run cmd/api/main.go
   ```

   The API will start running on `http://localhost:8080`.

## Usage

### Create a Person

**Endpoint:** `POST /persons`

**Request Body:**

```json
{
  "externalId": "abc123",
  "firstName": "John",
  "lastName": "Doe",
  "email": "john.doe@example.com",
  "phone": "1234567890",
  "category": "customer"
}
```

### Update a Person

**Endpoint:** `PUT /persons/{id}`

**Request Body:**

```json
{
  "firstName": "Jane",
  "email": "jane.doe@example.com"
}
```

### Delete a Person

**Endpoint:** `DELETE /persons/{id}`

### Get a Person by ID

**Endpoint:** `GET /persons/{id}`

### List Persons (with Pagination)

**Endpoint:** `GET /persons?page=1&pageSize=10`

**Response:**

```json
{
  "currentPage": 1,
  "data": [
    {
      "id": "abc123",
      "externalId": "abc123",
      "firstName": "John",
      "lastName": "Doe",
      "email": "john.doe@example.com",
      "phone": "1234567890",
      "category": "customer",
      "created": "2023-08-07T19:52:34Z",
      "updated": "2023-08-07T19:52:34Z"
    }
  ],
  "firstPageURL": "/persons?page=1",
  "from": 1,
  "lastPage": 100,
  "lastPageURL": "/persons?page=100",
  "nextPageURL": "/persons?page=2",
  "path": "/persons",
  "perPage": 10,
  "prevPageURL": null,
  "to": 10,
  "total": 1000
}
```

## Error Handling

The API uses custom error types to provide more meaningful error messages and HTTP status codes. The following error types are used:

- `ValidationError`: Returned for invalid input data (e.g., missing required fields).
- `ConflictError`: Returned when a resource (e.g., email) already exists.
- `NotFoundError`: Returned when a resource is not found.
- `InternalError`: Returned for internal server errors.

## Tech Stack

- Go (version 1.16 or higher)
- MongoDB
- Gin web framework
- go-uuid for generating UUIDs

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
