package report

import("bytes";"fmt";"artstock/model")
func CSV(rs []model.Record) string {var b bytes.Buffer;b.WriteString("id,name,category,quantity,status\n");for _,r:=range rs{fmt.Fprintf(&b,"%s,%s,%s,%d,%s\n",r.ID,r.Name,r.Category,r.Quantity,r.Status)};return b.String()}
func GroupByCategory(rs []model.Record) map[string][]model.Record{out:=map[string][]model.Record{};for _,r:=range rs{out[r.Category]=append(out[r.Category],r)};return out}
func Counts(rs []model.Record) map[string]int{out:=map[string]int{};for _,r:=range rs{out[r.Status]++};return out}
