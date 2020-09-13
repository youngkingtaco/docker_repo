#!/bin/sh

set -e

python manage.py collectstatic --noinput

uswgi --socket :8383 --master --enable-threads --module app.wsgi
