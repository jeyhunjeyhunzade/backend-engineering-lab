# 🏗️ Backend Engineering Lab — Fullstack Learning Journey

Welcome to the **Backend Engineering Lab** — a structured, hands-on roadmap to mastering backend development by **building small to mid-sized projects** for practicing backend engineering.

---

## 🎯 Goal

The goal of this lab is to gain **deep backend engineering experience** by:

- Designing and implementing focused backend projects.
- Building small to mid-sized applications to practice core backend concepts quickly.
- Building projects in your language(s) of choice based on learning goals and project requirements.
- Learning modern backend architecture, APIs, databases, CI/CD, and deployment practices.

---

## 🧩 Project Structure

Each project lives under its own folder inside `/backend-engineering-lab/`:

```
/backend-engineering-lab/
  ├── todo-app/
  │     ├── typescript/
  │     ├── golang/
  │     └── rust/           # Same project, multiple language implementations
  ├── url-shortener/
  │     ├── python/
  │     └── java/
  ├── ...
  ├── infra/
  │     ├── db/
  │     │     └── migrations/     # Shared SQL migrations across all implementations
  │     └── k8s/                  # Optional Kubernetes manifests (later)
  ├── .github/
  │     └── workflows/            # CI/CD automation (GitHub Actions)
  ├── templates/                  # Reusable Makefile / Dockerfile templates
  ├── tools/                      # Scripts, linters, formatters, etc.
  └── README.md                   # This file
```

---

## 🧱 Project Lifecycle

Each project follows this consistent cycle:

| Step             | Description                                                     |
| ---------------- | --------------------------------------------------------------- |
| **1. Design**    | Define endpoints, database schema, and architecture.            |
| **2. Implement** | Build in your chosen language(s).                               |
| **3. Test**      | Write unit/integration tests.                                   |
| **4. Dockerize** | Add Dockerfile and docker-compose for local setup.              |
| **5. CI/CD**     | Integrate build, test, and deploy pipelines via GitHub Actions. |
| **6. Deploy**    | Deploy containers to Render, Fly.io, or Railway.                |

---

## 🧰 Tech Stack

### Languages

Projects can be implemented in multiple languages. Examples include:

- **TypeScript / Node.js** — productivity and rich ecosystem
- **Go** — performance and concurrency
- **Python** — versatility and extensive libraries
- **Rust** — memory safety and performance
- **Java / Kotlin** — enterprise-grade JVM ecosystem
- **PHP / Laravel** — web-focused development
- **C# / .NET** — Microsoft stack

Choose based on your learning goals or implement the same project in multiple languages.

### Databases

- **PostgreSQL** (primary)

### Tooling

| Purpose          | Tool                                                            |
| ---------------- | --------------------------------------------------------------- |
| API              | Express / Fastify / Fiber / Gin / Flask / FastAPI / Axum / etc. |
| ORM/Database     | Prisma, GORM, sqlx, SQLAlchemy, Diesel, etc.                    |
| DB Migrations    | Shared SQL scripts in `/infra/db/migrations`                    |
| Containerization | Docker & docker-compose                                         |
| CI/CD            | GitHub Actions                                                  |
| Testing          | Language-specific testing frameworks                            |

---

## ⚙️ CI/CD Overview

The repository supports **per-project** and **per-language** CI workflows.

### Workflow Triggers

- Runs automatically on changes to relevant paths (e.g., `todo-app/**`).
- Can be triggered manually via **workflow_dispatch**.
- Builds, lints, and tests the service.
- Optionally builds and pushes Docker images.

### Example

```
.github/workflows/ci-todo-app-ts.yml
.github/workflows/ci-todo-app-go.yml
```

Each language implementation has its own CI pipeline.

---

## 🧾 Makefile Commands

Each language folder has a **Makefile** for standardized operations:

```bash
make install    # Install dependencies
make test       # Run tests
make build      # Build the project
make docker     # Build Docker image
```

---

## 🧪 Local Development

Each language implementation has its own `docker-compose.yml`:

```bash
cd todo-app/typescript
docker-compose up --build
```

This spins up:

- Application container
- Database (Postgres)
- Optional admin tools (pgAdmin, etc.)

---

## 🌍 Deployment

Once a project passes CI:

- Tag it (e.g., `v1.0.0`)
- GitHub Actions will build and push Docker images.
- Deploy to your preferred environment:

  - **Railway.app**
  - **Fly.io**
  - **Render**
  - or **Kubernetes** (when you reach that phase)

---

## 🧑‍💻 Learning Progression

| Stage                     | Focus        |
| ------------------------- | ------------ |
| **Beginner Projects**     | Task Tracker |
| **Intermediate Projects** |              |
| **Advanced Projects**     |              |
| **DevOps Expansion**      |              |

---

## 🧠 Learning Goals

By the end of this lab, you will:

- Be comfortable designing and building backend systems from scratch.
- Master backend development in your chosen language(s).
- Use Docker, Postgres, and CI/CD like a professional.
- Have a strong portfolio of production-ready backend projects.

---

## 📜 License

This repository is for **personal learning and portfolio purposes**.
You are free to modify and distribute your code under your preferred license.

---

## 🧩 Next Step

Pick your **first project** — recommended starting point:
👉 **`todo-app`**

Then:

1. Create its architecture plan.
2. Choose your implementation language.
3. Implement and test the project.
4. Commit and push — your CI/CD pipeline will handle the rest.

---
