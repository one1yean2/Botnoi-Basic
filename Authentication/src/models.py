from typing import List

from sqlalchemy import Column, ForeignKey, Integer, String
from sqlalchemy.orm import Mapped, relationship

from . import db


class User(db.Model):      
    __tablename__ = "user"

    id = Column(
        Integer,
        index=True,
        nullable=False,
        primary_key=True
    )
    username = Column(String)
    password = Column(String)
    
    
    # major_id = Column(Integer, ForeignKey(Major.id))
    # subjects: Mapped[List["Subject"]] = relationship(
    #     "Subject",
    #     back_populates="major"
    # )
    # students: Mapped[List["Student"]] = relationship(
    #     "Student",
    #     back_populates="major"
    # )

