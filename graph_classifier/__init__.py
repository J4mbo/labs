import json

data = {}

with open("./data", 'r') as file:
    data = file.read()
    data = json.loads(data)

print(data.get('data1'))