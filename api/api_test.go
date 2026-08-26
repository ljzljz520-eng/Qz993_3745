package api
import("net/http/httptest";"path/filepath";"testing";"artstock/store";"artstock/inventory")
func TestHealth(t *testing.T){s,_:=store.Open(filepath.Join(t.TempDir(),"a.db"));defer s.Close();h:=New(inventory.New(s)).Handler();r:=httptest.NewRecorder();h.ServeHTTP(r,httptest.NewRequest("GET","/health",nil));if r.Code!=200{t.Fatal(r.Code)}}
