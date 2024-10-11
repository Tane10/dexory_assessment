#!/usr/bin/bash

# Check if the correct number of arguments is provided
if [ "$#" -ne 2 ]; then
  echo "Usage: $0 <file_path> <file_type>"
  echo "Example: $0 path/to/file.json json"
  exit 1
fi

FILE_PATH="$1"  # The first argument is the file path
FILE_TYPE="$2"  # The second argument is the file type (e.g., "json" or "csv")

# Check if the file exists
if [ ! -f "$FILE_PATH" ]; then
  echo "Error: File '$FILE_PATH' not found!"
  exit 1
fi

# Upload the file using curl
curl -X POST "http://localhost:8080/upload?type=${FILE_TYPE}" \
  -F "file=@${FILE_PATH}"

# Check if the curl command was successful
if [ $? -eq 0 ]; then
  echo "Upload successful!"
else
  echo "Upload failed."
fi
