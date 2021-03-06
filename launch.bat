venv\scripts\python user_service/manage.py migrate 
start venv\scripts\python user_service/manage.py runserver 
venv\scripts\python comment_service/manage.py migrate
start venv\scripts\python comment_service/manage.py runserver 8003
cd ad_service
set GO111MODULE=auto
start go run .
cd ..\event_service
set GO111MODULE=auto
start go run .
cd ..