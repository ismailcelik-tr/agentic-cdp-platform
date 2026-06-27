# Agentic CDP Platform (Customer Data Platform)

A production-grade, microservices-based Agentic Customer Data Platform (CDP). This project demonstrates a modern architecture combining a robust Domain-Driven Design (DDD) backend in Go with advanced AI capabilities (RAG, ReAct Agents) in Python, bridged together using the Model Context Protocol (MCP).

## 🚀 Architecture Overview

This is a **Monorepo** containing three main microservices:

1. **`go-core-api/`** (The Core Backend)
   - **Language:** Go
   - **Architecture:** Domain-Driven Design (DDD) & Clean Architecture
   - **Role:** Manages customer segments, event ingestion, and core business logic.
   - **Storage:** PostgreSQL

2. **`ai-agent-service/`** (The AI Brain)
   - **Language:** Python (FastAPI + LangGraph)
   - **Role:** Handles RAG (Retrieval-Augmented Generation) via `pgvector` and executes ReAct agent loops.

3. **`mcp-gateway/`** (The Bridge)
   - **Protocol:** Model Context Protocol (MCP)
   - **Role:** Exposes the `go-core-api`'s REST endpoints as native "Tools" for the Python AI Agent, allowing the AI to interact with the CDP data seamlessly.

## 🐳 Infrastructure & DevOps
- **Containerization:** Docker (Multi-stage builds for all services).
- **Orchestration:** Kubernetes (K8s) deployments with Rolling Updates.
- **Cloud:** Designed for AWS (VPC, EC2 IAM profiles).

## 🏃‍♂️ Running Locally

The entire stack can be spun up locally using Docker Compose:

```bash
docker compose up -d
```

*(Detailed instructions will be added as services are implemented).*

## 📄 License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
