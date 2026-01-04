#!/bin/bash

if [ -z "$1" ]; then
    echo "Usage: $0 <new_module_name>"
    echo "Example: $0 github.com/myuser/myproject"
    exit 1
fi

OLD_MODULE="go-layout"
NEW_MODULE="$1"
OS="$(uname)"

echo "Replacing module name from '$OLD_MODULE' to '$NEW_MODULE'..."

# Find all files excluding the scripts themselves and .git directory
# Using LC_ALL=C to handle potential encoding issues
# Using xargs with sed for replacement
if [ "$OS" = "Darwin" ]; then
    # macOS requires empty string for -i
    grep -rl "$OLD_MODULE" . --exclude-dir=.git --exclude="rename_project.sh" --exclude="rename_project.bat" | xargs sed -i '' "s|$OLD_MODULE|$NEW_MODULE|g"
else
    # Linux/GNU sed
    grep -rl "$OLD_MODULE" . --exclude-dir=.git --exclude="rename_project.sh" --exclude="rename_project.bat" | xargs sed -i "s|$OLD_MODULE|$NEW_MODULE|g"
fi

echo "Done."
echo "Please run 'go mod tidy' to update dependencies."
