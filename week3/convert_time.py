import csv
import random
import json
from datetime import datetime

csvfile = open('./Timestamp.csv')
reader = csv.DictReader(csvfile)
time = []
for row in reader:
        if row['login']:
            timelogin = int(row['login'])/1000.00
        print(datetime.fromtimestamp(timelogin).strftime('%Y-%m-%d %H:%M:%S'), end =" ")
        if row['logout'] == '-' or not row['logout']:
            print('- -', end =" ")
        else:  
            timelogout = int(row['logout'])/1000.00
            print(datetime.fromtimestamp(timelogout).strftime('%Y-%m-%d %H:%M:%S'), end =" ")
        if row['lastseen']:
            timelastseen = int(row['lastseen'])/1000.00
        print(datetime.fromtimestamp(timelastseen).strftime('%Y-%m-%d %H:%M:%S'))

