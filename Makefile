OUTPUT_LINUX=gocommit-linux
OUTPUT_WINDOWS=gocommit.exe

# Default target
all: linux windows

# Build for Linux
linux:
	@echo "Building for Linux..."
	@GOOS=linux GOARCH=amd64 go build -ldflags='-s' -o $(OUTPUT_LINUX)

# Build for Windows
windows:
	@echo "Building for Windows..."
	@GOOS=windows GOARCH=amd64 go build -ldflags='-s' -o $(OUTPUT_WINDOWS)

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	@rm -f $(OUTPUT_LINUX) $(OUTPUT_WINDOWS)

# Rebuild everything
rebuild: 
	@echo "Rebuilding..."
	@clean all
