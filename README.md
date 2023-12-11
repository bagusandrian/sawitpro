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
| Field | Data type | Info | Mandatory |
| ------| --------- | ---- | --------- |
| fullname | string | full name of user. Minimum 3, Maximum 60 characters | YES |
| phonenumber | string | phone number of user. Minimum 10, Maximum 13 characters, Indonesian code +62 | YES |
| password | string | password of user. Minimum 6, Maximum 64 characters, at least 1 capital, 1 number, and 1 special character | YES |

#### example request curl: 
```
curl --location 'localhost:8002/registration' \
--form 'fullname="budi"' \
--form 'phonenumber="+62812345678"' \
--form 'password="I23a$6j890"'
```

#### example response: 
```
{
    "data": {
        "id": 0,
        "fullname": "budi",
        "phonenumber": "+62812345678",
        "error_message": null
    }
}
```

### 2. `/login` [POST]

On this endpoint, u need put `phonenumber` and `password` to get `JWTToken` and will be use for `/get_my_profile` and `/get_my_profile`. 

#### Request

| Field | Data type | Info | Mandatory |
| ------| --------- | ---- | --------- |
| phonenumber | string | phone number of user. Minimum 10, Maximum 13 characters, Indonesian code +62 | YES |
| password | string | password of user. Minimum 6, Maximum 64 characters, at least 1 capital, 1 number, and 1 special character | YES |

#### example request curl: 
```
curl --location 'localhost:8002/login' \
--form 'phonenumber="+62812345678"' \
--form 'password="I23a$6j890"'
```

#### example response: 
```
{
    "data": {
        "ID": 2,
        "JWTToken": "exampleOfJWTToken"
    }
}
```

### 3. `/get_my_profile` [GET]

#### Request

On this endpoint u need add `header` with key `Authorization` and value coming from `JWTToken` as mandatory field. 

#### example request curl: 
```
curl --location 'localhost:8002/get_my_profile' \
--header 'Authorization: Bearer tokenFromLoginEndpointJWTToken'
```

#### example response: 
```
{
    "data": {
        "FullName": "budi",
        "PhoneNumber": "+62812345678"
    }
}
```

### 4. `/update_profile` [POST]

On this endpoint, u can edit phone number and fullname of user.

#### Request

| Field | Data type | Info | Mandatory |
| ------| --------- | ---- | --------- |
| fullname | string | full name of user. Minimum 3, Maximum 60 characters | YES |
| phonenumber | string | phone number of user. Minimum 10, Maximum 13 characters, Indonesian code +62 | YES |

#### example request curl: 
```
curl --location 'localhost:8002/update_profile' \
--header 'Authorization: Bearer tokenFromLoginEndpointJWTToken' \
--form 'phonenumber="+62812345678"' \
--form 'fullname="samsul ganti neh"'
```

#### example response: 
```
{
    "data": {
        "FullName": "samsul ganti neh",
        "PhoneNumber": "+62812345678",
        "SuccessMessage": "success",
        "ErrorMessage": ""
    }
}
```
## Makefile Commands

- `make init`: Initialize the repository.
- `make test`: Run unit tests.
- `make run`: Run the application on your local machine.
- `docker-compose up`: Start the application, including the database.
