from django.shortcuts import render

from users.models import CustomUser


def activate_account(request, key):
    try:
        custom_user = CustomUser.objects.get(activation_link=key)
        custom_user.user.is_active = True
        custom_user.user.save()
        return render(request, 'users/activation.html', context={'success': True})
    except CustomUser.DoesNotExist:
        return render(request, 'users/activation.html', context={'success': False})