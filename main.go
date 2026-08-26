package main

import("log";"net/http";"artstock/api";"artstock/config";"artstock/inventory";"artstock/store")
func main(){c:=config.FromEnv();s,e:=store.Open(c.DataPath);if e!=nil{log.Fatal(e)};defer s.Close();i:=inventory.New(s);srv:=api.New(i);log.Printf("artstock listening on %s",c.Address);if e=http.ListenAndServe(c.Address,srv.Handler());e!=nil{log.Fatal(e)}}
