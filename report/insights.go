package report

import("sort";"artstock/model")
type Insight struct{Category string; Quantity int; Records int; Share float64}
func Insights(rs []model.Record)[]Insight{groups:=GroupByCategory(rs);out:=make([]Insight,0,len(groups));total:=0;for _,r:=range rs{total+=r.Quantity};for k,list:=range groups{q:=0;for _,r:=range list{q+=r.Quantity};share:=0.0;if total>0{share=float64(q)/float64(total)};out=append(out,Insight{Category:k,Quantity:q,Records:len(list),Share:share})};sort.Slice(out,func(i,j int)bool{return out[i].Quantity>out[j].Quantity});return out}
func TopCategory(rs []model.Record) string{is:=Insights(rs);if len(is)==0{return ""};return is[0].Category}
func LowStock(rs []model.Record,threshold int)[]model.Record{out:=[]model.Record{};for _,r:=range rs{if r.Quantity<=threshold&&r.Status!="archived"{out=append(out,r)}};return out}
