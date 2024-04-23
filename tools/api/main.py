from typing import Union

from fastapi import FastAPI
import os

app = FastAPI()


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
    header["file"] = file.split("/")[-1]
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


@app.get("/")
def read_root():
    return {"Hello": "World"}


@app.get("/items/{item_id}")
def read_item(item_id: int, q: Union[str, None] = None):
    return {"item_id": item_id, "q": q}