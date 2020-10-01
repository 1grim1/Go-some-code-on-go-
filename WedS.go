package main

import (

"fmt" // пакет для форматированного ввода вывода

"net/http" // пакет для поддержки HTTP протокола


"log" // пакет для логирования

"io/ioutil"
)

func hadler(w http.ResponseWriter, r *http.Request) {
    str,err:= ioutil.ReadFile("A:/js v/aa/public/html/creation.tx") 
    if err!= nil {
      panic(err)
    }
    fmt.Fprintf(w,string(str) ) // отправляем данные на клиентскую сторону
}

  func main() {

    http.HandleFunc("/",hadler)

    err := http.ListenAndServe(":4343", nil) // задаем слушать порт

    if err != nil {

        log.Fatal("ListenAndServe: ", err)

    }

}
