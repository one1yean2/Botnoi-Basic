from locust import HttpUser,task,between,TaskSet,User 


class MySet(TaskSet):
    @task
    def test(self):
        self.client.get(url="/")
    
    @task
    def test2(self):
        self.client.get(url="/")
        
class WebsiteUser(HttpUser):
    host = "http://127.0.0.1:5000"
    wait_time = between(1, 5)
    tasks = [MySet]
    @task
    def test(self):
        # self.client.get(url="/")
    
        self.client.post(url="/test",data={"name":"test"})
        
    @task 
    def test2(self):
        self.client.get(url="/")
        
    
        
        
        
        
        
        
        
    # def register(self):
    #     username = "testaaaaaa" 
    #     password = "test" 
    #     # x +=1
    #     self.client.post(url="/register",
    #     data={
    #         "username":username,
    #         "password":password,
    #         "email":"one1yean2@gmail.com"
    #     })
        
    # @task 
    # def login(self):
    #     self.client.get(url="/login")
        