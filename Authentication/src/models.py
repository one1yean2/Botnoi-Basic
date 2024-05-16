from sqlalchemy import Column, Integer, String

from . import db


class User(db.Model):      
    __tablename__ = "user"

    id = Column(
        Integer,
        index=True,
        nullable=False,
        primary_key=True
    )
    username = Column(String,unique=True)
    password = Column(String)
    email = Column(String)
    token = Column(String)
    
