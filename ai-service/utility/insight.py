from transformers import T5Tokenizer, T5ForConditionalGeneration



def buildCSV(coin_name: str, data):

	## iterates over ingested data 

	for d in data:
		print(d)

	return




def main():

	model_name = "t5-small"

	tokenizer = T5Tokenizer.from_pretrained(model_name)

	model = T5ForConditionalGeneration.from_pretrained(model_name)

	input_text = "summarize: Bitcoin surges to new high. Ethereum gains traction with new partnerships."

	inputs = tokenizer(input_text, return_tensors="pt")

	outputs = model.generate(**inputs)

	print(tokenizer.decode(outputs[0], skip_special_tokens=True))

	return



if __name__ == "__main__":
	main()