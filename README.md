# Commerce Hub

Commerce Hub is an enterprise-grade microservices-based commerce platform built with Go.

The system is designed following real-world industry standards:

- Clear service boundaries
- gRPC for internal service-to-service communication
- PostgreSQL with database-per-service isolation
- Clean Architecture principles
- Contract-first development using Protocol Buffers

## Services

- **Auth Service**
  - User authentication and authorization
  - Token issuance and validation
  - Role-based access control

- **Product Service**
  - Product and catalog management
  - Pricing and stock handling

- **Order Service**
  - Order creation and lifecycle management
  - Business workflow orchestration

## Tech Stack

- Language: Go (vanilla)
- Communication: gRPC
- Database: PostgreSQL
- Containerization: Docker
- Architecture: Clean Architecture + DDD-lite

## Principles

- Each service owns its database
- No shared data between services
- Services communicate only via gRPC
- Business logic is isolated from transport and infrastructure
- Production-minded defaults (timeouts, logging, context propagation)

## Status

Work in progress — developed incrementally following enterprise engineering practices.
