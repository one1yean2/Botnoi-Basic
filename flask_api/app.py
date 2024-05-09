from flask import Flask, jsonify , request

app = Flask(__name__)

data = [
        {
            "id": 1,
            "frameworks": "Django",
            "year": 2005
        },
        {
            "id": 2,
            "frameworks": "Flask",
            "year": 2010
        },
        {
            "id": 3,
            "frameworks": "Web2Py",
            "year": 2007
        }
    ]

@app.route('/')
def home():
    return "Hello World"
@app.route('/api/param', methods=['POST'])
def print_param():
    print("Args",request.args)
    return jsonify(request.args), 200
@app.route('/api/body', methods=['POST'])
def print_body():
    print("Body",request.json)
    return jsonify(request.json), 200
@app.route('/api/header', methods=['POST'])
def print_headers():
    print("Header",request.headers)
    return jsonify('success: print in terminal'), 200
@app.route('/api/form', methods=['POST'])
def print_form():
    print("Form",request.form)
    return jsonify(request.form), 200
@app.route('/api', methods=['GET'])
def get_api():
    return jsonify(data), 200



if __name__ == "__main__":
    app.run(debug=True)