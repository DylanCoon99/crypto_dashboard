from transformers import T5Tokenizer, T5ForConditionalGeneration, Trainer, TrainingArguments
from datasets import load_dataset

# Load model and tokenizer
model_name = "t5-small"
tokenizer = T5Tokenizer.from_pretrained(model_name)
model = T5ForConditionalGeneration.from_pretrained(model_name)

# Load dataset
dataset = load_dataset("csv", data_files="../data/training-data/train.csv")
def preprocess(examples):
    inputs = ["summarize: " + h for h in examples["headlines"]]
    return tokenizer(inputs, text_target=examples["insight"], padding="max_length", truncation=True, max_length=512)
tokenized_dataset = dataset.map(preprocess, batched=True)

# Training arguments
training_args = TrainingArguments(
    output_dir="./fine-tuned-t5",
    num_train_epochs=3,
    per_device_train_batch_size=8,
    save_strategy="epoch",
)

# Trainer
trainer = Trainer(
    model=model,
    args=training_args,
    train_dataset=tokenized_dataset["train"],
)

# Fine-tune
trainer.train()

# Save locally
model.save_pretrained("./fine-tuned-t5")
tokenizer.save_pretrained("./fine-tuned-t5")