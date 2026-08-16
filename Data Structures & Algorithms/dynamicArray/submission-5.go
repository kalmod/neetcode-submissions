type DynamicArray struct {
    capacity int
    size int
    array []int
}

func NewDynamicArray(capacity int) *DynamicArray {
    if capacity <= 0 {
        return nil
    }
    return &DynamicArray{capacity: capacity, size: 0, array: make([]int, capacity)}
}

func (da *DynamicArray) Get(i int) int {
    return da.array[i]
}

func (da *DynamicArray) Set(i int, n int) {

    da.array[i] = n
}

func (da *DynamicArray) Pushback(n int) {
    if da.size == da.capacity {
        da.Resize()
    }
    da.array[da.size] = n
    da.size++
}

func (da *DynamicArray) Popback() int {
    da.size--
    return da.array[da.size]
}

func (da *DynamicArray) Resize() {
    da.capacity *= 2
    newArray := make([]int, da.capacity)
    for i := 0; i < da.size; i++ {
        newArray[i] = da.array[i]
    }
    da.array = newArray
}

func (da *DynamicArray) GetSize() int {
    return da.size
}

func (da *DynamicArray) GetCapacity() int {
    return da.capacity
}
