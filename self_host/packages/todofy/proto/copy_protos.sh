#!/bin/bash
set -e # Exit on error

# Get Bazel workspace directory
WORKSPACE_DIR=$(bazel info workspace)
cd "${WORKSPACE_DIR}"

# Get bazel-bin directory
BAZEL_BIN=$(bazel info bazel-bin)
SOURCE_DIR="$BAZEL_BIN/self_host/packages/todofy/proto/self_host_packages_todofy_proto_go_grpc_pb/github.com/ziyixi/monorepo/self_host/packages/todofy/proto/self_host/packages/todofy/proto"
DEST_DIR="self_host/packages/todofy/proto"

# Copy the files
bazel clean
bazel build //self_host/packages/todofy/proto:self_host_packages_todofy_proto_go_grpc
cp -f $SOURCE_DIR/*.go $DEST_DIR/
