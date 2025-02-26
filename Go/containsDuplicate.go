func containsDuplicate(nums []int) bool {
    set := NewSet()

    for _, num := range nums {
        if(set.Has(num)) {
            return true;
        }

        set.Add(num)
    }

    return false
}

type Set struct {
    items map[int]struct{}
}

func NewSet() *Set {
    return &Set{
        items: make(map[int]struct{}),
    }
}

func (s *Set) Has(val int) bool {
    _, ok := s.items[val]

    return ok;
}

func (s *Set) Add(val int) {
    s.items[val] = struct{}{}
}
