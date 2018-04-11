package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItPrefixZerosAsNeeded(t *testing.T) {
	otp := "1234"
	assert.Equal(t, "001234", prefix0(otp))
}

func TestShouldNotPrefixZero(t *testing.T) {
	otp := "123456"
	assert.Equal(t, "123456", prefix0(otp))
}

func TestThatOTPGeneratedIsValid(t *testing.T) {
	secret := "dummySECRETdummy"
	interval := int64(50780342)
	otp := "971294"
	assert.Equal(t, otp, getHOTPToken(secret, interval))
}

func TestThatOTsNotGeneratedAndTestPanics(t *testing.T) {
	secret := "dummySECRETdumm$"
	interval := int64(50780342)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code panic")
		}
	}()
	getHOTPToken(secret, interval)
}
