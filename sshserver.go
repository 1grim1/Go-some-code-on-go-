package main

import (
    "golang.org/x/crypto/ssh/terminal"
    "log"
    "os"
    "github.com/gliderlabs/ssh"
    "fmt"
    "io/ioutil"
)


func LS(term *terminal.Terminal){
  xs,err := term.ReadLine()
  dir,_:= os.Getwd()
  if err!=nil {panic(err)}
  data,err := ioutil.ReadDir(dir + string(os.PathSeparator) + xs)
  if err!=nil {panic(err)}
  list:=""
  for _, k:= range data {
    list += k.Name() + "\n"
  }
  term.Write([]byte(list))
}

func MAKEDIR(term *terminal.Terminal){
  xs,err:= term.ReadLine()
  dir,_:= os.Getwd()
  if err!=nil{panic(err)}
  os.Mkdir(dir + string(os.PathSeparator) + xs,0777)
}

func RMDIR(term *terminal.Terminal){
  xs,err:= term.ReadLine()
  dir,_:= os.Getwd()
  if err!=nil{panic(err)}
  os.Remove(dir + string(os.PathSeparator) + xs)
}

func Q(){
  os.Exit(1)
}


func Command(line string,term *terminal.Terminal){
  switch line {
    case "ls":
      LS(term)
    case "mkdir":
      MAKEDIR(term)
    case "rmdir":
      RMDIR(term)
    case "quit":
      Q()
  }
}


func main() {
  ssh.Handle(func(sess ssh.Session) {
    term:= terminal.NewTerminal(sess,"> ")
    for {
      term.Write([]byte("Write command : "))
      line,err := term.ReadLine()
      if err!=nil {panic(err)}
      Command(line,term)
    }
    fmt.Println("terminal closed")
  })

  err := ssh.ListenAndServe(":7889",nil, ssh.PasswordAuth(func(ctx ssh.Context,pass string) bool { return pass == "coala"}),)
  if err!=nil {
    log.Fatal(err)
  } else {
    fmt.Printf("Server starting\n")
  }
}
