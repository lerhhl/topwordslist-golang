package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GetHtmlText is to get clean html text from a page
func GetHtmlText(url string) string {
	urlLink := url
	doc, err := goquery.NewDocument(urlLink)
	if err != nil {
		log.Fatal(err)
	}

	bodyText := ""

	orginalText := doc.Find("#Introduction").Children().Next().Text()
	cleanText := strings.Replace(orginalText, ".", "", -1)
	cleanText = strings.Replace(cleanText, "\"", "", -1)
	cleanText = strings.Replace(cleanText, "\n", "", -1)
	cleanText = strings.TrimSpace(cleanText)
	bodyText = bodyText + cleanText

	return bodyText
}

// WordFeqMap counts the word in a string
func WordFeqMap(s string) map[string]int {
	wordFeq := make(map[string]int)
	re := regexp.MustCompile(" +")
	sArr := re.Split(s, -1)

	for i := 0; i < len(sArr); i++ {
		if wordFeq[sArr[i]] == 0 {
			wordFeq[sArr[i]] = 1
		} else {
			wordFeq[sArr[i]]++
		}
	}
	return wordFeq
}

// Table structure
type Table struct {
	Key   string
	Value int
}

func (word Table) String() string {
	return fmt.Sprintf("%s: %d\n", word.Key, word.Value)
}

// SortTopWords sorts the WordFeqMap
func SortTopWords(WordFeqMap map[string]int) []Table {
	var newMap []Table
	for k, v := range WordFeqMap {
		newMap = append(newMap, Table{k, v})
	}

	sort.Slice(newMap, func(i, j int) bool {
		return newMap[i].Value > newMap[j].Value
	})

	// List top 10 most-used words
	topNum := 10
	if len(newMap) < topNum {
		topNum = len(newMap)
	}

	topSelectWords := newMap[:topNum]
	return topSelectWords
}

func main() {
	url := "http://www.themindofgod.net/samples.htm"
	bodyText := GetHtmlText(url)
	wordFeqMap := WordFeqMap(bodyText)
	fmt.Println(SortTopWords(wordFeqMap))
}
