package inventory
import("path/filepath";"testing";"artstock/store";"artstock/model")
func TestRegistrationAndSearch(t *testing.T){s,_:=store.Open(filepath.Join(t.TempDir(),"i.db"));defer s.Close();i:=New(s);if e:=i.Register(model.NewRecord("r","red","paint",2));e!=nil{t.Fatal(e)};rs,_,e:=i.Search("red");if e!=nil||len(rs)!=1{t.Fatal(e)}}
