load("@rules_go//go:def.bzl", "go_binary", "go_library")
load("@rules_oci//oci:defs.bzl", "oci_image", "oci_load")
load("@rules_pkg//pkg:tar.bzl", "pkg_tar")

go_library(
    name = "todofy_lib",
    srcs = [
        "router.go",
        "todofy.go",
    ],
    importpath = "github.com/ziyixi/monorepo/self_host/packages/todofy",
    visibility = ["//visibility:private"],
    deps = [
        "//self_host/packages/todofy/cloudmailin",
        "//self_host/packages/todofy/gemini",
        "//self_host/packages/todofy/mailjet",
        "//self_host/packages/todofy/utils",
        "@com_github_gin_gonic_gin//:gin",
    ],
)

go_binary(
    name = "todofy",
    embed = [":todofy_lib"],
    visibility = ["//visibility:public"],
)

pkg_tar(
    name = "todofy_tar",
    srcs = [":todofy"],
)

oci_image(
    name = "image",
    annotations = {
        "org.opencontainers.image.authors": "docker@ziyixi.science",
        "org.opencontainers.image.source": "https://github.com/ziyixi/Monorepo/self_host/packages/todofy",
        "org.opencontainers.image.description": "Package to handle cloudmailin's post on receiving emails, and update todoist tasks accordingly with Gemini summary",
        "org.opencontainers.image.licenses": "MIT",
    },
    base = "@distroless_static_debian12_linux_amd64",
    entrypoint = ["/todofy"],
    env = {
        "PORT": "8080",
    },
    exposed_ports = ["8080"],
    tars = [":todofy_tar"],
)

oci_load(
    name = "image_tarball",
    image = ":image",
    repo_tags = ["todofy:test"],
)

filegroup(
    name = "todofy.tar",
    srcs = [":image_tarball"],
    output_group = "tarball",
)
