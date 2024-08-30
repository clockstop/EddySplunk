#!/bin/bash

# Build the SAM application
echo "Building the SAM application..."
sam build

# Start the SAM local API in the background
sam local start-api > sam_local.log 2>&1 &
SAM_PID=$!
echo "Started SAM local API with PID $SAM_PID"

echo "Sleeping 10 seconds to allow API to start fully"
# Wait for SAM to start
sleep 10

# Ask user for number of invocations
echo "How many times do you want to invoke the Lambda function? (10, 100, 1000)"
read NUM_INVOCATIONS

# Validate input
if [[ "$NUM_INVOCATIONS" != "10" && "$NUM_INVOCATIONS" != "100" && "$NUM_INVOCATIONS" != "1000" ]]; then
    echo "Invalid input. Please enter 10, 100, or 1000."
    exit 1
fi

# File to record invocation times
TIME_FILE="invocation_times_${NUM_INVOCATIONS}.txt"

# Invoke Lambda function and record durations
echo "Invoking Lambda function $NUM_INVOCATIONS times..."

for ((i=1; i<=NUM_INVOCATIONS; i++))
do
    START=$(gdate +%s%3N)  # Get start time in milliseconds
    curl -X GET http://localhost:3000/invoke -d '{}' > /dev/null
    END=$(gdate +%s%3N)   # Get end time in milliseconds
    echo "Start : $START"
    echo "End   : $END"
    DURATION=$(( (END - START) ))  # Convert nanoseconds to milliseconds
    echo "Invocation $i: $DURATION ms" >> $TIME_FILE
done

# Stop the SAM local API
kill $SAM_PID
echo "Stopped SAM local API"

# Run Python script to visualize data
python3 visualize_times.py $TIME_FILE

