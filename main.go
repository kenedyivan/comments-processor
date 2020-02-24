package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
)

// Comment is a struct which represents comment and it's details
type Comment struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

// Comment is a struct which represents a single word and it's word count
type Word struct {
	Word  string
	Count int
}

type ByCount []Word

func (c ByCount) Len() int           { return len(c) }
func (c ByCount) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByCount) Less(i, j int) bool { return c[i].Count < c[j].Count }

func main() {
	ProcessComments()
}

//ProcessComments starts processing the comments data
func ProcessComments() {
	body := RequestCommentsData("https://jsonplaceholder.typicode.com/comments")
	comments, _ := GetCommentsFromString(body)
	commentBody := ExtractCommentBody(comments)
	words := GetWordsFromString(commentBody)
	wordFrequencyMap := GetWordFrequency(words)
	myWords := ParseToSortedWordSlice(wordFrequencyMap)
	fourLeastUsedWords := GetWords(myWords, 4)
	Display(fourLeastUsedWords)
}

//Display prints the words and their word count to console
func Display(w []Word) {

	for k, v := range w {
		fmt.Printf("%d, %s -> %d\n", k+1, v.Word, v.Count)
	}
}

// RequestCommentsData makes a Get Http request to the url and returns a json string with the comments
func RequestCommentsData(url string) string {
	log.Println("Fetching comments data...")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

// GetWords makes a Get Http request to the url and returns a json string with the comments
func GetWords(words []Word, count int) []Word {
	return words[:count]
}

// GetWordsFromString returns a slice of words in the given string
func GetWordsFromString(s string) []string {
	return strings.Fields(s)
}

// ExtractCommentBody extracts the body content from the Comment struct and returns a string with all comment bodies
func ExtractCommentBody(comments []Comment) string {
	var commentBody string
	for _, v := range comments {
		commentBody += " " + fmt.Sprint(strings.TrimSpace(v.Body))
	}
	return strings.TrimSpace(commentBody)
}

// GetCommentsFromString parses the json and returns a slice of Comment
func GetCommentsFromString(s string) ([]Comment, error) {

	var comments []Comment
	var errr error

	err := json.Unmarshal([]byte(s), &comments)
	if err != nil {
		errr = err
	}
	return comments, errr
}

// GetWordFrequency gets counts the number of times each word appears in the string and returns a map of each word and
// the number of times is appears in the string
func GetWordFrequency(s []string) map[string]int {
	m := make(map[string]int)
	for _, v := range s {
		if val, ok := m[v]; ok {
			m[v] = val + 1
		} else {
			m[v] = 1
		}
	}
	return m
}

// ParseToSortedWordSlice converts each item in the map to a Word Type
// Returns a slice of Word
func ParseToSortedWordSlice(m map[string]int) []Word {
	var w []Word
	for k, v := range m {
		w = append(w, Word{Word: k, Count: v})
	}

	return SortByCount(w)
}

// SortByCount sorts the items in the Word slice by Count
func SortByCount(w []Word) []Word {
	wordSlice := w
	sort.Sort(ByCount(wordSlice))
	return wordSlice
}
