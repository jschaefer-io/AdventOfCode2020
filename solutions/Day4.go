package solutions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func newPassport(batch []string) passport {
	pass := passport{}
	for _, s := range batch {
		kV := strings.Split(s, ":")
		kV[1] = strings.TrimSpace(kV[1])
		switch kV[0] {
		case "byr":
			pass.byr = kV[1]
			break
		case "iyr":
			pass.iyr = kV[1]
			break
		case "eyr":
			pass.eyr = kV[1]
			break
		case "hgt":
			pass.hgt = kV[1]
			break
		case "hcl":
			pass.hcl = kV[1]
			break
		case "ecl":
			pass.ecl = kV[1]
			break
		case "pid":
			pass.pid = kV[1]
			break
		case "cid":
			pass.cid = kV[1]
			break
		}
	}
	return pass
}

func (p passport) validateRequired() bool {
	valid := true
	valid = valid && p.byr != ""
	valid = valid && p.iyr != ""
	valid = valid && p.eyr != ""
	valid = valid && p.hgt != ""
	valid = valid && p.hcl != ""
	valid = valid && p.ecl != ""
	valid = valid && p.pid != ""
	return valid
}

func (p passport) validate() bool {
	if !p.validateRequired() {
		return false
	}
	var r *regexp.Regexp
	valid := true

	// Validate all dates
	checkYear := func(s string, min int, max int) bool {
		if len(s) != 4 {
			return false
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		return n >= min && n <= max
	}
	valid = valid && checkYear(p.byr, 1920, 2002)
	valid = valid && checkYear(p.iyr, 2010, 2020)
	valid = valid && checkYear(p.eyr, 2020, 2030)

	// Validate Height
	r = regexp.MustCompile("^(\\d+)(cm|in)$")
	valid = valid && (func() bool {
		res := r.FindStringSubmatch(p.hgt)
		if len(res) != 3 {
			return false
		}
		height, err := strconv.Atoi(res[1])
		if err != nil {
			return false
		}
		switch res[2] {
		case "cm":
			return height >= 150 && height <= 193
		case "in":
			return height >= 59 && height <= 76
		default:
			return false
		}
	})()

	// Validate Hair Color
	r = regexp.MustCompile("^#([a-f]|[0-9]){6}$")
	valid = valid && r.MatchString(p.hcl)

	// Validate Eye color
	valid = valid && (func() bool {
		switch p.ecl {
		case "amb",
			"blu",
			"brn",
			"gry",
			"grn",
			"hzl",
			"oth":
			return true
		}
		return false
	})()

	// Validate Passport ID
	r = regexp.MustCompile("^\\d{9}$")
	valid = valid && r.MatchString(p.pid)

	return valid
}

type Day4 struct {
	passports []passport
}

func (d *Day4) init(s string) {
	splitR := regexp.MustCompile("((|\\r)\\n){2}")
	evalR := regexp.MustCompile("\\w+:(?:#|\\d|\\w)+")
	batches := splitR.Split(s, -1)
	for _, batch := range batches {
		res := evalR.FindAllString(batch, -1)
		d.passports = append(d.passports, newPassport(res))
	}
}

func (d *Day4) executeA() int {
	count := 0
	for _, passport := range d.passports {
		if passport.validateRequired() {
			count++
		}
	}
	return count
}

func (d *Day4) executeB() int {
	count := 0
	for _, passport := range d.passports {
		if passport.validate() {
			count++
		}
	}
	return count
}

func (d *Day4) Handle(s string) ([]string, error) {
	d.init(s)

	results := make([]string, 0)
	results = append(results, fmt.Sprintf("%d", d.executeA()))
	results = append(results, fmt.Sprintf("%d", d.executeB()))

	return results, nil
}
