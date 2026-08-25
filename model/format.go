package model

import("fmt";"strings";"time")
func FormatRecord(r Record) string{return fmt.Sprintf("%s [%s] %s x%d",r.ID,r.Status,r.Name,r.Quantity)}
func Slug(value string) string{return strings.ToLower(strings.ReplaceAll(strings.TrimSpace(value)," ","-"))}
func Fresh(r Record, now time.Time, age time.Duration) bool{return !r.UpdatedAt.IsZero()&&now.Sub(r.UpdatedAt)<=age}
func StatusRank(status string) int{switch status{case "received":return 1;case "processing":return 2;case "available":return 3;case "reserved":return 4;case "archived":return 5};return 0}
func CompareStatus(a,b string) int{ra,rb:=StatusRank(a),StatusRank(b);if ra<rb{return -1};if ra>rb{return 1};return 0}
