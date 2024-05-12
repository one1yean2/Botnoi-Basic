from kafka import KafkaProducer
import json

producer = KafkaProducer(bootstrap_servers='localhost:9092')
