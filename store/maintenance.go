package store

import("fmt";"artstock/model";"go.etcd.io/bbolt")
func (s *Store) Count(bucket string)(int,error){s.mu.RLock();defer s.mu.RUnlock();if s.db==nil{return 0,fmt.Errorf("store closed")};n:=0;e:=s.db.View(func(tx *bbolt.Tx)error{b:=tx.Bucket([]byte(bucket));if b==nil{return bbolt.ErrBucketNotFound};return b.ForEach(func(_,v []byte)error{if v!=nil{n++};return nil})});return n,e}
func (s *Store) Snapshot() (map[string]int,error){out:=map[string]int{};for name:=range buckets{n,e:=s.Count(name);if e!=nil{return nil,e};out[name]=n};return out,nil}
func (s *Store) ReplaceRecords(rs []model.Record) error{s.mu.RLock();defer s.mu.RUnlock();if s.db==nil{return fmt.Errorf("store closed")};return s.db.Update(func(tx *bbolt.Tx)error{b:=tx.Bucket([]byte("records"));if e:=b.ForEach(func(k,_ []byte)error{return b.Delete(k)});e!=nil{return e};for _,r:=range rs{raw,e:=encode(r);if e!=nil{return e};if e=b.Put([]byte(r.ID),raw);e!=nil{return e}};return nil})}
func (s *Store) Health() bool{n,e:=s.Count("records");return e==nil&&n>=0}
