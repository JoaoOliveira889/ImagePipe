# ImagePipe

> **Article:** [Optimizing Images for the Web with Go and Docker](https://joaooliveira.net/en/blog/2026/01/imagepipe/)

**ImagePipe** is a high-performance, cross-platform CLI tool built with **Go** and Docker designed to streamline image optimization for the web. It automates the conversion of JPG/PNG files to **WebP**, handles resizing, and supports both single-file and batch processing.

---

## What This Project Covers

- **WebP Conversion:** Dramatically reduces file size while maintaining visual quality.
- **Smart Resizing:** Automatically scales images to a max width of 1600px.
- **Batch Processing:** Optimize entire directories with one command.
- **Interactive Mode:** Supports drag-and-drop directly into the terminal.

---

## Tech Stack

* Language: Go (Golang)
* Containerization: Docker (Multi-stage builds)
* Base OS: Alpine Linux
* Libraries: Standard Go image processing packages

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

## About

This repository is part of my technical writing and learning notes.  
If you found it useful, consider starring the repo and sharing feedback.

- Author: Joao Oliveira
- Blog: https://joaooliveira.net
- Topics: Go, Docker, Web Performance, CLI Tools

## Contributing

Issues and pull requests are welcome.  
If you plan a larger change, please open an issue first so we can align on scope.

## License

Licensed under the **MIT License**. See the `LICENSE` file for details.
