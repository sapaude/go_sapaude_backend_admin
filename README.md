# Sapaude - Go Backend Admin Server

1. 基于DDD目录布局，提供了快捷的`init-project.sh`脚手架初始化目录，支持扩展
2. 依赖`github.com/sapadue/go-shims`组件库(shim、x系列)支持灵活扩展
3. 使用wire把所有的依赖注入串起
4. 支持`conf/*.yml`配置
5. 支持通过Echo框架路由和中间件配置

## 目录职责

| Layer                  | Responsibilities                                                    |
|------------------------|---------------------------------------------------------------------|
| **api/**               | Maps HTTP → DTOs → Application services. No business logic here.    |
| **application/**       | Coordinates use cases, calls domain/repo methods. Stateless.        |
| **domain/entity/**     | Rich domain models (e.g., `User` with methods like `SetPassword`)   |
| **domain/repository/** | Interface definitions only (e.g., `IReposUserMode`)                 |
| **infra/**             | Postgres, JWT, logging, etc. Implements repo and service interfaces |

## DevOps Guide for User Admin Service

This document describes the CI/CD pipeline and deployment strategy for the `user-admin` service.

The goal is to automate testing, build, and deployment to a Tencent Cloud CVM server.

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