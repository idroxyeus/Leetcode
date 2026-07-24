MOD = 10**9 + 7
class Solution:
    def zigZagArrays(self, n: int, l: int, r: int) -> int:
        m = r - l + 1
        size = 2 * m
        
        def mat_mult(A: list[list[int]], B: list[list[int]]) -> list[list[int]]:
            res = [[0] * size for _ in range(size)]
            for i in range(size):
                for k in range(size):
                    if not A[i][k]:
                        continue
                    for j in range(size):
                        res[i][j] = (res[i][j] + A[i][k] * B[k][j]) % MOD
            return res

        def mat_pow(A: list[list[int]], p: int) -> list[list[int]]:
            res = [[1 if i == j else 0 for j in range(size)] for i in range(size)]
            base = A
            while p > 0:
                if p & 1:
                    res = mat_mult(res, base)
                base = mat_mult(base, base)
                p >>= 1
            return res

        T = [[0] * size for _ in range(size)]
        
        for x in range(m):
            down_state = x
            up_state = x + m
            
            for y in range(x + 1, m):
                T[y][up_state] = 1
                
            for y in range(x):
                T[y + m][down_state] = 1

        Tn = mat_pow(T, n - 1)
        start = [1] * size
        
        ans = 0
        for i in range(size):
            vec_val = sum(Tn[i][j] * start[j] for j in range(size)) % MOD
            ans = (ans + vec_val) % MOD
            
        return ans