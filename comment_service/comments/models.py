from django.db import models
from django.utils import timezone


class Comment(models.Model):
    created_on = models.DateTimeField(default=timezone.now)
    user_id = models.IntegerField(default=0)
    email = models.CharField(max_length=100)
    product_id = models.IntegerField(default=0)
    text = models.TextField(max_length=1000)
    parent = models.ForeignKey('self', null=True, blank=True, on_delete=models.CASCADE)


class Like(models.Model):
    created_on = models.DateTimeField(default=timezone.now)
    comment = models.ForeignKey(Comment, on_delete=models.CASCADE)
    user_id = models.IntegerField(default=0)
    dislike = models.BooleanField()

