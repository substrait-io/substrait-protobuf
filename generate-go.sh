#!/bin/sh

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

SPEC_VERSION="$1"
SPEC_REPO_URL="https://github.com/substrait-io/substrait.git"

# Use git ls-remote to check if the specific tag exists
echo "Checking if spec version $SPEC_VERSION exists in $SPEC_REPO_URL"
if git ls-remote --tags "$SPEC_REPO_URL" "refs/tags/$SPEC_VERSION" | grep -q .; then
    echo "✅ Spec version $SPEC_VERSION exists"
else
    echo "❌ Spec version $SPEC_VERSION does NOT exist"
    exit 3
fi

GO_VERSION_TAG="go/$SPEC_VERSION"

# Use git tags to check if Golang protobufs have already been published
echo "\nCheck if Golang protobufs exist for $SPEC_VERSION"
if git tag --list "$GO_VERSION_TAG" | grep -q .; then
    echo "✅ Protobufs for $SPEC_VERSION already exist"
    exit 0
else
    echo "✅ Protobufs do no exist for $SPEC_VERSION yet"
fi

BRANCH_NAME="releases/$GO_VERSION_TAG"
echo "\n🔨 Creating new branch: $BRANCH_NAME"
git checkout -B "$BRANCH_NAME"

TARGET="https://github.com/substrait-io/substrait.git#tag=$SPEC_VERSION"
echo "🔧 Executing Protobuf code generation for $TARGET"
buf generate "$TARGET"

echo "Committing generated code"

git add go/substraitpb
git commit -m "generated go/substraitpb for $SPEC_VERSION"
git push --set-upstream origin "$BRANCH_NAME"

git tag "$GO_VERSION_TAG" -m "Generated Go code for spec version $SPEC_VERSION"
git push origin "$GO_VERSION_TAG"

git checkout main