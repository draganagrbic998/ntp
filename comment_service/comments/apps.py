from django.apps import AppConfig
import os

from django.db import connection


class CommentsConfig(AppConfig):
    name = 'comments'

    def ready(self):
        delete_path = os.path.join(os.path.dirname(__file__), 'delete.sql')
        insert_path = os.path.join(os.path.dirname(__file__), 'insert.sql')
        with connection.cursor() as c:
            c.execute(open(delete_path).read())
            c.execute(open(insert_path).read())


