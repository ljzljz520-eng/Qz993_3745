package store
import("path/filepath";"testing";"artstock/model")
func TestStoreRoundTrip(t *testing.T){s,e:=Open(filepath.Join(t.TempDir(),"a.db"));if e!=nil{t.Fatal(e)};defer s.Close();r:=model.NewRecord("r","brush","brush",3);if e=s.PutRecord(r);e!=nil{t.Fatal(e)};if _,e=s.GetRecord("r");e!=nil{t.Fatal(e)}}
