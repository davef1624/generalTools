import os
import sys

# Count files in directory and return a dictionary with file sizes
def count_files(directory):
    file_count = 0
    file_sizes = {}

    for root, dirs, files in os.walk(directory):
        for file in files:
            file_path = os.path.join(root, file)
            file_count += 1
            file_sizes[file_path] = os.path.getsize(file_path)

    return file_count, file_sizes

def main():
    if len(sys.argv) < 2:
        print(f"Usage: {sys.argv[0]} <directory>")
        sys.exit(1)

    directory = sys.argv[1]

    if not os.path.isdir(directory):
        print(f"Error: {directory} is not a valid directory")
        sys.exit(1)

    file_count, file_sizes = count_files(directory)

    print(f"Number of files in directory '{directory}': {file_count}")

    for file_name, size in file_sizes.items():
        print(f"File: {file_name}, Size: {size} bytes")

if __name__ == "__main__":
    main()
