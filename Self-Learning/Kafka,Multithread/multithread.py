from threading import Thread
import time
from turtle import delay
# from concurrent_futures import ThreadPoolExecutor

done = False 

x=0
def increase(number,counter):
    global x
    
    time.sleep(0.5)
    for _ in range(counter):
        x += number
        # print(x)



class MyThread(Thread):

    def __init__(self, ThreadID):
        
        Thread.__init__(self)
        self.ThreadID = ThreadID
        
    def run(self):
        print("Hello, I am thread " + str(self.ThreadID))
def workerA():
    while True:
        print("worker A  3 sec")
        time.sleep(3)
def workerB():
    while True:
        print("worker B 2 sec")
        time.sleep(2)

# thr1 = MyThread(1).start()
# thr3 = Thread(target=workerB ,daemon=True)
# thr3.start()
# # thr3.join()

# thr2 = Thread(target=workerA ,daemon=True).start()
def main():
    global x
    x = 0
    t4 = Thread(target=increase, args=(5,10))
    t5 = Thread(target=increase, args=(6,5))

    t4.start()
    t5.start()

    t4.join()
    t5.join()

    # t4.start()
for i in range(5):
    main()
    print('x',x)