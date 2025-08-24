from google.cloud import storage
import json

storage_client = storage.Client()


bucket_name = "raw-digested-data"
bucket = storage_client.get_bucket(bucket_name)



def read_folder_from_bucket(folder_name):

	blobs = bucket.list_blobs(prefix=folder_name)

	data = []

	for n, blob in enumerate(blobs):
		## skip the folder itself; it is empty but listed as a blob
		if blob.name == folder_name:
			continue

		print(f"Reading file: {blob.name}")

		# blob.download_to_filename(destination_name + "_" + str(n) + ".jpg")
		contents = json.loads(blob.download_as_bytes())


		content = contents["data"]["content"]

		data.append(content)



	# return text block joining all data

	return "\r".join(data)

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