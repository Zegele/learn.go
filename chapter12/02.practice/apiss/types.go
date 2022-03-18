package apiss

type PersonalInfomation struct {
	Name   string  `json:"name"`
	Sex    string  `json:"sex"`
	Tall   float64 `json:"tall"`
	Weight float64 `json:"weight"`
	Age    int     `json:"age"`
}

type PersonalInfomationFatRate struct {
	Name    string
	FatRate float64
}

type PersonalRank struct {
	Name       string
	Sex        string
	RankNumber int
	FatRate    float64
}
