# filepath: /Users/mb/Projects/my-cli-app/ollama.Dockerfile
FROM ollama/ollama

# Expose the Ollama API port
EXPOSE 11434

# Start the Ollama server (model will be pulled at runtime)
CMD ["serve"]
