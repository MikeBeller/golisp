
def half_adder(a, b):
    return a ^ b, a & b

assert half_adder(0, 0) == (0, 0)
assert half_adder(0, 1) == (1, 0)
assert half_adder(1, 0) == (1, 0)
assert half_adder(1, 1) == (0, 1)

def full_adder(a, b, ci):
    s,c1 = half_adder(a, b)
    sm,c2 = half_adder(s, ci)
    return sm,c1 | c2

assert full_adder(0, 0, 0) == (0, 0)
assert full_adder(0, 0, 1) == (1, 0)
assert full_adder(0, 1, 0) == (1, 0)
assert full_adder(0, 1, 1) == (0, 1)
assert full_adder(1, 0, 0) == (1, 0)
assert full_adder(1, 0, 1) == (0, 1)
assert full_adder(1, 1, 0) == (0, 1)
assert full_adder(1, 1, 1) == (1, 1)

def add_nums(xs, ys):
    zs = []
    c = 0
    for i in range(max(len(xs), len(ys))):
        x = xs[i] if i < len(xs) else 0
        y = ys[i] if i < len(ys) else 0
        z,c = full_adder(x, y, c)
        zs.append(z)
    if c == 1:
        zs.append(c)
    return zs

import random

def test_add():
    for i in range(100):
        a = random.randint(0, 1000)
        b = random.randint(0, 1000)
        abin = to_binary(a)
        bbin = to_binary(b)
        cbin = add_nums(abin, bbin)
        c = to_integer(cbin)
        #print(a, b, c, abin, bbin, cbin)
        assert c == a + b

def to_binary(n):
    return [int(d) for d in reversed(bin(n)[2:])]

def to_integer(bs):
    r = 0
    for b in reversed(bs):
        r *= 2
        r += b
    return r


assert add_nums(to_binary(10), to_binary(24)) == to_binary(34)
assert to_integer(to_binary(34)) == 34

test_add()

