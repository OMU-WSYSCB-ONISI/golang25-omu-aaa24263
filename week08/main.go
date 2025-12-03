package main

import (
    "fmt"
    "net/http"
    "strconv"
)

func scoreHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET"
        fmt.Fprintf(w, `
            <h1>得点の平均と分布計算</h1>
            <form method="POST">
                <p>得点（0〜100）を複数入力してください:</p>
                得点1: <input type="number" name="score" min="0" max="100" required><br>
                得点2: <input type="number" name="score" min="0" max="100" required><br>
                得点3: <input type="number" name="score" min="0" max="100" required><br>
                得点4: <input type="number" name="score" min="0" max="100" required><br>
                得点5: <input type="number" name="score" min="0" max="100" required><br>
                <input type="submit" value="計算・分布表示">
            </form>
         `)
        return
    }

    scoresStr := r.Form["score"]

    if len(scoresStr) == 0 {
        http.Error(w, "得点が入力されていません。", http.StatusBadRequest)
        return
    }

    totalScore := 0.0
    validCount := 0
    distribution := make([]int, 11)

    for _, s := range scoresStr {
        score, err := strconv.Atoi(s)
        if err != nil || score < 0 || score > 100 {
            http.Error(w, "無効な得点値が含まれています (0〜100の整数)。", http.StatusBadRequest)
            return
        }

        totalScore += float64(score)
        validCount++

        if score == 100 {
            distribution[10]++
        } else {
            distribution[score/10]++
        }
    }

    average := totalScore / float64(validCount)

    fmt.Fprintf(w, "<h2>集計結果</h2>")
    fmt.Fprintf(w, "全データ数: %d<br>", validCount)
    fmt.Fprintf(w, "合計得点: %.0f<br>", totalScore)
    fmt.Fprintf(w, "平均値: **%.2f**<br><br>", average)

    fmt.Fprintf(w, "<h2>得点分布 </h2>")
    fmt.Fprintf(w, "<table border='1' style='border-collapse: collapse;'>")
    fmt.Fprintf(w, "<tr><th>区間</th><th>度数</th></tr>")

    for i := 0; i < 10; i++ {
        fmt.Fprintf(w, "<tr><td>%d〜%d点</td><td>%d</td></tr>", i*10, i*10+9, distribution[i])
    }
    fmt.Fprintf(w, "<tr><td>100点</td><td>%d</td></tr>", distribution[10])
    fmt.Fprintf(w, "</table>")
}
