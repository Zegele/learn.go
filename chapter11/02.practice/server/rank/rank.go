package rank

import (
	"math"
	"sort"
)

type RankItem struct {
	Name    string
	FatRate float64
}

type FatRateRank struct {
	items []RankItem
}

func NewFatRateRank() *FatRateRank { //实例化函数 要学会这样做
	return &FatRateRank{items: make([]RankItem, 0, 100)}
}

func (r *FatRateRank) InputRecord(name string, fatRate ...float64) {
	minFatRate := math.MaxFloat64 // math.MaxFloat64是一个常量，它是最大的浮点数
	for _, item := range fatRate {
		if minFatRate > item {
			minFatRate = item
		}
	}

	found := false
	for i, item := range r.items {
		if item.Name == name {
			if item.FatRate >= minFatRate {
				item.FatRate = minFatRate
			}
			r.items[i] = item
			found = true
			break
		}
	}

	if !found {
		r.items = append(r.items, RankItem{
			Name:    name,
			FatRate: minFatRate,
		})
	}
}

func (r *FatRateRank) GetRank(name string) (rank int, fatRate float64) {
	sort.Slice(r.items, func(i, j int) bool {
		return r.items[i].FatRate < r.items[j].FatRate //这时排序？
	})
	frs := map[float64]struct{}{}
	for _, item := range r.items {
		frs[item.FatRate] = struct{}{}
		if item.Name == name {
			fatRate = item.FatRate
		}
	}
	rankArr := make([]float64, 0, len(frs))
	for k := range frs {
		rankArr = append(rankArr, k)
	}
	sort.Float64s(rankArr)
	for i, frItem := range rankArr {
		if frItem == fatRate {
			rank = i + 1
			break
		}
	}
	return
}
