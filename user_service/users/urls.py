from django.urls import path

from users.views import activate_account

urlpatterns = [
    path('activate/<str:key>', activate_account)
]