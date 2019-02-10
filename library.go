package main

import (
	"path/filepath"
	"fmt"
	"os"
	"time"
	"github.com/dhowden/tag"
	"errors"
	"strings"
)

type Item struct {
	name string
	artist string
	album string
	year int
	path string
	size int64
}

type Library struct {
	path string
	listing map[string]Item
	modifiedAt int64
}

var validExtensions = map[string]int{ "mp3": 1 }

func CreateLibrary(path string) Library {
	return Library {
		path: path,
		listing: CreateListing(path),
		modifiedAt: time.Now().UnixNano(),
	}
}

func (lib * Library) UpdateLibrary() {
	lib.modifiedAt = time.Now().UnixNano()
	lib.listing = CreateListing(lib.path)
}

func (lib * Library) ToString() string {
	p := fmt.Sprintf("Library for path: %v\nLast updated at: %v", lib.path, lib.modifiedAt)
	for k, v := range lib.listing {
		p = fmt.Sprintf("%v\n%v\n%v", p, k, v.ToString())
	}

	return p
}

func (it * Item) ToString() string {
	return fmt.Sprintf("name: %v,\nartist: %v,\nalbum: %v,\nyear: %v,\npath: %v,\nsize: %v,\n",
		it.name, it.artist, it.album, it.year, it.path, it.size)
}

func CreateListing(path string) map[string]Item {
	listing := make(map[string]Item)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if (!info.IsDir()) {
			if _, ok := CheckExtension(path); ok {
				id, item, err := GenerateItem(info, path)
				if err != nil {
					return nil
				}
	
				listing[id] = item
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	return listing
}

func CheckExtension(path string) (string, bool) {
	splits := strings.Split(path, ".")
	extension := splits[len(splits) - 1]
	if _, ok := validExtensions[extension]; ok {
		return extension, true
	}

	return "", false
}

func GenerateItem(info os.FileInfo, path string) (string, Item, error) {
	var generated Item

	f, err := os.Open(path)
	if err != nil {
		return "", generated, errors.New(fmt.Sprintf("Unable to read file at path: %v", path))
	}
	defer f.Close()

	m, err := tag.ReadFrom(f)
	if err != nil {
		return "", generated, errors.New(fmt.Sprintf("Unable to read metadata from %v", path))
	}

	id, err := tag.Sum(f);
	if err != nil {
		return "", generated, errors.New(fmt.Sprintf("Unable to read metadata from %v", path))
	}

	generated = Item{
		name: m.Title(),
		artist: m.Artist(),
		album: m.Album(),
		year: m.Year(),
		path: path,
		size: info.Size(),
	}

	return id, generated, nil
}