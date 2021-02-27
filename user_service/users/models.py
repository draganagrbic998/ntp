from django.contrib.auth.models import User
from django.db import models


class CustomUser(models.Model):
    user = models.OneToOneField(User, on_delete=models.CASCADE)
    address = models.CharField(max_length=100)
    city = models.CharField(max_length=100)
    zip_code = models.CharField(max_length=100)
    activation_link = models.CharField(null=True, blank=True, max_length=100)

    def __str__(self):
        return f'{self.user.first_name} {self.user.last_name}'

    class Meta:
        verbose_name = 'custom user'
        verbose_name_plural = 'custom users'

