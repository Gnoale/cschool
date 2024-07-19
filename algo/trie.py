class Trie:

    def __init__(self):
        self.child = [None]*26
        self.end = False

    def insert(self, word: str) -> None:
        current = self
        for c in word:
            char = c.lower()
            idx = ord(char) - ord("a")
            if current.child[idx] is None:
                current.child[idx] = Trie()
            current = current.child[idx]
        current.end = True

    def search(self, word: str) -> bool:
        current = self
        for c in word:
            char = c.lower()
            idx = ord(char) - ord("a")
            if current.child[idx] is None:
                return False
            current = current.child[idx]
        return current.end

    def startsWith(self, prefix: str) -> bool:
        current = self
        for c in prefix:
            char = c.lower()
            idx = ord(char) - ord("a")
            if current.child[idx] is None:
                return False
            current = current.child[idx]

        return True


if __name__ == "__main__":

    words = ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]


    trie = Trie()
    trie.insert("apple")
    print(trie.search("apple"))
    print(trie.search("app"))
    print(trie.startsWith("app"))
    trie.insert("app")
    print(trie.search("app"))
