#!/bin/bash

# Path to the file containing the version
FILE="utils.go"

# Extract the current version
CURRENT_VERSION=$(grep -oP 'const version = "\K[0-9]+\.[0-9]+\.[0-9]+' $FILE)

# Split the version into its components
IFS='.' read -r -a VERSION_PARTS <<< "$CURRENT_VERSION"

# Increment the patch version
VERSION_PARTS[2]=$((VERSION_PARTS[2] + 1))

# Join the version parts back together
NEW_VERSION="${VERSION_PARTS[0]}.${VERSION_PARTS[1]}.${VERSION_PARTS[2]}"

# Replace the old version with the new version in the file
sed -i "s/const version = \"$CURRENT_VERSION\"/const version = \"$NEW_VERSION\"/" $FILE

# Add the updated file to the staging area
git add $FILE

echo "Version updated to $NEW_VERSION"