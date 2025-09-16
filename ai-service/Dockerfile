# Stage 1: Build dependencies
FROM python:3.10-slim AS builder

WORKDIR /app

# Install dependencies first for caching
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Stage 2: Final Image
FROM python:3.10-slim

WORKDIR /app

# Copy installed deps from builder
COPY --from=builder /usr/local/lib/python3.10/site-packages /usr/local/lib/python3.10/site-packages
COPY --from=builder /usr/local/bin /usr/local/bin

# Copy app code
COPY . .


## ENV GOOGLE_APPLICATION_CREDENTIALS=/app/credentials/key.json


# Expose fastapi port
EXPOSE 8000

# Run the app (use uvicorn)
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]
