from flask import Flask,request

app = Flask(__name__)


#Get Route
@app.route('/',methods=['GET'])
def home():
    return "GET METHOD"


@app.route('/test',methods=['POST'])
def test():
    name = request.form.get('name')
    print(name)
    return "POST METHOD"

if __name__ == "__main__":
    app.run(debug=True)