package rank

import (
	gobmi "github.com/armstrongli/go-bmi"
	"learn.go/chapter12/apiss"
	"learn.go/chapter12/frinterface"
	"log"
	"math"
	"sort"
	"sync"
)

var _ frinterface.ServeInterface = &FatRateRank{}

type RankItem struct {
	Name    string
	Sex     string
	FatRate float64
}

type FatRateRank struct {
	items     []RankItem
	itemsLock sync.Mutex
}

func (r *FatRateRank) RegisterPersonalInformation(pi *apiss.PersonalInformation) error {
	bmi, err := gobmi.BMI(float64(pi.Weight), float64(pi.Tall))
	if err != nil {
		log.Println("计算BMI失败：", err)
		return err
	}
	fr := gobmi.CalcFatRate(bmi*100, pi.Age, pi.Sex)
	r.InputRecord(pi.Name, pi.Sex, fr)
	return nil
}

func (r *FatRateRank) UpdatePersonalInformation(pi *apiss.PersonalInformation) (*apiss.PersonalInformationFatRate, error) {
	bmi, err := gobmi.BMI(float64(pi.Weight), float64(pi.Tall))
	if err != nil {
		log.Println("计算BMI失败：", err)
		return nil, err
	}
	fr := gobmi.CalcFatRate(bmi*100, pi.Age, pi.Sex)
	r.InputRecord(pi.Name, pi.Sex, fr)
	//rankID, fr := r.GetRank(pi.Name)
	return &apiss.PersonalInformationFatRate{
		Name:    pi.Name,
		FatRate: fr,
	}, nil
}

func (r *FatRateRank) GetFatRate(name string) (*apiss.PersonalRank, error) {
	rankID, sex, fr := r.getRank(name)
	return &apiss.PersonalRank{
		Name:       name,
		Sex:        sex,
		RankNumber: rankID,
		FatRate:    fr,
	}, nil
}

func (r *FatRateRank) GetTop() ([]*apiss.PersonalRank, error) {
	return r.getRankTop(), nil
}

func NewFatRateRank() *FatRateRank { //实例化函数 要学会这样做
	return &FatRateRank{items: make([]RankItem, 0, 100)}
}

func (r *FatRateRank) InputRecord(name, sex string, fatRate ...float64) {
	r.itemsLock.Lock()
	defer r.itemsLock.Unlock()
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
			Sex:     sex,
			FatRate: minFatRate,
		})
	}
}

func (r *FatRateRank) getRank(name string) (rank int, sex string, fatRate float64) {
	r.itemsLock.Lock()
	defer r.itemsLock.Unlock()

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

func (r *FatRateRank) getRankTop() []*apiss.PersonalRank {
	r.itemsLock.Lock() // 加锁
	defer r.itemsLock.Unlock()

	sort.Slice(r.items, func(i, j int) bool {
		return r.items[i].FatRate < r.items[j].FatRate //这时排序？
	})
	out := make([]*apiss.PersonalRank, 0, 10)
	for i := 0; i < 10 && i < len(r.items); i++ {
		item := r.items[i]
		out = append(out, &apiss.PersonalRank{
			Name:       item.Name,
			Sex:        item.Sex,
			RankNumber: i + 1,
			FatRate:    item.FatRate,
		})
	}
	return out
}
