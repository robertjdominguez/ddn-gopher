# DDN Gopher

![image](https://github.com/user-attachments/assets/e46e2cb8-07fd-4cf1-a919-6212291a4313)

This is a bundled suite of applications which you can run with Docker:

- [x] An authentication server, using Gin
- [x] A PostgreSQL database with a `users` table by default
- [x] A pre-loaded Hasura DDN project which can be run locally
- [x] A starter CLI application, using Bubble Tea

## Prerequisites

- [Docker](https://docs.docker.com/engine/install/)
- The [Hasura DDN CLI](https://hasura.io/docs/3.0/cli/installation)

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

### Step 3. Update the AuthConfig

Replace the `value` for the JWT key in `hasura/supergraph_globals/auth-config.hml`

With whatever value you used above, replace `3q2+7w==iQ==` in this ☝️ file:

```yaml
kind: AuthConfig
version: v2
definition:
  mode:
    jwt:
      claimsConfig:
        namespace:
          claimsFormat: Json
          location: /claims.jwt.hasura.io
      tokenLocation:
        type: BearerAuthorization
      key:
        fixed:
          algorithm: HS256
          key:
            value: <your-new-value-here>
```

### Step 3. Start Docker Compose

From the root of the project, run:

```bash
HASURA_DDN_PAT=$(ddn auth print-pat) docker compose --env-file hasura/.env up --build --watch
```
