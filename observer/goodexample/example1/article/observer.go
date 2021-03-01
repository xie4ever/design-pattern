package article

import "sync"

var (
	obs      observer
	obsMutex sync.Mutex
)

type observer struct {
	ProcessorMap sync.Map `json:"processor_map"`
}

func init() {
	rp := rankProcessor{
		ID:   1,
		Name: "rankProcessor",
	}
	pp := pointProcessor{
		ID:   2,
		Name: "pointProcessor",
	}
	obs.ProcessorMap.Store(rp.ID, rp)
	obs.ProcessorMap.Store(pp.ID, pp)
}

func (o observer) addProcessor(id int64, p processor) {
	obsMutex.Lock()
	o.ProcessorMap.Store(id, p)
	obsMutex.Unlock()
}

func (o observer) deleteProcessor(id int64) {
	obsMutex.Lock()
	o.ProcessorMap.Delete(id)
	obsMutex.Unlock()
}

func (o observer) postEvent(e Event) error {
	if e.ID == 0 {
		return nil
	}

	o.ProcessorMap.Range(
		func(k, v interface{}) bool {
			p := v.(processor)
			if e.Type == TypeEventAdd {
				if err := p.entryAdded(e); err != nil {
					return false
				}
			}
			if e.Type == TypeEventDelete {
				if err := p.entryDeleted(e); err != nil {
					return false
				}
			}
			if e.Type == TypeEventModify {
				if err := p.entryModified(e); err != nil {
					return false
				}
			}
			return true
		})

	return nil
}
