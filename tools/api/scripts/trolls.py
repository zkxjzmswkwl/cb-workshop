# Author: Carter
# Description: Kills Troll Chuckers in Burthorpe for combat experience.
# Revision Date: 4/22/2024

from time import sleep
from lennissa import console, world
from lennissa.player import attack

def log(message):
    console.writeln(f"<img=7> <img=8> <col=0xd74000>[Lennissa]:</col> <col=e0a3ff>{message}</col>", 1)

class Script:
    def __init__(self):
        log("Script ctor")

        # Testing
        entities = world.get_entities()
        for e in entities:
            log(f"{e.get_name()}: Lv. {str(e.combatLevel)}")
        
        attack("Troll")

    def on_server_tick(self):
        pass
    
    def on_exp_drop(self):
        log("UpdateStat")
        sleep(2)
        attack("Troll")