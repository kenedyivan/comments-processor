package main

import (
	"strings"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	var comments = `[
  {
    "postId": 1,
    "id": 1,
    "name": "id labore ex et quam laborum",
    "email": "Eliseo@gardner.biz",
    "body": "laudantium enim quasi est quidem magnam voluptate ipsam eos\ntempora quo necessitatibus\ndolor quam autem quasi\nreiciendis et nam sapiente accusantium"
  }
]`
	actual, _ := GetCommentsFromString(comments)
	if actual[0].PostID != 1 {
		t.Errorf("Expected post id to be  '1'. Got '%v'", actual[0].PostID)
	}
}

func TestUnmarshalFailed(t *testing.T) {
	var comments = `[
  {
    postId": 1,
    "id": 1,
    "name": "id labore ex et quam laborum",
    "email": "Eliseo@gardner.biz",
    "body": "laudantium enim quasi est quidem magnam voluptate ipsam eos\ntempora quo necessitatibus\ndolor quam autem quasi\nreiciendis et nam sapiente accusantium"
  }
]`
	_, err := GetCommentsFromString(comments)
	if err == nil {
		t.Errorf("Expected unmarshal error. Got '%v'", err)
	}
}

func TestExtractCommentsBody(t *testing.T) {
	comments := []Comment{
		{Body: "sint nostrum voluptatem reiciendis et"},
		{Body: "quia molestiae reprehenderit quasi aspernatur\naut expedita"},
	}

	expectedString := "sint nostrum voluptatem reiciendis et quia molestiae reprehenderit quasi aspernatur\naut expedita"

	commentBody := ExtractCommentBody(comments)
	if commentBody != expectedString {
		t.Errorf("Expected string to be '%v'. Got '%v'", expectedString, commentBody)
	}
}

func TestExtractCommentsBodyWithPadding(t *testing.T) {
	comments := []Comment{
		{Body: "  sint nostrum voluptatem reiciendis et  "},
		{Body: " quia molestiae reprehenderit quasi aspernatur\naut expedita  "},
	}

	expectedString := "  sint nostrum voluptatem reiciendis et quia molestiae reprehenderit quasi aspernatur\naut expedita   "

	commentBody := ExtractCommentBody(comments)
	if commentBody != strings.TrimSpace(expectedString) {
		t.Errorf("Expected string to be '%v'. Got '%v'", expectedString, commentBody)
	}
}

func TestGetCommentBodyWordSlice(t *testing.T) {

	input := "Hello world hello world hello world hello hello hello hello world world"

	generatedWordSlice := getWordsFromString(input)
	if len(generatedWordSlice) != 12 {
		t.Errorf("Expected slice size to be '12'. Got '%v'", len(generatedWordSlice))
	}
}

func TestGenerateWordFrequencyMap(t *testing.T) {

	wordSlice := []string{"Hello", "world", "hello", "world", "hello", "world", "hello", "hello", "hello", "hello", "world", "world"}

	generatedWordFrequencyMap := getWordFrequency(wordSlice)

	if generatedWordFrequencyMap["Hello"] != 1 {
		t.Errorf("Expected slice Hello frequency to be '1'. Got '%v'", generatedWordFrequencyMap["Hello"])
	}

	if generatedWordFrequencyMap["hello"] != 6 {
		t.Errorf("Expected slice hello frequency to be '6'. Got '%v'", generatedWordFrequencyMap["hello"])
	}

	if generatedWordFrequencyMap["world"] != 5 {
		t.Errorf("Expected slice world frequency to be '5'. Got '%v'", generatedWordFrequencyMap["world"])
	}

}

func TestParseWordFrequencyMapToWordTypeMapAndSort(t *testing.T) {

	w := map[string]int{"Hello": 1}

	wordSlice := ParseToSortedWordSlice(w)

	if wordSlice[0].Word != "Hello" {
		t.Errorf("Expected Word to be 'Hello' . Got '%v'", wordSlice[0].Word)
	}

	if wordSlice[0].Count != 1 {
		t.Errorf("Expected Count to be '1' . Got '%v'", wordSlice[0].Count)
	}

}

func TestSortByCount(t *testing.T) {

	w := []Word{
		{Word: "Aello", Count: 30},
		{Word: "mino", Count: 2},
		{Word: "bello", Count: 27},
		{Word: "facebook", Count: 6},
	}

	sortedWordSlice := SortByCount(w)

	if sortedWordSlice[0].Word != "mino" {
		t.Errorf("Expected Word to be 'mino' . Got '%v'", sortedWordSlice[0].Word)
	}

	if sortedWordSlice[1].Word != "facebook" {
		t.Errorf("Expected Word to be 'facebook' . Got '%v'", sortedWordSlice[1].Word)
	}

	if sortedWordSlice[2].Word != "bello" {
		t.Errorf("Expected Word to be 'bello' . Got '%v'", sortedWordSlice[2].Word)
	}

	if sortedWordSlice[3].Word != "Aello" {
		t.Errorf("Expected Word to be 'Aello' . Got '%v'", sortedWordSlice[3].Word)
	}

}

func TestGetWords(t *testing.T) {

	w := []Word{
		{Word: "Aello", Count: 30},
		{Word: "mino", Count: 5},
		{Word: "bello", Count: 8},
		{Word: "facebook", Count: 6},
		{Word: "google", Count: 1},
		{Word: "twitter", Count: 5},
		{Word: "instagram", Count: 4},
		{Word: "food", Count: 3},
	}

	wordSlice := GetWords(w, 4)
	if len(wordSlice) != 4 {
		t.Errorf("Expected size to be '4'. Got '%v'", len(wordSlice))
	}
}
