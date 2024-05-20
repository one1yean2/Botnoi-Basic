from locust import HttpUser,task,between,TaskSet

class MySet(TaskSet):
    
    @task(100)
    def test(self):
        print("executing my_task")
        
    @task(20)
    def test2(self):
        print("executing my_task2")
    @task(50)
    def test3(self):
        print("executing my_task3")
    @task(30)
    def test4(self):
        print("executing my_task4")
    
class MySet2(TaskSet):
    @task
    def get_route(self):
        res = self.client.get("/")
        print(res.status_code)
    @task
    def post_route(self):
        res = self.client.post("/test",data={"name":"test"})
        print(res.status_code)
        
class Users(HttpUser):
    tasks = [MySet2]
    wait_time = between(1, 5)