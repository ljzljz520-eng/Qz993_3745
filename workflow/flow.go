package workflow

import("fmt";"artstock/inventory";"artstock/model")
type Engine struct { Inventory *inventory.Service }
func New(i *inventory.Service)*Engine{return &Engine{Inventory:i}}
func (e *Engine) Receive(id,name,category string,qty int) error { return e.Inventory.Register(model.NewRecord(id,name,category,qty)) }
func (e *Engine) Process(id,actor string) error {if err:=e.Inventory.Review(id,actor);err!=nil{return err};return e.Inventory.Complete(id,actor)}
func (e *Engine) Retire(id,actor string) error{return e.Inventory.Archive(id,actor)}
func (e *Engine) FullCycle(id,name,category string,qty int,actor string) error {if err:=e.Receive(id,name,category,qty);err!=nil{return err};if err:=e.Process(id,actor);err!=nil{return err};return nil}
func (e *Engine) Require(id string) (model.Record,error){r,err:=e.Inventory.Get(id);if err!=nil{return r,err};if !r.Valid(){return r,fmt.Errorf("invalid record")};return r,nil}
