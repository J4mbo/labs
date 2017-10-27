import json
import pandas as pd
import numpy as np
import matplotlib.pyplot as plt

json_data = {}
# csv_data = pd.read_csv('cpu_usage2.csv')

with open('./data', 'r') as file:
    json_data = file.read()
    json_data = json.loads(json_data)

# print("======cpu_usage data======")
# print("file length : " + str(len(csv_data['value'])))

plt.figure(figsize=[100,3])
plt.subplot(2,1,1)
plt.plot(json_data['value1'])
plt.subplot(2,1,2)
plt.plot(json_data['value2'])
plt.show()