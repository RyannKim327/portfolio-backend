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
	// TODO: To set default values
	CONSONANTS := utils.BaybayinCharacters{
		"b": 5898,
		"k": 5891,
		"d": 5895,
		"g": 5892,
		"h": 5905,
		"l": 5902,
		"m": 5899,
		"n": 5896,
		"p": 5897,
		"s": 5904,
		"t": 5894,
		"w": 5903,
		"y": 5900,
	}

	NG := 5893

	VOWELS := utils.BaybayinCharacters{
		"a": 5906,
		"e": 5907,
		"o": 5908,
	}

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

	// TODO: To gather all the keys from each of interface
	consonantsKeys := []string{}
	vowelkeys := []string{}
	vowelDictriticsKeys := []string{}
	punctuationKeys := []string{}

	for k := range CONSONANTS {
		consonantsKeys = append(consonantsKeys, k)
	}
	for k := range VOWELS {
		vowelkeys = append(vowelkeys, k)
	}
	for k := range VOWELDICTRITICS {
		vowelDictriticsKeys = append(vowelDictriticsKeys, k)
	}
	for k := range PUNCTUATIONS {
		punctuationKeys = append(punctuationKeys, k)
	}

	return word
}

func process(text string) string {
	result := []string{}

	text = strings.ToLower(text)
	// TODO: Unnecessary characters removal
	text = regexp.MustCompile("-").ReplaceAllString(text, "")

	// TODO: String normalization
	text = regexp.MustCompile(`i`).ReplaceAllString(text, "e")
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
			result = append(result, transliteration(word))
		}
	}

	return strings.Join(result, " ")
}

func baybayin(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"response": process("Kamusta ka aking mahal"),
	})
}
