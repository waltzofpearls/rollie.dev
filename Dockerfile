# Build frontend
FROM node:18.17-bookworm-slim AS frontend-build

WORKDIR /frontend

COPY frontend/package.json frontend/package-lock.json ./
RUN npm install

COPY ./frontend .
RUN npm run build


# Build backend
FROM rust:1.72-slim-bookworm AS backend-build

RUN apt-get update; \
    apt-get install -y --no-install-recommends \
        pkg-config \
        libssl-dev \
        ; \
    rm -rf /var/lib/apt/lists/*

WORKDIR /backend

COPY backend/Cargo.toml backend/Cargo.lock ./
# dummy build to cache dependencies
RUN mkdir src && echo "fn main() {}" > src/main.rs && cargo build --release

COPY ./backend .
RUN cargo build --release


# Final image
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=frontend-build /frontend/dist /app/static
COPY --from=backend-build /backend/target/release/backend /app/backend
COPY --from=backend-build /backend/src/resume /app/resume

EXPOSE 3000

CMD ["/app/backend"]
