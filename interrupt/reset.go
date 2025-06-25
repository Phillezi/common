package interrupt

import "sync"

func resetInstance() {
	instance = nil
	once = sync.Once{}
}
