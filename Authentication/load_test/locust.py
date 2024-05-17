import time
from locust import HttpUser,task,between

class WebsiteUser(HttpUser):
    wait_time = between(1, 5)
    # global x 
    # x = 0
    @task
    def register(self):
        global x
        username = "test" 
        password = "test" 
        # x +=1
        self.client.post(url="/register",
        data={
            "username":username,
            "password":password,
            "email":"one1yean2@gmail.com"
        })
        
    # @task 
    # def login(self):
    #     self.client.get(url="/login")
        