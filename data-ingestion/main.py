## Cloud Functions

'''

source_folder_name = "../api_data/bitcoin_test_4.json"
destination_folder_name = "bitcoin_digested_data/"
destination_name = "bitcoin_test_4.json"


upload_file_to_bucket(source_folder_name, destination_folder_name, destination_name)

'''


import functions_framework


@functions_framework.http
def hello_get(request):

	return "Hello World!"