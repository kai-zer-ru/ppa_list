WORK_DIR=/opt/ppalist/
STATIC_DIR=./static
PAGES_DIR=./pages
CONF_FILE=main.cfg
PPA_LIST=/opt/ppalist/ppalist
SOFT_LIST=/opt/ppalist/softlist
SOURCE_LIST=/opt/ppalist/sourcelist
APP_NAME=ppa_list
INIT_D_FILE=ppa_list.sh

all:
	go build -o $(APP_NAME)

clean:
	rm $(APP_NAME)

install:
	mkdir -p $(WORK_DIR)
	cp -rf $(STATIC_DIR) $(WORK_DIR)
	cp -rf $(PAGES_DIR) $(WORK_DIR)
	cp $(APP_NAME) $(WORK_DIR)
	touch $(PPA_LIST)
	touch $(SOFT_LIST)
	touch $(SOURCE_LIST)
	cp example.conf $(WORK_DIR)/$(CONF_FILE)
	cp $(INIT_D_FILE) /etc/init.d/$(APP_NAME)
	chmod +x /etc/init.d/$(APP_NAME)
	update-rc.d ppa_list defaults

uninstall:
	service $(APP_NAME) stop
	update-rc.d -f $(APP_NAME) remove
	rm /etc/init.d/$(APP_NAME)
	rm -rf $(WORK_DIR)

reinstall:
	service $(APP_NAME) stop
	update-rc.d -f $(APP_NAME) remove
	rm /etc/init.d/$(APP_NAME)
	rm -rf $(WORK_DIR)
	mkdir -p $(WORK_DIR)
	cp -rf $(STATIC_DIR) $(WORK_DIR)
	cp -rf $(PAGES_DIR) $(WORK_DIR)
	cp $(APP_NAME) $(WORK_DIR)
	touch $(PPA_LIST)
	touch $(SOFT_LIST)
	touch $(SOURCE_LIST)
	cp example.conf $(WORK_DIR)/$(CONF_FILE)
	cp $(INIT_D_FILE) /etc/init.d/$(APP_NAME)
	chmod +x /etc/init.d/$(APP_NAME)
	update-rc.d ppa_list defaults