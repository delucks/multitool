// Naive implementation of a spell-checker for programs that are available to run in the bash shell.
// Works by computing 1 and 2 distance edits from the input word and comparing them against all available:
// - Shell executables
// - Aliases
// - Functions
package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"strings"
)

// Used to generate suggestions, should include all valid characters for CLI names
const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890._-"

func build_corpus() (map[string]bool, error) {
	// Get all available symbols that are available in my bash shell
	s := make(map[string]bool)
	process := exec.Command("bash", "-i", "-c", "compgen -A function -ac")
	stdout, err := process.StdoutPipe()
	if err != nil {
		return s, err
	}
	err = process.Start()
	if err != nil {
		return s, err
	}
	// Get the required data from the stdout of the bash invocation
	data, err := ioutil.ReadAll(stdout)
	if err != nil {
		return s, err
	}
	err = process.Wait()
	if err != nil {
		return s, err
	}
	for _, program := range strings.Split(string(data), "\n") {
		s[program] = true
	}
	return s, nil
}

func editDistance1(input string) []string {
	var splits [][]string
	var possibilities []string
	// Create a set of partial-words split by index
	for i := 0; i < len(input)+1; i++ {
		splits = append(splits, []string{input[:i], input[i:]})
	}
	// Now generate off-by-one possibilities for each combination
	for _, sp := range splits {
		// Delete one character
		if sp[1] != "" {
			possibilities = append(possibilities, sp[0]+sp[1][1:])
		}
		// Transpose one character
		if len(sp[1]) > 1 {
			possibilities = append(possibilities, fmt.Sprintf("%s%c%c%s", sp[0], sp[1][1], sp[1][0], sp[1][2:]))
		}
		// Replace one character
		if sp[1] != "" {
			for _, ch := range alphabet {
				possibilities = append(possibilities, fmt.Sprintf("%s%c%s", sp[0], ch, sp[1][1:]))
			}
		}
		// Insert one character
		for _, ch := range alphabet {
			possibilities = append(possibilities, fmt.Sprintf("%s%c%s", sp[0], ch, sp[1]))
		}
	}
	return possibilities
}

// Recursively generates the second edit distance from the input string
func editDistance2(distance1 []string) []string {
	var result []string
	for _, val := range distance1 {
		l := editDistance1(val)
		result = append(result, l...)
	}
	return result
}

func SpellCorrectCommand(args []string, _ io.Reader) error {
	/* Ideas for improvement:
	 *   Implement flags that select which kinds of things we should correct against (just executables, etc)
	 *   Create a weight for each program based on how frequently it's used
	 *   Add benchmarking/tests
	 *   TODO fix whatever is breaking "s"
	 */
	// Make sure we have one cli argument, the name to attempt correction on
	if len(args) < 2 {
		return fmt.Errorf("Usage: %s <token to spell correct>\n", args[0])
	}
	correctionTarget := args[1]
	// Load all possibilities
	executables, err := build_corpus()
	if err != nil {
		return err
	}
	// If the name is within the possibilities group, return immediately
	if _, ok := executables[correctionTarget]; ok {
		fmt.Println(correctionTarget)
		return nil
	}
	// Generate all off-by-1 and off-by-2 correction possibilities for name
	distance_one := editDistance1(correctionTarget)
	for _, val := range distance_one {
		// Blindly select the first match
		// TODO improve this later if necessary
		if _, ok := executables[val]; ok {
			fmt.Println(val)
			return nil
		}
	}
	// If we've gotten this far, none of the first-distance corrections worked. More, more I say!
	for _, val := range editDistance2(distance_one) {
		if _, ok := executables[val]; ok {
			fmt.Println(val)
			return nil
		}
	}
	// No match found in two distances, return an exit status of 1
	return errors.New("")
}
