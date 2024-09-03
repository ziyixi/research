load("@build_stack_rules_hugo//hugo:rules.bzl", "hugo_serve", "hugo_site", "hugo_theme")
load("@rules_pkg//pkg:tar.bzl", "pkg_tar")

hugo_theme(
    name = "hugo-coder",
    srcs = [
        "@com_github_luizdepra_hugo_coder//:files",
    ],
    theme_name = "hugo-coder",
)

hugo_site(
    name = "homepage",
    archetypes = glob(["archetypes/**"]),
    assets = glob(["assets/**"]),
    config = "config.toml",
    content = glob(["content/**"]),
    data = glob(["data/**"]),
    images = glob(["images/**"]),
    layouts = glob(["layouts/**"]),
    static = glob(["static/**"]),
    theme = ":hugo-coder",
)

hugo_serve(
    name = "serve",
    dep = [
        ":homepage",
    ],
    quiet = False,
)

# Create a tgz package
pkg_tar(
    name = "docs_tgz",
    srcs = [":docs"],
    package_file_name = "docs.tgz",
)
