package api

import("encoding/json";"net/http";"artstock/model")
func DecodeRecord(r *http.Request)(model.Record,error){var x model.Record;e:=json.NewDecoder(r.Body).Decode(&x);return x,e}
func Error(w http.ResponseWriter,status int,err error){http.Error(w,err.Error(),status)}
func MethodAllowed(w http.ResponseWriter,method string,r *http.Request)bool{if r.Method!=method{w.Header().Set("Allow",method);Error(w,405,model.ErrUnavailable);return false};return true}
