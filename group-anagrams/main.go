package main

import "fmt"

type AnagramRunes struct {
	runesAmount map[rune]int
	runes []rune
}

func main() {
	res := groupAnagrams([]string{"tea","and","ace","ad","eat","dans"})

	fmt.Println(res)
}

func groupAnagrams(strs []string) [][]string {
    res := make([][]string, 0)

	anagrams := make(map[string][]string, 0)
	anagramChars := make(map[string]AnagramRunes, 0)
	for _, str := range strs {
		strRunesMap := make(map[rune]int, len(str))
		for _, char := range str {
			strRunesMap[char]++
		}

		if _, exists := anagramChars[str]; exists {
			anagrams[str] = append(anagrams[str], str)
			continue
		} else {
			if len(anagrams) == 0 {
				addNewAnagram(anagrams, anagramChars, str, strRunesMap)

				continue
			}

			anagramFound := false
			for k, anagramRunes := range anagramChars {
				if k == str {
					anagrams[k] = append(anagrams[k], str)
					anagramFound = true
					break
				}

				if len(k) != len(str) {
					continue
				}

				anagramCharsCount := 0
				for _, r := range anagramRunes.runes {
					strRuneCount, ok := strRunesMap[r]
					if !ok {
						break
					}

					if strRuneCount != anagramRunes.runesAmount[r] {
						break
					}

					anagramCharsCount++
				}

				if anagramCharsCount > 0 && anagramCharsCount == len(anagramRunes.runes) {
					// Анаграмма из имеющихся к данному слову найдена - добавляем к нему в список
					anagrams[k] = append(anagrams[k], str)
					anagramFound = true
					break
				}
			}

			if !anagramFound {
				// Анаграммы из имеющихся к данному слову не найдено - добавляем с список анаграмм отдельно.
				// В дальнейшем будем проверять посимвольно, есть ли ещё какие-то слова-анаграммы для этого слова
				// anagrams[str] = []string{str}
				addNewAnagram(anagrams, anagramChars, str, strRunesMap)
			}
		}
	}

	for _, anagramGroup := range anagrams {
		res = append(res, anagramGroup)
	}

	return res
}

func addNewAnagram(anagrams map[string][]string, anagramChars map[string]AnagramRunes, str string, strRunesMap map[rune]int){
	anagrams[str] = []string{str}

	anagramRunes := AnagramRunes{
		runesAmount: make(map[rune]int, 0),
		runes: make([]rune, 0, len(strRunesMap)),
	}
	for k, v := range strRunesMap {
		anagramRunes.runes = append(anagramRunes.runes, k)

		if _, ok := anagramRunes.runesAmount[k]; ok {
			anagramRunes.runesAmount[k]++
		} else {
			anagramRunes.runesAmount[k] = v
		}
	}
	
	anagramChars[str] = anagramRunes
}
