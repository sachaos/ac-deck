FROM rust:1.42.0

WORKDIR /src
COPY . /src
RUN cargo build --release
RUN rm /src/src/main.rs
