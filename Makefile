WORK_DIR=/opt/ppalist/
STATIC_DIR=./static
PAGES_DIR=./pages
APP_NAME=ppa_list
INIT_D_FILE=ppa_list.sh
all:
	go build -o $(APP_NAME)

install:
	mkdir -p $(WORK_DIR)
	cp -rf $(STATIC_DIR) $(WORK_DIR)
	cp -rf $(PAGES_DIR) $(WORK_DIR)
	cp $(APP_NAME) $(WORK_DIR)
	cp $(INIT_D_FILE) /etc/init.d/$(APP_NAME)
	chmod +x /etc/init.d/$(APP_NAME)
	update-rc.d ppa_list defaults