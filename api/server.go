package api

import("encoding/json";"net/http";"artstock/inventory";"artstock/report")
type Server struct { Inventory *inventory.Service }
func New(i *inventory.Service)*Server{return &Server{Inventory:i}}
func (s *Server) Handler() http.Handler {m:=http.NewServeMux();m.HandleFunc("/health",s.health);m.HandleFunc("/records",s.records);return m}
func (s *Server) health(w http.ResponseWriter,r *http.Request){w.WriteHeader(http.StatusOK);w.Write([]byte("ok"))}
func (s *Server) records(w http.ResponseWriter,r *http.Request){rs,summary,e:=s.Inventory.Search(r.URL.Query().Get("q"));if e!=nil{http.Error(w,e.Error(),500);return};if r.Method==http.MethodPost{http.Error(w,"use workflow",405);return};json.NewEncoder(w).Encode(report.Build(rs,summary.Label()))}
func WriteJSON(w http.ResponseWriter,v any){w.Header().Set("Content-Type","application/json");json.NewEncoder(w).Encode(v)}
