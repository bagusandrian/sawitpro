# Sawit Pro

This repo provide API using golang. 

## Overview

This service provides 4 endpoints (registration, login, get profile, and update profile). 

## Prerequisites

- go with min version 1.16
- docker (if u want using docker)

## Getting Started

### Installation

1. Run `make init` to initialize the repository.
2. Run `make test` to run the unit tests.
3. Run `make run` to start the application on your local machine (make sure to run `make init` first) will be reserve port `:8002`, u can change the port on `files/etc/developmen/config.yaml`.
4. Run `docker-compose up` to start the application, including the database.

Your application should be accessible through the base URL [http://localhost:8080/](http://localhost:8080/).

## API Endpoints

### 1. `/registration` [POST]

#### Request

- **Full Name**: (Minimum 3, Maximum 60 characters)
- **Phone Number**: (Minimum 10, Maximum 13 characters, Indonesian code +62)
- **Password**: (Minimum 6, Maximum 64 characters, at least 1 capital, 1 number, and 1 special character)

#### Response

Include information about the expected response and any error codes.

### 2. `/login` [POST]

#### Request

- **Phone Number**
- **Password**

#### Response

Include information about the expected response and any error codes.

### 3. `/get_my_profile` [GET]

#### Request

- **Authorization Header**: Bearer Token (JWT)

#### Response

Include information about the expected response and any error codes.

### 4. `/update_profile` [POST]

#### Request

- **Authorization Header**: Bearer Token (JWT)
- **Phone Number OR Full Name**: Update the fields that exist in the request

#### Response

Include information about the expected response and any error codes.

## Makefile Commands

- `make init`: Initialize the repository.
- `make test`: Run unit tests.
- `make run`: Run the application on your local machine.
- `docker-compose up`: Start the application, including the database.
