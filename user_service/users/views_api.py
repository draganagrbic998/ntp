import datetime
import hashlib
import smtplib

from django.contrib.auth.models import User
from django.core.mail import send_mail
from rest_framework import status, permissions
from rest_framework.decorators import api_view, permission_classes
from rest_framework.response import Response
from rest_framework_jwt.serializers import jwt_payload_handler, jwt_encode_handler

from user_service.jwt import jwt_response_handler
from users.email import ACCOUNT_ACTIVATION_TITLE, ACCOUNT_ACTIVATION_TEXT
from users.models import CustomUser
from users.serializers import RegistrationSerializer, ProfileSerializer


def generate_link(email: str, timestamp: datetime.datetime):
    src = email + timestamp.strftime('%Y-%m-%d %H:%M')
    return hashlib.sha1(src.encode('utf-8')).hexdigest()


@api_view(['POST'])
@permission_classes([permissions.AllowAny])
def register(request):
    serializer = RegistrationSerializer(data=request.data)
    if serializer.is_valid():
        try:
            User.objects.get(email=request.data['email'])
            return Response(status=status.HTTP_409_CONFLICT)
        except User.DoesNotExist:
            user = User.objects.create_user(
                username=serializer.data['email'],
                email=serializer.data['email'],
                password=serializer.data['password'],
                is_active=False,
                is_staff=False,
                first_name=serializer.data['first_name'],
                last_name=serializer.data['last_name']
            )
            user.save()
            customer = CustomUser(
                user=user,
                address=serializer.data['address'],
                city=serializer.data['city'],
                zip_code=serializer.data['zip_code'],
                activation_link=generate_link(user.email, datetime.datetime.now())
            )
            customer.save()
            try:
                send_mail(ACCOUNT_ACTIVATION_TITLE,
                          ACCOUNT_ACTIVATION_TEXT % (
                              user.first_name,
                              user.last_name,
                              customer.activation_link
                          ),
                          'pcelarstvo.com',
                          [user.email],
                          fail_silently=True)

            except smtplib.SMTPException:
                pass

            return Response(data={'activation_link': customer.activation_link},
                            status=status.HTTP_200_OK, content_type='application/json')

    else:
        return Response(status=status.HTTP_400_BAD_REQUEST)


@api_view(['POST'])
def update_profile(request):
    serializer = ProfileSerializer(data=request.data)
    if serializer.is_valid():
        custom_user = CustomUser.objects.get(user_id=request.user.id)
        if serializer.data['new_password']:
            if not custom_user.user.check_password(serializer.data['old_password']):
                return Response(status=status.HTTP_400_BAD_REQUEST)
            custom_user.user.set_password(serializer.data['new_password'])

        custom_user.user.username = serializer.data['email']
        custom_user.user.email = serializer.data['email']
        custom_user.user.first_name = serializer.data['first_name']
        custom_user.user.last_name = serializer.data['last_name']
        custom_user.address = serializer.data['address']
        custom_user.city = serializer.data['city']
        custom_user.zip_code = serializer.data['zip_code']

        custom_user.user.save()
        custom_user.save()
        payload = jwt_payload_handler(custom_user.user)
        token = jwt_encode_handler(payload)

        return Response(data=jwt_response_handler(token, custom_user.user, request),
                        status=status.HTTP_200_OK, content_type='application/json')
    else:
        return Response(status=status.HTTP_400_BAD_REQUEST)
