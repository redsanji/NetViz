import csv
import random
import json
import datetime

csvfile = open('realtimetop.csv', 'r')
tree_jsonfile = open('realtimetop.json', 'w')
reader = csv.DictReader(csvfile,  delimiter=',', fieldnames=['date', 'user', 'IPsrc', 'IPdes', 'Port', 'hostname'])
js = []
for row in reader:
    data = {}
    data['date'] = row['date']
    data['user'] = row['user']
    data['IPsrc'] = row['IPsrc']
    data['IPdes'] = row['IPdes']
    data['Port'] = row['Port']
    data['hostname'] = row['hostname']
    js.append(data)

json.dump(js, tree_jsonfile)