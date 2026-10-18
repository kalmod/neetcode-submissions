class MyHashSet:

    def __init__(self):
        self.h = [[] for _ in range(1009)]
        self.size = len(self.h)

    def add(self, key: int) -> None:
        for n in self.h[self._hash(key)]:
            if n == key:
                return 
        self.h[self._hash(key)].append(key) 

    def remove(self, key: int) -> None:
        for i in range(len(self.h[self._hash(key)])): 
            if self.h[self._hash(key)][i] == key:
                self.h[self._hash(key)].pop(i)

    def contains(self, key: int) -> bool:
        for n in self.h[self._hash(key)]:
            if n == key:
                return True
        return False
    
    def _hash(self, key: int) -> None:
        return key % self.size
        


# Your MyHashSet object will be instantiated and called as such:
# obj = MyHashSet()
# obj.add(key)
# obj.remove(key)
# param_3 = obj.contains(key)