package model

type Summary struct { Total int; Available int; Processing int; Archived int; Quantity int }
func (s *Summary) Add(r Record) { s.Total++; s.Quantity+=r.Quantity; switch r.Status { case "available","reserved": s.Available++; case "processing": s.Processing++; case "archived": s.Archived++ } }
func (s Summary) Healthy() bool { return s.Total>0 && s.Available+s.Processing+s.Archived<=s.Total }
func (s Summary) Label() string { if s.Total==0 { return "empty" }; if s.Available==s.Total { return "ready" }; if s.Processing>0 { return "in-progress" }; return "mixed" }
