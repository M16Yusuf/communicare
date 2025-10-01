# Communicare Your Social App

![badge golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![badge postgresql](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![badge redis](https://img.shields.io/badge/redis-%23DD0031.svg?&style=for-the-badge&logo=redis&logoColor=white)

communicate, your communication app. This project is backend for social app "communicare" build with gin gonic, postgreSQL and redis cache

## Entity-Relationship Diagram Communicare

![erd communicare](/erd-communicare.png)

## 🔧 Tech Stack

- [Go](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Redis](https://redis.io/docs/latest/operate/oss_and_stack/install/archive/install-redis/install-redis-on-windows/)
- [JWT](https://github.com/golang-jwt/jwt)
- [argon2](https://pkg.go.dev/golang.org/x/crypto/argon2)
- [migrate](https://github.com/golang-migrate/migrate)
- [Docker](https://docs.docker.com/engine/install/ubuntu/#install-using-the-repository)
- [Swagger for API docs](https://swagger.io/) + [Swaggo](https://github.com/swaggo/swag)

## 🗝️ Environment

```bash
# database
DBUSER=<your_database_user>
DBPASS=<your_database_password>
DBNAME=<your_database_name
DBHOST=<your_database_host>
DBPORT=<your_database_port>

# JWT hash
JWT_SECRET=<your_secret_jwt>
JWT_ISSUER=<your_jwt_issuer>

# Redish
RDB_HOST=<your_redis_host>
RDB_PORT=<your_redis_port>
RDB_USER=<your_redis_user>
RDB_PWD=<your_redis_password>
```

## ⚙️ Installation From Clone Repository

1. Clone the project

```sh
$ https://github.com/M16Yusuf/communicare.git
```

2. Navigate to project directory

```sh
$ cd communicare
```

3. Install dependencies

```sh
$ go mod tidy
```

4. Setup your [environment](##-environment)

5. Install [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation) for DB migration

6. Do the DB Migration

```sh
$ migrate -database YOUR_DATABASE_URL -path ./db/migrations up
```

or if u install Makefile run command

```sh
$ make migrate-createUp
```

7. Run the project

```sh
$ go run ./cmd/main.go
```

## ⚙️ Installation From Image GHCR

1. Pull image

```sh
$ docker pull ghcr.io/m16yusuf/communicare:1.2
```

2. Run on your container and setup [.env](##-environment)

## 🚧 API Documentation

| Method | Endpoint                | Body                                                                                 | Description                  |
| ------ | ----------------------- | ------------------------------------------------------------------------------------ | ---------------------------- |
| POST   | /auth/register          | email:string, password:string                                                        | Register                     |
| POST   | /auth/login             | email:string, password:string                                                        | login registered user        |
| DELETE | /auth/logout            | header: Authorization (token jwt)                                                    | Logout                       |
| POST   | /social/follow/:user_id | header: Authorization (token jwt) user_id:string                                     | follow a user                |
| GET    | /social/post            | header: Authorization (token jwt)                                                    | get post from followed users |
| POST   | /social/post            | header: Authorization (token jwt) comment:string, is_like:boolean, post_id:string    | like or comment on a post    |
| POST   | /users                  | header: Authorization (token jwt), fullname:string, bio:string, profile_picture:file | update profile login user    |
| POST   | /users/post             | header: Authorization (token jwt), caption:string, photo:file                        | create post for login user   |

## 📄 LICENSE

MIT License

Copyright (c) 2025 Muhammad Yusuf @m16yusuf

## 📧 Contact Info

[![LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/m16yusuf/)
[![Instagram](https://img.shields.io/badge/Instagram-E4405F?style=for-the-badge&logo=Instagram&logoColor=white)](https://www.instagram.com/M16Yusuf/)
[![Twitter](https://img.shields.io/badge/Twitter-0077b5?style=for-the-badge&logo=Twitter&logoColor=white)](https://twitter.com/M16Yusuf)
[![Facebook](https://img.shields.io/badge/Facebook-1877F2?style=for-the-badge&logo=facebook&logoColor=white)](https://facebook.com/m16yusuff)

## 🎯 Related Project
