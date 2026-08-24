#!/bin/sh

set -eu

# Ensure exactly one argument is provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <version>"
    exit 1
fi

# Only run off of the main branch
CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
if [ "$CURRENT_BRANCH" != "main" ]; then
    echo "❌ ./generate-go must be invoked on main branch"
    exit 2
fi

SPEC_REPO_URL="https://github.com/substrait-io/substrait.git"
SPEC_VERSION="$1"

case "$SPEC_VERSION" in
    */*)
        echo "❌ Spec version must not contain '/': $SPEC_VERSION" >&2
        exit 1
        ;;
    v[0-9]*.[0-9]*.[0-9]*)
        ;;
    *)
        echo "❌ Spec version must be a release tag such as v0.102.0: $SPEC_VERSION" >&2
        exit 1
        ;;
esac

# Use git ls-remote to check if the specific tag exists
echo "Checking if spec version $SPEC_VERSION exists in $SPEC_REPO_URL"
if git ls-remote --exit-code --tags "$SPEC_REPO_URL" "refs/tags/$SPEC_VERSION" >/dev/null 2>&1; then
    echo "✅ Spec version $SPEC_VERSION exists"
else
    echo "❌ Spec version $SPEC_VERSION does NOT exist" >&2
    exit 3
fi

GO_VERSION_TAG="go/$SPEC_VERSION"
BRANCH_NAME="releases/$GO_VERSION_TAG"

# A tag is the publication marker. Do not regenerate or mutate an existing release.
if git rev-parse --verify --quiet "refs/tags/$GO_VERSION_TAG" >/dev/null 2>&1 || \
    git ls-remote --exit-code --tags origin "refs/tags/$GO_VERSION_TAG" >/dev/null 2>&1; then
    echo "✅ Protobufs for $SPEC_VERSION already exist"
    exit 0
fi

if git ls-remote --exit-code --heads origin "$BRANCH_NAME" >/dev/null 2>&1; then
    echo "❌ Release branch $BRANCH_NAME exists without tag $GO_VERSION_TAG" >&2
    exit 4
fi

if [ -n "$(git status --porcelain)" ]; then
    echo "❌ ./generate-go must be invoked with a clean working tree" >&2
    exit 4
fi

echo "🔨 Creating new branch: $BRANCH_NAME"
git checkout -B "$BRANCH_NAME" main

TARGET="https://github.com/substrait-io/substrait.git#tag=$SPEC_VERSION"
echo "🔧 Executing Protobuf code generation for $TARGET"
buf generate "$TARGET"

echo "Committing generated code"

git add go/substraitpb
git commit --allow-empty -m "generated go/substraitpb for $SPEC_VERSION"

git tag "$GO_VERSION_TAG" -m "Generated Go code for spec version $SPEC_VERSION"
# Publish branch and tag together so an interrupted run cannot leave a branch
# without its release tag.
git push --atomic origin "$BRANCH_NAME" "$GO_VERSION_TAG"

git checkout main
