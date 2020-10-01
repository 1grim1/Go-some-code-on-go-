package main

import (
    "fmt" // пакет для форматированного ввода вывод
    "github.com/jlaffaye/ftp"
    "time"
    "log"
    "bytes"
    "io/ioutil"
    "os"
    "bufio"
)
var quit bool


func main(){
  var path string
  fmt.Scanf("%s\n",&path)
  c,err := ftp.Dial(path,ftp.DialWithTimeout(5 * time.Second))
  if err != nil {log.Fatal(err)}
  Login(c)
  quit=true
  for quit {
    Com(c)
  }
}



func Com(c *ftp.ServerConn){
  var cmd string
  fmt.Scanf("%s\n",&cmd)
  switch cmd {
    case "makedir" :
      MakeDir(c)
    case "removedir" :
      RemoveDir(c)
    case "load" :
      SaveFile(c)
    case "ls" :
      LS(c)
    case "read" :
      Read(c)
    case "delete" :
      Delete(c)
    case "quit" :
      quit = false
      if err := c.Quit();err != nil {
          log.Fatal(err)
        }
    }
}

func Login(ftp *ftp.ServerConn){
  var login,password string
  fmt.Scanf("%s %s\n",&login,&password)
  err := ftp.Login(login,password)
  if err!=nil {log.Fatal(err)}
}

func MakeDir(ftp *ftp.ServerConn){
  var path string
  fmt.Scanf("%s\n",&path)
  err := ftp.MakeDir(path)
  if err!=nil {log.Fatal(err)}
}

func RemoveDir(ftp *ftp.ServerConn){
  var path string
  fmt.Scanf("%s\n",&path)
  err := ftp.RemoveDir(path)
  if err!=nil{log.Fatal(err)}
}

func SaveFile(ftp *ftp.ServerConn){
  var path string

  fmt.Scanln(&path)
  str := ""
  scanner := bufio.NewScanner(os.Stdin)

  for scanner.Scan() {
      str2 := scanner.Text()
      str += str2
  }
  data := bytes.NewBufferString(str)
  err := ftp.Stor(path,data)
  if err!=nil {panic(err)}
}

func LS(ftp *ftp.ServerConn){
  var path string
  fmt.Scanf("%s\n",&path)
  entries, err := ftp.List(path)
  if err!=nil {panic(err)}

  for _, item := range entries {
      fmt.Printf("%s %s\n",item.Name,item.Target)
  }



}

func  Read(ftp *ftp.ServerConn){
  var path ,ospath string
  fmt.Scanf("%s\n",&path)
  Response,err := ftp.Retr(path)
  if err!=nil {log.Fatal(err)}

  byte,error := ioutil.ReadAll(Response)
  if error!= nil {log.Fatal(error)}

  Response.Close()
  fmt.Scanf("%s\n",&ospath)
  save,err:=os.Create(ospath)
  if err!=nil {panic(err)}
  save.Write(byte)
  fmt.Printf(string(byte))
  save.Close()
}

func Delete(ftp *ftp.ServerConn){
  var path string
  fmt.Scanf("%s\n",&path)
  err := ftp.Delete(path)
  if err!=nil {panic(err)}
}
