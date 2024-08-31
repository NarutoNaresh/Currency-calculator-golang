package main

import (
    "fmt"
    "html/template"
    "net/http"
    "strconv"
)

const (
    eurRate = 1.1  // Example exchange rate from USD to EUR
    gbpRate = 0.8  // Example exchange rate from USD to GBP
    jpyRate = 140  // Example exchange rate from USD to JPY
    inrRate = 82   // Example exchange rate from USD to INR
    dinarRate = 0.3 // Example exchange rate from USD to Kuwaiti Dinar
)

func main() {
    http.HandleFunc("/", handleRequest)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    fmt.Println("Server started at http://localhost:5050")
    http.ListenAndServe(":5050", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        r.ParseForm()
        amountStr := r.FormValue("amount")
        currency := r.FormValue("currency")

        amount, err := strconv.ParseFloat(amountStr, 64)
        if err != nil {
            http.Error(w, "Invalid amount", http.StatusBadRequest)
            return
        }

        var convertedAmount float64
        var currencyName string

        switch currency {
        case "eur":
            convertedAmount = amount * eurRate
            currencyName = "EUR"
        case "gbp":
            convertedAmount = amount * gbpRate
            currencyName = "GBP"
        case "jpy":
            convertedAmount = amount * jpyRate
            currencyName = "JPY"
        case "inr":
            convertedAmount = amount * inrRate
            currencyName = "INR"
        case "dinar":
            convertedAmount = amount * dinarRate
            currencyName = "KWD" // Kuwaiti Dinar
        default:
            http.Error(w, "Invalid currency", http.StatusBadRequest)
            return
        }

        data := struct {
            Amount          float64
            ConvertedAmount float64
            CurrencyName    string
        }{
            Amount:          amount,
            ConvertedAmount: convertedAmount,
            CurrencyName:    currencyName,
        }

        tmpl := template.Must(template.ParseFiles("templates/index.html"))
        tmpl.Execute(w, data)
        return
    }

    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    tmpl.Execute(w, nil)
}
