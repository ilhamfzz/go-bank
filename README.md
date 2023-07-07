# Golang Authentication

## Description
Implementing Register and authentication in Golang with email otp and withouth using jwt. The idea is to use a refrence id and otp to validate the user. the refrence id, otp, and time otp expired will be stored in cache. This is useful for matching the data submitted for user validation. After user validated, you can generate a token, This feature is usually used to login user. You will get token and using it with Bearer Authorization to get user data.

## Frameworks
- HTTP framework : Fiber
- SQL builder : Goqu
- Database : PostgreSQL
- Cache : BigCache
- Email : Mailtrap

## Features
### Register
This will register a new user and send an OTP to the user's email. But in this repository environtment, i am using sandbox (testing) email from mailtrap. So, you can't get the OTP. But you can see the OTP in the console. You can change the credential with prefix `MAIL_` in `.env` file.
#### Endpoint
```
http://localhost:8080/user/register
```
#### Method
```
POST
```
#### Input Body JSON
```
{
    "full_name": "<VARCHAR(255)>",
    "phone": "<VARCHAR(255)>",
    "email": "<VARCHAR(255)>",
    "username": "<VARCHAR(255)>",
    "password": "<VARCHAR(255)>"
}
```

#### Expected Output Body JSON
```
{
    "id": "<your registered id>",
    "full_name": "<your registered full name>",
    "phone": "<your registered phone>",
    "email": "<your registered email>",
    "username": "<your input username>",
    "refrence_id": "<your refrerence id>",
}
```

### Validate OTP
This will validate your account using refrence id and OTP that you can get frome result of register endpoint and your email (in this repo, you can get otp from console).
#### Endpoint
```
http://localhost:8080/user/validate-otp
```

#### Method
```
POST
```

#### Input Body JSON
```
{
    "refrence_id": "<your reference id>",
    "otp": "<your otp>"
}
```
### Generate Token
This will return a `token` that you can use to validate the user.
#### Endpoint
```
http://localhost:8080/token/generate
```
#### Method
```
POST
```
#### Input Body JSON
```
{
    "username": "<your username>",
    "password": "<your password>"
}
```

#### Expected Output Body JSON
```
{
    "token": "<your token>"
}
```

### Validate Token
#### Endpoint
```
http://localhost:8080/token/validate
```
#### Method
```
GET
```
#### Authorization
```
Bearer <token>
```

### 

## Documentation
