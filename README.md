# DevOps Guide for User Admin Service

This document describes the CI/CD pipeline and deployment strategy for the `user-admin` service. The goal is to automate testing, build, and deployment to a Tencent Cloud CVM server.

---

## 🔁 CI/CD Flow Overview

```plaintext
Local Dev (Mac)
     │
     ▼
Push to GitHub (main)
     │
     ▼
GitHub Actions
 ├── build.yml   → Run tests + Build binary (all branches)
 └── deploy.yml  → Build + Deploy to CVM (only main)
     │
     ▼
CVM Server
 └── systemd restarts service
```

## 目录职责

| Layer                  | Responsibilities                                                    |
|------------------------|---------------------------------------------------------------------|
| **api/**               | Maps HTTP → DTOs → Application services. No business logic here.    |
| **application/**       | Coordinates use cases, calls domain/repo methods. Stateless.        |
| **domain/entity/**     | Rich domain models (e.g., `User` with methods like `SetPassword`)   |
| **domain/repository/** | Interface definitions only (e.g., `IReposUserMode`)                 |
| **infra/**             | Postgres, JWT, logging, etc. Implements repo and service interfaces |
