# Auth Service

Authentication microservice built with Go, PostgreSQL, gRPC, and Kafka following Clean Architecture principles.

---

## 📌 Responsibilities

- User registration
- User authentication
- JWT token generation
- gRPC-based synchronous communication
- Kafka-based asynchronous events (later)

---

## 🏗 Architecture

This service follows **Clean Architecture**:

## 🔐 Authentication Logic

### Register

- Checks if user exists
- Hashes password using bcrypt
- Stores user in PostgreSQL

### Login

- Verifies credentials
- Returns access token (JWT coming next)

### Password Security

- bcrypt with default cost
- Plain passwords are never stored
