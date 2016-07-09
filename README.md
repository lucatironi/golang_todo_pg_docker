# Go ToDo Microservice with Postgres DB and Docker

## Setup

Build the docker images and run them:

    docker-compose build
    docker-compose up

Run an one-off postgres container to create the `todos` table (password for db is `postgres`):

    docker-compose run postgres psql -h postgres -U postgres -d todo_json_api_db -c "CREATE TABLE todos (id serial, name text, completed boolean, due timestamp);"

## API Usage

Create a new todo via curl:

    curl -i -X POST -H "Content-Type:application/json" -d '{"name":"Complete microservice","due":"2016-07-20T22:00:00Z","completed":false}' 'http://localhost:3000/todos'

Retrieve all todos (index):

    curl -i -X GET -H "Content-Type:application/json" 'http://localhost:3000/todos'

Retrieve the todo with `id = 1`:

    curl -i -X GET -H "Content-Type:application/json" 'http://localhost:3000/todos/1'
