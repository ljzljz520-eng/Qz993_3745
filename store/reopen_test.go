package store
import("path/filepath";"testing";"artstock/model")
func TestPersistenceSurvivesReopen(t *testing.T){p:=filepath.Join(t.TempDir(),"persist.db");s,e:=Open(p);if e!=nil{t.Fatal(e)};if e=s.PutRecord(model.NewRecord("persist","ink","paint",4));e!=nil{t.Fatal(e)};s.Close();s,e=Open(p);if e!=nil{t.Fatal(e)};defer s.Close();r,e:=s.GetRecord("persist");if e!=nil||r.Name!="ink"{t.Fatalf("reopen %v",e)}}
