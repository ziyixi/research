"""This is the module docstring for the MonoRepo workspace."""

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
# HUGO

RULES_HUGO_COMMIT = "d9118796628e8033b968e6085350c6e1ce55461e"

http_archive(
    name = "build_stack_rules_hugo",
    strip_prefix = "rules_hugo-%s" % RULES_HUGO_COMMIT,
    url = "https://github.com/ziyixi/rules_hugo/archive/%s.zip" % RULES_HUGO_COMMIT,
)

load("@build_stack_rules_hugo//hugo:rules.bzl", "github_hugo_theme", "hugo_repository")

hugo_repository(
    name = "hugo",
    extended = True,
    version = "0.133.1",
)

github_hugo_theme(
    name = "com_github_luizdepra_hugo_coder",
    commit = "5c702174587c11abf3063117cdd8a8fade2d50df",
    owner = "luizdepra",
    repo = "hugo-coder",
)
