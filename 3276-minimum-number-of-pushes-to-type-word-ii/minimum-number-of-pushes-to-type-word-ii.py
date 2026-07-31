from collections import Counter

class Solution:
    def minimumPushes(self, word: str) -> int:
        char_counts = Counter(word)
        sorted_counts = sorted(char_counts.values(), reverse=True)
        
        total_pushes = 0
        for index, count in enumerate(sorted_counts):
            pushes_per_char = (index // 8) + 1
            total_pushes += count * pushes_per_char
            
        return total_pushes
