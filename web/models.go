package web

import "fmt"

type UserSummaryRequest struct {
	Ids []string `json:"ids"`
}

func (u UserSummaryRequest) Validate() error {
	if len(u.Ids) <= 0 {
		return fmt.Errorf("List is empty.")
	}

	return nil
}