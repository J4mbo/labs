import pymongo
import datetime

# DB connection
client = pymongo.MongoClient('mongodb://192.168.0.104:27017/')
db = client['REAL']
collection = db['user']


user = {
    'name': 'brown',
    'passwd': 'linePlus',
    'date': datetime.datetime.now()
}

# insert user
# collection.insert(user)

## count
# collection.count()

## find user
# collection.find(query)

for user in collection.find().sort('name'):
    print(user)
