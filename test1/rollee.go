package rollee

import "sync"

type ID = int

// We suppose L is always valid with len (l.Values) >= 1).
type List struct {
	ID     ID
	Values []int
}

func Fold(initialValue int, f func(int, int) int, l List) map[ID]int {
	result := make(map[int]int)

	//Recursive helper function
	var recursive func(int, []int) int
	recursive = func(finalValue int, values []int) int {
		if len(values) == 0 {
			return finalValue
		}
		return recursive(f(finalValue, values[0]), values[1:])
	}

	finalValue := recursive(initialValue, l.Values)
	result[l.ID] = finalValue

	return result
}

func FoldChan(initialValue int, f func(int, int) int, ch chan List) map[ID]int {
	result := make(map[int]int)

	// Recursive helper function
	var recursive func(int, []int) int
	recursive = func(finalValue int, values []int) int {
		if len(values) == 0 {
			return finalValue
		}
		return recursive(f(finalValue, values[0]), values[1:])
	}

	// Iterate over the lists in the channel and call the recursive function
	for list := range ch {
		currentValue, exists := result[list.ID]
		if exists {
			result[list.ID] = recursive(f(currentValue, initialValue), list.Values)
		} else {
			result[list.ID] = recursive(initialValue, list.Values)
		}
	}

	return result
}

func FoldChanX(initialValue int, f func(int, int) int, chs ...chan List) map[int]int {
	result := make(map[int]int)
	var mutex sync.Mutex // Mutex to protect concurrent access to the 'result' map

	// Recursive helper function
	var recursive func(int, []int) int
	recursive = func(finalValue int, values []int) int {
		if len(values) == 0 {
			return finalValue
		}
		return recursive(f(finalValue, values[0]), values[1:])
	}

	// Waitgroup to ensure main goroutine waits for all goroutines to finish before returning result
	wg := sync.WaitGroup{}
	wg.Add(len(chs))

	for _, ch := range chs {
		go func(c chan List) {
			defer wg.Done()
			for list := range c {
				currentValue, exists := func() (int, bool) {
					mutex.Lock()
					defer mutex.Unlock()
					val, ok := result[list.ID]
					return val, ok
				}()

				if exists {
					updatedValue := recursive(f(currentValue, initialValue), list.Values)

					mutex.Lock()
					result[list.ID] = updatedValue
					mutex.Unlock()
				} else {
					updatedValue := recursive(initialValue, list.Values)

					mutex.Lock()
					result[list.ID] = updatedValue
					mutex.Unlock()
				}
			}
		}(ch)
	}

	wg.Wait()

	return result
}
