<<<<<<< HEAD
# filepath: /Users/mb/Projects/my-cli-app/ollama.Dockerfile
=======
>>>>>>> f1f5527 (All changes: static assets, Docker, Go, and UI polish)
FROM ollama/ollama

# Expose the Ollama API port
EXPOSE 11434

# Start the Ollama server (model will be pulled at runtime)
CMD ["serve"]
