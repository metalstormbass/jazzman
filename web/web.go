package web

import (
	"encoding/json"

	"fmt"


	"jazzman/llm"
	"net/http"
	"path/filepath"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/static/jazzman.jpeg" {

		// Serve the static image file


		imgPath := filepath.Join("web", "static", "jazzman.jpeg")
		http.ServeFile(w, r, imgPath)
		return
	}


	if r.URL.Path == "/static/style.css" {
		cssPath := filepath.Join("web", "static", "style.css")
		http.ServeFile(w, r, cssPath)
		return
	}
	if r.URL.Path == "/static/index.html" {
		indexPath := filepath.Join("web", "static", "index.html")
		http.ServeFile(w, r, indexPath)
		return
	}

	if r.Method == http.MethodPost {
		quote := llm.GenerateBadJazzQuote()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"quote": quote})
		return
	}

	quote := llm.GenerateBadJazzQuote()
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Jazzman</title>
<style>
body {
  background: repeating-linear-gradient(135deg, #232526 0%%, #414345 10%%, #1a1a1a 20%%, #232526 30%%, #414345 40%%, #1a1a1a 50%%);
  color: #e0e0e0;
  font-family: 'Comic Sans MS', 'Papyrus', cursive, sans-serif;
  min-height: 100vh;
  margin: 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  animation: bgmove 8s linear infinite alternate;
}
@keyframes bgmove {
  0%% { filter: hue-rotate(0deg); }
  100%% { filter: hue-rotate(90deg); }
}
.container {
  background: rgba(30, 30, 40, 0.92);
  border-radius: 18px 2px 18px 2px;
  box-shadow: 0 0 40px 10px #ff00cc, 0 0 10px 2px #00ffcc;
  padding: 2.5rem 2.5rem 2rem 2.5rem;
  max-width: 420px;
  text-align: center;
  /* Removed border: 4px dashed #ff00cc; */
  transform: rotate(-2deg) scale(1.01);
}
h1 {
  font-family: 'Impact', 'Arial Black', cursive, sans-serif;
  font-size: 2.2rem;
  letter-spacing: 2px;
  margin-bottom: 1.2rem;
  color: #ff00cc;
  text-shadow: 2px 2px 8px #00ffcc, 0 0 2px #ff00cc;
  transform: rotate(2deg);
}
#quote-box {
  background: #1a1a1a;
  border-radius: 10px;
  box-shadow: 0 2px 8px #ff00cc, 0 0 8px #00ffcc;
  padding: 1.2rem 1rem;
  margin: 1.5rem 0 2rem 0;
  min-height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 2.5px dotted #00ffcc;
  transform: rotate(-1.5deg);
}
#quote {
  font-size: 1.1rem;
  color: #ffb347;
  transition: color 0.2s;
  font-family: 'Courier New', Courier, monospace;
  letter-spacing: 1px;
}
img.jazzman-img {
  width: 80px;
  height: 80px;
  object-fit: cover;
  border-radius: 50%%;
  margin-bottom: 1.1rem;
  box-shadow: 0 0 20px 4px #ff00cc, 0 0 8px 2px #00ffcc;
  border: 3px solid #ffb347;
  filter: grayscale(80%%) contrast(180%%) blur(0.5px) hue-rotate(30deg);
  transform: rotate(-8deg) scale(0.95);
  animation: imgshake 1.2s infinite alternate;
}
@keyframes imgshake {
  0%% { transform: rotate(-8deg) scale(0.95); }
  100%% { transform: rotate(8deg) scale(1.05); }
}
button {
  background: #e0e0e0;
  color: #3a3a3a;
  font-weight: bold;
  border: 2px solid #3a3a3a;
  padding: 12px 28px;
  border-radius: 8px;
  box-shadow: 0 2px 8px #222;
  transition: background 0.2s, color 0.2s;
  font-family: 'Comic Sans MS', 'Papyrus', cursive, sans-serif;
  margin-top: 0.5rem;
}
button:hover {
  background: #b0b0b0;
  color: #232526;
  border-color: #232526;
}
@media (max-width: 600px) {
  .container {
    padding: 1.2rem 0.5rem 1.2rem 0.5rem;
    max-width: 98vw;
  }
  h1 {
    font-size: 1.3rem;
  }
  #quote {
    font-size: 0.9rem;
  }
  img.jazzman-img {
    width: 50px;
    height: 50px;
  }
}
</style>
<script>
function getQuote() { const quoteElem = document.getElementById('quote'); quoteElem.style.color = '#ffb347'; fetch('/', {method: 'POST'}).then(r => r.json()).then(data => { quoteElem.textContent = data.quote; quoteElem.style.color = '#f8f8f8'; }); }
</script>
</head>
<body>
<div style="position:relative; display:flex; flex-direction:column; align-items:center; justify-content:center; min-height:100vh;">
  <div class="container">
    <img src="/static/jazzman.jpeg" alt="Unsettling Jazzman" class="jazzman-img" />
    <h1>Jazzman</h1>
    <div id="quote-box"><span id="quote">%s</span></div>
    <button onclick="getQuote()">Generate Another</button>
  </div>
</div>
</body>
</html>`, quote)

	// Serve static index.html for all other GET requests
	indexPath := filepath.Join("web", "static", "index.html")
	http.ServeFile(w, r, indexPath)

}
