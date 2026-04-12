# EPSE - Express Project Structure Generator

EPSE is a CLI tool designed to accelerate the development of Node.js/Express/TypeScript projects. It is built to generate project structures and boilerplate code, allowing developers to focus on business logic rather than setup.

---

## Overview

EPSE supports two project architectures:

- **Lite** — A minimal Express/TypeScript structure for straightforward APIs and prototypes.
- **Clean** — A full Clean Architecture structure with TSOA, dependency injection via tsyringe, and clear separation of concerns.

Both architectures are supported across all commands.

---

## Commands

### generate

Generates a new project structure.

```bash
epse generate <project-name> [destination]
epse generate <project-name> --lite [destination]
epse generate <project-name> --clean [destination]
```

### add route

Generates a controller and registers the route.

```bash
epse add route <domain> <route-url> --method <HTTP_METHOD> --controller <name>
epse add route <domain> <route-url> --crud
```

### add middleware

Generates a custom Express middleware.

```bash
epse add middleware <name>
```

### add auth

Generates a complete JWT authentication system including login, register, middleware and routes.

```bash
epse add auth
```

### add service

Generates a service class.

```bash
epse add service <name>
```

### add repository

Generates a repository class. For Clean architecture, also generates the gateway interface.

```bash
epse add repository <name>
```

---

## Project Architectures

### Lite

```
src/
├── config/
├── controllers/
├── middlewares/
├── repositories/
├── routes/
├── services/
├── types/
└── utils/
```

### Clean

```
src/
├── adapters/
│   ├── controllers/
│   ├── gateway/
│   ├── middlewares/
│   └── services/
├── entities/
├── frameworks/
│   └── tsoa/
├── useCases/
│   └── gateway/
└── utilities/
```

---

## Interactive Mode

All commands support an interactive mode. Simply omit the arguments and EPSE will prompt you for the required information.

```bash
epse add route
epse add middleware
epse add service
epse add repository
```

---

## Configuration

EPSE maintains an `epseconfig.json` file at the root of your project to track generated resources.

```json
{
  "projectName": "my-api",
  "projectType": "lite",
  "controllersPath": "src/controllers",
  "database": false,
  "auth": false,
  "routes": [],
  "customMiddlewares": []
}
```

---

## Requirements

- Go 1.21+
- Node.js 18+

Dependencies are always fetched at their latest versions at project generation time.

---

## License

MIT