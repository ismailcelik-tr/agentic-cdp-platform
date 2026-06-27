# Project Guidelines: Agentic CDP Platform

This file provides system-wide rules for AI agents working in this repository.

## Architecture
- **Monorepo:** All microservices are housed in this single repository.
- **Go Core API (`go-core-api`):** Must follow Domain-Driven Design (DDD). Use standard Go project layout. Emphasize `goroutines` for async tasks.
- **AI Agent Service (`ai-agent-service`):** Python based. Use `asyncio`, type hints, and FastAPI.
- **MCP Gateway (`mcp-gateway`):** Acts as the MCP Server.

## DevOps
- **Docker First:** Every new service or module must include a `Dockerfile` immediately. Use multi-stage builds.
- **Kubernetes:** Manifests should be placed in a `k8s/` directory at the root when created.
- **AWS:** Keep security (VPC, IAM) in mind when designing network interactions.

## Coding Standards
- **Go:** Follow `gofmt` and idiomatic Go practices. Return errors explicitly.
- **Python:** Use `ruff` or `flake8` standards. Strictly use type hints.
- **Documentation:** Every service must have its own internal `README.md` explaining its specific design decisions.
