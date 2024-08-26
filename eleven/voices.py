import requests
import os

ELEVENLABS_API_KEY = os.getenv("ELEVENLABS_API_KEY")

url = "https://api.elevenlabs.io/v1/voices"

headers = {
  "Accept": "application/json",
  "xi-api-key": ELEVENLABS_API_KEY,
  "Content-Type": "application/json"
}

response = requests.get(url, headers=headers)

if response.status_code == 200:
    voices = response.json().get("voices", [])
    
    print("Available Voices:")
    for voice in voices:
        print(f"Name: {voice.get('name')}, ID: {voice.get('voice_id')}")
else:
    print(f"Failed to fetch voices. Status code: {response.status_code}, Response: {response.text}")

