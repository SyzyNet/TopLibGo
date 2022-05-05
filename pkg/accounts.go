package toplib

import (
	"bufio"
	"os"
	"strings"
)

type Combo struct {
	Username string
	Password string
}

type Hit struct {
	UserPass string
	Capture  []string
}

func

//LoadCombos is a function to load each line of a file into a map of usernames
//and passwords defined by the Combo struct
//Returns a map of usernames and passwords and an error
func LoadCombos(file string) (map[string]Combo, error) {
	combos := make(map[string]Combo)
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.Split(line, ":")
		combos[parts[0]] = Combo{Username: parts[0], Password: parts[1]}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return combos, err
}
