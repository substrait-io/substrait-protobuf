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
esac
if ! expr "$SPEC_VERSION" : 'v[0-9][0-9]*\.[0-9][0-9]*\.[0-9][0-9]*$' >/dev/null; then
    echo "❌ Spec version must be a release tag such as v0.102.0: $SPEC_VERSION" >&2
    exit 1
fi

# Use git ls-remote to check if the specific tag exists
echo "Checking if spec version $SPEC_VERSION exists in $SPEC_REPO_URL"
if git ls-remote --exit-code --tags "$SPEC_REPO_URL" "refs/tags/$SPEC_VERSION" >/dev/null; then
    echo "✅ Spec version $SPEC_VERSION exists"
else
    spec_tag_status=$?
    if [ "$spec_tag_status" -ne 2 ]; then
        echo "❌ Unable to check spec release tag $SPEC_VERSION" >&2
        exit 4
    fi
    echo "❌ Spec version $SPEC_VERSION does NOT exist" >&2
    exit 3
fi

GO_VERSION_TAG="go/$SPEC_VERSION"
BRANCH_NAME="releases/$GO_VERSION_TAG"

# Only a remote tag confirms that publication succeeded.
if git ls-remote --exit-code --tags origin "refs/tags/$GO_VERSION_TAG" >/dev/null; then
    echo "✅ Protobufs for $SPEC_VERSION already exist"
    exit 0
else
    tag_status=$?
    if [ "$tag_status" -ne 2 ]; then
        echo "❌ Unable to check published release tags" >&2
        exit 4
    fi
fi

if git rev-parse --verify --quiet "refs/tags/$GO_VERSION_TAG" >/dev/null 2>&1; then
    echo "❌ Local tag $GO_VERSION_TAG exists but has not been published; inspect the previous release attempt before retrying" >&2
    exit 4
fi

if git ls-remote --exit-code --heads origin "$BRANCH_NAME" >/dev/null; then
    echo "❌ Release branch $BRANCH_NAME exists without tag $GO_VERSION_TAG" >&2
    exit 4
else
    branch_status=$?
    if [ "$branch_status" -ne 2 ]; then
        echo "❌ Unable to check published release branches" >&2
        exit 4
    fi
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
