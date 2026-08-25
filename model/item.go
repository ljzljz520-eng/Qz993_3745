package model

import "time"

type Record struct { ID string `json:"id"`; Name string `json:"name"`; Category string `json:"category"`; Quantity int `json:"quantity"`; Status string `json:"status"`; UpdatedAt time.Time `json:"updated_at"`; Notes string `json:"notes"` }
type User struct { ID string; Name string; Role string; Active bool }
type Event struct { ID string; RecordID string; Kind string; Actor string; At time.Time; Detail string }
type Audit struct { ID string; Action string; RecordID string; Actor string; At time.Time; Outcome string }

func NewRecord(id, name, category string, quantity int) Record { return Record{ID:id, Name:name, Category:category, Quantity:quantity, Status:"received", UpdatedAt:time.Now().UTC()} }
func (r Record) Valid() bool { return r.ID != "" && r.Name != "" && r.Category != "" && r.Quantity >= 0 }
func (r Record) Available() bool { return r.Status == "available" && r.Quantity > 0 }
func (r *Record) SetStatus(status string) { r.Status=status; r.UpdatedAt=time.Now().UTC() }
func (r Record) Clone() Record { return r }
func NewUser(id,name,role string) User { return User{ID:id,Name:name,Role:role,Active:true} }
func (u User) CanReview() bool { return u.Active && (u.Role=="manager" || u.Role=="reviewer") }
func NewEvent(id,recordID,kind,actor,detail string) Event { return Event{ID:id,RecordID:recordID,Kind:kind,Actor:actor,At:time.Now().UTC(),Detail:detail} }
func NewAudit(id,action,recordID,actor,outcome string) Audit { return Audit{ID:id,Action:action,RecordID:recordID,Actor:actor,At:time.Now().UTC(),Outcome:outcome} }
