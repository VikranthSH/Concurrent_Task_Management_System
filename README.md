# Concurrent Task Management System

A production-ready backend service for managing users, projects, and tasks with role-based access control, JWT authentication, and advanced MongoDB querying capabilities.

## ***Project Goals***

- Implement clean layered architecture (Handler → Service → Repository)
- Demonstrate JWT-based authentication and authorization
- Optimize MongoDB indexing and aggregation pipelines
- Build efficient role-based dashboard queries
- Follow Go best practices and conventions

## ***Table of Contents***

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [Authentication](#authentication)
- [User Roles](#user-roles)
- [API Documentation](#api-documentation)
- [Database & Indexing](#database--indexing)
- [Error Handling](#error-handling)
- [Testing with Postman](#testing-with-postman)

---

## ***Features***

- ***Role-Based Access Control*** - Three-tier role system (Super Admin, Admin, Employee)
- ***JWT Authentication*** - Secure token-based authentication
- ***MongoDB Integration*** - Optimized with indexes and aggregation pipelines
- ***Clean Architecture*** - Separated concerns with handler, service, and repository layers
- ***Dashboard API*** - Role-based aggregated data views
- ***RESTful APIs*** - Complete CRUD operations for Users, Projects, and Tasks

---

## ***Technology Stack***

| Layer | Technology |
|-------|-----------|
| **Language** | Go 1.22 |
| **Database** | MongoDB |
| **Router** | Gorilla Mux |
| **Authentication** | JWT (golang-jwt/v5) |
| **Architecture** | Clean Layered Architecture |
| **DB Features** | Indexes, Aggregation Pipeline |

---

## ***Project Structure***

```
Concurrent_Task_Management_System/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/                  # Configuration management
│   ├── dto/                     # Data Transfer Objects
│   │   ├── api_response.go
│   │   ├── dashboard_response.go
│   │   └── dashboard_user.go
│   ├── handlers/                # HTTP request handlers
│   │   ├── auth_handler.go
│   │   ├── dashboard_handler.go
│   │   ├── project_handler.go
│   │   ├── task_handler.go
│   │   └── user_handler.go
│   ├── models/                  # Data models
│   │   ├── project.go
│   │   ├── task.go
│   │   └── user.go
│   ├── repositories/            # Database access layer
│   │   ├── dashboard_repository.go
│   │   ├── project_repository.go
│   │   ├── task_repository.go
│   │   └── user_repository.go
│   ├── routes/                  # Route definitions
│   │   ├── auth_routes.go
│   │   ├── dashboard_routes.go
│   │   ├── project_routes.go
│   │   ├── task_routes.go
│   │   └── user_routes.go
│   ├── services/                # Business logic layer
│   │   ├── dashboard_service.go
│   │   ├── project_service.go
│   │   ├── task_service.go
│   │   └── user_service.go
│   └── utils/                   # Utility functions
│       ├── jwt.go
│       ├── mongo_indexes.go
│       └── response.go
├── go.mod                       # Go module definition
└── README.md                    # Documentation
```

### Architecture Layers

1. **Handlers** - HTTP layer handling request/response
2. **Services** - Business logic and validation
3. **Repositories** - MongoDB CRUD operations
4. **Models** - Database entity definitions
5. **DTOs** - Response structure definitions

---

## ***Prerequisites***

- **Go** 1.22 or higher
- **MongoDB** 4.4 or higher (running on `localhost:27017`)
- **Postman** (optional, for API testing)

### Installation

#### 1. Clone the Repository

```bash
git clone https://github.com/VikranthSH/Concurrent_Task_Management_System.git
cd Concurrent_Task_Management_System
```

#### 2. Install Dependencies

```bash
go mod download
go mod tidy
```

#### 3. Verify Installation

```bash
go run cmd/server/main.go --version
```

---

## ***Configuration***

### MongoDB Connection

Edit the MongoDB URI in [cmd/server/main.go](cmd/server/main.go#L23):

```go
mongoURI := "mongodb://localhost:27017"
db := client.Database("trello_lite")
```

**Environment Variables** (Optional Future Enhancement):
- `MONGO_URI` - MongoDB connection string
- `DB_NAME` - Database name
- `PORT` - Server port (default: 8080)

---

## ***Running the Application***

### 1. Start MongoDB

```bash
# Using MongoDB locally
mongod

# Or using Docker
docker run -d -p 27017:27017 --name mongodb mongo:latest
```

### 2. Start the Server

```bash
go run cmd/server/main.go
```

### Expected Output

```
MongoDB indexes ensured
Server running on port 8080
```

### Build for Production

```bash
go build -o ctms cmd/server/main.go
./ctms
```

---

## ***Authentication***

### JWT Implementation

The system uses **handler-level JWT validation** (not middleware):

1. **Login** → Generate JWT token
2. **Include token** in subsequent requests
3. **Validate token** inside handler functions
4. **Extract user** from database
5. **Check permissions** based on role

### JWT Token Format

```
Authorization: Bearer <JWT_TOKEN>
```

### JWT Flow Diagram

```
Client Request
    ↓
[Handler] - Decode JWT
    ↓
[Service] - Validate Permissions
    ↓
[Repository] - Fetch Data
    ↓
Response with Data
```

---

## ***User Roles***

| Role | Permissions | Access Level |
|------|------------|--------------|
| **super_admin** | Full system access | All users, projects, tasks |
| **admin** | Manage employees & their data | Employees, owned projects, assigned tasks |
| **employee** | View own data only | Own projects and tasks |

### Role-Based Behavior

```
Dashboard Access:
├── super_admin → All users + Projects + Tasks
├── admin → Employees + Owned Projects + Related Tasks
└── employee → Own Projects + Own Tasks
```

---

## ***API Documentation***

### Base URL

```
http://localhost:8080
```

### Response Format

All API responses follow a standard structure:

```json
{
  "status": "success",
  "description": "Operation completed successfully",
  "data": {}
}
```

### Error Response Format

```json
{
  "status": "error",
  "description": "Error message describing what went wrong",
  "data": null
}
```

---

### Authentication Endpoints

#### Login
```
POST /auth/login
Content-Type: application/json

{
  "user_id": "admin_001"
}

Response:
{
  "status": "success",
  "description": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

---

### User Endpoints

#### Create User
```
POST /users
Authorization: Bearer <JWT_TOKEN>
Content-Type: application/json

{
  "user_id": "emp_001",
  "name": "Employee One",
  "email": "emp1@company.com",
  "role": "employee"
}
```

#### Get All Users
```
GET /users
Authorization: Bearer <JWT_TOKEN>
```

#### Get User by ID
```
GET /users/{id}
Authorization: Bearer <JWT_TOKEN>
```

#### Update User
```
PUT /users/{id}
Authorization: Bearer <JWT_TOKEN>
Content-Type: application/json
```

#### Delete User
```
DELETE /users/{id}
Authorization: Bearer <JWT_TOKEN>
```

---

### Project Endpoints

#### Create Project
```
POST /projects
Authorization: Bearer <JWT_TOKEN>
Content-Type: application/json

{
  "name": "Backend API",
  "description": "Core backend services",
  "ownerId": "<OWNER_OBJECT_ID>",
  "memberIds": ["<MEMBER_OBJECT_ID>"]
}
```

#### Get All Projects
```
GET /projects
Authorization: Bearer <JWT_TOKEN>
```

#### Get Project by ID
```
GET /projects/{id}
Authorization: Bearer <JWT_TOKEN>
```

#### Get Projects by User
```
GET /projects/user/{userId}
Authorization: Bearer <JWT_TOKEN>
```

#### Update Project
```
PUT /projects/{id}
Authorization: Bearer <JWT_TOKEN>
```

#### Delete Project
```
DELETE /projects/{id}
Authorization: Bearer <JWT_TOKEN>
```

---

### Task Endpoints

#### Create Task
```
POST /tasks
Authorization: Bearer <JWT_TOKEN>
Content-Type: application/json

{
  "title": "Design DB Schema",
  "description": "Create MongoDB schema design",
  "projectId": "<PROJECT_OBJECT_ID>",
  "assignedTo": "<USER_OBJECT_ID>",
  "status": "Todo",
  "priority": "High"
}
```

#### Get All Tasks
```
GET /tasks
Authorization: Bearer <JWT_TOKEN>
```

#### Get Task by ID
```
GET /tasks/{id}
Authorization: Bearer <JWT_TOKEN>
```

#### Get Tasks by Project
```
GET /tasks/project/{projectId}
Authorization: Bearer <JWT_TOKEN>
```

#### Get Tasks by User
```
GET /tasks/user/{userId}
Authorization: Bearer <JWT_TOKEN>
```

#### Update Task
```
PUT /tasks/{id}
Authorization: Bearer <JWT_TOKEN>
```

#### Delete Task
```
DELETE /tasks/{id}
Authorization: Bearer <JWT_TOKEN>
```

---

### Dashboard Endpoint

#### Get Dashboard (Role-Based)
```
GET /dashboard
Authorization: Bearer <JWT_TOKEN>

Response for super_admin:
{
  "status": "success",
  "data": {
    "totalUsers": 10,
    "totalProjects": 5,
    "totalTasks": 42,
    "users": [...],
    "projects": [...],
    "tasks": [...]
  }
}
```

**Behavior by Role:**
- **super_admin** → All users, all projects, all tasks
- **admin** → Employees under them, owned projects, related tasks
- **employee** → Own projects only, own tasks only

---

## ***Database & Indexing***

### MongoDB Database

```
Database: trello_lite
Collections: users, projects, tasks, dashboards
```

### Indexed Fields

**Users Collection**
- `user_id` (unique)
- `email` (unique)
- `role` (non-unique)

**Projects Collection**
- `ownerId` (non-unique)
- `memberIds` (array index)

**Tasks Collection**
- `projectId` (non-unique)
- `assignedTo` (non-unique)
- `status` (non-unique)

### Why Indexing Matters

- ***Faster query execution*** - Reduces database scan time
- ***Optimized dashboard performance*** - Aggregation queries run faster
- ***Production-ready optimization*** - Handles scale efficiently
- ***Automatic index creation*** - Created on server startup

### Index Creation

Indexes are automatically created on server startup via:

```go
utils.EnsureMongoIndexes(db)
```

See [internal/utils/mongo_indexes.go](internal/utils/mongo_indexes.go) for implementation.

---

## ***Error Handling***

| HTTP Code | Status | Reason |
|-----------|---------|--------|
| **200** | success | Operation successful |
| **400** | error | Bad request / Invalid input |
| **401** | error | Missing or invalid JWT token |
| **403** | error | Insufficient permissions |
| **404** | error | Resource not found |
| **500** | error | Server error |

### Common Errors

| Error | Cause | Solution |
|-------|-------|----------|
| 401 Unauthorized | Missing/invalid token | Include valid JWT in Authorization header |
| 403 Forbidden | Insufficient permissions | User role lacks required access |
| 404 Not Found | Invalid ID or route | Verify ID format and endpoint URL |
| 400 Bad Request | Invalid data format | Check request body structure |
| 500 Internal Server | Server error | Check server logs |

---

## ***Testing with Postman***

### Step 1: Import Endpoints

Create a new Postman collection and import all endpoints above.

### Step 2: Login and Get Token

1. **Send** `POST /auth/login` with user_id
2. **Copy** the JWT token from response
3. **Save** in Postman environment variable: `{{token}}`

### Step 3: Add Authorization Header

For all protected endpoints, add:

```
Authorization: Bearer {{token}}
```

### Step 4: Test Endpoints

1. Create a user
2. Create a project
3. Create a task
4. Fetch dashboard
5. Verify role-based access

### Example Postman Flow

```
1. POST /auth/login → Get JWT
   ↓
2. POST /users → Create user (Admin role)
   ↓
3. POST /projects → Create project
   ↓
4. POST /tasks → Create task
   ↓
5. GET /dashboard → View aggregated data
```

---

## ***Example API Calls***

### 1. Login and Get Token

```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{"user_id": "admin_001"}'
```

### 2. Get Dashboard with Token

```bash
curl -X GET http://localhost:8080/dashboard \
  -H "Authorization: Bearer <JWT_TOKEN>"
```

### 3. Create Project

```bash
curl -X POST http://localhost:8080/projects \
  -H "Authorization: Bearer <JWT_TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Mobile App",
    "description": "React Native mobile application",
    "ownerId": "admin_id",
    "memberIds": ["emp_id"]
  }'
```

---

## ***Important Notes***

- JWT validation is **handler-level**, not middleware-based
- Indexes are created automatically on server startup
- MongoDB aggregation pipelines are used for dashboard queries
- The system supports scalability with proper indexing
- All timestamps are stored in MongoDB ObjectId

---

## ***Contributing***

Contributions are welcome! Please follow these guidelines:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'Add AmazingFeature'`)
4. Push to branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## ***License***

This project is licensed under the MIT License - see the LICENSE file for details.

---

## ***Author***

**Vikranth SH**  
GitHub: [@VikranthSH](https://github.com/VikranthSH)

---

## ***Quick References***

- [Go Documentation](https://golang.org/doc)
- [MongoDB Documentation](https://docs.mongodb.com)
- [Gorilla Mux Documentation](https://github.com/gorilla/mux)
- [JWT Best Practices](https://tools.ietf.org/html/rfc7519)

---

