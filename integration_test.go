package main
import("path/filepath";"testing";"artstock/store";"artstock/inventory";"artstock/model")
func TestRecordFlow25(t *testing.T){s,_:=store.Open(filepath.Join(t.TempDir(),"bug.db"));defer s.Close();i:=inventory.New(s);if e:=i.Register(model.NewRecord("25","watercolor","paint",5));e!=nil{t.Fatal(e)};if e:=i.Review("25","reviewer");e!=nil{t.Fatal(e)};if e:=i.Complete("25","reviewer");e!=nil{t.Fatal(e)};r,_:=i.Get("25");if r.Status!="available"{t.Fatalf("expected real-time available status, got %s",r.Status)}}
