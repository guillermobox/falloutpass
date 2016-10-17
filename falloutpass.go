package main

import (
	"fmt"
	"os"
	)

/* Calculate the score of a word pair */
func score(word, guess string) int {
	score := 0
	for i := 0; i < len(guess); i++ {
		if word[i] == guess[i] {
			score += 1
		}
	}
	return score
}

/* Is this word compatible with the given guess and score? */
func compatible(word, guess string, pairscore int) bool {
	return pairscore == score(word, guess)
}

/* How many compatible words are in the list for the guess/score? */
func compatibles(guess string, words []string, wordscore int) int {
	return len(words) - strokeout(guess, words, wordscore)
}

/* How many words I can strikeout with the guess/score combination */
func strokeout(guess string, words []string, wordscore int) int {
	count := 0
	for _, word := range words {
		if !compatible(guess, word, wordscore) {
			count += 1
		}
	}
	return count
}

/* What is the maximum number of words that will survive for this guess? */
func maxwords(guess string, words []string) int {
	max := 0
	for score := 0; score <= len(guess); score++ {
		valid := compatibles(guess, words, score)
		if valid > max {
			max = valid;
		}
	}
	return max
}

/* How many words can I strokeout for each possible solution for this guess? */
func strokeoutsum(guess string, words []string) int {
	total := 0
	for _, solution := range words {
		s := score(guess, solution)
		maximal := strokeout(guess, words, s)
		total += maximal
	}
	return total
}

/* Return the optimal guess that potentialy removes more words */
func optimalstrokeout(words []string) string {
	var optimal string
	optimalscore := 0

	for _, guess := range words {
		total := strokeoutsum(guess, words)

		if total > optimalscore {
			optimalscore = total
			optimal = guess
		}
	}
	return optimal
}

/* Return the optimal guess that allows the least ammount of words to survive */
func optimalmaxwords(words []string) string {
	var optimal string
	optimalscore := len(words[0])

	for _, guess := range words {
		total := maxwords(guess, words)

		if total < optimalscore {
			optimalscore = total
			optimal = guess
		}
	}
	return optimal
}

/* Filter the given word list for the know guess and score */
func filter(guess string, words []string, score int) []string {
	var list []string
	for _, word := range words {
		if compatible(word, guess, score) {
			list = append(list, word)
		}
	}
	return list
}

/* Analyse the effect of choosing one word */
func analyse(guess string, words []string) {
	fmt.Printf("    %s", guess)
	for score := 0; score <= len(guess); score++ {
		valid := compatibles(guess, words, score)
		if valid != 0 {
			fmt.Printf(" %2d ", valid)
		} else {
			fmt.Printf("  - ");
		}
	}
	fmt.Printf("\n")
}

func main() {
	var words []string

	if len(os.Args) < 2 {
		fmt.Println("Please provide the file with words as first argument")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Impossible to open file", os.Args[1])
		return
	}

	for {
		var word string
		_, err := fmt.Fscanln(file, &word)
		if err != nil {
			break
		}
		words = append(words, word)
	}
	file.Close()

	if len(words) == 0 {
		fmt.Println("At least provide one word!")
		return
	}

	for len(words) > 1 {
		var word string
		var score int
		var found bool

		fmt.Println()
		fmt.Println("With", len(words), "words left, I recommend you to chose", optimalmaxwords(words))
		fmt.Println()
		fmt.Print("What word and score did you play (empty to see info)? ")

		_, err := fmt.Scanln(&word, &score)
		if err != nil {
			fmt.Println()
			if err.Error() == "unexpected newline" {
				fmt.Println("    Word list", words)
				fmt.Println("    Optimal maxwords and strokeout are", optimalmaxwords(words), optimalstrokeout(words))
				fmt.Println()
				fmt.Println("    Distribution of surviving words")
				for _, word = range words {
					analyse(word, words)
				}
				continue
			} else {
				fmt.Println("Error found when reading your answer:", err)
				return
			}
		}

		found = false
		for _, w := range words {
			if w == word {
				found = true
				continue
			}
		}

		if !found {
			fmt.Println()
			fmt.Println("That word is not in the list")
			continue
		}

		if compatibles(word, words, score) == 0 {
			fmt.Println()
			fmt.Println("That word/score combination is not in the list")
			continue
		}

		words = filter(word, words, score)
	}

	fmt.Println()
	fmt.Println("That's it, you have finished the game! The word is", words[0])
}
