package main

import (
    "fmt"
    "net/smtp"
    "os"
)

func main() {
    var (
      from string
      to string
      host string
      password string
    )
    fmt.Printf("We are authorizing as: ")
    fmt.Scanf("%s\n",&from)

    fmt.Printf("We are sending email: ")
    fmt.Scanf("%s\n",&to)

    fmt.Printf("Ship via: ")
    fmt.Scanf("%s\n",&host)

    fmt.Printf("Write password: ")
    fmt.Scanf("%s\n",&password)
    auth := smtp.PlainAuth("", from,password,host)


    message:=To(to) +"From: <"+ from +">\n"+ Subject() + Message_body()

    fmt.Printf(message)

    if err := smtp.SendMail(host, auth, from, []string{to}, []byte(message)); err != nil {
        fmt.Println("Error SendMail: ", err)
        os.Exit(1)
    }
    fmt.Println("Email Sent!")
}


func To(to string) string {
  return ("To: <" + to + ">\n")
}

func Subject() string {
  var subject string
  fmt.Printf("Field Subject: ")
  fmt.Scanf("%s\n",&subject)
  return ("Subject: " + subject+ "\n")
}


func Message_body() string {
  var body string
    fmt.Printf("Field Message body: ")
  fmt.Scanf("%s\n",&body)
  return (body + "<CRLF>.<CRLF>")
}
