# web_opt_mangadownloader

## PRODUCTION: run on docker containers

-   `cd frontend`
-   change `.env.production.example` to `.env.production` and add your `<host_ip>`
-   `cd ..` back to root folder and run
    -   for raspberry pi 4 environment it uses an older mongodb version: `docker-compose -f docker-compose.raspi.yml up -d`
    -   for other environments: `docker-compose up -d`
-   after containers are up navigate to `http://<host_ip>:3000` to download chapters

## DEVELOPMENT: Backend - API written in GoLang using Standard HTTP Library

### run server

-   run `docker-compose -f docker-compose.dev.yml up -d` to create mongodb dev container
-   change directory to `/backend`
-   `go run main.go`
-   server should run on `http://localhost:8080`

### (optional: update Open API)

-   run `swag init`

## Frontend - React with NextJS + TypeScript

-   change directory to `/frontend`
-   `npm run dev`
-   frontend should run on `http://localhost:3000`
