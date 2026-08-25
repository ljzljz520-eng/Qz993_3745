package inventory

import("artstock/model";"artstock/store")
func (s *Service) Search(query string) ([]model.Record,model.Summary,error) { rs,e:=s.Store.ListRecords();if e!=nil{return nil,model.Summary{},e}; rs=store.FilterRecords(rs,query);rs=store.SortRecords(rs);return rs,store.Summarize(rs),nil }
func (s *Service) Available(query string)([]model.Record,error){rs,_,e:=s.Search(query);if e!=nil{return nil,e};out:=[]model.Record{};for _,r:=range rs{if r.Available(){out=append(out,r)}};return out,nil}
func (s *Service) Reserve(id,actor string) error {r,e:=s.Get(id);if e!=nil{return e};if !r.Available(){return model.ErrUnavailable};r.SetStatus("reserved");if e=s.Store.PutRecord(r);e!=nil{return e};return s.Store.PutEvent(model.NewEvent(id+"-reserve",id,"reserve",actor,"reserved"))}
