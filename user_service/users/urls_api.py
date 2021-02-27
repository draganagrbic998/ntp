from django.urls import path

from users.views_api import register, update_profile

urlpatterns = [
    path('register', register),
    path('update-profile', update_profile)
]