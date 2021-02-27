from django.contrib import admin
from django.contrib.auth.admin import UserAdmin
from django.contrib.auth.models import User

from users.models import CustomUser


class CustomUserInline(admin.StackedInline):
    model = CustomUser
    can_delete = False
    verbose_name = 'custom user'
    verbose_name_plural = 'custom users'


class CustomUserAdmin(UserAdmin):
    inlines = (CustomUserInline,)


admin.site.unregister(User)
admin.site.register(User, CustomUserAdmin)
