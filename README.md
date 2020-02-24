# Comments-processor

Comments-processor is a small tool written in go.

The purpose of this tool to fetch comments data from [https://jsonplaceholder.typicode.com/comments](https://jsonplaceholder.typicode.com/comments) and process the body field returning the four least used words based on their word count.

## Getting Started

## Install

## Run
Build
```bash
make build
```
Run
```bash
make run
```

## Functions
The processor starts by running a series of functions in sequence in the ProcessComments method

```go
//ProcessComments starts processing the comments data
func ProcessComments() {
	body := RequestCommentsData("https://jsonplaceholder.typicode.com/comments")
	comments, _ := GetCommentsFromString(body)
	commentBody := ExtractCommentBody(comments)
	words := getWordsFromString(commentBody)
	wordFrequencyMap := getWordFrequency(words)
	myWords := ParseToSortedWordSlice(wordFrequencyMap)
	fourLeastUsedWords := GetWords(myWords, 4)
	DisplayWords(fourLeastUsedWords)
}
```

A request is made to the URL that returns a json string with the comments array
```go
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
```
##GetCommentsFromString
The json string with the comments body is converted to
```go
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
```

### ExtractCommentBody
Extracts the body contents from all the comments in the comment slice
```go
// ExtractCommentBody extracts the body content from the Comment struct and returns a string with all comment bodies
func ExtractCommentBody(comments []Comment) string {
	var commentBody string
	for _, v := range comments {
		commentBody += " " + fmt.Sprint(strings.TrimSpace(v.Body))
	}
	return strings.TrimSpace(commentBody)
}

// getWordsFromString returns a slice of words in the given string
func getWordsFromString(s string) []string {
	return strings.Fields(s)
}
```
### GetWordFrequency

```go
// getWordFrequency gets counts the number of times each word appears in the string and returns a map of each word and
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
```
### ParseToSortedWordSlice
ParseToSortedWordSlice converts each item in the map to a Word Type
```go
// ParseToSortedWordSlice converts each item in the map to a Word Type
// Returns a slice of Word
func ParseToSortedWordSlice(m map[string]int) []Word {
	var w []Word
	for k, v := range m {
		w = append(w, Word{Word: k, Count: v})
	}

	return SortByCount(w)
}
```
### GetWords
GetWords returns a slice with a given number of words
```go
// GetWords returns a slice with a given number of words
func GetWords(words []Word, count int) []Word {
	return words[:count]
}
```
### Display
This function is used for printing the results of the process to the console
```go
//DisplayWords prints the words and their word count to console
func Display(w []Word) {

	for k, v := range w {
		fmt.Printf("%d, %s -> %d\n", k+1, v.Word, v.Count)
	}
}
```

## Testing
```bash
make test
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
