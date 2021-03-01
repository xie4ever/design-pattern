package article

import "log"

type processor interface {
	entryAdded(e Event) error
	entryDeleted(e Event) error
	entryModified(e Event) error
}

type rankProcessor struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (p rankProcessor) entryAdded(e Event) error {
	log.Printf("rankProcessor entryAdded, event id :%d", e.ID)
	return nil
}

func (p rankProcessor) entryDeleted(e Event) error {
	log.Printf("rankProcessor entryDeleted, event id :%d", e.ID)
	return nil
}

func (p rankProcessor) entryModified(e Event) error {
	log.Printf("rankProcessor entryModified, event id :%d", e.ID)
	return nil
}

type pointProcessor struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (p pointProcessor) entryAdded(e Event) error {
	log.Printf("pointProcessor entryAdded, event id :%d", e.ID)
	return nil
}

func (p pointProcessor) entryDeleted(e Event) error {
	log.Printf("pointProcessor entryDeleted, event id :%d", e.ID)
	return nil
}

func (p pointProcessor) entryModified(e Event) error {
	log.Printf("pointProcessor entryModified, event id :%d", e.ID)
	return nil
}
