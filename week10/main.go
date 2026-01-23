package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
)

const (
	saveDir  = "public"
	saveFile = "public/memo.txt"
)

func main() {
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		os.Mkdir(saveDir, 0755)
	}

	http.HandleFunc("/memo", memo)
	http.HandleFunc("/mwrite", mwrite)
	http.HandleFunc("/mdelete", mdelete)

	fmt.Println("Server launched on :8080")
	http.ListenAndServe(":8080", nil)
}

func memo(w http.ResponseWriter, r *http.Request) {
	text, err := os.ReadFile(saveFile)
	if err != nil {
		text = []byte("")
	}

	htmlText := html.EscapeString(string(text))

	s := `<html>
	<head>
		<title>Codespaces メモ帳</title>
		<style>
			body { font-family: sans-serif; padding: 20px; background: #f4f4f9; }
			.container { max-width: 800px; margin: auto; background: white; padding: 20px; border-radius: 8px; box-shadow: 0 2px 5px rgba(0,0,0,0.1); }
			textarea { width: 100%; height: 250px; padding: 10px; border: 1px solid #ddd; border-radius: 4px; font-size: 16px; box-sizing: border-box; }
			.info { margin: 10px 0; font-size: 14px; color: #666; text-align: right; }
			.btn-save { padding: 10px 25px; background: #007bff; color: white; border: none; border-radius: 4px; cursor: pointer; }
			.btn-save:hover { background: #0056b3; }
		</style>
	</head>
	<body>
		<div class="container">
			<h2> メモ帳</h2>
			<form method="get" action="/mwrite">
				<textarea name="text" id="memo-area" oninput="updateCount()">` + htmlText + `</textarea>
				<div class="info">現在の文字数: <span id="char-count">0</span> 文字</div>
				<input type="submit" value="保存する" class="btn-save" />
			</form>
		</div>

		<script>
			function updateCount() {
				const textarea = document.getElementById('memo-area');
				const countDisplay = document.getElementById('char-count');
				countDisplay.innerText = textarea.value.length;
			}
			window.onload = updateCount;
		</script>
	</body>
	</html>`
	w.Write([]byte(s))
}

func mwrite(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Form["text"]) > 0 {
		text := r.Form["text"][0]
		os.WriteFile(saveFile, []byte(text), 0644)
	}
	http.Redirect(w, r, "/memo", http.StatusSeeOther)
}

func mdelete(w http.ResponseWriter, r *http.Request) {
	os.WriteFile(saveFile, []byte(""), 0644)
	http.Redirect(w, r, "/memo", http.StatusSeeOther)
}
