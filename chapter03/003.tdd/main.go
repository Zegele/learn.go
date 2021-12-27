package main

import (
	"math"
	"sort"
)

var (
	personFatRate = map[string]float64{}
)

func inputRecord(name string, fatRate ...float64) { //一个人，可能有多个体脂率，要取最小的。
	minFatRate := math.MaxFloat64 //最大的数？
	for _, item := range fatRate {
		if minFatRate > item {
			minFatRate = item
		}
	}
	personFatRate[name] = minFatRate
}

func getRand(name string) (rand int, fatRate float64) {
	fatRate2PersonMap := map[float64][]string{}       //同一个体脂率，可能有多个人
	randArr := make([]float64, 0, len(personFatRate)) //对体脂率进行排名 ，需要有体脂率，和人数。
	for nameItem, frItem := range personFatRate {
		fatRate2PersonMap[frItem] = append(fatRate2PersonMap[frItem], nameItem)
		randArr = append(randArr, frItem)
	}
	sort.Float64s(randArr) //从大到小排序了？
	for i, frItem := range randArr {
		_nameS := fatRate2PersonMap[frItem]
		for _, _name := range _nameS {
			if _name == name {
				rand = i + 1
				fatRate = frItem
				return
			}
		}
	}

	return 0, 0
}
