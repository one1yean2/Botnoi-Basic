from concurrent.futures import ThreadPoolExecutor

def sum(num1, num2):
    if not isinstance(num1, int) or not isinstance(num2, int):
        raise ValueError
    
    return num1 + num2


with ThreadPoolExecutor(max_workers=2) as executor:
    future = executor.submit(sum, 5, 10)
    f2 = executor.submit(sum, 1, '10')
    f3 = executor.submit(sum, 2, 10)
    f4 = executor.submit(sum, 3, 10)
    f2_ex = f2.exception()
    if f2_ex is None:
        print(f2.result())
    fs = executor.map(sum, [1, 2, 3], [4, 5, 6])
    for f in fs:
        print(f)
    
    # print(future.result())
    # # print(f2.result())
    # print(f3.result())
    # print(f4.result())
    executor.shutdown()


# P'Tee Part
# Context manager
with ThreadPoolExecutor(max_workers=2) as ex:
    f1 = ex.submit(sum, 10, 5)
    f2 = ex.submit(sum, 60, "1")

    # Check future object for any raised exception before access result.
    f1_ex = f1.exception()
    if f1_ex is None:
        print(f1.result())

    f2_ex = f2.exception()
    if f2_ex is None:
        print(f2.result())

    # Map list of parameter to one function
    fs = ex.map(sum, [1, 5], [1, 4])
    for f in fs:
        print(f)

# Thread pool object
ex = ThreadPoolExecutor(max_workers=2)
f = ex.submit(sum, 10, 5)
f.result()
ex.shutdown()