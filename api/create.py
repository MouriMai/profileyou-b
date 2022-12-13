import requests
import sys

args = sys.argv

print(args[1])

print(args[1])

# ---------------- OPENAI -----------------------

url = "https://api.openai.com/v1/images/generations"

OPENAI_API_KEY="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"

payload = {
  "prompt": args[1],
  "n": 1,
  "size": '512x512'
}
# size must be one ogf ['256x256', '512x512', '1024x1024']

headers = {
    'Content-Type': 'application/json',
    'Authorization': 'Bearer {}'.format(OPENAI_API_KEY)
}

response = requests.post(url, headers=headers, json=payload)

print(response.text)

# image_url = response['data'][0]['url']
# print(image_url)

# {
#   "prompt": "A cute baby sea otter",
#   "n": 2,
#   "size": "1024x1024"
# }