from flask import Flask,request
# import threading

app = Flask(__name__)

@app.route('/',methods=['GET'])
def home():
    return "Hello"

@app.route('/test',methods=['POST'])
def test():
    name = request.form.get('name')
    print(name)
    return "test"

if __name__ == "__main__":
    app.run(debug=True)