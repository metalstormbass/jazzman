// filepath: /Users/mb/Projects/my-cli-app/test/test_llm.sh
#!/bin/zsh

LLM_API_URL=${LLM_API_URL:-http://localhost:11434}

RESPONSE=$(curl -s -X POST "$LLM_API_URL/api/generate" \
  -H 'Content-Type: application/json' \
  -d '{"model": "llama3.2", "prompt": "Generate a short, intentionally bad and mocking quote about jazz or jazz musicians. The quote should not hold jazz in high regard and should sound like a joke.", "stream": false}')

QUOTE=$(echo "$RESPONSE" | grep -o '"response"[ ]*:[ ]*"[^"}]*"' | sed 's/"response"[ ]*:[ ]*"\(.*\)"/\1/')

if [[ -z "$QUOTE" ]]; then
  echo "[LLM test failed] No quote returned!"
  echo "Raw response: $RESPONSE"
  exit 1
else
  echo "LLM responded with: $QUOTE"
  exit 0
fi
