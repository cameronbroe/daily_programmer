package cache

import (
	"../utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Cache struct {
	Entries  map[DifficultyType]map[int]Entry `json:"entries"`
	filename string
}

type Entry struct {
	Number      int            `json:"number"`
	Difficulty  DifficultyType `json:"difficulty"`
	Description string         `json:"description"`
	Url         string         `json:"url"`
}

func CreateEntry(number int, difficulty DifficultyType, description string, url string) Entry {
	return Entry{
		Number:      number,
		Difficulty:  difficulty,
		Description: description,
		Url:         url,
	}
}

func (e Entry) Display() {
	fmt.Printf("number: %d\n", e.Number)
	fmt.Printf("difficulty: %d\n", e.Difficulty)
	fmt.Printf("description: %s\n", e.Description)
	fmt.Printf("URL: %s\n", e.Url)
}

type DifficultyType int

const (
	DifficultyEasy   DifficultyType = 1
	DifficultyMedium DifficultyType = 2
	DifficultyHard   DifficultyType = 3
)

type CacheOptions struct {
	filename string
}

type CacheOption func(*CacheOptions)

func filename(filename string) CacheOption {
	return func(s *CacheOptions) { s.filename = filename }
}

func New(options ...CacheOption) Cache {
	args := &CacheOptions{
		filename: utils.DefaultCachePath(),
	}
	for _, option := range options {
		option(args)
	}

	return Cache{
		Entries:  make(map[DifficultyType]map[int]Entry),
		filename: args.filename,
	}
}

func FromFile(options ...CacheOption) Cache {
	args := &CacheOptions{
		filename: utils.DefaultCachePath(),
	}
	for _, option := range options {
		option(args)
	}

	data, err := ioutil.ReadFile(args.filename)
	if err != nil {
		panic(err)
	}
	cache := Cache{}
	json.Unmarshal(data, &cache)
	cache.filename = args.filename
	return cache
}

func (c *Cache) AddEntry(entry *Entry) {
	difficultyEntries, ok := c.Entries[entry.Difficulty]
	if !ok {
		difficultyEntries = make(map[int]Entry)
		c.Entries[entry.Difficulty] = difficultyEntries
	}
	difficultyEntries[entry.Number] = *entry
}

func (c *Cache) RemoveEntry(difficulty DifficultyType, dpId int) {
	delete(c.Entries[difficulty], dpId)
	if len(c.Entries[difficulty]) == 0 {
		delete(c.Entries, difficulty)
	}
}

func (c *Cache) GetEntryById(difficulty DifficultyType, dpId int) Entry {
	return c.Entries[difficulty][dpId]
}

func (c *Cache) EntryCount() int {
	counter := 0
	for _, difficultyEntries := range c.Entries {
		counter += len(difficultyEntries)
	}
	return counter
}

func (c *Cache) Display() {
	fmt.Printf("There are %d entries in this cache:\n", c.EntryCount())
	for _, difficultyEntries := range c.Entries {
		for _, entry := range difficultyEntries {
			entry.Display()
		}
	}
}

func (c *Cache) Save() {
	file, err := os.Create(c.filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cacheJson, err := json.Marshal(*c)
	if err != nil {
		panic(err)
	}

	file.Write(cacheJson)
}

func (c *Cache) SavePretty() {
	file, err := os.Create(c.filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cacheJson, err := json.MarshalIndent(*c, "", "    ")
	if err != nil {
		panic(err)
	}

	file.Write(cacheJson)
}
