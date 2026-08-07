class Solution:
    def smallestNumber(self, num: str, t: int) -> str:
        primes = [2, 3, 5, 7]
        req_p = [0] * 4
        for i, p in enumerate(primes):
            while t % p == 0:
                req_p[i] += 1
                t //= p
        
        if t > 1:
            return "-1"

        def get_min_digits(p_counts):
            c2, c3, c5, c7 = p_counts
            cnt = [0] * 10
            
            cnt[5] = max(0, c5)
            cnt[7] = max(0, c7)
            
            r3 = max(0, c3)
            cnt[9] = r3 // 2
            r3 %= 2
            
            r2 = max(0, c2)
            cnt[8] = r2 // 3
            r2 %= 3
            
            if r3 == 1 and r2 == 2:
                cnt[6] += 1
                cnt[2] += 1
            elif r3 == 1 and r2 == 1:
                cnt[6] += 1
            elif r3 == 1 and r2 == 0:
                cnt[3] += 1
            elif r3 == 0 and r2 == 2:
                cnt[4] += 1
            elif r3 == 0 and r2 == 1:
                cnt[2] += 1
                
            return cnt

        def build_suffix(cnt):
            return "".join(str(d) * cnt[d] for d in range(2, 10))

        init_digits = get_min_digits(req_p)
        if sum(init_digits) > len(num):
            return build_suffix(init_digits)

        n = len(num)
        prefix_p = [0] * 4
        first_zero = n
        
        for i, ch in enumerate(num):
            if ch == '0':
                first_zero = i
                break
            d = int(ch)
            for idx, p in enumerate(primes):
                while d % p == 0:
                    prefix_p[idx] += 1
                    d //= p

        if first_zero == n and all(prefix_p[i] >= req_p[i] for i in range(4)):
            return num

        for i in range(n - 1, -1, -1):
            if i > first_zero:
                continue
            
            current_digit = int(num[i])
            d_val = current_digit
            if d_val > 0:
                for idx, p in enumerate(primes):
                    while d_val % p == 0:
                        prefix_p[idx] -= 1
                        d_val //= p
            
            for nxt in range(current_digit + 1, 10):
                nxt_p = list(prefix_p)
                temp = nxt
                for idx, p in enumerate(primes):
                    while temp % p == 0:
                        nxt_p[idx] += 1
                        temp //= p
                
                needed_p = [max(0, req_p[k] - nxt_p[k]) for k in range(4)]
                suffix_cnt = get_min_digits(needed_p)
                suffix_len = sum(suffix_cnt)
                available_slots = n - 1 - i
                
                if suffix_len <= available_slots:
                    ones_padding = available_slots - suffix_len
                    return num[:i] + str(nxt) + "1" * ones_padding + build_suffix(suffix_cnt)

        final_digits = get_min_digits(req_p)
        ones_padding = (n + 1) - sum(final_digits)
        return "1" * ones_padding + build_suffix(final_digits)
