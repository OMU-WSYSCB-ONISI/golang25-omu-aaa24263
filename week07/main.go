package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func calpmhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, `
			<h1>簡易計算機 (四則演算)</h1>
			<form method="POST">
				数値X: <input type="number" name="x" step="any" required><br>
				数値Y: <input type="number" name="y" step="any" required><br>
				<p>演算子を選択してください:</p>
				<input type="radio" name="cal0" value="+" required> + (加算)<br>
				<input type="radio" name="cal0" value="-" required> - (減算)<br>
				<input type="radio" name="cal0" value="*" required> * (乗算)<br>
				<input type="radio" name="cal0" value="/" required> / (除算)<br>
				<input type="submit" value="計算">
			</form>
		`)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintln(w, "フォームの解析エラーです。")
		return
	}

	xStr := r.FormValue("x")
	yStr := r.FormValue("y")
	operator := r.FormValue("cal0")

	x, errX := strconv.ParseFloat(xStr, 64)
	y, errY := strconv.ParseFloat(yStr, 64)

	if errX != nil || errY != nil {
		fmt.Fprintln(w, "無効な数値入力です。")
		return
	}

	var result float64
	var calculationError error

	switch operator {
	case "+":
		result = x + y
	case "-":
		result = x - y
	case "*":
		result = x * y
	case "/":
		if y == 0 {
			calculationError = fmt.Errorf("0での除算はできません")
		} else {
			result = x / y
		}
	default:
		calculationError = fmt.Errorf("無効な演算子が選択されました")
	}

	if calculationError != nil {
		fmt.Fprintln(w, calculationError.Error())
	} else {
		fmt.Fprintf(w, "計算: %.2f %s %.2f = **%.2f**", x, operator, y, result)
	}
}

func main() {
	http.HandleFunc("/", calpmhandler)
}
