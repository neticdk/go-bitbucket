load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "access_tokens.go",
        "access_tokens_repos.go",
        "access_tokens_users.go",
        "bitbucket.go",
        "keys.go",
        "keys_repos.go",
        "projects.go",
        "projects_repos.go",
    ],
    importpath = "github.com/neticdk/go-bitbucket/bitbucket",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "access_tokens_repos_test.go",
        "access_tokens_users_test.go",
        "bitbucket_test.go",
        "keys_repos_test.go",
        "projects_repos_test.go",
    ],
    embed = [":go_default_library"],
    deps = ["@com_github_stretchr_testify//assert:go_default_library"],
)
