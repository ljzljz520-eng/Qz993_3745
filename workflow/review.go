package workflow

import("fmt";"artstock/model")
func ValidateReviewer(u model.User) error {if !u.CanReview(){return fmt.Errorf("reviewer not authorized")};return nil}
func Decision(status string,approved bool) string {if approved{return "approved"};if status=="processing"{return "returned"};return "rejected"}
func StepsComplete(steps []string) bool {if len(steps)<4{return false};for _,s:=range steps{if s==""{return false}};return true}
