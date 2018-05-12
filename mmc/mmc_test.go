package mmc

import "testing"

func TestMMCOf2(t *testing.T) {
	result := MMC(2)
	if result != 2 {
		t.Error("mmc(2) should be 2 and we got", result)
	}
}

func TestMMCOf3(t *testing.T) {
	result := MMC(3)
	if result != 3 {
		t.Error("mmc(3) should be 3 and we got", result)
	}
}

func TestMMCOf2And3(t *testing.T) {
	result := MMC(2, 3)
	if result != 6 {
		t.Error("mmc(2, 3) should be 6 and we got", result)
	}
}

func TestMMCOf4And6(t *testing.T) {
	result := MMC(4, 6)
	if result != 12 {
		t.Error("mmc(4, 6) should be 12 and we got", result)
	}
}

func TestMMCOf2And6And8(t *testing.T) {
	result := MMC(2, 8, 6)
	if result != 24 {
		t.Error("mmc(2, 8, 6) should be 24 and we got", result)
	}
}

func TestMMCOf2And5And3(t *testing.T) {
	result := MMC(2, 5, 3)
	if result != 30 {
		t.Error("mmc(2, 5, 3) should be 30 and we got", result)
	}
}

func TestMMCOf2And10And7And9(t *testing.T) {
	result := MMC(2, 10, 7, 9)
	if result != 630 {
		t.Error("mmc(2, 10, 7, 9) should be 630 and we got", result)
	}
}

func TestMMCOf1And1(t *testing.T) {
	result := MMC(1, 1)
	if result != 0 {
		t.Error("mmc(1, 1) should be 0 and we got", result)
	}
}
