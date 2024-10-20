FROM python:3.11-bookworm

WORKDIR /automation
COPY . /automation

# Install required packages
RUN pip install -r requirements.txt

# Install Chrome and Firefox
RUN apt-get update && apt-get install firefox-esr -y
RUN apt-get install -y wget
RUN wget -q https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
RUN apt-get install -y ./google-chrome-stable_current_amd64.deb
