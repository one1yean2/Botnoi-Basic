import queue
import threading

values =(        
        [10,10],
        [5,100],
        [40,2],
        [70,30]
)
def main():
    q = queue.Queue()
    
    
    for v in values:
        q.put(v)
    
    threads: list[threading.Thread] = []
    for _ in range(2):
        t = threading.Thread(target=sum_task, args=(q,))
        t.start()
        threads.append(t)
    
    for t in threads: 
        t.join()
    
    
    return

def sum_task(q: queue.Queue):
    t = threading.current_thread()
    
    while not q.empty():
    
        if q.empty():
            return
        
        values = q.get()
        if not isinstance(values, list):
            raise TypeError
        
        if len(values) < 2:
            raise ValueError
        
        print(t.name, values[0] + values[1])
    
def sum(num1:int,num2:int):
    return num1 + num2

if __name__ == '__main__':
    main()