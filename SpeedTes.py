import time
def fibo(n):
    if n == 0:
        return 0
    if n == 1:
        return 1
    return fibo(n - 1) + fibo(n - 2)


start = time.time()
print(fibo(45))
print("Time: ", time.time() - start)
