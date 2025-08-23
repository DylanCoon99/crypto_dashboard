from fastapi import FastAPI

app = FastAPI()


'''




'''


@app.get("/sentiment/{coin_name}")
async def root(coin_name):
	return {"message": "Hello World"}
