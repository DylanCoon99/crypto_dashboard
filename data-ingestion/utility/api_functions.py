from newsapi import NewsApiClient
from dotenv import load_dotenv
import os

load_dotenv()

api = NewsApiClient(api_key=os.getenv("NEWS_API_KEY"))

top_headlines = api.get_top_headlines(sources='bbc-news')
search_results = api.get_everything(q='bitcoin', language='en')