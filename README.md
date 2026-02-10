Concurrent Task Management System (Backend)

This project is a backend service built using Go, MongoDB, and JWT authentication.
It is designed to manage users, projects, and tasks with role-based access and a dashboard view for admins.

The main goal of this project is to:

-Understand clean backend layering (handler → service → repository)
-Implement JWT-based authentication
-Use MongoDB indexing and aggregation properly
-Build a real-world dashboard query efficiently


<!-- Tech Stack -->

-Language: Go (Golang)
-Database: MongoDB
-Router: Gorilla Mux
-Auth: JWT (golang-jwt)
-Architecture: Clean layered architecture
-DB Features Used:
    -Indexes
    -Aggregation pipeline ($lookup, $match, $project)

<!-- Project Structure -->
Concurrent_Task_Management_System/
│
├── cmd/

│   └── server/
│       └── main.go
│
├── internal/
│   ├── handlers/        // HTTP layer (request/response)
│   ├── services/        // Business logic
│   ├── repositories/   // MongoDB access
│   ├── models/          // DB models
│   ├── dto/             // Response structures
│   ├── routes/          // Route registration
│   └── utils/           // JWT, responses, indexes
│
└── go.mod



---

## User Roles

| Role        | Description |
|------------|-------------|
| super_admin | Full access |
| admin       | Can view employees, their projects, and tasks |
| employee    | Can view only their own data |

---

## Authentication (JWT)

This project **does not use auth middleware**.  
JWT validation is handled **inside handlers**.

### JWT Flow

1. User logs in
2. Server generates JWT
3. Client sends JWT in request header
4. JWT is decoded
5. User is fetched from DB
6. Access is validated by role

### Header Format

    Authorization: Bearer <JWT_TOKEN>


---

## MongoDB Indexing

Indexes are created automatically on server startup.

### Why indexing is used
- Faster queries
- Required for dashboard performance
- Production-level optimization

### Indexed Fields

**Users**
- `user_id` (unique)
- `email` (unique)
- `role`

**Projects**
- `ownerId`
- `memberIds`

**Tasks**
- `projectId`
- `assignedTo`
- `status`

Index creation code:
    internal/utils/EnsureMongoIndexes.go

POST /auth/login
### users

POST /users
GET /users
GET /users/{id}
PUT /users/{id}
DELETE /users/{id}

### Projects

POST /projects
GET /projects
GET /projects/{id}
GET /projects/user/{userId}
PUT /projects/{id}
DELETE /projects/{id}

### tasks

POST /tasks
GET /tasks
GET /tasks/{id}
GET /tasks/project/{projectId}
GET /tasks/user/{userId}
PUT /tasks/{id}
DELETE /tasks/{id}

<!-- GET /dashboard -->

---

## API Response Format

All responses follow a standard structure:

```json
{
  "status": "success",
  "description": "Dashboard fetched successfully",
  "data": {}
}


__________________________________________________________

How to Run

1. Start MongoDB
mongodb://localhost:27017

2. Run Server
go run cmd/server/main.go


Expected output:

MongoDB indexes ensured
Server running on port 8080

______________________________________________________________

Postman Testing

Step 1: Login
POST /auth/login


Copy the returned JWT.

Step 2: Access Protected APIs

Add header:

Authorization: Bearer <JWT_TOKEN>


Example:

GET /dashboard

⃣ Login (Get JWT Token)

Method: POST
URL:

http://localhost:8080/auth/login


Headers:

Content-Type: application/json


Body (JSON):

{
  "user_id": "admin_001"
}


 Response:

{
  "status": "success",
  "description": "Login successful",
  "data": {
    "token": "<JWT_TOKEN>"
  }
}


 Copy the token value.
This JWT represents the currently logged-in user.

️ Use JWT for Protected APIs

For all protected APIs, add this header in Postman:

Authorization: Bearer <JWT_TOKEN>


Example:

Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...

Dashboard API
 Get Dashboard (Role-Based)

Method: GET
URL:

http://localhost:8080/dashboard


Headers:
    Authorization: Bearer <JWT_TOKEN>

Behavior by Role
    super_admin
        Can see all users, projects, and tasks

    admin
        Can see employees under them
        Can see projects they own
        Can see tasks assigned to those employees

    employee
        Can see only their own projects and tasks

User APIs
Create User

Method: POST

http://localhost:8080/users


Body:

{
  "user_id": "emp_001",
  "name": "Employee One",
  "email": "emp1@company.com",
  "role": "employee"
}

Get All Users

Method: GET

http://localhost:8080/users


Headers:

Authorization: Bearer <JWT_TOKEN>

Project APIs
Create Project

Method: POST

http://localhost:8080/projects


Headers:

Authorization: Bearer <JWT_TOKEN>
Content-Type: application/json


Body:

{
  "name": "Backend API",
  "description": "Core backend services",
  "ownerId": "<ADMIN_OBJECT_ID>",
  "memberIds": [
    "<EMPLOYEE_OBJECT_ID>"
  ]
}

Get All Projects

Method: GET

http://localhost:8080/projects


Headers:

Authorization: Bearer <JWT_TOKEN>

Task APIs
Create Task

Method: POST

http://localhost:8080/tasks


Headers:

Authorization: Bearer <JWT_TOKEN>
Content-Type: application/json


Body:

{
  "title": "Design DB Schema",
  "projectId": "<PROJECT_OBJECT_ID>",
  "assignedTo": "<EMPLOYEE_OBJECT_ID>",
  "status": "Todo",
  "priority": "High"
}

Get Tasks by Project

Method: GET

http://localhost:8080/tasks/project/<PROJECT_OBJECT_ID>

Get Tasks by User

Method: GET

http://localhost:8080/tasks/user/<USER_OBJECT_ID>

Common Errors (Expected)
Error	Reason
401 Unauthorized	Missing / invalid JWT
403 Forbidden	Role-based access denied
404 Not Found	Invalid route or ID
400 Bad Request	Invalid input
