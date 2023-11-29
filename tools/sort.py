import json
import sys
import re

input_file = sys.argv[1]
output_file = sys.argv[2]

def split_id(id):
    match = re.match(r'([A-Za-z]+)([0-9]+)', id)
    if match:
        return match.group(1), int(match.group(2))
    return '', 0


with open(input_file, 'r', encoding='utf-8') as file:
    data = json.load(file)

sorted_data = sorted(data, key=lambda x: split_id(x['id']))

with open(output_file, 'w', encoding='utf-8') as file:
    json.dump(sorted_data, file, ensure_ascii=False, indent=4)
