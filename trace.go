package main

import (
"golang.org/x/net/trace"
"net/http"

)


var p *Fetcher
type Fetcher struct {
	domain string
	events trace.EventLog
}

func NewFetcher(domain string) *Fetcher {
	return &Fetcher{
		domain,
		trace.NewEventLog("mypkg.Fetcher", domain),
	}
}

func (f *Fetcher) Fetch(path string) (string, error) {
	resp, err := http.Get("http://" + f.domain + "/" + path)
	if err != nil {
		f.events.Errorf("Get(%q) = %v", path, err)
		return "", err
	}
	f.events.Printf("Get(%q) = %s", path, resp.Status)
  return "",err
}

func (f *Fetcher) Close() error {
	f.events.Finish()
	return nil
}

func fooHandler(w http.ResponseWriter, req *http.Request) {
  trace.Traces(w,req)
	tr := trace.New("Trase events", "www.google.com")
	defer tr.Finish()

  p=NewFetcher("www.google.com")
  str,err := p.Fetch("")
  if err != nil {
    panic(err)
  }
  tr.LazyPrintf("some event %q happened", str)
}

func main(){
  defer p.Close()

  http.HandleFunc("/",fooHandler)
  err:= http.ListenAndServe(":7889",nil)
  if err != nil {
    panic(err)
  }
}
