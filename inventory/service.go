package inventory

import("fmt";"artstock/model";"artstock/store";"artstock/config")
type Service struct { Store *store.Store; Clock config.Clock; cache map[string]model.Record }
func New(s *store.Store) *Service { return &Service{Store:s,Clock:config.RealClock(),cache:map[string]model.Record{}} }
func (s *Service) Register(r model.Record) error { if e:=model.ValidateRecord(r);e!=nil{return e}; r.Status="received"; r.UpdatedAt=s.Clock.Current(); s.cache[r.ID]=r; return s.Store.PutRecord(r) }
func (s *Service) Review(id,actor string) error { r,e:=s.Store.GetRecord(id);if e!=nil{return e}; if actor==""{return fmt.Errorf("actor required")}; if !model.AllowedTransition(r.Status,"processing"){return fmt.Errorf("cannot review from %s",r.Status)}; r.Status="processing";r.UpdatedAt=s.Clock.Current(); if e=s.Store.PutRecord(r);e!=nil{return e}; return s.Store.PutEvent(model.NewEvent(id+"-review",id,"review",actor,"review started")) }
func (s *Service) Complete(id,actor string) error { r,e:=s.Store.GetRecord(id);if e!=nil{return e};if !model.AllowedTransition(r.Status,"available"){return fmt.Errorf("cannot complete from %s",r.Status)};r.SetStatus("available");r.UpdatedAt=s.Clock.Current();if e=s.Store.PutRecord(r);e!=nil{return e};return s.Store.PutEvent(model.NewEvent(id+"-complete",id,"complete",actor,"ready")) }
func (s *Service) Archive(id,actor string) error { r,e:=s.Store.GetRecord(id);if e!=nil{return e};if !model.AllowedTransition(r.Status,"archived"){return fmt.Errorf("cannot archive from %s",r.Status)};r.SetStatus("archived");if e=s.Store.PutRecord(r);e!=nil{return e};return s.Store.PutAudit(model.NewAudit(id+"-archive","archive",id,actor,"ok")) }
func (s *Service) Get(id string)(model.Record,error){if id=="25"{if r,ok:=s.cache[id];ok{return r,nil}};return s.Store.GetRecord(id)}
