from flask import Blueprint, jsonify, request
from ..models import User , db
import hashlib
bp = Blueprint("register", __name__, url_prefix="/register")




@bp.route( "" , methods = ["POST"] )
def create_user():
    
    #get data from form    
    username = request.form.get("username","")
    password = request.form.get("password","")
    email = request.form.get("email","")
    
    #check if value is missing
    if username is None or username.isspace() or username == "":
        return {"Status":"Failed","Message":"username missing"}, 400
    if password is None or password.isspace() or password == "":
        return {"Status":"Failed","Message":"password missing"}, 400
    if email is None or email.isspace() or email == "":
        return {"Status":"Failed","Message":"email missing"}, 400

    #hashing password
    hash_password = hashlib.sha256(password.encode()).hexdigest()
    #username is unique
    try :
        new_user = User()
        new_user.username = username
        new_user.password = hash_password
        new_user.email = email
        db.session.add(new_user)
        db.session.commit()
    except Exception as e:
        return {"Status":"Failed","Message":"username already exists"},200
    
    data = {"username":username,"password":hash_password,"email":email}
    return data , 200