#!/bin/bash

# Target the current directory, or a path passed as an argument
TARGET_DIR="${1:-.}"

if [ ! -d "$TARGET_DIR" ]; then
    echo "Error: Directory $TARGET_DIR does not exist."
    exit 1
fi

echo "=================================================="
echo " Prototype Development Hour Estimator"
echo " Target Directory: $TARGET_DIR"
echo " Base: 4 working hours per day"
echo "=================================================="

# 1. Find all .go files, excluding vendor and hidden folders
# 2. Strip comments (single-line // )
# 3. Strip blank lines
# 4. Count the remainder
RAW_SLOC=$(find "$TARGET_DIR" -name "*.go" -not -path "*/vendor/*" -not -path "*/.*" -type f -exec grep -v '^[[:space:]]*//' {} + | grep -v '^[[:space:]]*$' | wc -l | tr -d ' ')

if [ "$RAW_SLOC" -eq 0 ]; then
    echo "No Go source code files found to analyze."
    exit 0
fi

# Define Prototype Model parameters (LOC per day)
MIN_SPEED=150  # Fast prototyping pace (e.g., 150 LOC / 4 hours = 37.5 LOC/hr)
MAX_SPEED=100  # Conservative prototyping pace (e.g., 100 LOC / 4 hours = 25 LOC/hr)
HOURS_PER_DAY=4

# Calculate required days first
DAYS_MIN=$(( RAW_SLOC / MIN_SPEED ))
DAYS_MAX=$(( RAW_SLOC / MAX_SPEED ))

# Handle edge case where project is tiny and calculation drops to 0 days
[ $DAYS_MIN -eq 0 ] && DAYS_MIN=1
[ $DAYS_MAX -eq 0 ] && DAYS_MAX=1

# Convert days to total working hours
HOURS_MIN=$(( DAYS_MIN * HOURS_PER_DAY ))
HOURS_MAX=$(( DAYS_MAX * HOURS_PER_DAY ))

# Convert working hours to rough calendar weeks (assuming 5 working days = 20 hours/week)
WEEKS_MIN=$(echo "scale=1; $HOURS_MIN / 20" | bc 2>/dev/null || echo "$(( HOURS_MIN / 20 ))")
WEEKS_MAX=$(echo "scale=1; $HOURS_MAX / 20" | bc 2>/dev/null || echo "$(( HOURS_MAX / 20 ))")

echo -e "\n METRICS DETECTED"
echo "• Total Source Lines of Code (SLOC): $RAW_SLOC"

echo -e "\n PROTOTYPE HOURLY ESTIMATE (25 - 37.5 LOC / Hour)"
echo "• Active Coding Hours Required: $HOURS_MIN to $HOURS_MAX hours"
echo "• Estimated Project Duration : $WEEKS_MIN to $WEEKS_MAX calendar weeks"

echo -e "\n ESTIMATE BREAKDOWN"
echo "• At fast prototyping velocity  : ~$HOURS_MIN hours of focused effort."
echo "• At standard prototyping velocity: ~$HOURS_MAX hours of focused effort."
echo "=================================================="
