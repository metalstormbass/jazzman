# Jazzman

Jazzman is a simple Go web app that generates intentionally bad jazz quotes using a local LLM API (such as Ollama).

## Project Structure

```
main.go         # Main Go application
Dockerfile      # Container build for Go app
README.md       # Project documentation
go.mod          # Go module definition
docker-compose.yml # Compose file for Go app and LLM service
```

## Prerequisites

- **Docker**: Required for running containers ([Install Docker](https://docs.docker.com/get-docker/))
- **Docker Compose**: Included with Docker Desktop ([Compose info](https://docs.docker.com/compose/))
- **Git** (optional, for cloning): [Install Git](https://git-scm.com/downloads)

## Running on Windows or Mac

### 1. Install Docker Desktop
- [Docker Desktop for Mac](https://docs.docker.com/desktop/install/mac/)
- [Docker Desktop for Windows](https://docs.docker.com/desktop/install/windows-install/)

### 2. Clone or Download the Repository

```
git clone <your-repo-url>
cd my-cli-app
```
Or download and unzip, then open a terminal in the project folder.

### 3. Build and Run the App

Open a terminal (Command Prompt, PowerShell, or Terminal on Mac) and run:

```
docker compose up --build
```

This will build and start both the Jazzman app and the Ollama LLM service. The first run may take a while as the LLM model is downloaded.

### 4. Access the App

Open your browser and go to: [http://localhost:8080](http://localhost:8080)

You should see the Jazzman web app. Click the button to generate a new quote.

---

## Environment Variables

- `LLM_API_URL` (optional): URL for the LLM API (default: http://llm:11434)
- `PORT` (optional): Port for the web server (default: 8080)

---

For more on Docker and Docker Compose, see the [official documentation](https://docs.docker.com/get-started/).
