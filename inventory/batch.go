package inventory

import("fmt";"artstock/model")
type BatchResult struct{Accepted int; Rejected int; Errors []string}
func (s *Service) RegisterBatch(rs []model.Record) BatchResult{out:=BatchResult{};for _,r:=range rs{if e:=s.Register(r);e!=nil{out.Rejected++;out.Errors=append(out.Errors,e.Error())}else{out.Accepted++}};return out}
func (s *Service) Transition(id,to,actor string) error{r,e:=s.Get(id);if e!=nil{return e};if !model.AllowedTransition(r.Status,to){return fmt.Errorf("transition %s to %s denied",r.Status,to)};r.SetStatus(to);if e=s.Store.PutRecord(r);e!=nil{return e};return s.Store.PutEvent(model.NewEvent(id+"-"+to,id,to,actor,"state changed"))}
func (s *Service) Reconcile(id string,expected int) error{r,e:=s.Get(id);if e!=nil{return e};if expected<0{return fmt.Errorf("expected quantity invalid")};if r.Quantity!=expected{r.Quantity=expected;r.UpdatedAt=s.Clock.Current();return s.Store.PutRecord(r)};return nil}
func (s *Service) Rename(id,name string) error{r,e:=s.Get(id);if e!=nil{return e};if name==""{return fmt.Errorf("name required")};r.Name=name;r.UpdatedAt=s.Clock.Current();return s.Store.PutRecord(r)}
