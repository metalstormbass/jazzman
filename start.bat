@echo off
REM Start all services in the background
ECHO Starting Docker Compose services...
docker compose up -d --build

REM Wait for the Ollama server to be ready
ECHO Waiting for Ollama server to be ready...
SET RETRIES=30
:waitloop
curl -sf http://localhost:11434/api/tags >nul 2>&1
IF %ERRORLEVEL% EQU 0 GOTO ready
SET /A RETRIES=%RETRIES%-1
IF %RETRIES% LEQ 0 (
  ECHO Ollama server did not become ready in time.
  EXIT /B 1
)
TIMEOUT /T 2 >nul
GOTO waitloop
:ready
ECHO Ollama server is up. Running the llama3.2 model...
docker compose exec llm ollama run llama3.2
ECHO All services are up and the model is running.
