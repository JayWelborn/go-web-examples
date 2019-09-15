package main

import (
    "html/template"
    "net/http"
    "fmt"
)

type ContactDetails struct {
    Email   string
    Subject string
    Message string
}

type Received struct {
    Success bool
    Details ContactDetails
}

func main() {
    tmpl := template.Must(template.ParseFiles("forms.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        details := ContactDetails{
            Email:   r.FormValue("email"),
            Subject: r.FormValue("subject"),
            Message: r.FormValue("message"),
        }

        // do something with details
        received := Received{
            Success: true,
            Details: details,
        }

        fmt.Println(details)

        tmpl.Execute(w, received)
    })

    http.ListenAndServe(":8080", nil)
}

