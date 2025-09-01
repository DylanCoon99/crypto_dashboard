from transformers import T5Tokenizer, T5ForConditionalGeneration

# Load fine-tuned model
model_path = "./model/fine-tuned-t5/"
tokenizer = T5Tokenizer.from_pretrained(model_path)
model = T5ForConditionalGeneration.from_pretrained(model_path)




def getInsight(headlines):

	# Generate insight
	input_text = f"summarize: {' '.join(headlines)}"
	inputs = tokenizer(input_text, return_tensors="pt", max_length=512, truncation=True)
	outputs = model.generate(**inputs, max_length=100)
	insight = tokenizer.decode(outputs[0], skip_special_tokens=True)

	return insight