package main

import (
  "fmt"
  "log"
  "net/smtp"
  "crypto/tls"
  "net"
  "bufio"
  "os"
)



func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
    c, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    gcm, err := cipher.NewGCM(c)
    if err != nil {
        return nil, err
    }

    nonceSize := gcm.NonceSize()
    if len(ciphertext) < nonceSize {
        return nil, errors.New("ciphertext too short")
    }

    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
    return gcm.Open(nil, nonce, ciphertext, nil)
}


func main(){
  var (
    host string
    port string
    password string
    to string
    from string
    subject string
    Message_body string
  )

  fmt.Printf("Write host: ")
  fmt.Scanf("%s\n",&host)

  host,port,err := net.SplitHostPort(host)
  //fmt.Printf("%s !! %s ",host,port)

  fmt.Printf("Send from: ")
  fmt.Scanf("%s\n",&from)

  key := []byte("there something you don't have t")
      pass, err := ioutil.ReadFile("password.txt")
      if err != nil {
          panic(err)
      }
      password, err := decrypt([]byte(pass), key)

  a:=smtp.PlainAuth("",from,password,host)




  fmt.Printf("Send to: ")
  fmt.Scanf("%s\n",&to)

  fmt.Printf("Write subject: ")
  str := ""
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    str += scanner.Text()+"\n"
  }
  fmt.Scanf("\n")
  subject=str


  fmt.Printf("Write message body: ")
  str=""
  scanner1 := bufio.NewScanner(os.Stdin)
  for scanner1.Scan(){
    str+= scanner1.Text()+"\n"
  }
  Message_body=str
  fmt.Scanf("\n")


  config := &tls.Config {
    InsecureSkipVerify: true,
    ServerName: host,
  }

  connection,err := tls.Dial("tcp",host+":"+port,config)
  if err != nil{
    log.Fatal(err)
  }

  c,err := smtp.NewClient(connection,host)
  if err!=nil{
    log.Fatal(err)
  }

  if err:= c.Auth(a);err!=nil{
    log.Fatal(err)
  }

  if err:= c.Mail(from); err!=nil{
    log.Fatal(err)
  }

  if err:= c.Rcpt(to);err!=nil{
    log.Fatal(err)
  }

  data,err:= c.Data()
  if err!=nil{
    log.Fatal(err)
  }


  message:="From: "+from +"\r\nTo: "+ to+"\r\nSubject: "+subject+"\r\n"+Message_body

  fmt.Printf("%s     ",[]byte(message))
  _,err = data.Write([]byte(message))
  if err!=nil{
    log.Fatal(err)
  } else {
    fmt.Printf("Mail sended!!\n")
  }

  err= data.Close()
  if err!=nil{
    log.Fatal(err)
  }

  c.Quit()

}
