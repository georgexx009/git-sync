#!/bin/bash

GO_FILE="."
OUTPUT_BINARY="git-sync"
DEST_DIR="$HOME/.config/bin"

go build -o $OUTPUT_BINARY $GO_FILE

if [ $? -eq 0 ]; then
    echo "Compilation successful. Moving binary to $DEST_DIR"
    
    # Create the destination directory if it doesn't exist
    mkdir -p $DEST_DIR

    # Move the binary and db files to the destination directory
    mv $OUTPUT_BINARY $DEST_DIR/$OUTPUT_BINARY
    echo "All done"
else
    echo "Compilation failed."
fi

