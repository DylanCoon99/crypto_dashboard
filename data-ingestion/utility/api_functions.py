from newsapi import NewsApiClient
from dotenv import load_dotenv
from datetime import date, timedelta, datetime, timezone
import json
import time
import os

load_dotenv()

api = NewsApiClient(api_key=os.getenv("NEWS_API_KEY"))


def get_articles(coin_name: str, n=5):
    # gets the top n (default = 5) articles for a coin the past week

    articles = api.get_everything(q=coin_name,
                                  language='en',
                                  from_param=str(date.today() - timedelta(days=7)),
                                  to=str(date.today()),
                                  sort_by='relevancy',
                                  page_size=n
    )

    # write the articles to the api_data/{coin_name} directory
    for article in articles["articles"]:

        ## build the json for this article

        article_dict = {
            "coin_name": coin_name,
            "ingested_date": str(datetime.now(timezone.utc)),
            "data": {
                "source_name": article["source"]["name"],
                "source_date": article["publishedAt"],
                "title": article["title"],
            }
        }


        with open(f"../api_data/{coin_name}/{coin_name}_{str(datetime.now(timezone.utc))}.json", "w") as f:
            json.dump(article_dict, f, indent=4)

        time.sleep(.05)

    return 


def main():

    ## simple test; api endpoint for getting articles

    get_articles("bitcoin")


    return


if __name__ == "__main__":
    main()