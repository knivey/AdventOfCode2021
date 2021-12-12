package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("\n\n### Part1")
	part1()
	fmt.Println("\n\n### Part2")
	part2()
}

const filename = "day1.txt"

func part1() {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var incs uint = 0
	var last int
	var first = true

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		current, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if !first && current > last {
			incs++
		}
		first = false
		last = current
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %v Increases\n", incs)
}

func arraySum(a []int) (out int) {
	for _, i := range a {
		out += i
	}
	return out
}

func part2() {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var incs uint = 0
	var lasts [3]int
	var line = 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		current, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		line++
		if line > 3 {
			if (current + arraySum(lasts[1:])) > arraySum(lasts[:]) {
				incs++
			}

			for k, v := range lasts {
				if k > 0 {
					lasts[k-1] = v
				}
			}
			lasts[len(lasts)-1] = current
		} else {
			lasts[line-1] = current
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %v Increases\n", incs)
}
