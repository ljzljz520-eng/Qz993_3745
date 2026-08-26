package model

import "strings"

func StatusLabel(status string) string{switch status{case "received":return "待接收";case "processing":return "处理中";case "available":return "可用";case "reserved":return "已预留";case "archived":return "已归档"};return "未知"}
func CategoryLabel(category string) string{switch strings.ToLower(category){case "paint":return "颜料";case "brush":return "画笔";case "paper":return "纸张";case "canvas":return "画布"};return category}
func IsConsumable(category string)bool{return category=="paint"||category=="paper"}
func DisplayFields(r Record)map[string]string{return map[string]string{"id":r.ID,"name":r.Name,"category":CategoryLabel(r.Category),"quantity":fmtInt(r.Quantity),"status":StatusLabel(r.Status)}}
func fmtInt(n int)string{if n==0{return "0"};negative:=n<0;if negative{n=-n};buf:="";for n>0{buf=string(rune('0'+n%10))+buf;n/=10};if negative{return "-"+buf};return buf}
