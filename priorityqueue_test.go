package graph

import "testing"

// item is a (name,distance) pair for the testcases.
type item struct {
	name     string
	distance float64
}

// TestPriorityQueue fills a priority queue with items,
// update some of them and then pop them all.
// It checks that items have been popped in the expected order.
func TestPriorityQueue(t *testing.T) {
	testcases := []struct {
		// human readable name for this testcase
		name string

		// items in the priority queue
		items []item

		// distances to update
		toUpdate []item

		// ordered list of item names
		ordered []string
	}{
		// empty queue
		{
			name: "empty",
		},

		// one item queue
		{
			name: "one_item",
			items: []item{
				item{"a", 0.5},
			},
			ordered: []string{"a"},
		},

		// one item queue with update
		{
			name: "one_item_update",
			items: []item{
				item{"a", 0.5},
			},
			toUpdate: []item{
				item{"a", 0.8},
			},
			ordered: []string{"a"},
		},

		// sorted items
		{
			name: "sorted_items",
			items: []item{
				item{"a", 0.5},
				item{"z", 0.7},
				item{"c", 0.9},
			},
			ordered: []string{"a", "z", "c"},
		},

		// reverse sorted items
		{
			name: "reverse_sorted_items",
			items: []item{
				item{"c", 0.9},
				item{"z", 0.7},
				item{"a", 0.5},
			},
			ordered: []string{"a", "z", "c"},
		},

		// random order items
		{
			name: "random_order_items",
			items: []item{
				item{"z", 0.7},
				item{"c", 0.9},
				item{"a", 0.5},
			},
			ordered: []string{"a", "z", "c"},
		},

		// update one item
		{
			name: "udpate_one_item",
			items: []item{
				item{"z", 0.7},
				item{"c", 0.9},
				item{"a", 0.5},
			},
			toUpdate: []item{
				item{"z", 0.4},
			},
			ordered: []string{"z", "a", "c"},
		},
	}

	for _, tc := range testcases {
		t.Run(
			tc.name,
			func(t *testing.T) {
				q := newPriorityQueue(make(distanceMap))

				// fill the priority queue
				for _, itm := range tc.items {
					q.push(itm.name, itm.distance)
				}
				// update items in the priority queue
				for _, itm := range tc.toUpdate {
					q.update(itm.name, itm.distance)
				}

				// pop all elements
				popped := make([]string, 0, q.Len())
				for q.Len() > 0 {
					name := q.pop()
					popped = append(popped, name)
				}

				// compare with expected popped elements
				if len(popped) != len(tc.ordered) {
					t.Errorf("wrong number of elements, %v instead of %v", len(popped), len(tc.ordered))
				}
				for i := range popped {
					if popped[i] != tc.ordered[i] {
						t.Errorf("wrong %d-th element: %v instead of %v", i, popped[i], tc.ordered[i])
					}
				}
			},
		)
	}

}
