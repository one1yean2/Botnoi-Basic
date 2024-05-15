from flask import Blueprint

bp = Blueprint("login", __name__, url_prefix="/login")

@bp.route( "" , methods = ["POST"] )
def login():
    auth = request.authorization