FROM alpine:latest
LABEL org.opencontainers.image.authors="docker@ziyixi.science"
LABEL org.opencontainers.image.source="https://github.com/ziyixi/monorepo"
LABEL org.opencontainers.image.description="Todofy (todo) is a subpackage of Todofy for email/API populating tasks."
LABEL org.opencontainers.image.licenses="MIT"

ENV PORT=50052
ENV MAILJET_API_KEY_PUBLIC=""
ENV MAILJET_API_KEY_PRIVATE=""
ENV TARGET_EMAIL=""

# Add CA certificates (already included in alpine but we ensure it's updated)
RUN apk --no-cache add ca-certificates tzdata

# Copy the Bazel-built Go binary from the local system into the image
COPY ./self_host/packages/todofy/todo/todo_binary /todo

# Expose the port your app will run on
EXPOSE 50052

# Run the app binary
ENTRYPOINT ["/todo", "-mailjet-api-key-public=${MAILJET_API_KEY_PUBLIC}", "-mailjet-api-key-private=${MAILJET_API_KEY_PRIVATE}", "-target-email=${TARGET_EMAIL}"]