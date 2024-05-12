from kafka import KafkaProducer
import json

producer = KafkaProducer(bootstrap_servers='localhost:9092')

producer.send('test3', json.dumps({'hello': 'world'}).encode('utf-8'))
producer.send('test2', json.dumps({'hello': 'world'}).encode('utf-8'))
producer.flush()