# technical-test-golang

## Context

This application is part of a technical exercise for an internship interview.

## Description

Simple REST API for ticketing.

## Get started

To develop this project, I used mysql, go and [gorm](http://gorm.io/)

### Env variables

|      Variable     |      Mandatory     |
|:-----------------:|:------------------:|
| DATABASE_NAME     |         YES        |
| DATABASE_PORT     |         YES        |
| DATABASE_HOST     |         YES        |
| DATABASE_USER     |         YES        |
| DATABASE_PASSWORD |         YES        |
| PORT              | NO (Default: 8080) |

### Run

```bash
# Get all dependencies
make deps

# Build
make build

# Build for linux
make build-linux
```


## Database schema

![Database schema](docs/mysql-schema.jpg)

## API Doc

[Postman API documentation](https://documenter.getpostman.com/view/3951542/RWEZSiKK)

