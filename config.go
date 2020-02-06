package util

import (
	"bufio"
	"flag"
	"os"
	"strings"
)

//InitEnvironmentIfNeeded set env variables listed in file
func InitEnvironmentIfNeeded(flagName string) error {
	pathToEnv := flag.String(flagName, "", "Path to file with env values")
	flag.Parse()

	if len(*pathToEnv) == 0 {
		return nil
	}

	file, err := os.Open(*pathToEnv)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := strings.SplitN(scanner.Text(), "=", 2)
		if len(s) < 2 {
			//skip empty lines
			continue
		}

		_ = os.Setenv(s[0], s[1])
	}

	return scanner.Err()
}
