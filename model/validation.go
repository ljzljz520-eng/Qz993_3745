package model

import "fmt"

func ValidateRecord(r Record) error { if r.ID=="" { return fmt.Errorf("record id required") }; if r.Name=="" { return fmt.Errorf("record name required") }; if r.Quantity<0 { return fmt.Errorf("quantity cannot be negative") }; return nil }
func NormalizeStatus(status string) string { switch status { case "received","processing","available","archived","reserved": return status; default: return "received" } }
func AllowedTransition(from,to string) bool { if from==to { return true }; switch from { case "received": return to=="processing"||to=="archived"; case "processing": return to=="available"||to=="received"; case "available": return to=="reserved"||to=="archived"; case "reserved": return to=="available"||to=="archived"; case "archived": return false }; return false }
