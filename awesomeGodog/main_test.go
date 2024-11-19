package main

import (
	"fmt"
	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"testing"
)

var number, result_fac int
var result_expansion string
var boolValue bool

func get_number(arg int) error {
	number = arg
	return nil
}

func get_expansion(arg string) error {
	var err error
	boolValue, err = strconv.ParseBool(arg)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func get_factorial() error {
	result_fac = factorial(number)
	fmt.Printf("getfactorial %v\n", result_fac)
	return nil
}

func result_req(expected int) error {
	assert.Equal(nil, expected, result_fac)
	return nil
}

func result_exp(expected string) error {
	result_expansion = formatNiceFactorial(number, boolValue)
	assert.Equal(nil, expected, result_expansion)
	return nil
}

func initialize_scenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^the number is (\d+)$`, get_number)
	ctx.Step(`^the expansion is "(.*?)"$`, get_expansion)
	ctx.Step(`^I calculate the factorial$`, get_factorial)
	ctx.Step(`^the result should be (\d+)$`, result_req)
	ctx.Step(`^the result should be shown like "(.*?)"$`, result_exp)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: initialize_scenario,
		Options: &godog.Options{
			Format: "pretty",
			Paths:  []string{"features/main.feature"},
		},
	}

	if suite.Run() != 0 {
		t.Fatal("Test failed")
	}
}
