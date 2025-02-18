FROM debian:slim
LABEL org.opencontainers.image.authors="docker@ziyixi.science"
LABEL org.opencontainers.image.source="https://github.com/ziyixi/monorepo"
LABEL org.opencontainers.image.description="Todofy (database) is a subpackage of Todofy for task database management."
LABEL org.opencontainers.image.licenses="MIT"

ENV PORT=50053

# Install SQLite and CA certificates
RUN apt-get update && apt-get install -y \
    ca-certificates \
    sqlite3 \
    tzdata \
    && rm -rf /var/lib/apt/lists/*

# Copy the Bazel-built Go binary from the local system into the image
COPY ./self_host/packages/todofy/database/_database_binary /database

# Expose the port your app will run on
EXPOSE 50053

# Run the app binary
ENTRYPOINT ["/database"]