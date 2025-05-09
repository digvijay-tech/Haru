import time


start = time.time()

i = 1

while i < 100000:
    print(i)
    i = i + 1

end = time.time()
print(f"Python took: {(end - start) * 1000:.3f} ms")
