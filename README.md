# web_opt_mangadownloader

## PRODUCTION: run on docker containers

-   `cd frontend`
-   change `.env.production.example` to `.env.production` and add your `<host_ip>`
-   `docker-compose up -d` in root
-   after containers are up navigate to `http://<host_ip>:3000` to download chapters

## DEVELOPMENT: Backend - API written in GoLang using Standard HTTP Library

### run server

-   change directory to `/backend`
-   `go run main.go`
-   server should run on `http://localhost:8080`

### (optional: update Open API)

-   run `swag init`

## Frontend - React with NextJS + TypeScript

-   change directory to `/frontend`
-   `npm run dev`
-   frontend should run on `http://localhost:3000`
