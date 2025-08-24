import nltk
from nltk.sentiment.vader import SentimentIntensityAnalyzer



# add error handling

nltk.download('vader_lexicon', quiet=True)


def AnalyzeSentiment(data):

	analyzer = SentimentIntensityAnalyzer()

	sentiment_scores = analyzer.polarity_scores(data)

	return sentiment_scores