package persons

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type FakePerson struct {
	Name              string
	Surname           string
	Sname             string
	Birthday          time.Time
	PassportSerial    string
	PassportNumber    string
	PassportIssueDate time.Time
	PassportIssuer    string
}

func GetPerson() FakePerson {
	issuers := passportIssuers()
	return FakePerson{
		randLine(`./names_ru`),
		randLine(`./families_ru`),
		randLine(`./snames_ru`),
		randate(1975, 1989),
		randPassportSerial(),
		randPassportNumber(),
		randate(1994, 2010),
		issuers[rand.Intn(len(issuers)-1)]}
}

func randLine(path string) string {
	items, err := readLines(path)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return items[rand.Intn(len(items)-1)]
}

func randPassportNumber() string {
	number1 := [2]int{7, 8}
	number2 := [2]int{5, 6}
	number3 := [3]int{0, 1, 2}
	number4 := [4]int{2, 3, 6, 8}

	return fmt.Sprintf("51%d%d%d%d",
		number1[rand.Intn(len(number1)-1)],
		number2[rand.Intn(len(number2)-1)],
		number3[rand.Intn(len(number3)-1)],
		number4[rand.Intn(len(number4)-1)])
}

func randPassportSerial() string {
	serial1 := [2]int{0, 1}
	serial2 := [3]int{7, 8, 9}

	return fmt.Sprintf("65%d%d",
		serial1[rand.Intn(len(serial1)-1)],
		serial2[rand.Intn(len(serial2)-1)])
}

func randate(startY int, endY int) time.Time {
	min := time.Date(startY, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(endY, 1, 0, 0, 0, 0, 0, time.UTC).Unix()

	sec := rand.Int63n(max-min) + min
	return time.Unix(sec, 0)
}

func readLines(path string) ([]string, error) {
	var lines []string

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func passportIssuers() [4]string {
	issuers := [4]string{"ОУФМС России по Ханты-мансийскому автономному округу - Югре в Сургутском районе",
		"ОУФМС России по Нижегородской обл. в Ленинском р-не гор. Нижнего Новгорода",
		"Октябрьским районным отделом внутренних дел города Самары",
		"Отделом внутренних дел города Кулебаки Нижегородской области"}
	return issuers
}
