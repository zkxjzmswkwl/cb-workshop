# Author: Carter
# Description: Finds every "Troll chucker" entity and kills them.
# Revision Date: 4/25/2024
from time import sleep
from cockbot5 import console
from cockbot5 import scene

# External Python libraries can be utilized. woo.
import requests


def log(message):
    console.writeln(
        f"<img=7> <img=8> <col=0xd74000>[COCKBOT5]:</col> <col=FFFFFF>{message}</col>"
    )


def webhook(message):
    r = requests.post(
        "https://canary.discord.com/api/webhooks/removed",
        json={
            "content": message,
            "username": "CB REV. 5",
            "avatar_url": "https://pbs.twimg.com/profile_images/1652097774475952128/hOYRq8xi_400x400.jpg",
        },
        headers={"Content-Type": "application/json"},
    )


class Script:
    def __init__(self):
        log("Script __init__")
        self._load_entities()

    def _load_entities(self):
        # Get a list of all entities within the scene.
        self._entities = scene.get_entities()
        # List comprehension to filter out entities that aren't Troll chuckers.
        self._entities = [e for e in self._entities if "chucker" in e.get_name()]
        for e in self._entities:
            log(f"Found {e.get_name()}:{e.get_id()}")
        # pop() removes the last element from the list and returns it.
        # we call this here to start the killing process, since on_exp_drop() needs to be called to start the murder.
        self._entities.pop().attack()

    def on_server_tick(self):
        """
        This function is called every server tick (600ms).
        """
        log("on_server_tick")

    def on_exp_drop(self):
        """
        This function is called whenever the player receives experience.
        """
        log("on_exp_drop")
        # If we're out of Troll chuckers to kill, get more.
        if len(self._entities) == 0:
            self._load_entities()
        # If not, kill them.
        next_troll = self._entities.pop()
        next_troll.attack()
        webhook(f"Killed {next_troll.get_name()} (Lv. {next_troll.get_combat_level()})")
