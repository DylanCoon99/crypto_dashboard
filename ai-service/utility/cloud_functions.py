from google.cloud import storage
import json
from datetime import date
import os

storage_client = storage.Client()

bucket_name = "raw-digested-data"
bucket = storage_client.get_bucket(bucket_name)



def read_folder_from_bucket(folder_name):

	today = date.today()

	print(f"FILE EXIST: {os.path.exists(f"./data/{folder_name}{today}")}")

	if not os.path.exists(f"./data/{folder_name}{today}"):
		blobs = bucket.list_blobs(prefix=folder_name)

		data = []


		for n, blob in enumerate(blobs):
			## skip the folder itself; it is empty but listed as a blob
			if blob.name in [folder_name, folder_name + ".DS_Store"]:
				continue

			print(f"Reading file: {blob.name}")

			# blob.download_to_filename(destination_name + "_" + str(n) + ".jpg")
			contents = json.loads(blob.download_as_bytes())


			content = contents["data"]["title"]

			data.append(content)

		# write the data to a local folder to store for future requests
		try:
			with open(f"./data/{folder_name}{today}", 'w', encoding='utf-8') as file:
				json.dump(data, file, indent=4)

		except IOError as e:
			print(f"Error writing to file")

		# return text block joining all data

		return data
	else:
		# we can save time by reading from local directory
		data = []

		print("Reading locally from cached data.")

		try:
			with open(f"./data/{folder_name}{today}", 'r') as file:
				data = json.load(file)

		except IOError as e:
			print(f"Error reading from file")

		return data







'''

def main():

	## simple test; reading blobs from an existing folder in bucket

	folder_name = "bitcoin_digested_data/"
	#destination = "./image"

	data = read_folder_from_bucket(folder_name)

	print(data)

	return



if __name__ == "__main__":
	main()

'''