from kafka import KafkaConsumer
import json
group_id = "test"

consumer = KafkaConsumer(
    'test',
    bootstrap_servers=['0.0.0.0:9092'],
    group_id=group_id
)

consumer.subscribe(['test3', 'test2'])
while True:
    for message in consumer:
        print(message)
