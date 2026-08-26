package workflow
import("path/filepath";"testing";"artstock/store";"artstock/inventory")
func TestWorkflowOne(t *testing.T){s,_:=store.Open(filepath.Join(t.TempDir(),"w.db"));defer s.Close();e:=New(inventory.New(s));if err:=e.Receive("one","blue","paint",1);err!=nil{t.Fatal(err)};r,_:=e.Require("one");if r.Status!="received"{t.Fatal(r.Status)}}
func TestWorkflowTwo(t *testing.T){s,_:=store.Open(filepath.Join(t.TempDir(),"w.db"));defer s.Close();e:=New(inventory.New(s));if err:=e.FullCycle("two","paper","paper",2,"u");err!=nil{t.Fatal(err)};if err:=e.Retire("two","u");err!=nil{t.Fatal(err)}}
func TestWorkflowThree(t *testing.T){s,_:=store.Open(filepath.Join(t.TempDir(),"w.db"));defer s.Close();e:=New(inventory.New(s));if err:=e.Receive("three","canvas","canvas",2);err!=nil{t.Fatal(err)};n,err:=Notify(e.Inventory,"three","u");if err!=nil||!n.Delivered{t.Fatal(err)}}
