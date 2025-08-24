from newsapi import NewsApiClient
from dotenv import load_dotenv
from datetime import date, timedelta, datetime, timezone
import json
import time
import os

load_dotenv()

api = NewsApiClient(api_key=os.getenv("NEWS_API_KEY"))


def get_articles(coin_name: str, dir_prefix, n=5):
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


        with open(f"{dir_prefix}{coin_name}/{coin_name}_{str(datetime.now(timezone.utc))}.json", "w") as f:
            json.dump(article_dict, f, indent=4)

        time.sleep(.05)

    return 


def clean_local_directory(coin_name: str, directory):

    ## removes already transferred api data from local directory
    dir_path = f"{directory}{coin_name}"

    for filename in os.listdir(dir_path):
        file_path = os.path.join(dir_path, filename)
        if os.path.isfile(file_path):
            os.remove(file_path)
            print(f"Deleted: {file_path}")

    return




def main():

    ## simple test; api endpoint for getting articles

    #get_articles("bitcoin")

    clean_local_directory("bitcoin", "../api_data/")

    return


if __name__ == "__main__":
    main()

