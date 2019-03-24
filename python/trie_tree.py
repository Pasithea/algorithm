"""
    Trie tree.
"""


class TrieNode:
    def __init__(self, data: str):
        self.data = data
        self.children = [None] * 26
        self.is_ending_char = False


class TrieTree:
    def __init__(self):
        self._root = TrieNode('/')

    def insert(self, word: str) -> None:
        p = self._root
        a_index = ord('a')
        for i in range(len(word)):
            index = ord(word[i]) - a_index
            if p.children[index] is None:
                np = TrieNode(word[i])
                p.children[index] = np
            p = p.children[index]
        p.is_ending_char = True

    def find(self, pattern: str) -> bool:
        p = self._root
        a_index = ord('a')
        for i in range(len(pattern)):
            index = ord(pattern[i]) - a_index
            if p.children[index] is None:
                return False
            p = p.children[index]
        return p.is_ending_char


if __name__ == '__main__':
    t = TrieTree()
    text = ['work', 'audio', 'zodic', 'element', 'world']
    for w in text:
        t.insert(w)
    print(t.find('work'), t.find('audi'), t.find('zodicx'), t.find('ele'), t.find('world'))
