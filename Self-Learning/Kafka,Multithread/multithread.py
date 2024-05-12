import threading
import time

done = False 

def worker():
    counter = 0
    while True:
        time.sleep(1)
        counter += 1
        print(counter)    
        
        
threading.Thread(target=worker ,daemon=True).start()
threading.Thread(target=worker ,daemon=False).start()



input("Press Enter to quit...")
done = True
print("you did it")