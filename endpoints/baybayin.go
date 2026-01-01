package endpoints

import (
	"regexp"
	"strings"

	"portfolio-backend/utils"

	"github.com/gin-gonic/gin"
)

/*
 * TODO: This is just a template for the other endpoint
 */

var GetBaybayin = utils.Route{
	Path:    "/baybayin",
	Handler: baybayin,
}

// TODO: Core functions
func transliteration(word string) string {
	var result strings.Builder

	CONSONANTS := utils.BaybayinCharacters{
		"b": 5898, "k": 5891, "d": 5895, "g": 5892,
		"h": 5905, "l": 5902, "m": 5899, "n": 5896,
		"p": 5897, "s": 5904, "t": 5894,
		"w": 5903, "y": 5900,
	}

	VOWELS := utils.BaybayinCharacters{
		"a": 5888,
		"e": 5889,
		"o": 5890,
	}

	NG := 5893

	VOWELDICTRITICS := utils.BaybayinCharacters{
		"e":       5906,
		"o":       5907,
		"default": 5908,
	}

	PUNCTUATIONS := utils.BaybayinCharacters{
		".": 5942,
		"?": 5942,
		"!": 5942,
		",": 5941,
	}

	characters := []rune(word)

	for i := 0; i < len(characters); i++ {
		c := string(characters[i])

		// TODO: CONSONANT CHECK (equivalent to Set.has)
		if code, ok := CONSONANTS[c]; ok {
			next := ""

			if i+1 < len(characters) {
				next = string(characters[i+1])
			}

			// TODO: NG to NANG
			if i+1 < len(characters) {
				if strings.EqualFold(c, "n") && strings.EqualFold(next, "g") {
					result.WriteRune(rune(NG))
					i += 2
					continue
				}
			}

			// TODO: Removal of H with virama
			if strings.EqualFold(c, "h") {
				if next != "" {
					result.WriteRune(rune(code))
				}
				if next == "" {
					continue
				}
			} else {
				result.WriteRune(rune(code))
			}

			// TODO: Vowel diacritics
			if v, ok := VOWELDICTRITICS[next]; ok {
				result.WriteRune(rune(v))
				i++
			} else if !strings.EqualFold(next, "a") {
				result.WriteRune(rune(VOWELDICTRITICS["default"]))
			} else {
				i++
			}
			continue

		} else if code, ok := VOWELS[c]; ok {
			result.WriteRune(rune(code))
			continue
		}

		// TODO: PUNCTUATIONS
		if code, ok := PUNCTUATIONS[c]; ok {
			result.WriteRune(rune(code))
			continue
		}
	}

	if len(result.String()) > 0 {
		return result.String()
	}
	return ""
}

// TODO: Regex string normalization
func normalize(source string, regex string, tc string) string {
	return regexp.MustCompile(regex).ReplaceAllString(source, tc)
}

// TODO: Process for making string as readable by the transliterator
func process(text string) string {
	result := []string{}

	text = strings.ToLower(text)

	// TODO: Unnecessary characters removal
	text = regexp.MustCompile("-").ReplaceAllString(text, "")

	// TODO: String normalization
	text = normalize(text, `i`, "e") // regexp.MustCompile(`i`).ReplaceAllString(text, "e")
	text = regexp.MustCompile(`u`).ReplaceAllString(text, "o")
	text = regexp.MustCompile(`r`).ReplaceAllString(text, "d")
	text = regexp.MustCompile(`\bmga\b`).ReplaceAllString(text, "manga")
	text = regexp.MustCompile(`f`).ReplaceAllString(text, "p")
	text = regexp.MustCompile(`v`).ReplaceAllString(text, "b")
	text = regexp.MustCompile(`x`).ReplaceAllString(text, "s")
	text = regexp.MustCompile(`c`).ReplaceAllString(text, "k")
	text = regexp.MustCompile(`j`).ReplaceAllString(text, "dy")
	text = regexp.MustCompile(`\bng\b`).ReplaceAllString(text, "nang")

	// TODO: Transliteration Peocess
	split := strings.Split(text, " ")

	// TODO: Actual transliteration process
	for w := range split {
		word := split[w]
		if !strings.EqualFold(word, "") {
			baybay := transliteration(word)
			if baybay != "" {
				result = append(result, baybay)
			}
		}
	}

	return strings.Join(result, " ")
}

func baybayin(ctx *gin.Context) {
	if ctx.Query("text") != "" {
		ctx.JSON(200, gin.H{
			"original": ctx.Query("text"),
			"response": process(ctx.Query("text")),
		})
		return
	}
	ctx.JSON(403, gin.H{
		"error": "Required parameter \"text\"",
	})
}
