package store

import (
 "encoding/json"
 "fmt"
 "path/filepath"
 "sync"
 "go.etcd.io/bbolt"
 "artstock/model"
)
var buckets = map[string][]byte{"records":[]byte("records"),"users":[]byte("users"),"events":[]byte("events"),"audits":[]byte("audits")}
type Store struct { db *bbolt.DB; mu sync.RWMutex; path string }
func Open(path string) (*Store,error) { if path=="" { return nil,fmt.Errorf("path required") }; db,err:=bbolt.Open(filepath.Clean(path),0600,nil); if err!=nil{return nil,err}; s:=&Store{db:db,path:path}; err=db.Update(func(tx *bbolt.Tx) error { for _,b:=range buckets { if _,e:=tx.CreateBucketIfNotExists(b);e!=nil{return e} }; return nil }); if err!=nil { db.Close(); return nil,err }; return s,nil }
func (s *Store) Close() error { s.mu.Lock(); defer s.mu.Unlock(); if s.db==nil{return nil}; err:=s.db.Close(); s.db=nil; return err }
func (s *Store) Path() string { return s.path }
func encode(v any)([]byte,error){return json.Marshal(v)}
func decode(data []byte,v any)error{return json.Unmarshal(data,v)}
func (s *Store) put(bucket,key string,v any) error { raw,e:=encode(v); if e!=nil{return e}; s.mu.RLock(); defer s.mu.RUnlock(); if s.db==nil{return fmt.Errorf("store closed")}; return s.db.Update(func(tx *bbolt.Tx)error{return tx.Bucket([]byte(bucket)).Put([]byte(key),raw)}) }
func (s *Store) get(bucket,key string,v any) error { s.mu.RLock(); defer s.mu.RUnlock(); if s.db==nil{return fmt.Errorf("store closed")}; return s.db.View(func(tx *bbolt.Tx)error{d:=tx.Bucket([]byte(bucket)).Get([]byte(key)); if d==nil{return bbolt.ErrBucketNotFound}; return decode(append([]byte(nil),d...),v)}) }
func (s *Store) delete(bucket,key string) error { s.mu.RLock(); defer s.mu.RUnlock(); if s.db==nil{return fmt.Errorf("store closed")}; return s.db.Update(func(tx *bbolt.Tx)error{return tx.Bucket([]byte(bucket)).Delete([]byte(key))}) }
func (s *Store) PutRecord(r model.Record) error{return s.put("records",r.ID,r)}
func (s *Store) GetRecord(id string)(model.Record,error){var r model.Record; e:=s.get("records",id,&r); return r,e}
func (s *Store) DeleteRecord(id string) error{return s.delete("records",id)}
