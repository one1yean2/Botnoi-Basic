from flask import Blueprint, jsonify

bp = Blueprint("register", __name__, url_prefix="/register")


@bp.route( "" , methods = ["GET"] )
def register():
    return jsonify("PLEASE USE POST METHOD MY MAN"), 200

@bp.route( "" , methods = ["POST"] )
def create_register_request():
    return "OK"


