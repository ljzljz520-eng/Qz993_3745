package workflow

import("fmt";"artstock/inventory")
type Notice struct { RecordID string; Message string; Delivered bool }
func Notify(i *inventory.Service,id,actor string)(Notice,error){r,e:=i.Get(id);if e!=nil{return Notice{},e};n:=Notice{RecordID:id,Message:fmt.Sprintf("%s is %s",r.Name,r.Status)};if actor!=""{n.Delivered=true};return n,nil}
func Track(i *inventory.Service,id string)([]string,error){ev,e:=i.Store.ListEvents(id);if e!=nil{return nil,e};out:=[]string{};for _,x:=range ev{out=append(out,x.Kind+":"+x.Detail)};return out,nil}
