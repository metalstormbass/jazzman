#!/bin/zsh

# Start all services in the background
echo "Starting Docker Compose services..."
docker compose up -d --build

# Wait for the Ollama server to be ready
echo "Waiting for Ollama server to be ready..."
RETRIES=30
until curl -sf http://localhost:11434/api/tags > /dev/null; do
  ((RETRIES--))
  if [[ $RETRIES -le 0 ]]; then
    echo "Ollama server did not become ready in time."
    exit 1
  fi
  sleep 2
done

echo "Ollama server is up. Running the llama3.2 model..."
docker compose exec llm ollama run llama3.2

echo "All services are up and the model is running."
