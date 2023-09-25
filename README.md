# Orders API and CLI

A Golang backend with redis database to manage orders and a CLI to retrieve orders information

The API is based on **Net Ninja** tutorial on youtube [Build a Microservice with Go](https://www.youtube.com/watch?v=wpnN3RIRSxs&list=PL4cUxeGkcC9iImF8w9FbFOc2UntutL9Wv&ab_channel=NetNinja)

## Project purpose

The main purpose is learning doing backend stuff in Golang and leverage Chi and Cobra to build an api and a cli client

## Usage

### run the api
```bash
# you need to run redis-server before running the api
cd api
go run .
```

base endpoint : `http://localhost:3000/orders`

| Method | description |
| ---- | ---- |
|`POST` | Create order |
| `GET` /{id} | get order by id |
| `GET`  | list all orders  |
| `PUT`| update order |
| `DELETE` /{id} | delete order by id |

### use the cli
```bash
cd cli

# list the orders id
go run . list

# get order info by id
go run . get $id
```