#!/bin/bash

# Script to create LeetCode problem directory and files with templates
# Usage: ./mk_problem.sh --problem <number>

# Get the actual script location (resolve symlinks)
SCRIPT_PATH="$(readlink -f "${BASH_SOURCE[0]}")"
SCRIPT_DIR="$(dirname "$SCRIPT_PATH")"
# Get the root directory (parent of scripts)
ROOT_DIR="$(dirname "$SCRIPT_DIR")"
# Template directory
TEMPLATE_DIR="$SCRIPT_DIR/problem_template"
# LeetCode directory
LEETCODE_DIR="$ROOT_DIR/leetcode"

# Function to display usage
usage() {
    echo "Usage: $0 [OPTIONS]"
    echo ""
    echo "Options:"
    echo "  -p, --problem <number>    Create a new problem directory with templates"
    echo "  -r, --remove <number>     Remove an existing problem directory"
    echo "  -h, --help               Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0 -p 200                Create problem 200"
    echo "  $0 --problem 200         Create problem 200"
    echo "  $0 -r 200                Remove problem 200"
    echo "  $0 --remove 200          Remove problem 200"
    exit 1
}

# Parse command line arguments
PROBLEM_NUM=""
ACTION="create"

while [[ $# -gt 0 ]]; do
    case $1 in
        -p|--problem)
            PROBLEM_NUM="$2"
            ACTION="create"
            shift 2
            ;;
        -r|--remove)
            PROBLEM_NUM="$2"
            ACTION="remove"
            shift 2
            ;;
        -h|--help)
            usage
            ;;
        *)
            echo "Unknown option: $1"
            usage
            ;;
    esac
done

# Check if problem number is provided
if [[ -z "$PROBLEM_NUM" ]]; then
    echo "Error: Problem number is required"
    usage
fi

# Validate that problem number is a positive integer
if ! [[ "$PROBLEM_NUM" =~ ^[0-9]+$ ]]; then
    echo "Error: Problem number must be a positive integer"
    exit 1
fi

# Handle different actions
if [[ "$ACTION" == "create" ]]; then
    # Check if directory already exists
    if [[ -d "$LEETCODE_DIR/$PROBLEM_NUM" ]]; then
        echo "Warning: Directory '$PROBLEM_NUM' already exists"
        read -p "Do you want to continue? (y/N): " -n 1 -r
        echo
        if [[ ! $REPLY =~ ^[Yy]$ ]]; then
            echo "Operation cancelled"
            exit 0
        fi
    fi

    # Create directory and files
    echo "Creating LeetCode problem $PROBLEM_NUM..."

    # Create directory
    mkdir -p "$LEETCODE_DIR/$PROBLEM_NUM"

    # Copy and customize Python template
    if [[ -f "$TEMPLATE_DIR/template.py" ]]; then
        cp "$TEMPLATE_DIR/template.py" "$LEETCODE_DIR/$PROBLEM_NUM/$PROBLEM_NUM.py"
        # Replace PROBLEM_NUMBER placeholder
        sed -i "s/PROBLEM_NUMBER/$PROBLEM_NUM/g" "$LEETCODE_DIR/$PROBLEM_NUM/$PROBLEM_NUM.py"
    else
        echo "Warning: Python template not found, creating empty file"
        touch "$LEETCODE_DIR/$PROBLEM_NUM/$PROBLEM_NUM.py"
    fi

    # Copy and customize Go template
    if [[ -f "$TEMPLATE_DIR/template.go" ]]; then
        cp "$TEMPLATE_DIR/template.go" "$LEETCODE_DIR/$PROBLEM_NUM/$PROBLEM_NUM.go"
        # Replace PROBLEM_NUMBER placeholder
        sed -i "s/PROBLEM_NUMBER/$PROBLEM_NUM/g" "$LEETCODE_DIR/$PROBLEM_NUM/$PROBLEM_NUM.go"
    else
        echo "Warning: Go template not found, creating empty file"
        touch "$LEETCODE_DIR/$PROBLEM_NUM/$PROBLEM_NUM.go"
    fi

    echo "✅ Successfully created:"
    echo "   📁 leetcode/$PROBLEM_NUM/"
    echo "   🐍 leetcode/$PROBLEM_NUM/$PROBLEM_NUM.py (from template)"
    echo "   🐹 leetcode/$PROBLEM_NUM/$PROBLEM_NUM.go (from template)"

elif [[ "$ACTION" == "remove" ]]; then
    # Check if directory exists
    if [[ ! -d "$LEETCODE_DIR/$PROBLEM_NUM" ]]; then
        echo "Error: Directory 'leetcode/$PROBLEM_NUM' does not exist"
        exit 1
    fi

    # Show what will be removed
    echo "The following will be removed:"
    echo "   📁 leetcode/$PROBLEM_NUM/"
    if [[ -f "$LEETCODE_DIR/$PROBLEM_NUM/$PROBLEM_NUM.py" ]]; then
        echo "   🐍 leetcode/$PROBLEM_NUM/$PROBLEM_NUM.py"
    fi
    if [[ -f "$LEETCODE_DIR/$PROBLEM_NUM/$PROBLEM_NUM.go" ]]; then
        echo "   🐹 leetcode/$PROBLEM_NUM/$PROBLEM_NUM.go"
    fi
    
    # Confirm removal
    read -p "Are you sure you want to remove problem $PROBLEM_NUM? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo "Operation cancelled"
        exit 0
    fi

    # Remove directory
    rm -rf "$LEETCODE_DIR/$PROBLEM_NUM"
    echo "🗑️  Successfully removed leetcode/$PROBLEM_NUM/"
fi
