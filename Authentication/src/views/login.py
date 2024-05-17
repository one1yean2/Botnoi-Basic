import hashlib
import random
from flask import Blueprint,  request,current_app
from sqlalchemy import update
from ..models import User , db
from flask_mail import Message
from .. import mail , redis_cache 
from ..config import Config
import jwt
import threading

bp = Blueprint("login", __name__, url_prefix="/login")




@bp.route( "/token" , methods = ["GET"] )
def token():
    try:
        headers =request.headers
        bearer = headers.get('Authorization')    # Bearer YourTokenHere
        print(bearer)
        token = bearer.split()[1]
        print(token)
        user = User.query.filter_by(token = token).first()
        if user.token != token:
            return {"Status":"Failed","Message":"Unauthorized"}, 401
        else:
            return {"Status":"Success","Message":"Authorized"}, 200 
    except Exception as e:
        return {"Status":"Failed","Message":"No token"}, 400
    


@bp.route( "/otp" , methods = ["POST"] )
def user_otp():
    
    #get user otp from form
    username = request.form.get("username","")
    user_otp = request.form.get("otp","")
    #get otp from cache
    otp = redis_cache.get("OTP"+username)
    print(otp)
    
    #check if value is missing
    if username is None or username.isspace() or username == "":
        return {"Status":"Failed","Message":"username Missing"},  400
    if user_otp is None or user_otp.isspace() or user_otp == "":
        return {"Status":"Failed","Message":"OTP Missing"},  400
    if otp is None:
        return {"Status":"Failed","Message":"OTP Timed out"}, 400
    if user_otp != otp:
        return {"Status":"Failed","Message":"OTP mismatch"}, 400
    
    # TODO : gen token
    token = jwt.encode({"username":username,"userOTP":user_otp},Config.SECRET_KEY)
    
    #update user token in database
    db.session.execute(update(User).where(User.username==username).values({
        "token": token
    }))
    db.session.commit()

    return {'Status':'Success','token' : token}, 200

@bp.route( "" , methods = ["POST"] )
def user_login():
    
    #get data from form    
    username = request.form.get("username","")
    password = request.form.get("password","")
    
    #check if value is missing
    if username is None or username.isspace() or username == "":
        return {"Status":"Failed","Message":"username Missing"},  400
    if password is None or password.isspace() or password == "":
        return {"Status":"Failed","Message":"password Missing"},  400

    hash_password = hashlib.sha256(password.encode()).hexdigest()
    #username is unique
    try:
        user = User.query.filter_by(username = username).first()
        print(user)
        if user is None:
            return {"Status":"Failed","Message":"user not found please register"},  400
        if user.password != hash_password:
            return {"Status":"Failed","Message":"password mismatch"},  400
    
    except Exception as e:
        return {"Status":"Failed"},  400
    
    
    #otp generator
    otp = otp_generator()   
    #send otp via email
    app = current_app._get_current_object()
    t1 = threading.Thread(target=sending_otp_via_email,args=(app,user.email,otp))
    t1.start()
    # t1.join()
    # sending_otp_via_email(user.email,otp)
    #set otp in cache
    redis_cache.set("OTP"+user.username,otp) 
    redis_cache.expire("OTP"+user.username,60)
    print("SET > OTP"+user.username,otp)
    
    return {"Status":"Success","Message":"please check your email for OTP"}, 200

def sending_otp_via_email(app,email,otp):
    with app.app_context():
        msg_title = "[OTP-Yean]"
        sender = "noreply@gmail.com"
        msg = Message(msg_title, sender=sender,recipients=[email])
        msg.body = "YOUR OTP IS : "+ otp
        mail.send(msg)

def otp_generator():
    num = f'{random.randrange(1, 10**6):06}'
    return num
