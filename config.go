package util

import (
	"bufio"
	"flag"
	"os"
	"strings"
)

//InitEnvironmentIfNeeded set env variables listed in file
func InitEnvironmentIfNeeded(flagName string) error {
	flagSet := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	pathToEnv := flagSet.String(flagName, "", "Path to file with env values")
	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return nil
	}

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
