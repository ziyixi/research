#!/bin/bash
set -euo pipefail

# Check if TAG is provided
if [ -z "$1" ]; then
    echo "Error: TAG parameter is required"
    echo "Usage: $0 <tag>"
    exit 1
fi

TAG=$1

# Define arrays
bazel_targets=(
    "//self_host/packages/todofy"
    "//self_host/packages/todofy/llm:llm"
    "//self_host/packages/todofy/todo:todo"
    "//self_host/packages/todofy/database:database"
)

bazel_copy_from=(
    "bazel-bin/self_host/packages/todofy/todofy_/todofy"
    "bazel-bin/self_host/packages/todofy/llm/llm_/llm"
    "bazel-bin/self_host/packages/todofy/todo/todo_/todo"
    "bazel-bin/self_host/packages/todofy/database/database_/database"
)

bazel_copy_to=(
    "self_host/packages/todofy/_todofy_binary"
    "self_host/packages/todofy/llm/_llm_binary"
    "self_host/packages/todofy/todo/_todo_binary"
    "self_host/packages/todofy/database/_database_binary"
)

docker_targets=(
    "ghcr.io/ziyixi/todofy:${TAG}"
    "ghcr.io/ziyixi/todofy-llm:${TAG}"
    "ghcr.io/ziyixi/todofy-todo:${TAG}"
    "ghcr.io/ziyixi/todofy-database:${TAG}"
)

dockerfile_paths=(
    "self_host/packages/todofy/Dockerfile"
    "self_host/packages/todofy/llm/Dockerfile"
    "self_host/packages/todofy/todo/Dockerfile"
    "self_host/packages/todofy/database/Dockerfile"
)

# Verify arrays have the same length
if [ ${#bazel_copy_from[@]} -ne ${#bazel_copy_to[@]} ] ||
    [ ${#docker_targets[@]} -ne ${#dockerfile_paths[@]} ]; then
    echo "Error: Array lengths don't match"
    exit 1
fi

# Prepare essencial protobuf files
bash self_host/packages/todofy/proto/copy_protos.sh

echo "Step 1: Building Go binaries using Bazel..."
for target in "${bazel_targets[@]}"; do
    echo "Building $target..."
    bazel build "$target" || {
        echo "Error: Bazel build for $target failed"
        exit 1
    }
done

echo "Step 2: Copying binaries to Docker build context..."
for i in "${!bazel_copy_from[@]}"; do
    echo "Copying ${bazel_copy_from[$i]} -> ${bazel_copy_to[$i]}"
    cp -f "${bazel_copy_from[$i]}" "${bazel_copy_to[$i]}" || {
        echo "Error: Copying binary failed: ${bazel_copy_from[$i]} -> ${bazel_copy_to[$i]}"
        exit 1
    }
done

echo "Step 3: Building Docker images..."
for i in "${!docker_targets[@]}"; do
    echo "Building ${docker_targets[$i]}..."
    docker build -t "${docker_targets[$i]}" -f "${dockerfile_paths[$i]}" . || {
        echo "Error: Docker build for ${docker_targets[$i]} failed"
        exit 1
    }
done

echo "Step 4: Cleaning up temporary files..."
for i in "${!bazel_copy_to[@]}"; do
    echo "Removing ${bazel_copy_to[$i]}"
    # ensure the file exists before attempting to remove it
    [ -f "${bazel_copy_to[$i]}" ] &&
        rm -f "${bazel_copy_to[$i]}" || {
        echo "Error: Cleanup failed for ${bazel_copy_to[$i]}"
        exit 1
    }
done

echo "Success: All Docker images built successfully"
echo "Built images:"
for target in "${docker_targets[@]}"; do
    echo "  - $target"
done
