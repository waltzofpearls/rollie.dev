FROM rust:1.72-slim-bookworm

ARG GITHUB_TOKEN
ENV GITHUB_TOKEN $GITHUB_TOKEN

RUN apt-get update; \
    apt-get install -y --no-install-recommends \
        pkg-config \
        libssl-dev \
        ; \
    rm -rf /var/lib/apt/lists/*

RUN cargo install cargo-watch

WORKDIR /backend

COPY backend/Cargo.toml backend/Cargo.lock ./
# dummy build to cache dependencies
RUN mkdir src && echo "fn main() {}" > src/main.rs && cargo build --release

RUN mkdir -p static

CMD cargo watch -x run
