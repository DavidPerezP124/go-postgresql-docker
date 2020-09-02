# Description

This project is a *Go language* learning project with simple RestFul services. It uses postgres db inside with docker-compose. You can compose with dockerfile or create your own postgres database without it.

For run docker-compose, you need to write following commands. In your project folder,

```console
      cd docker
      docker-compose up
```

then PostgreSQL works on 32300 Port (32300 -> 5432). You can access with database IDE (DataGrip, Intellij etc.) with configure port 32300.

If you want to conncect from your host system type the following command to terminal.

```console
      psql -h localhost -p 32300 -d docker -U docker --password
```

For more information about it,

[Dockerize PostgreSQL](https://docs.docker.com/engine/examples/postgresql_service/#connecting-from-your-host-system)

## Database access configuration inside code

Create a .env file and use the followig parameters. Note that if you change this you should also change it in the docker/Dockerfile-Postgres to be able to access the connection.

```text
      DB_USER     = "docker"
      DB_PASSWORD = "changethis"
      DB_NAME     = "example"
      PORT = "32770"
```

## Testing API

You can start the CRUD operations with following URL.

```console
    curl  http://127.0.0.1:3000
```

## URL's and Example

List all of user (Need To Use GET method)

```console
    curl http://127.0.0.1:3000/getAll
```

Add new User with JSON type ((Need To Use POST method))

```console 

curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"name": "mockName","surname": "mockSurname","age": 30}' \
 http://127.0.0.1:3000/newUser

```

List one user with the given Id (Need To Use GET method)

```console
     curl http://127.0.0.1:3000/users/1
```

Update one user with the given Id (Need To Use PUT method)

```console
curl -X PUT -H "Content-Type: application/json"\
  --data '{"name": "newName","surname": "mockSurname","age": 30}' \
      http://127.0.0.1:3000/users/1
```

Delete one user with the given Id (Need To Use DELETE method)

```console
    curl -X DELETE http://127.0.0.1:3000/users/1
```
