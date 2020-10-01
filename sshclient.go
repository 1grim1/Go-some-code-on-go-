package main

import (
  "golang.org/x/crypto/ssh"
  "fmt"
  "io"
  "os"
  "net"
  "bufio"
)


func Sessions(connection *ssh.Client,cmd string){
  session,err := connection.NewSession()
  if err!=nil {
    panic(err)}


  stdin, err := session.StdinPipe()
  if err != nil {
     fmt.Errorf("Unable to setup stdin for session: %v", err)
   }
  go io.Copy(stdin, os.Stdin)

  stdout, err := session.StdoutPipe()
  if err != nil {
     fmt.Errorf("Unable to setup stdout for session: %v", err)
  }
  go io.Copy(os.Stdout, stdout)

  stderr, err := session.StderrPipe()
  if err != nil {
     fmt.Errorf("Unable to setup stderr for session: %v", err)
   }
  go io.Copy(os.Stderr, stderr)

  defer session.Close()

  fmt.Printf("Your command: %s\n",cmd)
  session.Run(cmd)
}



func main(){

  sshConfig := &ssh.ClientConfig{
  	User: "iu9_31_07",
  	Auth: []ssh.AuthMethod{
  		ssh.Password("Dimonchi1")},

    HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
                return nil
            },
  }


  var host_port string
  fmt.Scanf("%s\n",&host_port)

  connection,err := ssh.Dial("tcp",host_port,sshConfig)
  if err!=nil {
    panic(err)}


  q:=true
  var str string
  for q {
          fmt.Printf("Write command: ")
          scanner := bufio.NewScanner(os.Stdin)
          str = ""
          for scanner.Scan() {
              str += scanner.Text()
          }
          if str == "quit" {
              q = false
          } else {
              Sessions(connection, str)
          }
      }

}
