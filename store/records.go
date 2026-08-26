package store

import (
 "fmt"
 "artstock/model"
 "go.etcd.io/bbolt"
)
func (s *Store) ListRecords() ([]model.Record,error) { s.mu.RLock(); defer s.mu.RUnlock(); if s.db==nil{return nil,fmt.Errorf("store closed")}; out:=[]model.Record{}; e:=s.db.View(func(tx *bbolt.Tx)error{return tx.Bucket([]byte("records")).ForEach(func(_,v []byte)error{var r model.Record; if e:=decode(v,&r);e!=nil{return e}; out=append(out,r); return nil})}); return out,e }
func (s *Store) PutUser(u model.User) error{return s.put("users",u.ID,u)}
func (s *Store) GetUser(id string)(model.User,error){var u model.User;e:=s.get("users",id,&u);return u,e}
func (s *Store) PutEvent(ev model.Event) error{return s.put("events",ev.ID,ev)}
func (s *Store) ListEvents(recordID string)([]model.Event,error){s.mu.RLock();defer s.mu.RUnlock(); if s.db==nil{return nil,fmt.Errorf("store closed")}; out:=[]model.Event{};e:=s.db.View(func(tx *bbolt.Tx)error{return tx.Bucket([]byte("events")).ForEach(func(_,v []byte)error{var x model.Event;if e:=decode(v,&x);e!=nil{return e};if recordID==""||x.RecordID==recordID{out=append(out,x)};return nil})});return out,e}
func (s *Store) PutAudit(a model.Audit) error{return s.put("audits",a.ID,a)}
