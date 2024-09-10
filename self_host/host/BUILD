load("//tools:make_executable.bzl", "make_executable")

# Group all data and env files
filegroup(
    name = "data_group",
    srcs = glob(["data/**"]),
)

filegroup(
    name = "env_group",
    srcs = glob(["env/**"]),
)

# Group docker-compose.yml and other necessary files
filegroup(
    name = "compose_group",
    srcs = ["docker_compose.yml"],
)

# Make the update.sh script executable
make_executable(
    name = "update_executable",
    src = "update.sh",
)

# Define the sh_binary target
sh_binary(
    name = "update",
    srcs = [":update_executable"],
    data = [
        ":compose_group",  # Include docker-compose.yml in the sandbox
        ":data_group",  # Include data directory in the sandbox
        ":env_group",  # Include env directory in the sandbox
    ],
    visibility = ["//visibility:public"],
)
