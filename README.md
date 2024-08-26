# go-bun-rest-api

- 💻 Example rest api in module structure .

- 💻 Use httprouter with Gin framework.

- 💻 Connect PostgresSQL with Bun.

- 💻 Simple Deploy service on Docker with docker-compose.

## 🚀 Get Started
```bash
# init project
go mod vendor

# create file .env then copy data in .env.example to .env

# migrate product table to postgresSQL
go run . migrate-product

# run http with cobra command
go run . http
```

## 🚀 Route api

- **POST** `/product/create`
  - **Description:** Create a new product.

- **PATCH** `/product/:id`
  - **Description:** Update an existing product by ID.

- **DELETE** `/product/:id`
  - **Description:** Delete a product by ID.

- **GET** `/product/:id`
  - **Description:** Get a product by ID.

- **GET** `/product/list`
  - **Description:** List all products.

## 🚀 Deploy on docker

```bash
# run container with docker-compose
docker-compose up -d
```