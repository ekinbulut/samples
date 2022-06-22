package main

import (
	aservice "samples/app/arraylist"
	hservice "samples/app/hashtable"
	t "samples/app/tree"
)

type App struct {
	Name      string
	Version   string
	ArrayList *aservice.ArrayList
	HashTable *hservice.HashTable
}

func NewApp() *App {
	return &App{
		Name:      "Samples App",
		Version:   "1.0.0",
		ArrayList: aservice.NewArrayList(),
		HashTable: hservice.NewHashTable(10),
	}
}

func main() {

	t := t.NewTree()

	t.Add(5)
	t.Add(3)
	t.Add(10)

	t.Traversal()

	// app := NewApp()
	// app.ArrayList.Add(1)
	// app.ArrayList.Print()
}
