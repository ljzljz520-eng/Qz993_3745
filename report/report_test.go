package report
import("testing";"artstock/model")
func TestReportExport(t *testing.T){r:=model.NewRecord("r","ink","paint",2);x:=Build([]model.Record{r},"now");if !x.Ready(){t.Fatal("not ready")};if CSV([]model.Record{r})==""{t.Fatal("csv")}}
