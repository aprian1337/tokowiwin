package helper

import "tokowiwin/repositories/model"

func GetTransactionIDs(txs []*model.Transactions) []int64 {
	var ids = make([]int64, 0)
	for _, v := range txs {
		ids = append(ids, v.ID)
	}
	return ids
}
