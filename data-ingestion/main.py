from fastapi import FastAPI, HTTPException, status
from utility.api_functions import get_articles, clean_local_directory
from utility.cloud_functions import upload_folder_to_bucket
from datetime import datetime, timezone 

app = FastAPI()

@app.get("/ingest/{coin_name}")
async def root(coin_name: str):

	try:
		source_folder_name = "./api_data/"

		## ingests api data
		get_articles(coin_name, source_folder_name) # working

		## uploads api data to cloud storage bucket
		
		destination_folder_name = "bitcoin_digested_data/"
		upload_folder_to_bucket(source_folder_name + coin_name + "/", destination_folder_name)

		## clean local api directory
		# ./api_data/
		clean_local_directory(coin_name, source_folder_name)

		return {"message": f"Successfully digested data on {datetime.now(timezone.utc)}"}
	except Exception as e:
		print(f"An error occured: {e}")
		raise HTTPException(
			status_code=status.HTTP_500_INTERNAL_SERVER_ERROR,
			detail=f"An unexpected internal server error occured on {datetime.now(timezone.utc)}."
		)


@app.get("/health")
async def health_check():
    return {"status": "healthy"}