from flask import Flask
from redis import Redis
from flask_sqlalchemy import SQLAlchemy
from sqlalchemy.orm import DeclarativeBase
from flask_cors import CORS
from flask_mail import Mail
from .config import Config
import os
class Base(DeclarativeBase):
    pass
redis_host = os.getenv('REDIS_HOST', 'localhost')
redis_port = os.getenv('REDIS_PORT', 6379)
db = SQLAlchemy(model_class=Base)
redis_cache = Redis(host=redis_host,port=redis_port, decode_responses=True)
mail = Mail()
cors = CORS()


def create_app() -> Flask:
    app = Flask(__name__)
    app.config.from_object(Config)
    #init extension
    db.init_app(app)
    # redis_cache.init_app(app)
    mail.init_app(app)
    cors.init_app(app)
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