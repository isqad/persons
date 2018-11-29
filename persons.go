package persons

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type FakePerson struct {
	name     string
	surname  string
	sname    string
	birthday time.Time
}

func GetPerson() FakePerson {
	return FakePerson{randLine(`./names_ru`), randLine(`./families_ru`), randLine(`./snames_ru`), randate()}
}

func randLine(path string) string {
	items, err := readLines(path)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return items[rand.Intn(len(items)-1)]
}

func randate() time.Time {
	min := time.Date(1975, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(1989, 1, 0, 0, 0, 0, 0, time.UTC).Unix()

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
