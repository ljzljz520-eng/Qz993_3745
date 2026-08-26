package report

import("fmt";"artstock/model")
func ValidateReport(r Report)error{if len(r.Records)!=r.Summary.Total{return fmt.Errorf("summary mismatch")};for _,x:=range r.Records{if !x.Valid(){return fmt.Errorf("invalid record %s",x.ID)}};return nil}
func StatusBreakdown(rs []model.Record)map[string]int{out:=map[string]int{};for _,r:=range rs{out[r.Status]++};return out}
func QuantityByCategory(rs []model.Record)map[string]int{out:=map[string]int{};for _,r:=range rs{out[r.Category]+=r.Quantity};return out}
func HasArchived(rs []model.Record)bool{for _,r:=range rs{if r.Status=="archived"{return true}};return false}
