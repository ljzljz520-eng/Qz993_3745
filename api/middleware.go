package api

import("log";"net/http";"time")
func RequestLog(next http.Handler)http.Handler{return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){start:=time.Now();next.ServeHTTP(w,r);log.Printf("%s %s %s",r.Method,r.URL.Path,time.Since(start))})}
func CORS(next http.Handler)http.Handler{return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){w.Header().Set("Access-Control-Allow-Origin","*");if r.Method==http.MethodOptions{w.WriteHeader(204);return};next.ServeHTTP(w,r)})}
func Chain(h http.Handler)http.Handler{return RequestLog(CORS(h))}
