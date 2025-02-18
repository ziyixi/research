FROM alpine:latest
LABEL org.opencontainers.image.authors="docker@ziyixi.science"
LABEL org.opencontainers.image.source="https://github.com/ziyixi/monorepo"
LABEL org.opencontainers.image.description="Todofy is a self-hosted task management tool that helps you organize and prioritize your tasks efficiently."
LABEL org.opencontainers.image.licenses="MIT"

ENV PORT=8080
ENV ALLOWED_USERS=""
ENV DATABASE_PATH=""

# Add CA certificates (already included in alpine but we ensure it's updated)
RUN apk --no-cache add ca-certificates tzdata

# Copy the Bazel-built Go binary from the local system into the image
COPY ./self_host/packages/todofy/_todofy_binary /todofy

# Expose the port your app will run on
EXPOSE 8080

# Run the app binary
ENTRYPOINT ["/todofy","-allowed-users=${ALLOWED_USERS}", "-database-path=${DATABASE_PATH}"]