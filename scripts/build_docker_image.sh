#!/bin/bash

TAG=test

# Step 1: Build the Go binary using Bazel
bazel build //self_host/packages/todofy || {
    echo "Bazel build failed"
    exit 1
}

# Step 2: Copy the Bazel-built binary into the Docker build context (temporary location)
cp -f bazel-bin/self_host/packages/todofy/todofy_/todofy self_host/packages/todofy/_todofy_binary || {
    echo "Copying binary failed"
    exit 1
}

# Step 3: Build the Docker image
docker build -t todofy:${TAG} -f self_host/packages/todofy/Dockerfile .

# Step 4: Clean up (optional)
rm self_host/packages/todofy/_todofy_binary

# Optional: Print success message
echo "Docker image 'todofy:${TAG}' built successfully"
