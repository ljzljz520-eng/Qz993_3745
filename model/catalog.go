package model

type Category struct { Name string; Description string; AllowedStatuses []string; MinQuantity int }
type Location struct { Code string; Room string; Shelf string; Capacity int; Occupied int }
type Reservation struct { ID string; RecordID string; UserID string; Quantity int; State string }
func DefaultCategories() []Category{return []Category{{Name:"paint",Description:"pigments and tubes",AllowedStatuses:[]string{"received","processing","available","reserved","archived"},MinQuantity:0},{Name:"brush",Description:"brushes",AllowedStatuses:[]string{"received","available","reserved","archived"},MinQuantity:0},{Name:"paper",Description:"paper stock",AllowedStatuses:[]string{"received","processing","available","archived"},MinQuantity:0},{Name:"canvas",Description:"canvas",AllowedStatuses:[]string{"received","processing","available","reserved","archived"},MinQuantity:0}}}
func (c Category) Accepts(status string) bool{for _,x:=range c.AllowedStatuses{if x==status{return true}};return false}
func (l Location) Free() int{if l.Capacity<=l.Occupied{return 0};return l.Capacity-l.Occupied}
func (l *Location) Assign(n int) bool{if n<0||n>l.Free(){return false};l.Occupied+=n;return true}
func (l *Location) Release(n int) bool{if n<0||n>l.Occupied{return false};l.Occupied-=n;return true}
func NewReservation(id,record,user string,q int) Reservation{return Reservation{ID:id,RecordID:record,UserID:user,Quantity:q,State:"pending"}}
func (r *Reservation) Confirm() bool{if r.State!="pending"||r.Quantity<1{return false};r.State="confirmed";return true}
func (r *Reservation) Cancel() bool{if r.State=="cancelled"||r.State=="fulfilled"{return false};r.State="cancelled";return true}
