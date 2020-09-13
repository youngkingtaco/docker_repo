#!/bin/sh

set -e

python manage.py collectstatic --noinput

uwsgi --socket :8484 --master --enablethreads --module app.wsgi
