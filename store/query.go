package store

import("artstock/model";"strings")
func FilterRecords(rs []model.Record, query string) []model.Record { q:=strings.ToLower(strings.TrimSpace(query)); if q=="" { return append([]model.Record(nil),rs...) }; out:=[]model.Record{}; for _,r:=range rs { if strings.Contains(strings.ToLower(r.Name),q)||strings.Contains(strings.ToLower(r.Category),q)||strings.Contains(strings.ToLower(r.Status),q){out=append(out,r)} }; return out }
func Summarize(rs []model.Record) model.Summary { var s model.Summary; for _,r:=range rs{s.Add(r)}; return s }
func SortRecords(rs []model.Record) []model.Record { out:=append([]model.Record(nil),rs...); for i:=0;i<len(out);i++{for j:=i+1;j<len(out);j++{if out[j].UpdatedAt.Before(out[i].UpdatedAt){out[i],out[j]=out[j],out[i]}}};return out }
