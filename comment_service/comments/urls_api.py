from django.urls import path

from comments.views_api import manage_comments, like, replies, delete_comment, comments_statistic, likes_statistic, \
    dislikes_statistic

urlpatterns = [
    path('comments', manage_comments),
    path('comments/<int:key>', delete_comment),
    path('comments/<int:key>/like', like),
    path('comments/<int:key>/replies', replies),
    path('statistic-comments/<int:start>/<int:end>', comments_statistic),
    path('statistic-likes/<int:start>/<int:end>', likes_statistic),
    path('statistic-dislikes/<int:start>/<int:end>', dislikes_statistic)
]
