## Assignment Project - REST API
The assignment project is to create CRUD endpoints for two tables, which are scores and students. The assigment requires ORM and Web Server.

## Environment
In my local environment, I used **PostgreSQL** for the database. Thus, the package that I imported in `db.go` is `"gorm.io/driver/postgres"`. In cases where other database is used, please change accordingly.

For connecting the database, the program is retrieving the following variables from the .env file:
* HOST=&lt;host_name&gt;
* DB_USER=&lt;database_user&gt;
* DB_PASSWORD=&lt;database_password&gt;
* DB_NAME=&lt;database_name&gt;
* DB_PORT=&lt;database_port&gt;
* WEB_SERVER_PORT=&lt;web_server_port&gt;

**The .env file is ignored from being submitted into the repository, so please create a new .env file before executing the program**

## Endpoints
The postman collection is provided in `postmanCollection.json` file. The collection used `localhost:8080` as host and web server port, so please change accordingly if other port or host is used.

### Create Student
**URI**
`POST /student`

**Example Body**
```
{
    "name": "Tara",
    "age": 25,
    "scores": [
        {
            "assignment_title": "Assignment Project 1",
            "description": "Create simple API without middleware",
            "score": 95
        }
    ]
}
```

### Get All Students
**URI:**
` GET /students`

**Response Example**
```
{
    "data": [
        {
            "id": 1,
            "name": "Tara",
            "age": 25,
            "Scores": [
                {
                    "ID": 1,
                    "assignment_title": "Assignment Project 1",
                    "description": "Create simple API without middleware",
                    "score": 95,
                    "StudentID": 1,
                    "CreatedAt": "2023-08-24T21:06:01.691714+07:00",
                    "UpdatedAt": "2023-08-24T21:06:01.691714+07:00"
                }
            ],
            "CreatedAt": "2023-08-24T21:06:01.690518+07:00",
            "UpdatedAt": "2023-08-24T21:06:01.690518+07:00"
        }
    ],
    "success": true
}
```

### Update Student
**URI:**
` PUT /student/:id`

**Response Body**
```
{
    "name": "Tara",
    "age": 27,
    "scores": [
        {
            "assignment_title": "Assignment Project 2",
            "description": "Create simple API without middleware 2",
            "score": 100
        },
        {
            "assignment_title": "Assignment Project 3",
            "description": "Create simple API without middleware 3",
            "score": 75
        },
        {
            "assignment_title": "Assignment Project 4",
            "description": "Create simple API without middleware 4",
            "score": 80
        }
    ]
}
```

**Response Example**
```
{
    "data": {
        "id": 1,
        "name": "Tara",
        "age": 27,
        "Scores": [
            {
                "ID": 1,
                "assignment_title": "Assignment Project 2",
                "description": "Create simple API without middleware 2",
                "score": 100,
                "StudentID": 1,
                "CreatedAt": "2023-08-24T21:06:01.691714+07:00",
                "UpdatedAt": "2023-08-24T21:09:40.3180852+07:00"
            },
            {
                "ID": 3,
                "assignment_title": "Assignment Project 3",
                "description": "Create simple API without middleware 3",
                "score": 75,
                "StudentID": 1,
                "CreatedAt": "2023-08-24T21:09:00.568065+07:00",
                "UpdatedAt": "2023-08-24T21:09:40.3186297+07:00"
            },
            {
                "ID": 4,
                "assignment_title": "Assignment Project 4",
                "description": "Create simple API without middleware 4",
                "score": 80,
                "StudentID": 1,
                "CreatedAt": "2023-08-24T21:09:40.3191803+07:00",
                "UpdatedAt": "2023-08-24T21:09:40.3191803+07:00"
            }
        ],
        "CreatedAt": "2023-08-24T21:06:01.690518+07:00",
        "UpdatedAt": "2023-08-24T21:09:40.3191803+07:00"
    },
    "success": true
}
```

### Delete student
**URI:**
` DELETE /student/:id`

**Response Example**
```
{
    "data": [
        {
            "id": 1,
            "name": "Tara",
            "age": 25,
            "Scores": [
                {
                    "ID": 1,
                    "assignment_title": "Assignment Project 1",
                    "description": "Create simple API without middleware",
                    "score": 95,
                    "StudentID": 1,
                    "CreatedAt": "2023-08-24T21:06:01.691714+07:00",
                    "UpdatedAt": "2023-08-24T21:06:01.691714+07:00"
                }
            ],
            "CreatedAt": "2023-08-24T21:06:01.690518+07:00",
            "UpdatedAt": "2023-08-24T21:06:01.690518+07:00"
        }
    ],
    "success": true
}
```