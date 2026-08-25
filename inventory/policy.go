package inventory

import("fmt";"artstock/model")
type Policy struct{MaxQuantity int; RequireReviewer bool; AllowedCategories map[string]bool}
func DefaultPolicy()Policy{return Policy{MaxQuantity:10000,RequireReviewer:true,AllowedCategories:map[string]bool{"paint":true,"brush":true,"paper":true,"canvas":true}}}
func (p Policy) Check(r model.Record)error{if r.Quantity>p.MaxQuantity{return fmt.Errorf("quantity exceeds policy")};if len(p.AllowedCategories)>0&&!p.AllowedCategories[r.Category]{return fmt.Errorf("category denied")};return nil}
func (s *Service) ApplyPolicy(r model.Record,p Policy)error{return p.Check(r)}
func (s *Service) CanArchive(r model.Record)bool{return r.Status=="available"||r.Status=="reserved"||r.Status=="received"}
func (s *Service) CanEdit(r model.Record)bool{return r.Status!="archived"}
