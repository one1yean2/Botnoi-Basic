from locust import HttpUser,task,between,TaskSet,User,constant

class MySet(TaskSet):
    @task
    def test(self):
        print("THIS IS TEST")
        
    @task
    def test2(self):
        print("THIS IS TEST2")
        
    @task
    def test3(self):
        print("THIS IS TEST3")
    
class HttpSet(TaskSet):
    
    @task(100)
    def get_route(self):
        self.client.get(url="/")
        print("THIS IS GET")
    @task(20)
    def post_route(self):
        res = self.client.post(url="/test",data={"name":"test"})
        print(res.status_code)
        print("THIS IS POST")
        
    
class MyUser(HttpUser):
    tasks = [HttpSet]
    wait_time = between(1, 5)
