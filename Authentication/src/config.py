class Config(object):

    # Flask
    SECRET_KEY = "84b9ebfa0b5bab16df6601da5c35b0f8ebb5d062ebdddd14bff178960dfeea30"
    SERVER_NAME = "127.0.0.1:5000"
    # SERVER_NAME = "3ba5-171-103-249-230.ngrok-free.app"
    APPLICATION_ROOT = "/"
    PREFERRED_URL_SCHEME = "http"


    #Flask Mail
    MAIL_SERVER = "smtp.gmail.com"
    MAIL_PORT = 587
    MAIL_USERNAME = "yeeninw55@gmail.com"
    MAIL_PASSWORD = "gqdg lvgj eygn rpma"
    MAIL_USE_TLS = True
    
    # Flask Cors
    CORS_RESOURCES = {r"/*": {"origins": "*"}}
    CORS_SUPPORTS_CREDENTIALS = True

    # Flask SQLAlchemy
    SQLALCHEMY_DATABASE_URI = "sqlite:///demo.db"
