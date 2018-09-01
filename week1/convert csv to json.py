import csv
import random
import json

csvfile = open('./Domestic Exchange.csv')
reader = csv.DictReader(csvfile)
nodes = []
edges = []
nodeDict = {}
for row in reader:
    colorgen =  "%06x" % random.randint(0, 0xFFFFFF)
    if row['ASN'] not in nodeDict:
        node = {
            "color": '#'+colorgen,
            "label": row['Name'],
            "attributes": {},
            "x": random.randint(-1000,1000),
            "y": random.randint(-1000,1000),
            "id": row['ASN'],
        }
        nodes.append(node)
        nodeDict[row['ASN']] = 1
    else:
        nodeDict[row['ASN']] += 1
    if row['ASN-source'] not in nodeDict:
        nodeDict[row['ASN-source']] = 1
    else:
        nodeDict[row['ASN-source']] += 1
    edge = {
        "sourceID": row['ASN-source'],
        "attributes": {},
        "targetID": row['ASN'],
        "size": float(row['Bandwidth'])/10,
    }
    edges.append(edge)
for node in nodes:
    node["size"] = nodeDict[node['id']]*2
    
data = {
    "nodes": nodes,
    "edges": edges,
}
print(json.dumps(data))

