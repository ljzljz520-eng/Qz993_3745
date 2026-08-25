package workflow

import("sort";"artstock/inventory";"artstock/model")
type Dashboard struct{Summary model.Summary; Recent []model.Record; Alerts []string}
func BuildDashboard(i *inventory.Service) (Dashboard,error){rs,s,e:=i.Search("");if e!=nil{return Dashboard{},e};sort.Slice(rs,func(a,b int)bool{return rs[a].UpdatedAt.After(rs[b].UpdatedAt)});d:=Dashboard{Summary:s};if len(rs)>5{d.Recent=rs[:5]}else{d.Recent=rs};for _,r:=range rs{if r.Quantity==0&&r.Status!="archived"{d.Alerts=append(d.Alerts,r.ID+": empty")}};return d,nil}
func (d Dashboard) NeedsAttention() bool{return len(d.Alerts)>0||d.Summary.Processing>0}
func (d Dashboard) StatusText() string{if d.NeedsAttention(){return "attention"};if d.Summary.Total==0{return "empty"};return "normal"}
