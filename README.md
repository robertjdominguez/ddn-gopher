# DDN Gopher

This is a bundled suite of applications which you can run with Docker:

- [x] An authentication server, using Gin
- [x] A PostgreSQL database with a `users` table by default
- [x] A pre-loaded Hasura DDN project which can be run locally
- [ ] A starter CLI application, using Bubble Tea

## Prerequisites

- Docker

## Let's get up and running

### Step 1. Clone the repo

Clone the repo:

```bash
git clone https://github.com/robertjdominguez/ddn-gopher
```

### Step 2. Create a .env in the root

Move into the repository:

```bash
cd ddn-gopher
```

Then, create a new `.env` file with a JWT secret key:

```bash
touch ./auth-server/.env && echo "JWT_SECRET=somethingSuperSecureGoesHere!" > ./auth-server/.env
```

### Step 3. Start Docker Compose

From the root of the project, run:

```bash
 HASURA_DDN_PAT=$(ddn auth print-pat) docker compose up -d
```

### Step 4. Log into the CLI

## Applications

### Authentication server

### PostgreSQL database

### Hasura DDN project

### CLI application

## Deployment

### Auth server

### Hasura DDN

### CLI

## Contributing
