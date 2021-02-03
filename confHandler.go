package gometer

import (
	"fmt"
	"sync"
)


func Handle(c *Conf) {
	keeper := make(map[string]string)

	var wg sync.WaitGroup

	fmt.Println("-----Before-----")

	// execute "before" requests
	wg.Add(1)
	handleTests(c.Before, keeper, &wg)
	wg.Wait()

	//we iterate through test sets
	for i := 0; i < len(c.TestSets); i++ {

		fmt.Println("-----Test Set", i+1, "Begin-----")

		for retry := 0; retry < c.TestSets[i].Retries; retry++ {
			wg.Add(1)
			go handleTests(c.TestSets[i].Tests, keeper, &wg)
		}

		wg.Wait()
	}

	fmt.Println("-----After-----")
	// execute "after" requests
	wg.Add(1)
	handleTests(c.After, keeper, &wg)
	wg.Wait()
}
