from google.cloud import storage
import json
import os

storage_client = storage.Client()


bucket_name = "raw-digested-data"
bucket = storage_client.get_bucket(bucket_name)



def upload_folder_to_bucket(source_folder_name, destination_folder_name):


	for filename in os.listdir(source_folder_name):

		blob = bucket.blob(destination_folder_name + filename)
		
		blob.upload_from_filename(source_folder_name + filename)

		print(f"File {source_folder_name + filename} uploaded to {destination_folder_name}.")
	
	return 

'''

def main():

	## simple test; reading blobs from an existing folder in bucket

	source_folder_name = "../api_data/bitcoin_test_4.json"
	destination_folder_name = "bitcoin_digested_data/"
	destination_name = "bitcoin_test_4.json"


	upload_file_to_bucket(source_folder_name, destination_folder_name, destination_name)


	return



if __name__ == "__main__":
	main()

'''
