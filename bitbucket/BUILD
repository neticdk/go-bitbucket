load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "access_tokens.go",
        "access_tokens_repos.go",
        "access_tokens_users.go",
        "bitbucket.go",
        "events.go",
        "keys.go",
        "keys_repos.go",
        "projects.go",
        "projects_repos.go",
        "projects_repos_branches.go",
        "projects_repos_commits_builds.go",
        "projects_repos_prs.go",
        "projects_repos_webhooks.go",
        "users.go",
        "webhook.go",
    ],
    importpath = "github.com/neticdk/go-bitbucket/bitbucket",
    visibility = ["//visibility:public"],
    deps = ["@com_github_google_go_querystring//query:go_default_library"],
)

go_test(
    name = "go_default_test",
    srcs = [
        "access_tokens_repos_test.go",
        "access_tokens_users_test.go",
        "bitbucket_test.go",
        "events_test.go",
        "keys_repos_test.go",
        "projects_repos_branches_test.go",
        "projects_repos_commits_builds_test.go",
        "projects_repos_test.go",
        "projects_repos_webhooks_test.go",
        "users_test.go",
        "webhook_test.go",
    ],
    embed = [":go_default_library"],
    deps = ["@com_github_stretchr_testify//assert:go_default_library"],
)
