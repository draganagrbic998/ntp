from django.urls import path

from comments.views_api import manage_comments, like, replies, delete_comment

urlpatterns = [
    path('comments', manage_comments),
    path('comments/<int:key>', delete_comment),
    path('comments/<int:key>/like', like),
    path('comments/<int:key>/replies', replies)
]
