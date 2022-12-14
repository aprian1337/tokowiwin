package helper

import "tokowiwin/repositories/model"

func GetProductIDs(carts []*model.Carts) []int64 {
	var ids = make([]int64, 0)
	for _, v := range carts {
		if v.ProductID.Valid {
			ids = append(ids, v.ProductID.Int64)
		}
	}
	return ids
}
