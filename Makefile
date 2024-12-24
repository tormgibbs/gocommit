OUTPUT_LINUX=gocommit-linux
OUTPUT_WINDOWS=gocommit.exe

# Default target
all: linux windows

# Build for Linux
linux:
	GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o $(OUTPUT_LINUX)

# Build for Windows
windows:
	GOOS=windows GOARCH=amd64 go build -ldflags='-s' -o $(OUTPUT_WINDOWS)

# Clean up build artifacts
clean:
	rm -f $(OUTPUT_LINUX) $(OUTPUT_WINDOWS)

# Rebuild everything
rebuild: clean all
