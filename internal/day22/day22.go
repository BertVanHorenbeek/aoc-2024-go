package day22

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type pricing struct {
	price int
	diff  int
}

type priceAtSequence struct {
	price    int
	sequence string
}

func CalculateSecretNumberSum(input string, n int) int {
	lines := strings.Split(input, "\n")
	startNumbers := lo.Map(lines, func(item string, _ int) int {
		atoi, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		return atoi
	})
	return lo.Sum(lo.Map(startNumbers, func(item int, _ int) int {
		return calculateSecretNumberIterations(item, n)
	}))
}

func CalculateMostBananasYouCanEarn(input string) int {
	lines := strings.Split(input, "\n")
	startNumbers := lo.Map(lines, func(item string, _ int) int {
		atoi, err := strconv.Atoi(item)
		if err != nil {
			panic(err)
		}
		return atoi
	})
	prices := lo.Map(startNumbers, func(item int, _ int) []pricing {
		return readPricing(item)
	})
	priceAtSeq := lo.FlatMap(prices, func(item []pricing, _ int) []priceAtSequence {
		sequences := make([]priceAtSequence, len(item))
		for i := 3; i < len(item); i++ {
			sequence := fmt.Sprintf("%d,%d,%d,%d", item[i-3].diff, item[i-2].diff, item[i-1].diff, item[i].diff)
			if !lo.ContainsBy(sequences, func(item priceAtSequence) bool {
				return item.sequence == sequence
			}) {
				sequences = append(sequences, priceAtSequence{
					price:    item[i].price,
					sequence: sequence,
				})
			}
		}
		return sequences
	})
	grouped := lo.GroupBy(priceAtSeq, func(item priceAtSequence) string {
		return item.sequence
	})
	summed := lo.MapValues(grouped, func(value []priceAtSequence, key string) int {
		return lo.SumBy(value, func(item priceAtSequence) int {
			return item.price
		})
	})
	return lo.Max(lo.Values(summed))
}

func readPricing(number int) []pricing {
	prices := make([]pricing, 2000)
	for i := 0; i < 2000; i++ {
		next := calculateNext(number)
		diff := next%10 - number%10
		number = next
		prices[i] = pricing{number % 10, diff}
	}
	return prices
}

func calculateSecretNumberIterations(number int, n int) int {
	for i := 0; i < n; i++ {
		number = calculateNext(number)
	}
	return number
}

func calculateNext(number int) int {
	step1 := prune(mix(number*64, number))
	step2 := prune(mix(step1/32, step1))
	step3 := prune(mix(step2*2048, step2))
	return step3
}

func mix(n1, n2 int) int {
	return n1 ^ n2
}

func prune(n int) int {
	return n % 16777216
}
