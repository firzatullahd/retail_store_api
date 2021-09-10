# Retail Store API

Retail Store Ecommerce API
- [Swagger Documentation](https://app.swaggerhub.com/apis-docs/a4443/retail_store_api/2.0#/)

## Overview

Retail Store API, a simple minimalistic ecommerce REST API written in Golang and built with Echo, Gorm, and MySQL, showcasing five major functionalities:

1. Authentication (Register and Login User)
2. Products listing, and products listing by category
3. Shopping cart feature (add and delete products)
4. Order placements and payment transaction
5. Access restrictions (User and, Admin only routes)

Database Schema Design
![Database Schema Design](/docs/erd.png)

## Getting Started

### 1.1 Prerequisites

To get started, ensure that you have the following installed on your local machine:

- [Golang](https://golang.org/dl/)
- [MySQL](https://www.mysql.com/downloads/)

### 1.2. Run locally

- Clone repository or clone your own fork

  ```bash
  git clone https://https://github.com/nurjamil/retail_store_api.git
  ```

- Make a duplicate of `.env.example` and rename to `.env`, then configure your credentials.
- Install dependencies by running `go mod tidy` on your terminal.
- Run command: `go run main.go` to start the server on `localhost:3000`
  -. run command `go test ./controller/ -cover` to run unit testing

## HTTP requests

There are 4 basic HTTP requests that you can use in this API:

- `POST` Create a resource
- `PUT` Update a resource
- `GET` Get a resource or list of resources
- `DELETE` Delete a resource

## HTTP Responses

Each response will include a code(repsonse code),message,status and data object that can be single object or array depending on the query.

## HTTP Response Codes

Each response will be returned with one of the following HTTP status codes:

- `200` `OK` The request was successful
- `400` `Bad Request` There was a problem with the request (security, malformed)
- `401` `Unauthorized` The supplied API credentials are invalid
- `403` `Forbidden` The credentials provided do not have permissions to access the requested resource
- `404` `Not Found` An attempt was made to access a resource that does not exist in the API
- `500` `Server Error` An error on the server occurred
