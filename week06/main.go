package main

import (
    "fmt"
    "net/http"
    "strconv"
)

func bmiHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        fmt.Fprintf(w, `
            <h1>BMI計算</h1>
            <form method="POST">
                体重(kg, 整数): <input type="number" name="weight" required><br>
                身長(cm, 整数): <input type="number" name="height" required><br>
                <input type="submit" value="計算">
            </form>
        `)
        return
    }

    weightStr := r.FormValue("weight")
    heightStr := r.FormValue("height")

    weight, errW := strconv.ParseFloat(weightStr, 64)
    height, errH := strconv.ParseFloat(heightStr, 64)

    if errW != nil || errH != nil || weight <= 0 || height <= 0 {
        http.Error(w, "無効な入力です。正の整数を入力してください。", http.StatusBadRequest)
        return
    }

    heightM := height / 100.0
    bmi := weight / (heightM * heightM)

    fmt.Fprintf(w, "体重: %.0f kg, 身長: %.0f cm<br>", weight, height)
    fmt.Fprintf(w, "あなたの **BMI値** は: **%.2f** です。", bmi)
}

func main() {
    http.HandleFunc("/bmi", bmiHandler)
}
