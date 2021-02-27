from django.urls import path, include

urlpatterns = [
    path('api/', include('comments.urls_api'))
]
