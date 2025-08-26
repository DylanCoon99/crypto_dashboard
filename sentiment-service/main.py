from fastapi import FastAPI
from utility.sentiment_functions import AnalyzeSentiment
from utility.cloud_functions import read_folder_from_bucket
from datetime import datetime, timezone 

app = FastAPI()


@app.get("/sentiment/{coin_name}")
async def root(coin_name: str):
	try:
		folder_name = f"{coin_name}_digested_data/"

		data = read_folder_from_bucket(folder_name)

		sentiment = AnalyzeSentiment(data)

		sentiment_label = (
			"Positive" if sentiment["compound"] > 0.05
			else "Negative" if sentiment["compound"] < -0.05
			else "Neutral"
		)

		current_utc_time = datetime.now(timezone.utc)

		return {
			"coin_name": coin_name,
			"sentiment_score": sentiment,
			"sentiment_label": sentiment_label,
			"time": current_utc_time
		}
	except Exception as e:
		print(f"An error occured: {e}")
		raise HTTPException(
			status_code=status.HTTP_500_INTERNAL_SERVER_ERROR,
			detail=f"An unexpected internal server error occured on {datetime.now(timezone.utc)}."
		)



