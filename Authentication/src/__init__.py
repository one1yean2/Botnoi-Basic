from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from sqlalchemy.orm import DeclarativeBase

from .views import register

from .config import Config
class Base(DeclarativeBase):
    pass

db = SQLAlchemy(model_class=Base)
def create_app() -> Flask:
    app = Flask(__name__)
    app.config.from_object(Config)
    #init extension
    db.init_app(app)
    
    # create db table
    from . import models as _
    with app.app_context():
        db.create_all()
    
    # register blueprint
    from .views import login , register
    app.register_blueprint(register.bp)
    app.register_blueprint(login.bp)
    
    
    for route in app.url_map.iter_rules():
        print(f"{route.methods} {route}")
    
    return app 