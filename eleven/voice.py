import os
import uuid
from elevenlabs import VoiceSettings
from elevenlabs.client import ElevenLabs

ELEVENLABS_API_KEY = os.getenv("ELEVENLABS_API_KEY")
client = ElevenLabs(
    api_key=ELEVENLABS_API_KEY,
)


def text_to_speech_file(text: str) -> str:
    response = client.text_to_speech.convert(
        voice_id="XjLkpWUlnhS8i7gGz3lZ", # tiggy
        output_format="mp3_22050_32",
        text=text,
        model_id="eleven_turbo_v2_5", # use the turbo model for low latency
        voice_settings=VoiceSettings(
            stability=0.0,
            similarity_boost=1.0,
            style=0.0,
            use_speaker_boost=True,
        ),
    )

    # uncomment the line below to play the audio back
    # play(response)

    save_file_path = f"third.mp3"

    with open(save_file_path, "wb") as f:
        for chunk in response:
            if chunk:
                f.write(chunk)

    print(f"{save_file_path}: A new audio file was saved successfully!")

    return save_file_path


text_to_speech_file("<speak>You can re-write your nervous system. That's right, maybe you don't like when your nervous system is stressed. Dr Leah Lagos and Chris Williamson discus how to break free from your default reptillian brain programming. And go from this <break time=\"1.5s\" />  to this.</speak>")

