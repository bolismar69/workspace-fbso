# syntax=docker/dockerfile:1

FROM python:3.12-slim AS builder
WORKDIR /app

ENV VIRTUAL_ENV=/opt/venv
RUN python -m venv "$VIRTUAL_ENV"
ENV PATH="$VIRTUAL_ENV/bin:$PATH"

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY . ./

FROM python:3.12-slim
WORKDIR /app

ENV VIRTUAL_ENV=/opt/venv
ENV PATH="$VIRTUAL_ENV/bin:$PATH"
COPY --from=builder /opt/venv /opt/venv
COPY --from=builder /app /app

EXPOSE 8080
ENV PORT=8080
CMD ["python", "app.py"]
