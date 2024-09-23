FROM alpine:latest
LABEL org.opencontainers.image.authors="docker@ziyixi.science"
LABEL org.opencontainers.image.source="https://github.com/ziyixi/monorepo"
LABEL org.opencontainers.image.description="Package to handle cloudmailin's post on receiving emails, and update todo tasks accordingly with Gemini's summary"
LABEL org.opencontainers.image.licenses="MIT"

ENV PORT=8080

# Add CA certificates (already included in alpine but we ensure it's updated)
RUN apk --no-cache add ca-certificates

# Copy the Bazel-built Go binary from the local system into the image
COPY ./self_host/packages/todofy/_todofy_binary /todofy

# Expose the port your app will run on
EXPOSE 8080

# Run the app binary
ENTRYPOINT ["/todofy"]