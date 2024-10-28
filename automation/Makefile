APP_NAME=automation
TAG=0.0.1

.PHONY: test
test:
	docker compose up

.PHONY: test_api
test_api:
	pytest tests/api

.PHONY: test_ui_f
test_ui_f:
	pytest tests/ui --browser firefox

.PHONY: test_ui_hl
test_ui_c:
	pytest tests/ui --browser chrome

.PHONY: clean
clean:
	find . | grep -E "(__pycache__|.\.pyc|.\.pyo)" | xargs rm -rf

.PHONY: clean_reports
clean_reports:
	find . | grep -E "reports\/(.*\.html|screenshots\/.*\.png)"
