from google.cloud import storage

storage_client = storage.Client()


bucket_name = "raw-digested-data"
bucket = storage_client.get_bucket(bucket_name)



def read_folder_from_bucket(folder_name, destination_name):

	blobs = bucket.list_blobs(prefix=folder_name)


	for n, blob in enumerate(blobs):
		## skip the folder itself; it is empty but listed as a blob
		if blob.name == folder_name:
			continue

		print(f"Reading file: {blob.name}")

		blob.download_to_filename(destination_name + "_" + str(n) + ".jpg")


	return



def main():

	## simple test; reading blobs from an existing folder in bucket

	folder_name = "test_folder/"
	destination = "./image"

	read_folder_from_bucket(folder_name, destination)

	return



if __name__ == "__main__":
	main()