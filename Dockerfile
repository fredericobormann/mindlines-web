FROM golang:1.24-alpine AS backend_build

WORKDIR /app

COPY mindlines-backend/go.mod mindlines-backend/go.sum ./
RUN go mod download

COPY mindlines-backend/main.go ./
COPY mindlines-backend/scene scene
COPY mindlines-backend/helper helper
RUN CGO_ENABLED=0 GOOS=linux go build -o mindlines-backend


FROM node:23-alpine AS frontend_build

WORKDIR /app
COPY mindlines-frontend/package.json mindlines-frontend/package-lock.json ./
RUN npm ci

COPY mindlines-frontend/ .
RUN npm run build


FROM node:23-alpine

WORKDIR /app
COPY --from=backend_build /app/mindlines-backend .
COPY --from=frontend_build /app/.output .
COPY ./start.sh .

ENV GIN_MODE=release

EXPOSE 8080
EXPOSE 3000

CMD ["sh", "start.sh"]
