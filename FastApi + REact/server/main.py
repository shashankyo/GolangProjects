from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

from pydantic import BaseModel
from typing import Union
import requests

app = FastAPI()
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

PROJECT_ID = "f4b8a3cc-a808-4874-b4c4-41122a2b6fba"
PRIVATE_KEY = "bd989bb0-6c17-4b80-a2ba-3b15bff53bbd" 

class User(BaseModel):
    username: str
    secret: str
    email: Union[str, None] = None
    first_name: Union[str, None] = None
    last_name: Union[str, None] = None

@app.post("/login")
async def root(user: User):
    response = requests.get("https://api.chatengine.io/users/me/",
        headers={
            "Project-ID": PROJECT_ID,
            "User-Name": user.username,
            "User-Secret": user.secret
        }
    )
    return response.json()

@app.post("/signup/")
async def root(user: User):
    response = requests.post("https://api.chatengine.io/users",
        data={
            "username": user.username,
            "secret": user.secret,
            "email": user.email,
            "first_name": user.first_name,
            "last_name": user.last_name
        },
       
        headers={"Private-Key": PRIVATE_KEY}
    )
    return response.json