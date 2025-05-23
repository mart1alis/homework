from collections import Counter

result = []

a = list(map(int, open('A.txt').read().split()))
b = list(map(int, open('B.txt').read().split()))

S = set(a) & set(b)
print(S)
count_a = Counter(a)
count_b = Counter(b)

result = []
for num in sorted(S):
    result += [num] * max(count_a[num], count_b[num])

open('C.txt', 'w').write(' '.join(map(str, result)))
