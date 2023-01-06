load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:build_file_name BUILD,BUILD.bazel
# gazelle:go_naming_convention go_default_library
# gazelle:prefix github.com/neticdk/go-bitbucket
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-build_file_proto_mode=disable_global",
        "-prune",
    ],
    command = "update-repos",
)
