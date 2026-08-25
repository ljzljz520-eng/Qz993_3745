package report

import("encoding/json";"artstock/model";"artstock/store")
type Report struct { Summary model.Summary; Records []model.Record; Generated string }
func Build(rs []model.Record,generated string) Report{return Report{Summary:store.Summarize(rs),Records:store.SortRecords(rs),Generated:generated}}
func (r Report) JSON()([]byte,error){return json.Marshal(r)}
func (r Report) Ready() bool{return r.Summary.Healthy()&&len(r.Records)==r.Summary.Total}
