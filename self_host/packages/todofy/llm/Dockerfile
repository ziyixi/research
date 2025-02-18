FROM alpine:latest
LABEL org.opencontainers.image.authors="docker@ziyixi.science"
LABEL org.opencontainers.image.source="https://github.com/ziyixi/monorepo"
LABEL org.opencontainers.image.description="Todofy (llm) is a subpackage of Todofy for task management using large language models."
LABEL org.opencontainers.image.licenses="MIT"

ENV PORT=50051
ENV GEMINI_API_KEY=""

# Add CA certificates (already included in alpine but we ensure it's updated)
RUN apk --no-cache add ca-certificates tzdata

# Copy the Bazel-built Go binary from the local system into the image
COPY ./self_host/packages/todofy/llm/_llm_binary /llm
COPY ./self_host/packages/todofy/llm/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Expose the port your app will run on
EXPOSE 50051

# Run the app binary
ENTRYPOINT ["/entrypoint.sh"]