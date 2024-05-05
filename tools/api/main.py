from typing import Union
from fastapi import FastAPI
from fastapi.staticfiles import StaticFiles
import hashlib
import os

app = FastAPI()
app.mount("/static", StaticFiles(directory="static"), name="static")


def get_files_in_directory(directory: str) -> list[str]:
    files = []
    for file in os.listdir(directory):
        if os.path.isfile(os.path.join(directory, file)):
            files.append(file)
    return files


def parse_script_header(file: str) -> dict:
    tags = ["Author", "Description", "Revision Date"]

    with open(file, "r") as f:
        lines = f.readlines()
        header = {}
        for line in lines:
            if line.startswith("#") and any(substring in line for substring in tags):
                key, value = line.split(": ")
                header[key.replace("#", "").replace(" ", "")] = value.strip()
    header["File"] = file.split("/")[-1]
    return header


@app.get("/scripts")
def get_files():
    directory = "scripts/"
    files = get_files_in_directory(directory)
    ret = {"scripts": []}
    for f in files:
        ret["scripts"].append(parse_script_header(directory + f))
    return ret


@app.get("/scripts/{script}")
def get_script(script: str):
    with open(f"scripts/{script}", "r") as f:
        return {"script": f.read()}


@app.get("/module")
def get_module_hash():
    with open("static/COCKBOT.dll", "rb") as f:
        return {"module": hashlib.md5(f.read()).hexdigest()}
