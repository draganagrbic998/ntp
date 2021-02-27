from django.urls import path

from comments.views_api import manage_comments, DeleteComment, like, replies

urlpatterns = [
    path('comments', manage_comments),
    path('comments/<int:pk>', DeleteComment.as_view()),
    path('comments/<int:key>/like', like),
    path('comments/<int:key>/replies', replies)
]
