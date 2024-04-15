#!/bin/bash

# Set the path to your Go binary (if not in PATH)
# You can remove this line if go is already in your PATH
export GO_BIN=$(which go)

# Try to find the buf command
command -v buf >/dev/null 2>&1

# Check the exit code
if [ $? -eq 0 ]; then
  buf generate
else
  echo "buf-cli is not installed."
fi

# Build the Go main.go file
"$GO_BIN" build -o protoc-gen-custom

# Check if build was successful (exit code 0 indicates success)
if [[ $? -eq 0 ]]; then
  # Copy the built binary to the target directory
  cp protoc-gen-custom ${GOPATH}/bin/


# Run the protoc command with custom options
  protoc --custom_out=internal --custom_opt=paths=source_relative,dbtag=sql proto/user/*.proto

else
  echo "Error: Building Go application failed."
  exit 1
fi

echo "Successfully built, copied, and ran the command."
