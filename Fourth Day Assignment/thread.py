import threading
import queue
x = 0


def main():
    global x
    x = 0

    # lock = threading.Lock()
    # t1 = threading.Thread(target=task, args=(lock, ))
    # t2 = threading.Thread(target=task, args=(lock, ))
    t1 = threading.Thread(target=task)
    t2 = threading.Thread(target=task)

    t1.start()
    t2.start()

    t1.join()
    t2.join()
    
    # q = queue.Queue()


def increase():
    global x
    x += 1


# def task(lock: threading.Lock):
def task():
    for _ in range(500000):
        # lock.acquire()
        increase()
        # lock.release()


if __name__ == '__main__':
    for i in range(5):
        main()
        print(f"Loop number {i + 1}, X = {x}")