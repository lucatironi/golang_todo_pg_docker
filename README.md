# Go ToDo Microservice with Postgres DB and Docker

## Setup

    docker-compose build
    docker-compose up

## API Usage

    curl -i -X POST -H "Content-Type:application/json" -d '{"name":"Complete microservice","due":"2016-07-20T22:00:00Z","completed":false}' 'http://localhost:3000/todos'

    curl -i -X GET -H "Content-Type:application/json" 'http://localhost:3000/todos'

    curl -i -X GET -H "Content-Type:application/json" 'http://localhost:3000/todos/1'
