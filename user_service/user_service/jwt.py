from users.models import CustomUser


def jwt_response_handler(token, user=None, request=None):
    try:
        custom_user = CustomUser.objects.get(user_id=user.id)
        return {
            'id': custom_user.id,
            'token': token,
            'email': custom_user.user.email,
            'first_name': custom_user.user.first_name,
            'last_name': custom_user.user.last_name,
            'address': custom_user.address,
            'city': custom_user.city,
            'zip_code': custom_user.zip_code
        }
    except CustomUser.DoesNotExist:
        return None
