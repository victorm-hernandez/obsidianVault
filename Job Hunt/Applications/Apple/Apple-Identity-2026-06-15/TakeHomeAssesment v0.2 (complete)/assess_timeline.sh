#!/bin/bash

# Target the current directory, or a path passed as an argument
TARGET_DIR="${1:-.}"

if [ ! -d "$TARGET_DIR" ]; then
    echo "Error: Directory $TARGET_DIR does not exist."
    exit 1
fi

echo "=================================================="
echo " Analyzing Go Project Timeline: $TARGET_DIR"
echo " (Using macOS Native File System Metadata)"
echo "=================================================="

# Find all .go and go.mod files, ignoring hidden directories like .git
# %B = Birth (creation) time in seconds since epoch
# %m = Modification time in seconds since epoch
# %Sm / %SB = Human-readable strings for modification / birth times

echo -e "\n[1] Scanning File Creation (Birth) Dates..."
OLDEST_FILE=$(find "$TARGET_DIR" -type f \( -name "*.go" -o -name "go.mod" \) -not -path '*/.*' -exec stat -f "%B %SB %N" {} + | sort -n | head -n 1)

if [ -z "$OLDEST_FILE" ]; then
    echo "No Go source files found."
else
    OLDEST_DATE=$(echo "$OLDEST_FILE" | awk '{print $2, $3, $4, $5}')
    OLDEST_NAME=$(echo "$OLDEST_FILE" | awk '{print $6}')
    echo "• Development likely STARTED on: $OLDEST_DATE"
    echo "  (Oldest file created: $OLDEST_NAME)"
fi

echo -e "\n[2] Scanning File Modification Dates..."
NEWEST_FILE=$(find "$TARGET_DIR" -type f \( -name "*.go" -o -name "go.mod" \) -not -path '*/.*' -exec stat -f "%m %Sm %N" {} + | sort -n | tail -n 1)

if [ -z "$NEWEST_FILE" ]; then
    echo "No Go source files found."
else
    NEWEST_DATE=$(echo "$NEWEST_FILE" | awk '{print $2, $3, $4, $5}')
    NEWEST_NAME=$(echo "$NEWEST_FILE" | awk '{print $6}')
    echo "• Development likely PAUSED/ENDED on: $NEWEST_DATE"
    echo "  (Most recently modified file: $NEWEST_NAME)"
fi

echo "=================================================="
