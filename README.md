<h1 align="center"> CRUD Book API </h1>

## About

This is a simple API written in GO to practice the language syntax and way of coding it. The implementation is 
based on a video of freeCodeCamp tutorial.

---

## Libs

- Mysql
- gorm
- mux

--- 

## Database Setup

The database is mySQL. In this project there's a docker-compose file that creates a database instance. 
Database's default port: 3306.
To create a mySQL database instance, runs one of the commands below:

```bash
    # Makefile command
    $ make docker

    # docker command
    $ docker compose up -d
```

--- 

## Run

WebServer's default port: 9010.
To run the web server, execute one of the command below. 
Create a database instance before running the server.

```bash
    # Run with Makefile
    $ make

    # Run main.go file
    $ go run src/main.go
```

---

## Curls

Copy and paste the curls below to make requests to the server.
If the url has a {bookID}, replace it with one of the book id in the database.

```bash
    # Post method
    $ curl -X POST -H "Content-Type: application/json" \
    -d '{"name": "Test", "author": "John Doe", "publication": "01/01/2020"}' \
    localhost:9010/books

    # Delete method
    $ curl -X DELETE -H "Content-Type: application/json" \
    localhost:9010/books/{bookID}

    # Put method
    $ curl -X PUT -H "Content-Type: application/json" \
    -d '{"name": "Jane Doe", "author": "Test Doe", "publication": "07/07/2021"}' \
    localhost:9010/books/{bookdID}

    # Get Method
    $ curl localhost:9010/books

    # Get Method 
    $ curl localhost:9010/books/{bookID}
``` 