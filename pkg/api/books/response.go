package books

import "github.com/felipear89/go-examples/pkg/model"

type ListResponse struct {
	Items []*model.Book `json:"items"`
}
