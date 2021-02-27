from django.db import models


class Comment(models.Model):
    created_on = models.DateTimeField(auto_now_add=True, blank=True)    #jel ovo ok?
    user_id = models.IntegerField(default=0)    #jel mora ovo default?
    email = models.CharField(max_length=100)
    product_id = models.IntegerField(default=0) #jel mora ovo default?
    text = models.TextField(max_length=1000)
    parent = models.ForeignKey('self', null=True, blank=True, on_delete=models.CASCADE)


class Like(models.Model):
    comment = models.ForeignKey(Comment, on_delete=models.CASCADE)
    user_id = models.IntegerField(default=0)    #jel mora ovo default?
    dislike = models.BooleanField()

