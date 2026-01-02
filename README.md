# ImagePipe

A high-performance, **cross-platform** CLI tool built with **Go** and **Docker** to optimize images for the web. It converts JPG/PNG files to **WebP**, handles automatic resizing, and supports both single-file and batch processing.

## Features

- **WebP Conversion:** Dramatically reduces file size while maintaining visual quality.
- **Smart Resizing:** Automatically scales images to a max width of 1600px.
- **Batch Processing:** Optimize entire directories with one command.
- **Interactive Mode:** Supports drag-and-drop directly into the terminal.

---

## Installation

### 1. Build the Docker Image

Ensure Docker is installed and running, then execute:

```bash
git clone https://github.com/JoaoOliveira889/ImagePipe.git
cd ImagePipe
docker build -t imagepipe .
```

## 2. Global Shell Integration (Zsh/Bash)

To use imagepipe from any folder on your machine, add this function to your shell configuration:

Open your config:

```bash
nvim ~/.zshrc
```

Add the following function:

```bash
imagepipe() {
    docker run --rm -it -v "$(pwd)":/data imagepipe "$1" "$2"
}
```

Apply changes: 

```bash
source ~/.zshrc
```

## Usage

| Goal           | Command                |
| -------------- | ---------------------- |
| Single Image   | imagepipe photo.jpg    |
| Custom Quality | imagepipe photo.jpg 90 |
| Batch Folder   | imagepipe .            |
| Without docker | go run main.go         |

## Architecture

The tool uses a multi-stage Docker build to keep the final image extremely lightweight (under 50MB).

- Base: Alpine Linux (for security and size).
- Architecture: Automatically detects arm64 vs amd64 during build.
- Logic: Volume mapping (-v) connects your local Mac folders to the container's internal processing directory.

## Notes

- Timestamps: Output files include a timestamp (e.g., image_123045.webp) to prevent overwriting your files.
- Originals: Your original JPG/PNG files are never modified or deleted.
