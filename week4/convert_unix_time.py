import csv
import json
from datetime import datetime
# from time import strftime

csvfile = open('./realtimetop.csv', 'r')
reader = csv.DictReader(csvfile)
time = []
for row in reader:
    a = int(row['time'])/1000000.00
    
    print(datetime.fromtimestamp(a).strftime('%Y-%m-%d %H:%M:%S'), end ="\n")
