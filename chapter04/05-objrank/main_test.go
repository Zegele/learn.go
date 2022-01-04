package main

import "testing"

func TestCase1Part1(t *testing.T) {
	r := &FatRateRank{}
	r.inputRecord("小强", 0.38)
	r.inputRecord("小强", 0.32)
	{
		randOfXQ, fatRateOfXQ := r.getRank("小强")
		if randOfXQ != 1 {
			t.Fatalf("预期 小强 第一，但是得到的是：%d", randOfXQ)
		}
		if fatRateOfXQ != 0.32 {
			t.Fatalf("预期 小强 的体制是 0.32，但得到的是：%f", fatRateOfXQ)
		}
	}
}

func TestCase1(t *testing.T) {
	r := &FatRateRank{}
	r.inputRecord("小强", 0.38)
	r.inputRecord("小强", 0.32)
	{
		randOfXQ, fatRateOfXQ := r.getRank("小强")
		if randOfXQ != 1 {
			t.Fatalf("预期 小强 第一，但是得到的是：%d", randOfXQ)
		}
		if fatRateOfXQ != 0.32 {
			t.Fatalf("预期 小强 的体脂是 0.32，但得到的是：%f", fatRateOfXQ)
		}
	}
	r.inputRecord("李静", 0.28)
	{
		randOfXQ, fatRateOfXQ := r.getRank("小强")
		if randOfXQ != 2 {
			t.Fatalf("预期 小强 第2，但是得到的是：%d", randOfXQ)
		}
		if fatRateOfXQ != 0.32 {
			t.Fatalf("预期 小强 的体脂是0.32，但得到的是：%f", fatRateOfXQ)
		}
	}
	{
		randOfLJ, fatRateOfLJ := r.getRank("李静")
		if randOfLJ != 1 {
			t.Fatalf("预期 李静 第一，但是得到的是：%d", randOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期 李静 的体脂是0.28，但得到的是：%f", fatRateOfLJ)
		}
	}
}

func TestCase2(t *testing.T) {
	r := &FatRateRank{}
	r.inputRecord("小强", 0.38)
	r.inputRecord("张伟", 0.38)
	r.inputRecord("李静", 0.28)

	{
		randOfLJ, fatRateOfLJ := r.getRank("李静")
		if randOfLJ != 1 {
			t.Fatalf("预期 李静 第一，但是得到的是：%d", randOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期 李静 的体脂是0.28，但得到的是：%f", fatRateOfLJ)
		}
	}
	{
		randOfXQ, fatRateOfXQ := r.getRank("小强")
		if randOfXQ != 2 {
			t.Fatalf("预期 小强 第2，但是得到的是：%d", randOfXQ)
		}
		if fatRateOfXQ != 0.38 {
			t.Fatalf("预期 小强 的体脂是 0.38，但得到的是：%f", fatRateOfXQ)
		}
	}
	{
		randOfZW, fatRateOfZW := r.getRank("张伟")
		if randOfZW != 2 {
			t.Fatalf("预期 张伟 第2，但是得到的是：%d", randOfZW)
		}
		if fatRateOfZW != 0.38 {
			t.Fatalf("预期 张伟 的体脂是0.38，但得到的是：%f", fatRateOfZW)
		}
	}
}

func TestCase3(t *testing.T) {
	r := &FatRateRank{}
	r.inputRecord("小强", 0.38)
	r.inputRecord("李静", 0.28)
	r.inputRecord("张伟")
	{
		randOfLJ, fatRateOfLJ := r.getRank("李静")
		if randOfLJ != 1 {
			t.Fatalf("预期 李静 第一，但是得到的是：%d", randOfLJ)
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期 李静 的体脂是0.28， 但得到的是：%f", fatRateOfLJ)
		}
	}
	{
		randOfXQ, fatRateOfXQ := r.getRank("小强")
		if randOfXQ != 2 {
			t.Fatalf("预期 小强 第二， 但是得到的是：%d", randOfXQ)
		}
		if fatRateOfXQ != 0.38 {
			t.Fatalf("预期 小强 的体脂是 0.38，但得到的是：%f", fatRateOfXQ)
		}
	}
	{
		randOfZW, _ := r.getRank("张伟")
		if randOfZW != 3 {
			t.Fatalf("预期 张伟 第三，但得到的是：%d", randOfZW)
		}
	}
}
