# 
# Makefile module to test the CLI 
# 
# Copyright (c) 2016-2018 Roozbeh Farahbod
# Distributed under the MIT License.
#

REPO ?= $(shell git rev-parse --show-toplevel)
TEST_DIR ?= $(realpath test)

DHT22_MOCK_BIN ?= dht22-mock 
JUPITER_BIN ?= jupiter 
TEST_BIN ?= ./bin/home-weather
TEST_TEMP_FILE=.test-result

.PHONY: test-get-dep test-clean test-setup test-sensor-count test-step-list test-step-reading-one-sensor test-step-reading-list-of-sensors

test-get-dep:
	@echo "Skipping downloading of tools."

test-clean: 
	@echo stop all Jupiter services
	@killall jupiter || exit 0
	@echo stop all sensor mock services
	@killall dht22-mock || exit 0
	@sleep 3

test-setup: test-clean
	@echo Setting up test environment...
	@dht22-mock --port 8081 > /dev/null 2>&1 &
	@dht22-mock --port 8082 > /dev/null 2>&1 & 
	@dht22-mock --port 8083 > /dev/null 2>&1 &
	@sleep 3
	@jupiter --port 8080 -c test/3mocks.yml > /dev/null 2>&1 &
	@sleep 3

test-sensor-count: 
	$(MAKE) test-setup
	sleep 3
	curl http://localhost:8080/sensors | tr '"' '\n' | grep id | wc -l > $(TEST_TEMP_FILE)
	$(MAKE) test-clean
	read i < $(TEST_TEMP_FILE); if [ "$$i" != "3" ]; then echo "Expected 3 sensors, found $$i."; rm $(TEST_TEMP_FILE); exit 1; fi
	rm $(TEST_TEMP_FILE)

test-step-list:
	@rm -f $(TEST_TEMP_FILE)
	@touch $(TEST_TEMP_FILE)
	@echo Test list of sensors... 
	@$(TEST_BIN) list | wc -l > $(TEST_TEMP_FILE)
	@read i < $(TEST_TEMP_FILE); \
		if [ "$$i" != "4" ]; \
			then echo "List of sensors must show 3 sensors ($$i)."; rm $(TEST_TEMP_FILE); exit 1; \
			else echo "Passed."; \
		fi
	@rm -f $(TEST_TEMP_FILE)

test-step-reading-one-sensor:
	@rm -f $(TEST_TEMP_FILE)
	@touch $(TEST_TEMP_FILE)
	@echo Test reading one sensor... 
	@$(TEST_BIN) read livingroom | grep temperature | sed 's/^.*: //g' > $(TEST_TEMP_FILE)
	@read i < $(TEST_TEMP_FILE); \
		if [ -z "$$i" ]; \
			then echo "Failed reading a temperature value."; rm $(TEST_TEMP_FILE); exit 1; \
			else echo "Passed."; \
		fi
	@rm -f $(TEST_TEMP_FILE)

test-step-reading-list-of-sensors:
	@rm -f $(TEST_TEMP_FILE)
	@touch $(TEST_TEMP_FILE)
	@echo Test reading list of sensors... 
	@$(TEST_BIN) read -f kitchen,livingroom,bedroom | grep count | sed 's/^.*: //g' > $(TEST_TEMP_FILE)
	@read i < $(TEST_TEMP_FILE); \
		if [ $$i -eq 3 ]; \
			then echo "Passed."; \
			else echo "Failed reading all 3 sensors."; rm $(TEST_TEMP_FILE); exit 1; \
		fi
	@rm -f $(TEST_TEMP_FILE)

test: go-build test-setup
	@$(TEST_BIN) config set jupiter http://localhost:8080
	@$(MAKE) test-step-list
	@$(MAKE) test-step-reading-one-sensor
	@$(MAKE) test-step-reading-list-of-sensors
	@$(MAKE) test-clean

