# Go gin-gonic

Example of using [Go gin-gonic framework](https://github.com/gin-gonic/gin) and mysql to create REST API.

##### Requests: 
```
POST api/v1/todos/
GET api/v1/todos/
GET api/v1/todos/:id
PUT api/v1/todos/:id
DELETE api/v1/todos/:id 
```

##### Schema: Database: **demogin**, Table: **todos**.

| name         | description |
| -----------  | ----------- |
| id           | int(11) primary key not null auto increment |
| title        | varchar(255)     |
| completed    | int(11)          |
| created_at   | timestamp        |
| updated_at   | timestamp        |
| deleted_at   | timestamp        |
| created_at   | timestamp        |
