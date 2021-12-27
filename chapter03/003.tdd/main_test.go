package main

import "testing"

func TestCase1Part1(t *testing.T) {
	inputRecord("xiao", 0.38)
	inputRecord("xiao", 0.32)
	{
		randOfX, fatRateOfX := getRand("xiao")
		if randOfX != 1 {
			t.Fatalf("预期 xiao 第一，但是得到的是：%d", randOfX)
		}
		if fatRateOfX != 0.32 {
			t.Fatalf("预期 xiao 的体脂是 0.32，但得到的是：%f", fatRateOfX)
		}
	}
}

func TestCase1(t *testing.T) {
	inputRecord("xiao", 0.38)
	inputRecord("xiao", 0.32)
	{
		randOfX, fatRateOfX := getRand("xiao")
		if randOfX != 1 {
			t.Fatalf("预期 xiao 第一，但是得到的是：%d", randOfX)
		}
		if fatRateOfX != 0.32 {
			t.Fatalf("预期 xiao 的体脂是 0.32，但得到的是：%f", fatRateOfX)
		}
	}
	inputRecord("li", 0.28)
	{

		randOfX, fatRateOfX := getRand("xiao")
		if randOfX != 2 {
			t.Fatalf("预期 xiao 第2，但得到的是：%d", randOfX)
		}
		if fatRateOfX != 0.32 {
			t.Fatalf("预期 xiao 的体脂是0.32，但得到的是：%f", fatRateOfX)
		}

	}
	{
		randOfL, fatRateOfL := getRand("li")
		if randOfL != 1 {
			t.Fatalf("预期 li 第一，但是得到的是: %d", randOfL)
		}
		if fatRateOfL != 0.28 {
			t.Fatalf("预期 li 的体脂是0.28，但是得到的是：%f", fatRateOfL)
		}
	}
}

func TestCase2(t *testing.T) {
	inputRecord("xiao", 0.38)
	inputRecord("zhang", 0.38)
	inputRecord("li", 0.28)

	{
		randOfL, fatRateOfL := getRand("li")
		if randOfL != 1 {
			t.Fatalf("预期 li 是第一，但是得到的是：%d", randOfL)
		}
		if fatRateOfL != 0.28 {
			t.Fatalf("预期 li 的体脂是0.28，但是得到的是：%f", fatRateOfL)
		}
	}

	{
		randOfX, fatRateOfX := getRand("xiao")
		if randOfX != 2 {
			t.Fatalf("预期 xiao 第2，但得到是：%d", randOfX)
		}
		if fatRateOfX != 0.38 {
			t.Fatalf("预期 xiao 的体脂是0.32，但是得到的是：%f", fatRateOfX)
		}
	}

	{
		randOfZ, fatRateOfZ := getRand("zhang")
		if randOfZ != 2 {
			t.Fatalf("预期 zhang 第2，但是得到的是：%d", randOfZ)
		}
		if fatRateOfZ != 0.38 {
			t.Fatalf("预期 zhang 的体脂是0.38，但得到的是：%f", fatRateOfZ)
		}
	}
}

func TestCase3(t *testing.T) {
	inputRecord("xiao", 0.38)
	inputRecord("li", 0.28)
	inputRecord("zhang")

	{
		randOfL, fatRateOfL := getRand("li")
		if randOfL != 1 {
			t.Fatalf("预期 li 第1，但是得到的是：%d", randOfL)
		}
		if fatRateOfL != 0.28 {
			t.Fatalf("预期 li 的体脂是0.28，但得到的是：%f", fatRateOfL)
		}
	}

	{
		randOfX, fatRateOfX := getRand("xiao")
		if randOfX != 2 {
			t.Fatalf("预期 xiao 第2，但是得到的是：%d", randOfX)
		}
		if fatRateOfX != 0.38 {
			t.Fatalf("预期 xiao 的体脂是 0.38，但得到的是： %f", fatRateOfX)
		}
	}

	{
		randOfZ, _ := getRand("zhang")
		if randOfZ != 3 {
			t.Fatalf("预期 zhang 第3，但是得到的是：%d", randOfZ)
		}
	}
}
